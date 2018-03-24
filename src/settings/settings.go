package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fraenky8/tables-to-go/src/database"
	"github.com/fraenky8/tables-to-go/src/tagger"
)

var (
	// supportedDbTypes represents the supported databases
	supportedDbTypes = map[string]bool{
		"pg":    true,
		"mysql": true,
	}

	// supportedOutputFormats represents the supported output formats
	supportedOutputFormats = map[string]bool{
		"c": true,
		"o": true,
	}

	// dbDefaultPorts maps the database type to the default ports
	dbDefaultPorts = map[string]string{
		"pg":    "5432",
		"mysql": "3306",
	}

	// map of Tagger used
	// key is a ascending sequence of i*2 to determine easily which tags to generate later
	taggers = map[int]tagger.Tagger{
		1: new(dbTag),
		2: new(stblTag),
		4: new(SQLTag),
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

	if !supportedDbTypes[settings.DbType] {
		return fmt.Errorf("type of database %q not supported! %v", settings.DbType, PrettyPrintSupportedDbTypes())
	}

	if !supportedOutputFormats[settings.OutputFormat] {
		return fmt.Errorf("output format %q not supported", settings.OutputFormat)
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

func (settings *Settings) CreateEffectiveTags() {
	if settings.TagsNoDb {
		settings.effectiveTags = 0
	}
	if settings.TagsMastermindStructable {
		settings.effectiveTags |= 2
	}
	if settings.TagsMastermindStructableOnly {
		settings.effectiveTags = 0
		settings.effectiveTags |= 2
	}
	if settings.TagsSQL {
		settings.effectiveTags |= 4
	}
	if settings.TagsSQLOnly {
		settings.effectiveTags = 0
		settings.effectiveTags |= 4
	}
	// last tag-"ONLY" wins if multiple specified
}

func (settings *Settings) GenerateTags(db database.Database, column database.Column) (tags string) {
	for t := 1; t <= settings.effectiveTags; t *= 2 {
		shouldTag := settings.effectiveTags&t > 0
		if shouldTag {
			tags += taggers[t].GenerateTag(db, column) + " "
		}
	}
	if len(tags) > 0 {
		tags = " `" + strings.TrimSpace(tags) + "`"
	}
	return tags
}

func (settings *Settings) PrettyPrintSupportedDbTypes() string {
	names := make([]string, len(supportedDbTypes))
	i := 0
	for name := range supportedDbTypes {
		names[i] = name
		i++
	}
	return fmt.Sprintf("%v", names)
}
