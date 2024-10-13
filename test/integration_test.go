//go:build integration

package test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/internal/cli"
	"github.com/fraenky8/tables-to-go/v2/pkg/database"
	"github.com/fraenky8/tables-to-go/v2/pkg/output"
	"github.com/fraenky8/tables-to-go/v2/pkg/settings"
)

const (
	testdataDirectoryName = "testdata"
	expectedDirectoryName = "expected"
	outputDirectoryName   = "output"
)

type dbSettings struct {
	*settings.Settings

	// root filepath where the testdata and expected directories live
	dataFilepath string

	dockerImage string
	version     string
	env         []string
}

func (s *dbSettings) setSettings(ss *settings.Settings) {
	s.Settings = ss
	s.Settings.OutputFilePath = filepath.Join(s.dataFilepath, outputDirectoryName)
}

// nopLogger is used to silence MySQL logs of "packets.go:36: unexpected EOF".
type nopLogger struct{}

func (nopLogger) Print(...any) {}

func TestIntegration(t *testing.T) {
	log.Println("running Tables-to-Go integration tests")

	tests := []struct {
		desc     string
		settings *dbSettings
	}{
		{
			desc: "mysql 8",
			settings: func() *dbSettings {
				s := settings.New()
				s.DbType = settings.DBTypeMySQL
				s.User = "root"
				s.Pswd = "mysecretpassword"
				s.DbName = "public"
				s.Host = "localhost"
				s.Port = "3306"
				//s.Verbose = true
				//s.VVerbose = true

				dbs := &dbSettings{
					Settings: s,

					dataFilepath: "mysql8",

					dockerImage: "mysql",
					version:     "8",
					env: []string{
						"MYSQL_DATABASE=" + s.DbName,
						"MYSQL_ROOT_PASSWORD=" + s.Pswd,
					},
				}

				dbs.setSettings(s)

				// Suppress logs of "packets.go:36: unexpected EOF"
				_ = mysql.SetLogger(nopLogger{})

				return dbs
			}(),
		},
		{
			desc: "postgres 10",
			settings: func() *dbSettings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				s.User = "postgres"
				s.Pswd = "mysecretpassword"
				s.DbName = "postgres"
				s.Schema = "public"
				s.Host = "localhost"
				s.Port = "5432"
				s.SSLMode = "disable"
				//s.Verbose = true
				//s.VVerbose = true

				dbs := &dbSettings{
					Settings: s,

					dataFilepath: "postgres",

					dockerImage: "postgres",
					version:     "10",
					env: []string{
						"POSTGRES_DB=" + s.DbName,
						"POSTGRES_PASSWORD=" + s.Pswd,
					},
				}

				dbs.setSettings(s)

				return dbs
			}(),
		},
		{
			desc: "postgres 11",
			settings: func() *dbSettings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				s.User = "postgres"
				s.Pswd = "mysecretpassword"
				s.DbName = "postgres"
				s.Schema = "public"
				s.Host = "localhost"
				s.Port = "5432"
				s.SSLMode = "disable"
				//s.Verbose = true
				//s.VVerbose = true

				dbs := &dbSettings{
					Settings: s,

					dataFilepath: "postgres",

					dockerImage: "postgres",
					version:     "11",
					env: []string{
						"POSTGRES_DB=" + s.DbName,
						"POSTGRES_PASSWORD=" + s.Pswd,
					},
				}

				dbs.setSettings(s)

				return dbs
			}(),
		},
		{
			desc: "postgres 12",
			settings: func() *dbSettings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				s.User = "postgres"
				s.Pswd = "mysecretpassword"
				s.DbName = "postgres"
				s.Schema = "public"
				s.Host = "localhost"
				s.Port = "5432"
				s.SSLMode = "disable"
				//s.Verbose = true
				//s.VVerbose = true

				dbs := &dbSettings{
					Settings: s,

					dataFilepath: "postgres",

					dockerImage: "postgres",
					version:     "12",
					env: []string{
						"POSTGRES_DB=" + s.DbName,
						"POSTGRES_PASSWORD=" + s.Pswd,
					},
				}

				dbs.setSettings(s)

				return dbs
			}(),
		},
		{
			desc: "postgres 17",
			settings: func() *dbSettings {
				s := settings.New()
				s.DbType = settings.DBTypePostgresql
				s.User = "postgres"
				s.Pswd = "mysecretpassword"
				s.DbName = "postgres"
				s.Schema = "public"
				s.Host = "localhost"
				s.Port = "5432"
				s.SSLMode = "disable"
				//s.Verbose = true
				//s.VVerbose = true

				dbs := &dbSettings{
					Settings: s,

					dataFilepath: "postgres",

					dockerImage: "postgres",
					version:     "17",
					env: []string{
						"POSTGRES_DB=" + s.DbName,
						"POSTGRES_PASSWORD=" + s.Pswd,
					},
				}

				dbs.setSettings(s)

				return dbs
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			s := test.settings

			db, purgeFn, err := setupDatabase(s)
			if err != nil {
				if purgeFn != nil {
					t.Log(err)
					t.Fatal(purgeFn())
				}
				t.Fatal(err)
			}
			defer func() {
				if purgeFn != nil {
					if err := purgeFn(); err != nil {
						log.Fatal(err)
					}
				}

				// TODO need flag for not removing generated output but
				//  save it into the expected directory
				if !t.Failed() {
					_ = os.RemoveAll(s.Settings.OutputFilePath)
				}
			}()

			err = createTestData(db.SQLDriver(), s)
			if err != nil {
				t.Logf("could not create test data: %v", err)
				t.Fail()
				return
			}

			err = os.MkdirAll(s.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Logf("could not create output file path: %v", err)
				t.Fail()
				return
			}

			version, err := db.Version()
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				fmt.Printf("running tests against database %s\n", version)
			}

			writer := output.NewFileWriter(s.Settings.OutputFilePath)

			err = cli.Run(s.Settings, db, writer)
			assert.NoError(t, err)

			checkFiles(t, s)
		})
	}
}

func checkFiles(t *testing.T, s *dbSettings) {
	expectedPattern := filepath.Join(s.dataFilepath, expectedDirectoryName, s.Settings.Prefix+"*")
	expectedFiles, err := filepath.Glob(expectedPattern)
	assert.NoError(t, err)

	outputPattern := filepath.Join(s.Settings.OutputFilePath, s.Settings.Prefix+"*")
	outputFiles, err := filepath.Glob(outputPattern)
	assert.NoError(t, err)

	if len(expectedFiles) != len(outputFiles) {
		t.Fatalf("number of expected and output files differ: %d vs. %d",
			len(expectedFiles), len(outputFiles))
	}

	sort.Strings(expectedFiles)
	sort.Strings(outputFiles)

	for i := range expectedFiles {
		expectedFile, err := os.ReadFile(expectedFiles[i])
		assert.NoError(t, err)
		outputFile, err := os.ReadFile(outputFiles[i])
		assert.NoError(t, err)
		assert.Equal(t, string(expectedFile), string(outputFile), "file %q", expectedFiles[i])
	}
}

func setupDatabase(s *dbSettings) (database.Database, func() error, error) {
	log.Printf("spinning up database %s:%s ...\n", s.dockerImage, s.version)
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to Docker: %v", err)
	}
	pool.MaxWait = 5 * time.Minute

	resource, err := pool.Run(s.dockerImage, s.version, s.env)
	if err != nil {
		return nil, nil, fmt.Errorf("could not start resource: %s", err)
	}

	purgeFn := func() error {
		if err := pool.Purge(resource); err != nil {
			return fmt.Errorf("could not purge resource: %v", err)
		}
		return nil
	}

	var db database.Database

	// give docker some time to spin up the database
	// also reduce unnecessary output of
	// > packets.go:36: unexpected EOF errors when spinning up mysql
	if s.DbType == settings.DBTypeMySQL {
		time.Sleep(25 * time.Second)
	}

	if err = pool.Retry(func() error {
		newSettings := s.Settings
		port := resource.GetPort(s.Port + "/tcp")
		if port != "" {
			newSettings.Port = port
		}
		db = database.New(newSettings)
		err := db.Connect()
		if err != nil {
			if newSettings.Verbose {
				fmt.Println(err.Error())
			}
			return err
		}
		return nil
	}); err != nil {
		return nil, purgeFn, fmt.Errorf("could not connect to Docker: %v", err)
	}

	return db, purgeFn, nil
}

func createTestData(db *sqlx.DB, s *dbSettings) error {
	testDataPattern := filepath.Join(s.dataFilepath, testdataDirectoryName, "*.sql")
	files, err := filepath.Glob(testDataPattern)
	if err != nil {
		return fmt.Errorf("could not find sql testdata: %v", err)
	}

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("could not read %q: %v", f, err)
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
				return fmt.Errorf("could not create testdata %q: %v", q, err)
			}
		}
	}

	return nil
}
