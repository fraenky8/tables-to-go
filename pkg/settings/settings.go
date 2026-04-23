package settings

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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
	tags ResolvedTags

	DbType DBType

	User    string
	Pswd    string
	DbName  string
	Schema  string
	Host    string
	Port    string
	SSLMode string
	Socket  string

	OutputFilePath string
	OutputFormat   OutputFormat

	FileNameFormat FileNameFormat
	PackageName    string
	Prefix         string
	Suffix         string
	Null           NullType

	GeneratorVersion string

	Tables StringsFlag
	Tags   StringsFlag

	NoInitialism bool

	TagsNoDb bool

	TagsMastermindStructable       bool
	TagsMastermindStructableOnly   bool
	IsMastermindStructableRecorder bool
	IsGormModel                    bool

	Verbose   bool
	VVerbose  bool
	Force     bool // continue through errors
	GenHeader bool
}

// New constructs Settings with default values.
func New() *Settings {

	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}

	return &Settings{
		tags: ResolvedTags{TagDB},

		Verbose:          false,
		VVerbose:         false,
		Force:            false,
		GenHeader:        false,
		GeneratorVersion: "",

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
		IsGormModel:                    false,
	}
}

// Verify verifies the Settings and checks the given output paths.
func (s *Settings) Verify() (err error) {

	if err = s.verifyOutputPath(); err != nil {
		return err
	}

	if s.OutputFilePath, err = s.prepareOutputPath(); err != nil {
		return err
	}

	if s.Port == "" {
		s.Port = dbDefaultPorts[s.DbType]
	}

	if s.SSLMode == "" {
		s.SSLMode = "disable"
	}

	if s.PackageName == "" {
		return errors.New("name of package can not be empty")
	}

	if err = s.tags.Validate(); err != nil {
		return err
	}

	if s.VVerbose {
		s.Verbose = true
	}

	return err
}

// ResolvedTags returns already resolved tags.
func (s *Settings) ResolvedTags() ResolvedTags {
	return slices.Clone(s.tags)
}

func (s *Settings) verifyOutputPath() error {

	info, err := os.Stat(s.OutputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("output file path %q does not exists", s.OutputFilePath)
	}

	if !info.Mode().IsDir() {
		return fmt.Errorf("output file path %q is not a directory", s.OutputFilePath)
	}

	return err
}

func (s *Settings) prepareOutputPath() (string, error) {
	outputFilePath, err := filepath.Abs(s.OutputFilePath)
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
func (s *Settings) IsNullTypeSQL() bool {
	return s.Null == NullTypeSQL
}

// ShouldInitialism returns whether column names should be converted
// to initialisms or not.
func (s *Settings) ShouldInitialism() bool {
	return !s.NoInitialism
}

// IsOutputFormatCamelCase returns if the type given by command line args is of
// camel-case format.
func (s *Settings) IsOutputFormatCamelCase() bool {
	return s.OutputFormat == OutputFormatCamelCase
}

// IsFileNameFormatSnakeCase returns if the type given by the command line args
// is snake-case format.
func (s *Settings) IsFileNameFormatSnakeCase() bool {
	return s.FileNameFormat == FileNameFormatSnakeCase
}
