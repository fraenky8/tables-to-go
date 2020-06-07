// +build sqlite3

// Package database/sqlite_driver.go contains only the driver for the sqlite3
// database. It will get only included in the build if the tag `sqlite3` is
// specified.
//
// Default build of tables-to-go does NOT include sqlite3 support.
//
// Support for sqlite3 can be enabled by specifying the tag while
// building tables-to-go:
//
//		go {install/build} -mod=vendor -tags sqlite3 .
//
// Alternative the Makefile can be used which is an alias for the go command
// above:
//
//		make sqlite3
//
package database

import (
	// sqlite3 database driver
	_ "github.com/mattn/go-sqlite3"
)
