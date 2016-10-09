// main file
package main

import (
	"flag"
	"fmt"

	"github.com/fraenky8/tables-to-go/src"
)

// supported command line args
type CmdArgs struct {
	Help bool
	*tablestogo.Settings
}

// main function to run the transformations
func main() {

	cmdArgs := prepareCmdArgs()

	if cmdArgs.Help {
		flag.Usage()
		return
	}

	if err := tablestogo.Run(cmdArgs.Settings); err != nil {
		fmt.Println(err)
	}
}

// helper function to handle and prepare the command line arguments with default values
func prepareCmdArgs() (cmdArgs *CmdArgs) {

	cmdArgs = &CmdArgs{
		Settings: tablestogo.NewSettings(),
	}

	flag.BoolVar(&cmdArgs.Help, "?", false, "shows help and usage")
	flag.BoolVar(&cmdArgs.Help, "help", false, "shows help and usage")
	flag.BoolVar(&cmdArgs.Verbose, "v", cmdArgs.Verbose, "verbose output")
	flag.StringVar(&cmdArgs.DbType, "t", cmdArgs.DbType, fmt.Sprintf("type of database to use, currently supported: %v", tablestogo.SupportedDbTypes))
	flag.StringVar(&cmdArgs.User, "u", cmdArgs.User, "user to connect to the database")
	flag.StringVar(&cmdArgs.Pswd, "p", cmdArgs.Pswd, "password of user")
	flag.StringVar(&cmdArgs.DbName, "d", cmdArgs.DbName, "database name")
	flag.StringVar(&cmdArgs.Schema, "s", cmdArgs.Schema, "schema name")
	flag.StringVar(&cmdArgs.Host, "h", cmdArgs.Host, "host of database")
	flag.StringVar(&cmdArgs.Port, "port", cmdArgs.Port, "port of database host, if not specified, it will be the default ports for the supported databases")

	flag.StringVar(&cmdArgs.OutputFilePath, "of", cmdArgs.OutputFilePath, "output file path")
	flag.StringVar(&cmdArgs.OutputFormat, "format", cmdArgs.OutputFormat, "camelCase (c) or original (o)")
	flag.StringVar(&cmdArgs.Prefix, "pre", cmdArgs.Prefix, "prefix for file- and struct names")
	flag.StringVar(&cmdArgs.Suffix, "suf", cmdArgs.Suffix, "suffix for file- and struct names")
	flag.StringVar(&cmdArgs.PackageName, "pn", cmdArgs.PackageName, "package name")

	flag.BoolVar(&cmdArgs.TagsNoDb, "tags-no-db", cmdArgs.TagsNoDb, "do not create `db`-tags")

	flag.BoolVar(&cmdArgs.TagsMastermindStructable, "tags-structable", cmdArgs.TagsMastermindStructable, "generate struct for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&cmdArgs.TagsMastermindStructableOnly, "tags-structable-only", cmdArgs.TagsMastermindStructableOnly, "generate struct ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&cmdArgs.IsMastermindStructableRecorder, "structable-recorder", cmdArgs.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	flag.Parse()

	return cmdArgs
}
