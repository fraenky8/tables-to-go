package main

import (
	"database/sql"
	"flag"
	"fmt"

	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db  *sqlx.DB
	err error

	// command line args
	help   bool
	dbType string
	user   string
	pswd   string
	dbName string
	schema string
	host   string
	port   string
	//output         string
	//outputFilePath string
	//outputFormat   string
	//prefix         string
	//suffix         string
	//packageName    string

	//output         = flag.String("o", "file", "output, currently supported: file (default), stdout")
	//outputFilePath = flag.String("of", "./output", "output file path, default ./output")
	//outputFormat   = flag.String("format", "c", "camelCase (c) or under_scored (u), default c")
	//prefix         = flag.String("pre", "", "prefix for file- and struct name")
	//suffix         = flag.String("suf", "", "suffix for file- and struct name")
	//packageName    = flag.String("pn", "dto", "package name, default dto")
)

type Table struct {
	TableName string `db:"table_name"`
	Columns   []Column
}

type Column struct {
	OrdinalPosition        int            `db:"ordinal_position"`
	ColumnName             string         `db:"column_name"`
	DataType               string         `db:"data_type"`
	ColumnDefault          sql.NullString `db:"column_default"`
	IsNullable             string         `db:"is_nullable"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
}

type Database interface {
	GetTables() (tables []Table, err error)
	GetColumnsOfTable(tableName string) (columns []Column, err error)
}

type PostgreDatabase string

func (pg *PostgreDatabase) GetTables() (tables []Table, err error) {

	err = db.Select(&tables, `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_type = 'BASE TABLE'
		AND table_schema = ?
		ORDER BY table_name
	`, schema)

	return tables, err
}

func (pg *PostgreDatabase) GetColumnsOfTable(tableName string) (columns []Column, err error) {

	err = db.Select(&columns, `
		SELECT
		  ordinal_position,
		  column_name,
		  data_type,
		  column_default,
		  is_nullable,
		  character_maximum_length,
		  numeric_precision
		FROM information_schema.columns
		WHERE table_name = ?
		AND table_schema = ?
		ORDER BY ordinal_position
	`, tableName, schema)

	return columns, err
}

type MySQLDatabase string

func main() {

	// mysecretpassword

	prepareCmdArgs()

	err = handleCmdArgs()
	if err != nil {
		flag.Usage()
		return
	}

	err = connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// TODO lets go!
	fmt.Println("let's go!")

}

func prepareCmdArgs() {
	flag.BoolVar(&help, "?", false, "shows help and usage")
	flag.StringVar(&dbType, "t", "pg", "type of database to use, currently supported: pg, mysql") // TODO make map of supported db types -> can also be used in handleCMdArgs
	flag.StringVar(&user, "u", "postgres", "user to connect to the database, default for Postgres 'postgres'")
	flag.StringVar(&pswd, "p", "", "password of user")
	flag.StringVar(&dbName, "d", "postgres", "database name, default for Postgres 'postgres'")
	flag.StringVar(&schema, "s", "public", "schema name, default for Postgres 'public'")
	flag.StringVar(&host, "h", "127.0.0.1", "host of database, if not specified, it will be 127.0.0.1/localhost")
	flag.StringVar(&port, "port", "5432", "port of database host, if not specified, it will be the default ports for the supported databases")

	flag.Parse()
}

func handleCmdArgs() (err error) {

	if help {
		return errors.New("help called")
	}



	return err
}

func connect() (err error) {
	db, err = sqlx.Connect("postgres", fmt.Sprintf("host=%v user=%v dbname=%v password=%v sslmode=disable", host, user, dbName, pswd))
	if err != nil {
		usingPswd := "no"
		if pswd != "" {
			usingPswd = "yes"
		}
		return errors.New(fmt.Sprintf("Connection to Database (type=%q, user=%q, database=%q, host='%v:%v' (using password: %v) failed!", dbType, user, dbName, host, port, usingPswd))
	}
	return db.Ping()
}
