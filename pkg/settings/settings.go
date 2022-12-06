package settings

import (
	"fmt"
	"os"
	"path/filepath"
)

// DBType represents a type of a database.
type DBType string

// These database types are supported.
const (
	DBTypePostgresql DBType = "pg"
	DBTypeMySQL      DBType = "mysql"
	DBTypeSQLite     DBType = "sqlite3"
	DBTypeSQLServer  DBType = "sqlserver"
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

var (
	// SupportedDbTypes represents the supported databases
	SupportedDbTypes = map[DBType]bool{
		DBTypePostgresql: true,
		DBTypeMySQL:      true,
		DBTypeSQLite:     true,
		DBTypeSQLServer:  true,
	}

	// supportedOutputFormats represents the supported output formats
	supportedOutputFormats = map[OutputFormat]bool{
		OutputFormatCamelCase: true,
		OutputFormatOriginal:  true,
	}

	// dbDefaultPorts maps the database type to the default ports
	dbDefaultPorts = map[DBType]string{
		DBTypePostgresql: "5432",
		DBTypeMySQL:      "3306",
		DBTypeSQLite:     "",
		DBTypeSQLServer:  "1433",
	}

	// supportedNullTypes represents the supported types of NULL types
	supportedNullTypes = map[NullType]bool{
		NullTypeSQL:       true,
		NullTypeNative:    true,
		NullTypePrimitive: true,
	}

	// supportedFileNameFormats represents the supported filename formats
	supportedFileNameFormats = map[FileNameFormat]bool{
		FileNameFormatCamelCase: true,
		FileNameFormatSnakeCase: true,
	}
)

// Settings stores the supported settings / command line arguments.
type Settings struct {
	Verbose  bool
	VVerbose bool
	Force    bool // continue through errors

	DbType DBType

	User   string
	Pswd   string
	DbName string
	Schema string
	Host   string
	Port   string
	Socket string

	OutputFilePath string
	OutputFormat   OutputFormat

	FileNameFormat FileNameFormat
	PackageName    string
	Prefix         string
	Suffix         string
	Null           NullType

	NoInitialism bool

	TagsNoDb bool

	TagsMastermindStructable       bool
	TagsMastermindStructableOnly   bool
	IsMastermindStructableRecorder bool

	// TODO not implemented yet
	TagsGorm bool
}

// New constructs Settings with default values.
func New() *Settings {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = "."
	}

	return &Settings{
		Verbose:  false,
		VVerbose: false,
		Force:    false,

		DbType:         DBTypePostgresql,
		User:           "",
		Pswd:           "",
		DbName:         "postgres",
		Schema:         "public",
		Host:           "127.0.0.1",
		Port:           "", // left blank, automatically determined if not set
		Socket:         "",
		OutputFilePath: dir,
		OutputFormat:   OutputFormatCamelCase,
		FileNameFormat: FileNameFormatCamelCase,
		PackageName:    "dto",
		Prefix:         "",
		Suffix:         "",
		Null:           NullTypeSQL,

		NoInitialism: false,

		TagsNoDb: false,

		TagsMastermindStructable:       false,
		TagsMastermindStructableOnly:   false,
		IsMastermindStructableRecorder: false,

		TagsGorm: false,
	}
}

// Verify verifies the Settings and checks the given output paths.
func (settings *Settings) Verify() (err error) {

	if err = settings.verifyOutputPath(); err != nil {
		return err
	}

	if settings.OutputFilePath, err = settings.prepareOutputPath(); err != nil {
		return err
	}

	if settings.Port == "" {
		settings.Port = dbDefaultPorts[settings.DbType]
	}

	if settings.PackageName == "" {
		return fmt.Errorf("name of package can not be empty")
	}

	if settings.VVerbose {
		settings.Verbose = true
	}

	return err
}

func (settings *Settings) verifyOutputPath() (err error) {

	info, err := os.Stat(settings.OutputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("output file path %q does not exists", settings.OutputFilePath)
	}

	if !info.Mode().IsDir() {
		return fmt.Errorf("output file path %q is not a directory", settings.OutputFilePath)
	}

	return err
}

func (settings *Settings) prepareOutputPath() (outputFilePath string, err error) {
	outputFilePath, err = filepath.Abs(settings.OutputFilePath)
	outputFilePath += string(filepath.Separator)
	return outputFilePath, err
}

// SprintfSupportedDbTypes returns a slice of strings as names of the supported
// database types
func SprintfSupportedDbTypes() string {
	names := make([]string, 0, len(SupportedDbTypes))
	for name := range SupportedDbTypes {
		names = append(names, string(name))
	}
	return fmt.Sprintf("%v", names)
}

// SprintfSupportedNullTypes returns a slice of strings as names of the
// supported null types
func SprintfSupportedNullTypes() string {
	names := make([]string, 0, len(supportedNullTypes))
	for name := range supportedNullTypes {
		names = append(names, string(name))
	}
	return fmt.Sprintf("%v", names)
}

// IsNullTypeSQL returns true if the type given by the command line args is of
// null type SQL
func (settings *Settings) IsNullTypeSQL() bool {
	return settings.Null == NullTypeSQL
}

// ShouldInitialism returns whether column names should be converted
// to initialisms or not.
func (settings *Settings) ShouldInitialism() bool {
	return !settings.NoInitialism
}

// IsOutputFormatCamelCase returns if the type given by command line args is of
// camel-case format.
func (settings *Settings) IsOutputFormatCamelCase() bool {
	return settings.OutputFormat == OutputFormatCamelCase
}

// IsFileNameFormatSnakeCase returns if the type given by the command line args
// is snake-case format.
func (settings *Settings) IsFileNameFormatSnakeCase() bool {
	return settings.FileNameFormat == FileNameFormatSnakeCase
}
