package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// These are the output format command line parameter
const (
	OutputFormatCamelCase = "c"
	OutputFormatOriginal  = "o"
)

// NullType represents a null type.
type NullType string

// These null types are supported. The types native and primitve map to the same
// underlying builtin golang type.
const (
	NullTypeSQL      NullType = "sql"
	NullTypeNative   NullType = "native"
	NullTypePrimitve NullType = "primitive"
)

var (
	// supportedDbTypes represents the supported databases
	supportedDbTypes = map[string]bool{
		"pg":    true,
		"mysql": true,
	}

	// supportedOutputFormats represents the supported output formats
	supportedOutputFormats = map[string]bool{
		OutputFormatCamelCase: true,
		OutputFormatOriginal:  true,
	}

	// dbDefaultPorts maps the database type to the default ports
	dbDefaultPorts = map[string]string{
		"pg":    "5432",
		"mysql": "3306",
	}

	// supportedNullTypes represents the supported types of NULL types
	supportedNullTypes = map[NullType]bool{
		NullTypeSQL:      true,
		NullTypeNative:   true,
		NullTypePrimitve: true,
	}
)

// Settings stores the supported settings / command line arguments
type Settings struct {
	Verbose  bool
	VVerbose bool

	DbType         string
	User           string
	Pswd           string
	DbName         string
	Schema         string
	Host           string
	Port           string
	OutputFilePath string
	OutputFormat   string
	PackageName    string
	Prefix         string
	Suffix         string
	Null           string

	TagsNoDb bool

	TagsMastermindStructable       bool
	TagsMastermindStructableOnly   bool
	IsMastermindStructableRecorder bool

	// TODO not implemented yet
	TagsGorm bool

	// experimental
	TagsSQL     bool
	TagsSQLOnly bool
}

// NewSettings constructs settings with default values
func NewSettings() *Settings {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = "."
	}

	return &Settings{
		Verbose:  false,
		VVerbose: false,

		DbType:         "pg",
		User:           "postgres",
		Pswd:           "",
		DbName:         "postgres",
		Schema:         "public",
		Host:           "127.0.0.1",
		Port:           "", // left blank -> is automatically determined if not set
		OutputFilePath: dir,
		OutputFormat:   OutputFormatCamelCase,
		PackageName:    "dto",
		Prefix:         "",
		Suffix:         "",
		Null:           string(NullTypeSQL),

		TagsNoDb: false,

		TagsMastermindStructable:       false,
		TagsMastermindStructableOnly:   false,
		IsMastermindStructableRecorder: false,

		TagsGorm: false,

		TagsSQL:     false,
		TagsSQLOnly: false,
	}
}

// Verify verifies the settings and checks the given output paths
func (settings *Settings) Verify() (err error) {

	if !supportedDbTypes[settings.DbType] {
		return fmt.Errorf("type of database %q not supported! supported: %v", settings.DbType, settings.SupportedDbTypes())
	}

	if !supportedOutputFormats[settings.OutputFormat] {
		return fmt.Errorf("output format %q not supported", settings.OutputFormat)
	}

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

	if !supportedNullTypes[NullType(settings.Null)] {
		return fmt.Errorf("null type %q not supported! supported: %v", settings.Null, settings.SupportedNullTypes())
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

// SupportedDbTypes returns a slice of strings as names of the supported database types
func (settings *Settings) SupportedDbTypes() string {
	names := make([]string, 0, len(supportedDbTypes))
	for name := range supportedDbTypes {
		names = append(names, name)
	}
	return fmt.Sprintf("%v", names)
}

// SupportedNullTypes returns a slice of strings as names of the supported null types
func (settings *Settings) SupportedNullTypes() string {
	names := make([]string, 0, len(supportedNullTypes))
	for name := range supportedNullTypes {
		names = append(names, string(name))
	}
	return fmt.Sprintf("%v", names)
}

// IsNullTypeSQL returns if the type given by command line args is of null type SQL
func (settings *Settings) IsNullTypeSQL() bool {
	return settings.Null == string(NullTypeSQL)
}
