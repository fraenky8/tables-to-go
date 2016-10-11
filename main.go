// Tables-to-Go
//
// convert your database tables to structs easily
//
// A small and helpful tool which helps during developing with a changing database schema.
//
// Example
//
// Assuming you have the following table definition (PostgreSQL):
//
//      CREATE TABLE some_user_info  (
//          id SERIAL NOT NULL PRIMARY KEY,
//          first_name VARCHAR(20),
//          last_name  VARCHAR(20) NOT NULL,
//          height DECIMAL
//      );
//
// Run the following command (default local PostgreSQL instance):
//
//      go run main.go
//
// The following file SomeUserInfo.go with default package dto (data transfer object) will be created:
//
//      package dto
//
//      import (
//          "database/sql"
//      )
//
//      type SomeUserInfo struct {
//          Id        int             `db:"id"`
//          FirstName sql.NullString  `db:"first_name"`
//          LastName  string          `db:"last_name"`
//          Height    sql.NullFloat64 `db:"height"`
//      }
//
// Commandline Flags
//
//       go run main.go -help
//          -?    shows help and usage
//          -d string
//              database name (default "postgres")
//          -format string
//              camelCase (c) or original (o) (default "c")
//          -h string
//              host of database (default "127.0.0.1")
//          -help
//              shows help and usage
//          -of string
//              output file path (default "./output")
//          -p string
//              password of user
//          -pn string
//              package name (default "dto")
//          -port string
//              port of database host, if not specified, it will be the default ports for the supported databases
//          -pre string
//              prefix for file- and struct names
//          -s string
//              schema name (default "public")
//          -structable-recorder
//              generate a structable.Recorder field
//          -suf string
//              suffix for file- and struct names
//          -t string
//              type of database to use, currently supported: [pg mysql] (default "pg")
//          -tags-no-db
//              do not create db-tags
//          -tags-structable
//              generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)
//          -tags-structable-only
//              generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)
//          -u string
//              user to connect to the database (default "postgres")
//          -v    verbose output
//
//
// For more details & exmaples refer to https://github.com/fraenky8/tables-to-go/blob/master/README.md
//
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

	flag.BoolVar(&cmdArgs.TagsNoDb, "tags-no-db", cmdArgs.TagsNoDb, "do not create db-tags")

	flag.BoolVar(&cmdArgs.TagsMastermindStructable, "tags-structable", cmdArgs.TagsMastermindStructable, "generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&cmdArgs.TagsMastermindStructableOnly, "tags-structable-only", cmdArgs.TagsMastermindStructableOnly, "generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)")
	flag.BoolVar(&cmdArgs.IsMastermindStructableRecorder, "structable-recorder", cmdArgs.IsMastermindStructableRecorder, "generate a structable.Recorder field")

	flag.BoolVar(&cmdArgs.TagsSql, "experimental-tags-sql", cmdArgs.TagsSql, "generate struct with sql-tags")
	flag.BoolVar(&cmdArgs.TagsSqlOnly, "experimental-tags-sql-only", cmdArgs.TagsSqlOnly, "generate struct with ONLY sql-tags")

	flag.Parse()

	return cmdArgs
}
