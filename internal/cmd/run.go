package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"

	"github.com/fraenky8/tables-to-go/v2/internal/cli"
	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/output"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

const legacyTagsDeprecationWarning = "warning: -tags-structable and -tags-structable-only are deprecated; use -tags instead"

var (
	// ErrFlagParse is returned in case the parsing of flags returns an error.
	ErrFlagParse = errors.New("error parsing flags")
)

// Cmd represents the actual cli cmd to parse the args and set up the dependencies
// to run the actual cli.App.
type Cmd struct {
	db   database.Database
	info VersionInfo
}

// New creates a Cmd.
func New(info VersionInfo, db database.Database) Cmd {
	return Cmd{
		info: info,
		db:   db,
	}
}

// VersionInfo holds build and version information.
type VersionInfo struct {
	Revision       string
	VersionTag     string
	BuildTimestamp string
}

// Args represents the supported command line args.
type Args struct {
	usage func()

	*settings.Settings
	Version bool
	Help    bool
}

// Run runs the command with the provided arguments and IO streams.
func (c *Cmd) Run(ctx context.Context, args []string, stdout, stderr io.Writer) (err error) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cmdArgs, err := NewArgs(args, stderr)
	if err != nil {
		return err
	}

	if cmdArgs.Help {
		cmdArgs.usage()
		return nil
	}

	if cmdArgs.Version {
		printVersion(stdout, c.info)
		return nil
	}

	if err := cmdArgs.Verify(); err != nil {
		return err
	}

	if c.db == nil {
		c.db = database.New(cmdArgs.Settings)
	}

	if err := c.db.Connect(ctx); err != nil {
		return err
	}
	defer func(db database.Database) {
		if cErr := db.Close(); cErr != nil {
			if err != nil {
				err = fmt.Errorf("could not close database: %w, previous error: %w", cErr, err)
				return
			}
			err = fmt.Errorf("could not close database: %w", cErr)
		}
	}(c.db)

	writer := output.NewFileWriter(cmdArgs.OutputFilePath)

	app := cli.New(cmdArgs.Settings, c.db, writer, stderr)

	if err := app.Run(ctx); err != nil {
		return fmt.Errorf("run error: %w", err)
	}

	return nil
}

// NewArgs creates and prepares the command line arguments with default values.
func NewArgs(args []string, stderr io.Writer) (*Args, error) {
	if len(args) == 0 {
		args = []string{"tables-to-go"}
	}

	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)
	fs.SetOutput(stderr)

	a := Args{
		usage:    fs.Usage,
		Settings: settings.New(),
	}

	fs.BoolVar(&a.Help, "?", false, "shows help and usage")
	fs.BoolVar(&a.Help, "help", false, "shows help and usage")
	fs.BoolVar(&a.Verbose, "v", a.Verbose, "verbose output")
	fs.BoolVar(&a.VVerbose, "vv", a.VVerbose, "more verbose output")
	fs.BoolVar(&a.Version, "version", a.Version, "show version and build information")
	fs.BoolVar(&a.Force, "f", a.Force, "force; skip tables that encounter errors")

	fs.Var(&a.DbType, "t", fmt.Sprintf("type of database to use, currently supported: %v", settings.SprintfSupportedDbTypes()))
	fs.StringVar(&a.User, "u", a.User, "user to connect to the database")
	fs.StringVar(&a.Pswd, "p", a.Pswd, "password of user")
	fs.StringVar(&a.DbName, "d", a.DbName, "database name; for sqlite3, URL query params '_pragma=<fn()>' can be added, e.g. ?_pragma=busy_timeout(5000)")
	fs.StringVar(&a.Schema, "s", a.Schema, "schema name")
	fs.StringVar(&a.Host, "h", a.Host, "host of database")
	fs.StringVar(&a.Port, "port", a.Port, "port of database host, if not specified, it will be the default ports for the supported databases")
	fs.StringVar(&a.SSLMode, "sslmode", a.SSLMode, "Connect to database using secure connection. (default \"disable\")\nThe value will be passed as is to the underlying driver.\nRefer to this site for supported values: https://www.postgresql.org/docs/current/libpq-ssl.html")
	fs.StringVar(&a.Socket, "socket", a.Socket, "The socket file to use for connection. If specified, takes precedence over host:port.")
	fs.Var(&a.Tables, "table", "Filter for the specified table(s). Can be used multiple times or with comma separated values without spaces. Example: -table foobar -table foo,bar,baz")

	fs.StringVar(&a.OutputFilePath, "of", a.OutputFilePath, "output file path, default is current working directory")
	fs.Var(&a.OutputFormat, "format", "format of struct fields (columns): camelCase (c) or original (o)")

	fs.Var(&a.FileNameFormat, "fn-format", "format of the filename: camelCase (c, default) or snake_case (s)")
	fs.StringVar(&a.Prefix, "pre", a.Prefix, "prefix for file- and struct names")
	fs.StringVar(&a.Suffix, "suf", a.Suffix, "suffix for file- and struct names")
	fs.StringVar(&a.PackageName, "pn", a.PackageName, "package name")
	fs.Var(&a.Null, "null", "representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive)")

	fs.BoolVar(&a.NoInitialism, "no-initialism", a.NoInitialism, "disable the conversion to upper-case words in column names")

	fs.Var(&a.Tags, "tags", "List of struct tags. Can be used multiple times or with comma separated values without spaces. Example: -tags db -tags sqlx,json")
	fs.BoolVar(&a.TagsNoDb, "tags-no-db", a.TagsNoDb, "do not create db-tags")
	fs.BoolVar(&a.TagsMastermindStructable, "tags-structable", a.TagsMastermindStructable, "DEPRECATED: use -tags structable")
	fs.BoolVar(&a.TagsMastermindStructableOnly, "tags-structable-only", a.TagsMastermindStructableOnly, "DEPRECATED: use -tags structable")
	fs.BoolVar(&a.IsMastermindStructableRecorder, "structable-recorder", a.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	// NOOP to disable the print of usage when an error occurs.
	fs.Usage = func() {}

	err := fs.Parse(args[1:])
	if err != nil {
		// Note that we ignore the original error here and return our sentinel
		// value to be detected in main. Reason is that the flag package itself
		// prints any parsing error already to stderr and hence by not returning
		// it here we avoid printing it twice to stderr.
		return nil, ErrFlagParse
	}

	printLegacyTagsWarning(stderr, a.Settings)

	return &a, nil
}

func printLegacyTagsWarning(w io.Writer, s *settings.Settings) {
	if !s.TagsMastermindStructable && !s.TagsMastermindStructableOnly {
		return
	}

	_, _ = fmt.Fprintln(w, legacyTagsDeprecationWarning)
}

func printVersion(w io.Writer, info VersionInfo) {
	var (
		goOS, goArch = runtime.GOOS, runtime.GOARCH
	)
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		if info.VersionTag == "" {
			info.VersionTag = buildInfo.Main.Version
		}
		for _, s := range buildInfo.Settings {
			switch s.Key {
			case "vcs.revision":
				info.Revision = s.Value[:min(8, len(s.Value))]
			case "GOOS":
				goOS = s.Value
			case "GOARCH":
				goArch = s.Value
			}
		}
	}

	_, _ = fmt.Fprintf(w, "tables-to-go/%s-%s %s/%s built with %s",
		info.VersionTag, info.Revision, goOS, goArch, runtime.Version())

	if info.BuildTimestamp != "" {
		_, _ = fmt.Fprintf(w, " on %s", info.BuildTimestamp)
	}
	_, _ = fmt.Fprintln(w)
}
