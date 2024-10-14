package settings

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// SupportedDbTypes represents the supported databases
	SupportedDbTypes = map[DBType]bool{
		DBTypePostgresql: true,
		DBTypeMySQL:      true,
		DBTypeSQLite:     true,
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

	User    string
	Pswd    string
	DbName  string
	Schema  string
	Host    string
	Port    string
	SSLMode string
	Socket  string
	Tables  StringsFlag

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
		SSLMode:        "", // left blank, will set the default for Postgres to 'disable'
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

	if settings.SSLMode == "" {
		settings.SSLMode = "disable"
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
