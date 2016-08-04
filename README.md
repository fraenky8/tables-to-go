# Tables-to-Go
> convert your tables to structs easily

A small and helpful tool which helps during developing with a changing database schema.

**Tables change, run the tool, get your structs! Easy!**

## Installing

```
go get github.com/fraenky8/tables-to-go
```

## Getting started

```
go run main.go -v -of ../path/to/my/models
```

This grabs all tables of a local running PostgreSQL database. Therefore it uses the database `postgres`, schema `public` and user `postgres` with no password.
Flag `-v` is verbose mode, `-of` is the output file path in which the structs are created.

## Features

* convert your tables easily to structs
* table with name `a_foo_bar` will become file `AFooBar.go` with struct `AFooBar`
* properly formated files with imports
* automatically typed struct fields
* struct fields with tags for ready to use in database code
* currently supported: PostgreSQL, MySQL
* currently the following basic data types are supported:
  * numeric: integer, serial, double, real, float
  * character: varying, text, char, varchar, binary, varbinary, blob
  * date/time: timestamp, date, datetime, year, time with time zone, timestamp with time zone, time without time zone, timestamp without time zone
  * others: boolean

## Examples

Assuming you have the following table definition:

```sql
CREATE TABLE some_user_info  (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(20),
  last_name  VARCHAR(20) NOT NULL,
  height DECIMAL
);
```

Run the following command:

```
go run main.go
```

The following file `SomeUserInfo.go` with default package `dto` (data transfer object) will be created:

```go
package dto

import (
	"database/sql"
)

type SomeUserInfo struct {
	Id        int             `db:"id"`
	FirstName sql.NullString  `db:"first_name"`
	LastName  string          `db:"last_name"`
	Height    sql.NullFloat64 `db:"height"`
}
```

Running on remote database server (eg. Mysql@Docker)

```
go run main.go -v -t mysql -h 192.168.99.100 -d testdb -u root -p mysecretpassword
```

PostgreSQL exmple with different default schema but default database `postgres`:

```
go run main.go -v -t pg -h 192.168.99.100 -s test -u postgres -p mysecretpassword
```

Note: since database type `pg` is default, so following command will be equivalent:

```
go run main.go -v -h 192.168.99.100 -s test -u postgres -p mysecretpassword
```

You can also specify the package or prefix and suffix.

```
go run main.go -v -t mysql -h 192.168.99.100 -d testdb -u root -p mysecretpassword -pn models -pre model_ -suf _model
```

With same table given above, following file with Name `ModelSomeUserInfoModel.go` will be created:

```go
package models

import (
	"database/sql"
)

type ModelSomeUserInfoModel struct {
	Id        int             `db:"id"`
	FirstName sql.NullString  `db:"first_name"`
	LastName  string          `db:"last_name"`
	Height    sql.NullFloat64 `db:"height"`
}
```

### Commandline Flags

Print usage with `-?` or `-help`

```
go run main.go -help
  -?    shows help and usage
  -d string
        database name (default "postgres")
  -format string
        camelCase (c) or original (o) (default "c")
  -h string
        host of database (default "127.0.0.1")
  -help
        shows help and usage
  -of string
        output file path (default "./output")
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
  -suf string
        suffix for file- and struct names
  -t string
        type of database to use, currently supported: [pg mysql] (default "pg")
  -u string
        user to connect to the database (default "postgres")
  -v    verbose output
```

## Contributing

If you find any issues or missing a feature, feel free to contribute or make suggestions! 
You can fork the repository and use a feature branch too. Feel free to send me a pull request.

## Licensing

The code in this project is licensed under MIT license.
