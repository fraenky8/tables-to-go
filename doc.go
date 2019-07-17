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
//      go run tables-to-go.go
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
//          ID        int             `db:"id"`
//          FirstName sql.NullString  `db:"first_name"`
//          LastName  string          `db:"last_name"`
//          Height    sql.NullFloat64 `db:"height"`
//      }
//
// Commandline Flags
//
//       go run tables-to-go.go -help
//          -?	shows help and usage
//          -d string
//            	database name (default "postgres")
//          -f
//            	force, skip tables that encounter errors but construct all others
//          -format string
//            	format of struct fields (columns): camelCase (c) or original (o) (default "c")
//          -h string
//            	host of database (default "127.0.0.1")
//          -help
//            	shows help and usage
//          -no-initialism
//      	  	disable the conversion to upper-case words in column names
//          -null string
//       	  	representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive)  (default "sql")
//          -of string
//            	output file path (default "current working directory")
//          -p string
//            	password of user
//          -pn string
//            	package name (default "dto")
//          -port string
//            	port of database host, if not specified, it will be the default ports for the supported databases
//          -pre string
//            	prefix for file- and struct names
//          -s string
//            	schema name (default "public")
//          -structable-recorder
//            	generate a structable.Recorder field
//          -suf string
//            	suffix for file- and struct names
//          -t string
//            	type of database to use, currently supported: [pg mysql] (default "pg")
//          -tags-no-db
//            	do not create db-tags
//          -tags-structable
//            	generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)
//          -tags-structable-only
//            	generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)
//          -u string
//            	user to connect to the database (default "postgres")
//          -v	verbose output
//          -vv
//            	more verbose output
//
//
// For more details & exmaples refer to https://github.com/fraenky8/tables-to-go/blob/master/README.md
//
package main
