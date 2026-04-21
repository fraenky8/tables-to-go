// Tables-to-Go
//
// convert your database tables to structs easily
//
// A small and helpful tool which helps during developing with a changing database schema.
//
// Example.
//
// Assuming you have the following table definition (PostgreSQL):
//
//	CREATE TABLE some_user_info  (
//	    id SERIAL NOT NULL PRIMARY KEY,
//	    first_name VARCHAR(20),
//	    last_name  VARCHAR(20) NOT NULL,
//	    height DECIMAL
//	);
//
// Run the following command (default local PostgreSQL instance):
//
//	go run tables-to-go.go
//
// The following file SomeUserInfo.go with default package dto (data transfer object) will be created:
//
//	package dto
//
//	import (
//	    "database/sql"
//	)
//
//	type SomeUserInfo struct {
//	    ID        int             `db:"id"`
//	    FirstName sql.NullString  `db:"first_name"`
//	    LastName  string          `db:"last_name"`
//	    Height    sql.NullFloat64 `db:"height"`
//	}
//
// Commandline Flags
//
//	go run tables-to-go.go -help
//		-?	shows help and usage
//		-d string
//		  	database name; for sqlite3, URL query params '_pragma=<fn()>' can be added, e.g. ?_pragma=busy_timeout(5000) (default "postgres")
//		-f	force; skip tables that encounter errors
//		-fn-format value
//		  	format of the filename: camelCase (c, default) or snake_case (s) (default c)
//		-format value
//		  	format of struct fields (columns): camelCase (c) or original (o) (default c)
//		-h string
//		  	host of database (default "127.0.0.1")
//		-help
//		  	shows help and usage
//		-no-initialism
//		  	disable the conversion to upper-case words in column names
//		-null value
//		  	representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive) (default sql)
//		-of string
//		  	output file path, default is current working directory
//		-p string
//		  	password of user
//		-pn string
//		  	package name (default "dto")
//		-port string
//		  	port of database host, if not specified, it will be the default ports for the supported databases
//		-pre string
//		  	prefix for file- and struct names
//		-s string
//		  	schema name (default "public")
//		-socket string
//		  	The socket file to use for connection. If specified, takes precedence over host:port.
//		-sslmode string
//		  	Connect to database using secure connection. (default "disable")
//		  	The value will be passed as is to the underlying driver.
//		  	Refer to this site for supported values: https://www.postgresql.org/docs/current/libpq-ssl.html
//		-structable-recorder
//		  	generate a structable.Recorder field
//		-suf string
//		  	suffix for file- and struct names
//		-t value
//		  	type of database to use, currently supported: [pg mysql sqlite3] (default pg)
//		-table value
//		  	Filter for the specified table(s). Can be used multiple times or with comma separated values without spaces. Example: -table foobar -table foo,bar,baz
//		-tags value
//		  	List of struct tags. Can be used multiple times or with comma separated values without spaces. Example: -tags db -tags sqlx,json
//		  	Aliases: stbl => structable, sqlx => db
//		  	Any provided tag name is emitted as a struct tag, e.g. -tags json
//		-tags-no-db
//		  	do not create db-tags
//		-tags-structable
//		  	DEPRECATED: use -tags structable
//		-tags-structable-only
//		  	DEPRECATED: use -tags structable with -tags-no-db (legacy only semantics still override extra custom tags)
//		-u string
//		  	user to connect to the database
//		-v	verbose output
//		-version
//		  	show version and build information
//		-vv
//		  	more verbose output
//
// For more details & examples refer to https://github.com/fraenky8/tables-to-go/blob/master/README.md
package main
