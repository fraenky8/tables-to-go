package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fraenky8/tables-to-go/src"
	"github.com/fraenky8/tables-to-go/src/settings"
)

// CmdArgs represents the supported command line args
type CmdArgs struct {
	Help bool
	*settings.Settings
}

// main function to run the transformations
func main() {

	cmdArgs := prepareCmdArgs()

	if cmdArgs.Help {
		flag.Usage()
		os.Exit(0)
	}

	if err := settings.VerifySettings(cmdArgs.Settings); err != nil {
		fmt.Printf("settings verification error: %v", err)
		os.Exit(1)
	}

	if err := tablestogo.Run(cmdArgs.Settings); err != nil {
		fmt.Printf("run error %v", err)
		os.Exit(1)
	}
}

// helper function to handle and prepare the command line arguments with default values
func prepareCmdArgs() (cmdArgs *CmdArgs) {

	cmdArgs = &CmdArgs{
		Settings: settings.NewSettings(),
	}

	flag.BoolVar(&cmdArgs.Help, "?", false, "shows help and usage")
	flag.BoolVar(&cmdArgs.Help, "help", false, "shows help and usage")
	flag.BoolVar(&cmdArgs.Verbose, "v", cmdArgs.Verbose, "verbose output")
	flag.StringVar(&cmdArgs.DbType, "t", cmdArgs.DbType, fmt.Sprintf("type of database to use, currently supported: %v", cmdArgs.PrettyPrintSupportedDbTypes()))
	flag.StringVar(&cmdArgs.User, "u", cmdArgs.User, "user to connect to the database")
	flag.StringVar(&cmdArgs.Pswd, "p", cmdArgs.Pswd, "password of user")
	flag.StringVar(&cmdArgs.DbName, "d", cmdArgs.DbName, "database name")
	flag.StringVar(&cmdArgs.Schema, "s", cmdArgs.Schema, "schema name")
	flag.StringVar(&cmdArgs.Host, "h", cmdArgs.Host, "host of database")
	flag.StringVar(&cmdArgs.Port, "port", cmdArgs.Port, "port of database host, if not specified, it will be the default ports for the supported databases")

	flag.StringVar(&cmdArgs.OutputFilePath, "of", cmdArgs.OutputFilePath, "output file path, default is current working directory")
	flag.StringVar(&cmdArgs.OutputFormat, "format", cmdArgs.OutputFormat, "camelCase (c) or original (o)")
	flag.StringVar(&cmdArgs.Prefix, "pre", cmdArgs.Prefix, "prefix for file- and struct names")
	flag.StringVar(&cmdArgs.Suffix, "suf", cmdArgs.Suffix, "suffix for file- and struct names")
	flag.StringVar(&cmdArgs.PackageName, "pn", cmdArgs.PackageName, "package name")

	flag.BoolVar(&cmdArgs.TagsNoDb, "tags-no-db", cmdArgs.TagsNoDb, "do not create db-tags")

	flag.BoolVar(&cmdArgs.TagsMastermindStructable, "tags-structable", cmdArgs.TagsMastermindStructable, "generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&cmdArgs.TagsMastermindStructableOnly, "tags-structable-only", cmdArgs.TagsMastermindStructableOnly, "generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&cmdArgs.IsMastermindStructableRecorder, "structable-recorder", cmdArgs.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	flag.BoolVar(&cmdArgs.TagsSQL, "experimental-tags-sql", cmdArgs.TagsSQL, "generate struct with sql-tags")
	flag.BoolVar(&cmdArgs.TagsSQLOnly, "experimental-tags-sql-only", cmdArgs.TagsSQLOnly, "generate struct with ONLY sql-tags")

	flag.Parse()

	return cmdArgs
}
