package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"

	"github.com/fraenky8/tables-to-go/v2/internal/cli"
	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/output"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

var (
	revision       = "master"
	versionTag     = ""
	buildTimestamp = ""
)

// cmdArgs represents the supported command line args
type cmdArgs struct {
	usage func()
	*settings.Settings
	Version bool
	Help    bool
}

// newCmdArgs creates and prepares the command line arguments with default values
func newCmdArgs(args []string) *cmdArgs {
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)

	a := cmdArgs{
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

	fs.BoolVar(&a.TagsNoDb, "tags-no-db", a.TagsNoDb, "do not create db-tags")

	fs.BoolVar(&a.TagsMastermindStructable, "tags-structable", a.TagsMastermindStructable, "generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	fs.BoolVar(&a.TagsMastermindStructableOnly, "tags-structable-only", a.TagsMastermindStructableOnly, "generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	fs.BoolVar(&a.IsMastermindStructableRecorder, "structable-recorder", a.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	// NOOP to disable the print of usage when an error occurs.
	fs.Usage = func() {}

	// Ignore error since we are using flag.ExitOnError
	_ = fs.Parse(args[1:])

	return &a
}

// main function to run the transformations
func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Stderr); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stderr io.Writer) (err error) {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cmdArgs := newCmdArgs(args)

	if cmdArgs.Help {
		cmdArgs.usage()
		return nil
	}

	if cmdArgs.Version {
		printVersion(stderr)
		return nil
	}

	if err := cmdArgs.Verify(); err != nil {
		return err
	}

	db := database.New(cmdArgs.Settings)
	defer func(db database.Database) {
		if cErr := db.Close(); cErr != nil {
			if err != nil {
				err = fmt.Errorf("could not close database: %w, previous error: %w", cErr, err)
				return
			}
			err = fmt.Errorf("could not close database: %w", cErr)
		}
	}(db)

	if err := db.Connect(ctx); err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	writer := output.NewFileWriter(cmdArgs.OutputFilePath)

	app := cli.New(cmdArgs.Settings, db, writer)

	if err := app.Run(ctx); err != nil {
		return fmt.Errorf("run error: %w", err)
	}

	return nil
}

func printVersion(stderr io.Writer) {
	var (
		goOS, goArch = runtime.GOOS, runtime.GOARCH
	)
	info, ok := debug.ReadBuildInfo()
	if ok {
		if versionTag == "" {
			versionTag = info.Main.Version
		}
		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value[:8]
			case "GOOS":
				goOS = s.Value
			case "GOARCH":
				goArch = s.Value
			}
		}
	}

	_, _ = fmt.Fprintf(stderr, "tables-to-go/%s-%s %s/%s built with %s",
		versionTag, revision, goOS, goArch, runtime.Version())

	//goland:noinspection GoBoolExpressions
	if buildTimestamp != "" {
		_, _ = fmt.Fprintf(stderr, " on %s", buildTimestamp)
	}
	_, _ = fmt.Fprintln(stderr)
}
