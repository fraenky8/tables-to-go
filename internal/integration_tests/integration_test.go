//go:build integration

package integration_tests

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/moby/moby/api/types/container"
	"github.com/ory/dockertest/v4"
	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/internal/cli"
	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/output"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

const (
	testdataDirectoryName = "testdata"
	outputDirectoryName   = "output"
)

var (
	pool dockertest.ClosablePool
)

// nopLogger is used to silence MySQL logs of "packets.go:36: unexpected EOF".
type nopLogger struct{}

func (nopLogger) Print(...any) {}

var (
	isCI bool
)

func init() {
	isCI, _ = strconv.ParseBool(os.Getenv("CI"))

	// Suppress logs of "packets.go:36: unexpected EOF"
	_ = mysql.SetLogger(nopLogger{})
}

type testSettings struct {
	*settings.Settings

	// root filepath where the test can store its testdata and any (expected) output
	filepath string
	// the actual directory for that particular test under the filepath root
	testDirectory string

	dockerImage string
	version     string
	tmpfs       map[string]string
	cmd         []string
	env         []string
}

func newMySQLSettings(version, path, testDirectory string) *testSettings {
	s := settings.New()
	s.DbType = settings.DBTypeMySQL
	s.User = "root"
	s.Pswd = "mysecretpassword"
	s.DbName = "public"
	s.Host = "localhost"
	s.Port = "3306"
	s.OutputFilePath = filepath.Join(path, testDirectory, outputDirectoryName)

	return &testSettings{
		Settings:      s,
		filepath:      path,
		testDirectory: testDirectory,
		dockerImage:   "mysql",
		version:       version,
		tmpfs:         map[string]string{"/var/lib/mysql": ""},
		cmd: []string{
			"--skip-log-bin",
			"--innodb_flush_log_at_trx_commit=2",
			"--sync_binlog=0",
		},
		env: []string{
			"MYSQL_DATABASE=" + s.DbName,
			"MYSQL_ROOT_PASSWORD=" + s.Pswd,
		},
	}
}

func newPostgresSettings(version, path, testDirectory string) *testSettings {
	s := settings.New()
	s.DbType = settings.DBTypePostgresql
	s.User = "postgres"
	s.Pswd = "mysecretpassword"
	s.DbName = "postgres"
	s.Schema = "public"
	s.Host = "localhost"
	s.Port = "5432"
	s.SSLMode = "disable"
	s.OutputFilePath = filepath.Join(path, testDirectory, outputDirectoryName)

	return &testSettings{
		Settings:      s,
		filepath:      path,
		testDirectory: testDirectory,
		dockerImage:   "postgres",
		version:       version,
		tmpfs:         map[string]string{"/var/lib/postgresql": ""},
		cmd: []string{
			"postgres",
			"-c", "fsync=off",
			"-c", "full_page_writes=off",
			"-c", "synchronous_commit=off",
		},
		env: []string{
			"POSTGRES_DB=" + s.DbName,
			"POSTGRES_PASSWORD=" + s.Pswd,
		},
	}
}

func newSQLiteSettings(path, testDirectory, params string) *testSettings {
	s := settings.New()
	s.DbType = settings.DBTypeSQLite
	s.Schema = filepath.Join(path, "database.db") // Only used to keep track of the actual file name without any query params.
	s.DbName = s.Schema + params
	s.OutputFilePath = filepath.Join(path, testDirectory, outputDirectoryName)

	return &testSettings{
		Settings:      s,
		filepath:      path,
		testDirectory: testDirectory,
	}
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	logMsg := "running Tables-to-Go integration tests"
	if isCI {
		logMsg += " on CI"
	}
	log.Println(logMsg)

	log.Println("creating Docker pool...")

	var err error
	pool, err = dockertest.NewPool(ctx, "")
	if err != nil {
		log.Fatalf("error connecting to Docker: %v", err)
	}

	ctx = registerCleanupSignalHandler(ctx)

	code := m.Run()

	err = pool.Close(ctx)
	if err != nil {
		// No need to log.Fatal here, this is just informative.
		log.Printf("warning: error closing Docker pool: %v", err)
	}

	os.Exit(code)
}

func registerCleanupSignalHandler(ctx context.Context) context.Context {
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	done, stop := signal.NotifyContext(ctx, signals...)
	go func() {
		defer stop()
		select {
		case <-done.Done():
			log.Println("got signal:", context.Cause(done))
			log.Println("removing container...")
			_ = pool.Close(ctx)
			// Ignoring error here because it might be called multiple times due
			// to multiple signals arriving. The first of them will remove the
			// container already leading subsequent calls error. But we are not
			// interested in an error saying that the container does not exist (anymore).
		}
	}()
	return done
}

func TestIntegrationDefaultSettings(t *testing.T) {
	const testDirectory = "defaultsettings"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc:     "mysql 5",
			settings: newMySQLSettings("5", "mysql5", testDirectory),
		},
		{
			desc:     "mysql 8",
			settings: newMySQLSettings("8", "mysql8", testDirectory),
		},
		{
			desc:     "postgres 10",
			settings: newPostgresSettings("10", "postgres", testDirectory),
		},
		{
			desc:     "postgres 11",
			settings: newPostgresSettings("11", "postgres", testDirectory),
		},
		{
			desc:     "postgres 12",
			settings: newPostgresSettings("12", "postgres", testDirectory),
		},
		{
			desc:     "postgres 17",
			settings: newPostgresSettings("17", "postgres", testDirectory),
		},
		{
			desc:     "postgres 18",
			settings: newPostgresSettings("18", "postgres", testDirectory),
		},
		{
			desc:     "sqlite 3",
			settings: newSQLiteSettings("sqlite3", testDirectory, ""),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationNullTypePrimitive(t *testing.T) {
	const testDirectory = "nulltypeprimitive"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 5",
			settings: func() *testSettings {
				s := newMySQLSettings("5", "mysql5", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "postgres 10",
			settings: func() *testSettings {
				s := newPostgresSettings("10", "postgres", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "postgres 11",
			settings: func() *testSettings {
				s := newPostgresSettings("11", "postgres", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "postgres 12",
			settings: func() *testSettings {
				s := newPostgresSettings("12", "postgres", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "postgres 17",
			settings: func() *testSettings {
				s := newPostgresSettings("17", "postgres", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "postgres 18",
			settings: func() *testSettings {
				s := newPostgresSettings("18", "postgres", testDirectory)
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
		{
			desc: "sqlite 3",
			settings: func() *testSettings {
				s := newSQLiteSettings("sqlite3", testDirectory, "")
				s.Null = settings.NullTypePrimitive
				return s
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationTablesFlag(t *testing.T) {
	const testDirectory = "tablesflag"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 5",
			settings: func() *testSettings {
				s := newMySQLSettings("5", "mysql5", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "int_table", "varchar_table"}
				return s
			}(),
		},
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "int_table", "varchar_table", "user"}
				return s
			}(),
		},
		{
			desc: "postgres 10",
			settings: func() *testSettings {
				s := newPostgresSettings("10", "postgres", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"date", "float", "int_table", "varchar"}
				return s
			}(),
		},
		{
			desc: "postgres 11",
			settings: func() *testSettings {
				s := newPostgresSettings("11", "postgres", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"date", "float", "int_table", "varchar"}
				return s
			}(),
		},
		{
			desc: "postgres 12",
			settings: func() *testSettings {
				s := newPostgresSettings("12", "postgres", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"date", "float", "int_table", "varchar"}
				return s
			}(),
		},
		{
			desc: "postgres 17",
			settings: func() *testSettings {
				s := newPostgresSettings("17", "postgres", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"date", "float", "int_table", "varchar"}
				return s
			}(),
		},
		{
			desc: "postgres 18",
			settings: func() *testSettings {
				s := newPostgresSettings("18", "postgres", testDirectory)
				// Note: int_table non-existing
				s.Tables = settings.StringsFlag{"date", "float", "int_table", "varchar"}
				return s
			}(),
		},
		{
			desc: "sqlite 3",
			settings: func() *testSettings {
				s := newSQLiteSettings("sqlite3", testDirectory, "")
				s.Tables = settings.StringsFlag{"numeric_table", "text_table", "strict_types"}
				return s
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationOutputFormatOriginal(t *testing.T) {
	const testDirectory = "outputformatoriginal"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.OutputFormat = settings.OutputFormatOriginal

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationFileNameFormatSnakeCase(t *testing.T) {
	const testDirectory = "filenameformatsnakecase"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.FileNameFormat = settings.FileNameFormatSnakeCase

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationPackageName(t *testing.T) {
	const testDirectory = "packagename"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.PackageName = "models"

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationPrefix(t *testing.T) {
	const testDirectory = "prefix"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.Prefix = "Prefix_"

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationSuffix(t *testing.T) {
	const testDirectory = "suffix"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.Suffix = "_Suffix"

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationPrefixSuffix(t *testing.T) {
	const testDirectory = "prefixsuffix"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.Prefix = "Prefix_"
				s.Suffix = "_Suffix"

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationNoInitialism(t *testing.T) {
	const testDirectory = "noinitialism"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.NoInitialism = true
				s.Tables = settings.StringsFlag{"user"}
				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationTagsNoDB(t *testing.T) {
	const testDirectory = "tagsnodbflag"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.TagsNoDb = true

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationTagsMastermindStructable(t *testing.T) {
	const testDirectory = "tagsmastermindstructable"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.TagsMastermindStructable = true

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationTagsMastermindStructableOnly(t *testing.T) {
	const testDirectory = "tagsmastermindstructableonly"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.TagsMastermindStructableOnly = true

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func TestIntegrationIsMastermindStructableRecorder(t *testing.T) {
	const testDirectory = "ismastermindstructablerecorder"

	tests := []struct {
		desc     string
		settings *testSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *testSettings {
				s := newMySQLSettings("8", "mysql8", testDirectory)
				s.TagsMastermindStructableOnly = true
				s.IsMastermindStructableRecorder = true

				// Only set to reduce the amount of files
				s.Tables = settings.StringsFlag{"datetime_table", "float_table", "integer_table", "varchar_table", "user"}

				return s
			}(),
		},
		// Skipping all other DB types since it's not related to the type itself,
		// and testing for one type covers all others.
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			db := setupDatabase(t, test.settings)
			defer func() {
				if !t.Failed() {
					_ = os.RemoveAll(test.settings.Settings.OutputFilePath)
				}
			}()

			loadTestData(t, db.SQLDriver(), test.settings)

			err := os.MkdirAll(test.settings.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Fatalf("could not create output file path: %v", err)
			}

			version, err := db.Version(t.Context())
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(test.settings.Settings.OutputFilePath)

			app := cli.New(test.settings.Settings, db, writer)
			err = app.Run(t.Context())
			assert.NoError(t, err)

			checkFiles(t, test.settings)
		})
	}
}

func checkFiles(t *testing.T, s *testSettings) {
	expectedPattern := filepath.Join(s.filepath, s.testDirectory, "*.go")
	expectedFiles, err := filepath.Glob(expectedPattern)
	assert.NoError(t, err)

	outputPattern := filepath.Join(s.Settings.OutputFilePath, "*")
	outputFiles, err := filepath.Glob(outputPattern)
	assert.NoError(t, err)

	if len(expectedFiles) != len(outputFiles) {
		t.Fatalf("number of expected and output files differ: %d vs. %d",
			len(expectedFiles), len(outputFiles))
	}

	expectedByName := make(map[string]string, len(expectedFiles))
	for _, expectedFile := range expectedFiles {
		fileName := filepath.Base(expectedFile)
		if _, ok := expectedByName[fileName]; ok {
			t.Fatalf("duplicate expected file %q", fileName)
		}
		expectedByName[fileName] = expectedFile
	}

	for _, outputFile := range outputFiles {
		fileName := filepath.Base(outputFile)

		expectedFile, ok := expectedByName[fileName]
		if !ok {
			t.Fatalf("unexpected output file %q", outputFile)
		}
		delete(expectedByName, fileName)

		expectedInfo, err := os.Stat(expectedFile)
		assert.NoError(t, err)
		outputInfo, err := os.Stat(outputFile)
		assert.NoError(t, err)

		assert.Equal(t, expectedInfo.Size(), outputInfo.Size(), "file %q differs in size", fileName)

		expectedContent, err := os.ReadFile(expectedFile)
		assert.NoError(t, err)
		outputContent, err := os.ReadFile(outputFile)
		assert.NoError(t, err)

		assert.Equal(t, string(expectedContent), string(outputContent), "file %q differs in content", fileName)
	}

	if len(expectedByName) > 0 {
		missingFiles := make([]string, 0, len(expectedByName))
		for fileName := range expectedByName {
			missingFiles = append(missingFiles, fileName)
		}
		sort.Strings(missingFiles)
		t.Fatalf("missing output files for expected files: %v", missingFiles)
	}
}

func setupDatabase(t *testing.T, s *testSettings) database.Database {
	if s.Settings.DbType == settings.DBTypeSQLite {
		return setupSQLite(t, s)
	}

	t.Logf("spinning up database %s:%s", s.dockerImage, s.version)

	containerName := fmt.Sprintf("tables_to_go_%s_%s_integration", s.dockerImage, s.version)

	// Using pool.Run instead of pool.RunT here to be able to reuse containers.
	// Otherwise, t.Cleanup would have been run already and removed the container.
	resource, err := pool.Run(t.Context(), s.dockerImage,
		dockertest.WithCmd(s.cmd),
		dockertest.WithTag(s.version),
		dockertest.WithName(containerName),
		dockertest.WithEnv(s.env),
		dockertest.WithHostConfig(func(config *container.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = container.RestartPolicy{
				Name: container.RestartPolicyDisabled,
			}
			config.Tmpfs = s.tmpfs
		}),
	)
	if err != nil {
		t.Fatalf("could not start resource: %v", err)
	}

	var db database.Database

	if err := pool.Retry(t.Context(), 0, func() error {
		port := resource.GetPort(s.Port + "/tcp")
		if port != "" {
			s.Settings.Port = port
		}
		db = database.New(s.Settings)
		err := db.Connect(t.Context())
		if err != nil {
			if s.Settings.Verbose {
				t.Log(err.Error())
			}
			return err
		}
		return nil
	}); err != nil {
		t.Fatalf("could not connect to database: %v", err)
	}
	t.Cleanup(func() {
		_ = db.Close()
	})

	resetDatabase(t, db, s)

	return db
}

func setupSQLite(t *testing.T, s *testSettings) database.Database {
	db := database.New(s.Settings)
	err := db.Connect(t.Context())
	if err != nil {
		t.Fatalf("could not create sqlite: %v", err)
	}
	t.Cleanup(func() {
		_ = db.Close()
		err := os.Remove(s.Schema)
		if err != nil {
			t.Log(err)
		}
	})
	return db
}

func loadTestData(t *testing.T, db *sqlx.DB, s *testSettings) {
	testDataPattern := filepath.Join(s.filepath, testdataDirectoryName, "*.sql")
	files, err := filepath.Glob(testDataPattern)
	if err != nil {
		t.Fatalf("could not find sql testdata: %v", err)
	}

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("could not read %q: %v", f, err)
		}

		queries := bytes.Split(data, []byte(";"))

		for _, query := range queries {
			query = bytes.TrimSpace(query)
			q := string(query)
			if q == "" {
				continue
			}

			_, err = db.Exec(q)
			if err != nil {
				t.Fatalf("could not create testdata %q: %v", f, err)
			}
		}
	}
}

func resetDatabase(t *testing.T, db database.Database, s *testSettings) {
	start := time.Now()

	dbx := db.SQLDriver()

	// For the sake of integration testing and not to expose a DROP method
	// at the database.Database interface, we type switch here.
	switch tdb := db.(type) {
	case *database.MySQL:
		query := `DROP DATABASE ` + s.DbName
		if _, err := dbx.ExecContext(t.Context(), query); err != nil {
			t.Fatalf("could not drop database %q: %v", s.DbName, err)
		}
		query = `CREATE DATABASE ` + s.DbName
		if _, err := dbx.ExecContext(t.Context(), query); err != nil {
			t.Fatalf("could not create database %q: %v", s.DbName, err)
		}
		query = `USE ` + s.DbName
		if _, err := dbx.ExecContext(t.Context(), query); err != nil {
			t.Fatalf("could not use database %q: %v", s.DbName, err)
		}
	case *database.Postgresql:
		query := `DROP SCHEMA ` + s.Schema + ` CASCADE`
		if _, err := dbx.ExecContext(t.Context(), query); err != nil {
			t.Fatalf("could not drop schema %q: %v", s.Schema, err)
		}
		query = `CREATE SCHEMA ` + s.Schema
		if _, err := dbx.ExecContext(t.Context(), query); err != nil {
			t.Fatalf("could not create schema %q: %v", s.Schema, err)
		}
	case *database.SQLite:
		t.Log("not implemented since never reached")
	default:
		// MUST never happen
		t.Fatalf("unknown database %v", tdb)
	}

	t.Logf("resetting database (%s)", time.Since(start))
}
