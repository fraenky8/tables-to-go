package tablestogo

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// SupportedDbTypes represents the supported databases
	SupportedDbTypes = []string{"pg", "mysql"}
	// SupportedOutputFormats represents the supported output formats
	SupportedOutputFormats = []string{"c", "o"}

	// dbDefaultPorts maps the database type to the default ports
	dbDefaultPorts = map[string]string{
		"pg":    "5432",
		"mysql": "3306",
	}
)

// Settings stores the supported settings / command line arguments
type Settings struct {
	Verbose        bool
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

	TagsNoDb bool

	TagsMastermindStructable       bool
	TagsMastermindStructableOnly   bool
	IsMastermindStructableRecorder bool

	// TODO not implemented yet
	TagsGorm bool

	// experimental
	TagsSQL     bool
	TagsSQLOnly bool

	effectiveTags int
}

// NewSettings constructs settings with default values
func NewSettings() *Settings {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		dir = "."
	}

	return &Settings{
		Verbose:        false,
		DbType:         "pg",
		User:           "postgres",
		Pswd:           "",
		DbName:         "postgres",
		Schema:         "public",
		Host:           "127.0.0.1",
		Port:           "", // left blank -> is automatically determined if not set
		OutputFilePath: dir,
		OutputFormat:   "c",
		PackageName:    "dto",
		Prefix:         "",
		Suffix:         "",

		TagsNoDb: false,

		TagsMastermindStructable:       false,
		TagsMastermindStructableOnly:   false,
		IsMastermindStructableRecorder: false,

		TagsGorm: false,

		TagsSQL:     false,
		TagsSQLOnly: false,

		effectiveTags: 1,
	}
}

// VerifySettings verifies the settings and checks the given output paths
func VerifySettings(settings *Settings) (err error) {

	if !IsStringInSlice(settings.DbType, SupportedDbTypes) {
		return fmt.Errorf("type of database %q not supported! %v", settings.DbType, SupportedDbTypes)
	}

	if !IsStringInSlice(settings.OutputFormat, SupportedOutputFormats) {
		return fmt.Errorf("output format %q not supported! %v", settings.OutputFormat, SupportedOutputFormats)
	}

	if err = verifyOutputPath(settings.OutputFilePath); err != nil {
		return err
	}

	if settings.OutputFilePath, err = prepareOutputPath(settings.OutputFilePath); err != nil {
		return err
	}

	if settings.Port == "" {
		settings.Port = dbDefaultPorts[settings.DbType]
	}

	if settings.PackageName == "" {
		return fmt.Errorf("name of package can not be empty")
	}

	return err
}

func verifyOutputPath(outputFilePath string) (err error) {

	info, err := os.Stat(outputFilePath)

	if os.IsNotExist(err) {
		return fmt.Errorf("output file path %q does not exists", outputFilePath)
	}

	if !info.Mode().IsDir() {
		return fmt.Errorf("output file path %q is not a directory", outputFilePath)
	}

	return err
}

func prepareOutputPath(ofp string) (outputFilePath string, err error) {
	outputFilePath, err = filepath.Abs(ofp)
	outputFilePath += string(filepath.Separator)
	return outputFilePath, err
}
