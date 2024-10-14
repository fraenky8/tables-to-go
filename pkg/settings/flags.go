package settings

import (
	"fmt"
	"strings"
)

// DBType represents a type of a database.
type DBType string

// These database types are supported.
const (
	DBTypePostgresql DBType = "pg"
	DBTypeMySQL      DBType = "mysql"
	DBTypeSQLite     DBType = "sqlite3"
)

// Set sets the datatype for the custom type for the flag package.
func (db *DBType) Set(s string) error {
	*db = DBType(s)
	if *db == "" {
		*db = DBTypePostgresql
	}
	if !SupportedDbTypes[*db] {
		return fmt.Errorf("database type %q not supported, must be one of: %v",
			*db, SprintfSupportedDbTypes())
	}
	return nil
}

// String is the implementation of the Stringer interface needed for flag.Value interface.
func (db DBType) String() string {
	return string(db)
}

// These null types are supported. The types native and primitive map to the same
// underlying builtin golang type.
const (
	NullTypeSQL       NullType = "sql"
	NullTypeNative    NullType = "native"
	NullTypePrimitive NullType = "primitive"
)

// NullType represents a null type.
type NullType string

// Set sets the datatype for the custom type for the flag package.
func (t *NullType) Set(s string) error {
	*t = NullType(s)
	if *t == "" {
		*t = NullTypeSQL
	}
	if !supportedNullTypes[*t] {
		return fmt.Errorf("null type %q not supported, must be one of: %v",
			*t, SprintfSupportedNullTypes())
	}
	return nil
}

// String is the implementation of the Stringer interface needed for
// flag.Value interface.
func (t NullType) String() string {
	return string(t)
}

// OutputFormat represents an output format option.
type OutputFormat string

// These are the OutputFormat command line parameter.
const (
	OutputFormatCamelCase OutputFormat = "c"
	OutputFormatOriginal  OutputFormat = "o"
)

// Set sets the datatype for the custom type for the flag package.
func (of *OutputFormat) Set(s string) error {
	*of = OutputFormat(s)
	if *of == "" {
		*of = OutputFormatCamelCase
	}
	if !supportedOutputFormats[*of] {
		return fmt.Errorf("output format %q not supported", *of)
	}
	return nil
}

// String is the implementation of the Stringer interface needed for
// flag.Value interface.
func (of OutputFormat) String() string {
	return string(of)
}

// FileNameFormat represents a output filename format.
type FileNameFormat string

// These are the FileNameFormat command line parameter.
const (
	FileNameFormatCamelCase FileNameFormat = "c"
	FileNameFormatSnakeCase FileNameFormat = "s"
)

// Set sets the datatype for the custom type for the flag package.
func (of *FileNameFormat) Set(s string) error {
	*of = FileNameFormat(s)
	if *of == "" {
		*of = FileNameFormatCamelCase
	}
	if !supportedFileNameFormats[*of] {
		return fmt.Errorf("filename format %q not supported", *of)
	}
	return nil
}

func (of FileNameFormat) String() string {
	return string(of)
}

// StringsFlag can be used to specify multiple occurrences of a flag and hence
// multiple values without having to split anything by a delimiter.
type StringsFlag []string

// String is the implementation of the Stringer interface needed for
// flag.Value interface.
func (s *StringsFlag) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set sets the value for the StringsFlag.
func (s *StringsFlag) Set(val string) error {
	vals := strings.Split(val, ",")
	*s = append(*s, vals...)
	return nil
}
