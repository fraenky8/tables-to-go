# Tables-to-Go
> convert your database tables to structs easily

A small and convenient tool supporting development against a changing 
database schema.

**Tables change, run the tool, get your structs!**

[![Go Report Card](https://goreportcard.com/badge/github.com/fraenky8/tables-to-go/v2)](https://goreportcard.com/report/github.com/fraenky8/tables-to-go/v2)
[![GoDoc](https://godoc.org/github.com/fraenky8/tables-to-go/v2?status.svg)](https://godoc.org/github.com/fraenky8/tables-to-go/v2)
[![Build & Test](https://github.com/fraenky8/tables-to-go/actions/workflows/go.yml/badge.svg)](https://github.com/fraenky8/tables-to-go/actions)
[![codecov](https://codecov.io/github/fraenky8/tables-to-go/graph/badge.svg?token=EACZFL68U9)](https://codecov.io/github/fraenky8/tables-to-go)

## Requirements

- Go 1.26+

## Install

This project provides a make file but can also simply be installed with the
go-install command.

Get the latest stable release version:

```
go install github.com/fraenky8/tables-to-go/v2@latest
```

Get the latest changes from master:

```
go install github.com/fraenky8/tables-to-go/v2@master
```

## Getting Started

```
tables-to-go -v -of ../path/to/my/models
```

This gets all tables of a local running PostgreSQL database. Therefore, it uses 
the database `postgres`, schema `public` and user `postgres` with no password.
Flag `-v` is verbose mode, `-of` is the output file path where the go files 
containing the structs will get created (default: current working directory).

## Features

* convert your tables to structs
* table with name `a_foo_bar` will become file `AFooBar.go` with struct `AFooBar`
* properly formatted files with imports
* optional generated file header via `-gen-header` following Go's `Code generated ... DO NOT EDIT.` convention
* automatically typed struct fields, either with `sql.Null*` or primitive `*builtinType`
pointer types
* struct fields with `db`-tags for ready to use in database code
* **currently supported**:
  * PostgreSQL (10, 11, 12, 17, 18 tested)
  * MySQL (5.5+, 8 tested)
  * SQLite (3 tested)
* currently, the following basic data types are supported:
  * numeric: integer, serial, double, real, float
  * character: varying, text, char, varchar, binary, varbinary, blob
  * date/time: timestamp, date, datetime, year, time with time zone, timestamp
    with time zone, time without time zone, timestamp without time zone
  * others: boolean
* **partial support for [GORM](https://gorm.io/)**
  * struct fields with `gorm` tags via `-tag gorm`
    * We only generate a subset of [available tags](https://gorm.io/docs/models.html#Fields-Tags)
    since we only ever read existing tables and assume to not create those tables
    via the generated structs by this tool. If you need full GORM support,
    use their [gen-tool](https://gorm.io/gen/gen_tool.html).
  * optional `gorm.Model` embedding via `-gorm-model`
* **partial support for [Masterminds/structable](https://github.com/Masterminds/structable)**
  * only primary key & auto increment columns supported
  * struct fields with `stbl` tags
  * ability to generate structs only for Masterminds/structable:
    * without `db`-tags
    * with or without `structable.Recorder`

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

The column `id` got automatically converted to upper-case to follow the 
idiomatic Go guidelines. 
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

Filter for specific tables via (multiple) `-table` flags:

```
tables-to-go -v -of ../path/to/my/models -table foobar -table foo,bar,baz
```

Generate files with a Go-style generated header:

```
tables-to-go -v -gen-header -table user
```

### Command-line Flags

Print usage with `-?` or `-help`

```
Usage of tables-to-go:
  -?	shows help and usage
  -d string
    	database name; for sqlite3, URL query params '_pragma=<fn()>' can be added, e.g. path/to/database.db?_pragma=busy_timeout(5000) (default "postgres")
  -f	force; skip tables that encounter errors
  -fn-format value
    	format of the filename: camelCase (c, default) or snake_case (s) (default c)
  -format value
    	format of struct fields (columns): camelCase (c) or original (o) (default c)
  -gen-header
    	add generated code header to output files
  -gorm-model
    	generate an embedded gorm.Model field
  -h string
    	host of database (default "127.0.0.1")
  -help
    	shows help and usage
  -no-initialism
    	disable the conversion to upper-case words in column names
  -null value
    	representation of NULL columns: sql.Null* (sql) or primitive pointers (native|primitive) (default sql)
  -of string
    	output file path, default is current working directory
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
  -socket string
    	The socket file to use for connection. If specified, takes precedence over host:port.
  -sslmode string
    	Connect to database using secure connection. (default "disable")
    	The value will be passed as is to the underlying driver.
    	Refer to this site for supported values: https://www.postgresql.org/docs/current/libpq-ssl.html
  -structable-recorder
    	generate a structable.Recorder field
  -suf string
    	suffix for file- and struct names
  -t value
    	type of database to use, currently supported: [pg mysql sqlite3] (default pg)
  -table value
    	Filter for the specified table(s). Can be used multiple times or with comma separated values without spaces. Example: -table foobar -table foo,bar,baz
  -tag value
     	List of struct tags. Can be used multiple times or with comma separated values without spaces. Example: -tag db -tag sqlx,json,gorm
     	Aliases: stbl => structable, sqlx => db
      	Any provided valid tag key is emitted as a struct tag, e.g. -tag json
  -tags-no-db
    	do not create db-tags
  -tags-structable
    	DEPRECATED: use -tag structable
  -tags-structable-only
    	DEPRECATED: use -tag structable with -tags-no-db (legacy only semantics still override extra custom tags)
  -u string
    	user to connect to the database
  -v	verbose output
  -version
    	show version and build information
  -vv
    	more verbose output
```

### Tags Behavior

Long term goal is to replace the single `-tags-*` flags with a single `-tag`
flag. To not break backwards compatibility in v2 the following behavior applies:

| Current Tag flags        | Equivalent                                         | Notes                                                          |
|--------------------------|----------------------------------------------------|----------------------------------------------------------------|
| `-tags-structable`       | `-tag structable` or explicit `-tag db,structable` | The `db` tag is still added implicitly (current v2 behavior)   |
| `-tags-structable-only`  | Combination of `-tag structable -tags-no-db`       | Legacy "only" semantics win and can override extra custom tags |
| `-tags-no-db`            | no change                                          | Explicitly disables `db` tag generation                        |
| `-tag db,json`           | new                                                | Any provided **valid** tag key is emitted                      |
| `-tag stbl,sqlx`         | new                                                | Aliases are applied                                            |

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
