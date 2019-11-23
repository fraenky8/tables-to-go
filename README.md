# Tables-to-Go
> convert your database tables to structs easily

A small and convenient tool supporting development against a changing 
database schema.

**Tables change, run the tool, get your structs!**

[![Go Report Card](https://goreportcard.com/badge/github.com/fraenky8/tables-to-go)](https://goreportcard.com/report/github.com/fraenky8/tables-to-go)
[![GoDoc](https://godoc.org/github.com/fraenky8/tables-to-go?status.svg)](https://godoc.org/github.com/fraenky8/tables-to-go)
[![Build & Test](https://github.com/fraenky8/tables-to-go/workflows/Go/badge.svg)](https://github.com/fraenky8/tables-to-go/actions)
[![Code Coverage](https://scrutinizer-ci.com/g/fraenky8/tables-to-go/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/fraenky8/tables-to-go/?branch=master)

## Install

This project provides a make file but can also simply installed with the usual
go-get command:

```
go get github.com/fraenky8/tables-to-go
```

To enable SQLite3 support, run the make file:

```
make sqlite3
```

See [this PR](https://github.com/fraenky8/tables-to-go/pull/23) why it's 
disabled by default.

## Getting started

```
tables-to-go -v -of ../path/to/my/models
```

This gets all tables of a local running PostgreSQL database. Therefore it uses 
the database `postgres`, schema `public` and user `postgres` with no password.
Flag `-v` is verbose mode, `-of` is the output file path where the go files 
containing the structs will get created (default: current working directory).

## Features

* convert your tables to structs
* table with name `a_foo_bar` will become file `AFooBar.go` with struct `AFooBar`
* properly formated files with imports
* automatically typed struct fields, either with `sql.Null*` or primitve 
pointer types
* struct fields with `db`-tags for ready to use in database code
* **partial support for [Masterminds/structable](https://github.com/Masterminds/structable)**
  * only primary key & auto increment columns supported
  * struct fields with `stbl` tags
  * ability to generate structs only for Masterminds/structable:
    * without `db`-tags
    * with or without `structable.Recorder` 
* **currently supported**:
  * PostgreSQL (9.5 tested)
  * MySQL (5.5+ tested)
  * SQLite (3 tested)
* currently the following basic data types are supported:
  * numeric: integer, serial, double, real, float
  * character: varying, text, char, varchar, binary, varbinary, blob
  * date/time: timestamp, date, datetime, year, time with time zone, timestamp 
  with time zone, time without time zone, timestamp without time zone
  * others: boolean

## Examples

Assuming you have the following table definition (PostgreSQL):

```sql
CREATE TABLE some_user_info  (
  id SERIAL NOT NULL PRIMARY KEY,
  first_name VARCHAR(20),
  last_name  VARCHAR(20) NOT NULL,
  height DECIMAL
);
```

Run the following command (default local PostgreSQL instance):

```
tables-to-go
```

The following file `SomeUserInfo.go` with default package `dto` (data transfer 
object) will be created:

```go
package dto

import (
	"database/sql"
)

type SomeUserInfo struct {
	ID        int             `db:"id"`
	FirstName sql.NullString  `db:"first_name"`
	LastName  string          `db:"last_name"`
	Height    sql.NullFloat64 `db:"height"`
}
```

The column `id` got autmatically converted to upper-case to follow the idiomatic
go guidelines. 
See [here](https://github.com/golang/go/wiki/CodeReviewComments#initialisms) 
for more details. 
Words which gets converted can be found 
[here](https://github.com/fraenky8/tables-to-go/blob/master/internal/cli/tables-to-go-cli.go#L31).
<br>
This behaviour can be disabled by providing the command-line flag `-no-initialism`.

Running on remote database server (eg. Mysql@Docker)

```
tables-to-go -v -t mysql -h 192.168.99.100 -d testdb -u root -p mysecretpassword
```

PostgreSQL example with different default schema but default database `postgres`:

```
tables-to-go -v -t pg -h 192.168.99.100 -s test -u postgres -p mysecretpassword
```

Note: since database type `pg` is default, following command will be equivalent:

```
tables-to-go -v -h 192.168.99.100 -s test -u postgres -p mysecretpassword
```

You can also specify the package or prefix and suffix.

```
tables-to-go -v -t mysql -h 192.168.99.100 -d testdb -u root -p mysecretpassword -pn models -pre model_ -suf _model
```

With same table given above, following file with Name `ModelSomeUserInfoModel.go`
will be created:

```go
package models

import (
	"database/sql"
)

type ModelSomeUserInfoModel struct {
	ID        int             `db:"id"`
	FirstName sql.NullString  `db:"first_name"`
	LastName  string          `db:"last_name"`
	Height    sql.NullFloat64 `db:"height"`
}
```

### Command-line Flags

Print usage with `-?` or `-help`

```
tables-to-go -help
  -?	shows help and usage
  -d string
    	database name (default "postgres")
  -f
        force, skip tables that encounter errors but construct all others
  -format string
    	format of struct fields (columns): camelCase (c) or original (o) (default "c")
  -h string
    	host of database (default "127.0.0.1")
  -help
    	shows help and usage
  -no-initialism
      	disable the conversion to upper-case words in column names
  -null string
       	representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive)  (default "sql")
  -of string
    	output file path (default "current working directory")
  -p string
    	password of user
  -pn string
    	package name (default "dto")
  -port string
    	port of database host, if not specified, it will be the default ports for the supported databases
  -pre string
    	prefix for file- and struct names
  -s string
    	schema name (default "public")
  -structable-recorder
    	generate a structable.Recorder field
  -suf string
    	suffix for file- and struct names
  -t string
    	type of database to use, currently supported: [pg mysql] (default "pg")
  -tags-no-db
    	do not create db-tags
  -tags-structable
    	generate struct with tags for use in Masterminds/structable (https://github.com/Masterminds/structable)
  -tags-structable-only
    	generate struct with tags ONLY for use in Masterminds/structable (https://github.com/Masterminds/structable)
  -u string
    	user to connect to the database (default "postgres")
  -v	verbose output
  -vv
    	more verbose output
```

## Contributing

If you find any issues or missing a feature, feel free to contribute or make 
suggestions! You can fork the repository and use a feature branch too. Feel free
to send me a pull request. The PRs have to come with appropriate unit tests,
documentation of the added functionality and updated README with optional
examples.

To start developing clone via `git` or use go's get command to fetch this 
project.

This project uses [go modules](https://github.com/golang/go/wiki/Modules) so
make sure when adding new dependencies to update the `go.mod` file and the 
vendor directory:

```
go mod tidy
go mod vendor
```

## Licensing

The code in this project is licensed under MIT license.
