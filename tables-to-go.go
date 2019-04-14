package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fraenky8/tables-to-go/internal/cli"
	"github.com/fraenky8/tables-to-go/pkg/config"
	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/output"
)

// CmdArgs represents the supported command line args
type CmdArgs struct {
	Help bool
	*config.Settings
}

// NewCmdArgs creates and prepares the command line arguments with default values
func NewCmdArgs() (args *CmdArgs) {

	args = &CmdArgs{
		Settings: config.NewSettings(),
	}

	flag.BoolVar(&args.Help, "?", false, "shows help and usage")
	flag.BoolVar(&args.Help, "help", false, "shows help and usage")
	flag.BoolVar(&args.Verbose, "v", args.Verbose, "verbose output")
	flag.BoolVar(&args.VVerbose, "vv", args.VVerbose, "more verbose output")

	flag.Var(&args.DbType, "t", fmt.Sprintf("type of database to use, currently supported: %v", config.SupportedDbTypes()))
	flag.StringVar(&args.User, "u", args.User, "user to connect to the database")
	flag.StringVar(&args.Pswd, "p", args.Pswd, "password of user")
	flag.StringVar(&args.DbName, "d", args.DbName, "database name")
	flag.StringVar(&args.Schema, "s", args.Schema, "schema name")
	flag.StringVar(&args.Host, "h", args.Host, "host of database")
	flag.StringVar(&args.Port, "port", args.Port, "port of database host, if not specified, it will be the default ports for the supported databases")

	flag.StringVar(&args.OutputFilePath, "of", args.OutputFilePath, "output file path, default is current working directory")
	flag.Var(&args.OutputFormat, "format", "format of struct fields (columns): camelCase (c) or original (o)")

	flag.StringVar(&args.Prefix, "pre", args.Prefix, "prefix for file- and struct names")
	flag.StringVar(&args.Suffix, "suf", args.Suffix, "suffix for file- and struct names")
	flag.StringVar(&args.PackageName, "pn", args.PackageName, "package name")
	flag.Var(&args.Null, "null", "representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive)")

	flag.BoolVar(&args.NoInitialism, "no-initialism", args.NoInitialism, "disable the conversion to upper-case words in column names")

	flag.BoolVar(&args.TagsNoDb, "tags-no-db", args.TagsNoDb, "do not create db-tags")

	flag.BoolVar(&args.TagsMastermindStructable, "tags-structable", args.TagsMastermindStructable, "generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&args.TagsMastermindStructableOnly, "tags-structable-only", args.TagsMastermindStructableOnly, "generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&args.IsMastermindStructableRecorder, "structable-recorder", args.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	flag.BoolVar(&args.TagsSQL, "experimental-tags-sql", args.TagsSQL, "generate struct with sql-tags")
	flag.BoolVar(&args.TagsSQLOnly, "experimental-tags-sql-only", args.TagsSQLOnly, "generate struct with ONLY sql-tags")

	// disable the print of usage when an error occurs
	flag.CommandLine.Usage = func() {}

	flag.Parse()

	return args
}

// main function to run the transformations
func main() {

	cmdArgs := NewCmdArgs()

	if cmdArgs.Help {
		flag.Usage()
		os.Exit(0)
	}

	if err := cmdArgs.Verify(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	db := database.New(cmdArgs.Settings)

	if err := db.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writer := output.NewFileWriter(cmdArgs.OutputFilePath)

	if err := cli.Run(cmdArgs.Settings, db, writer); err != nil {
		fmt.Printf("run error: %v", err)
		os.Exit(1)
	}
}
