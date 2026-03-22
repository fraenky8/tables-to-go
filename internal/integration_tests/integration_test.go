//go:build integration

package integration_tests

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strconv"
	"syscall"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
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

const (
	// Note: The more integration tests, the higher we have to set this time.
	// Otherwise, the resources might be purged before your tests are finished.
	resourceExpirationSeconds = 300
)

var (
	pool *dockertest.Pool
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

var (
	isCI bool
)

func init() {
	isCI, _ = strconv.ParseBool(os.Getenv("CI"))

	// Suppress logs of "packets.go:36: unexpected EOF"
	_ = mysql.SetLogger(nopLogger{})
}

func TestMain(m *testing.M) {
	logMsg := "running Tables-to-Go integration tests"
	if isCI {
		logMsg += " on CI"
	}
	log.Println(logMsg)

	log.Println("creating Docker pool...")

	var err error
	pool, err = newPool()
	if err != nil {
		log.Fatalf("error connecting to Docker: %v", err)
	}

	os.Exit(m.Run())
}

func newPool() (*dockertest.Pool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("os.UserHomeDir failed: %w", err)
	}

	endpoints := []string{
		"", // first try dockertest's default based on env vars
		"unix://" + filepath.Join(home, ".docker/run/docker.sock"),     // In case the symlink is missing.
		"unix://" + filepath.Join(home, ".rd/docker.sock"),             // Rancher Desktop
		"unix://" + filepath.Join(home, ".colima/default/docker.sock"), // Colima
	}
	for _, endpoint := range endpoints {
		pool, err := dockertest.NewPool(endpoint)
		if err != nil {
			return nil, fmt.Errorf("dockertest.NewPool failed: %w", err)
		}

		// Our "ping" function
		_, err = pool.NetworksByName("none")
		if err != nil {
			log.Println("docker:", err)
			continue
		}

		log.Printf("using %q", endpoint)
		return pool, nil
	}

	return nil, fmt.Errorf("could not create pool from any given endpoint")
}

func registerCleanupSignalHandler(t *testing.T, container string) chan struct{} {
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL}
	terminate := make(chan os.Signal, len(signals))
	done := make(chan struct{})
	signal.Notify(terminate, signals...)
	go func() {
		select {
		case <-done:
			return
		case s := <-terminate:
			t.Log()
			t.Log("got signal:", s.String())
			t.Logf("removing container %q", container)
			_ = pool.RemoveContainerByName(container)
			// Ignoring error here because it might be called multiple times due
			// to multiple signals arriving. The first of them will remove the
			// container already leading subsequent calls error. But we are not
			// interested in an error saying that the container does not exist (anymore).
		}
	}()
	return done
}

func TestIntegration(t *testing.T) {
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
				// s.Verbose = true
				// s.VVerbose = true

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
				// s.Verbose = true
				// s.VVerbose = true

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
				// s.Verbose = true
				// s.VVerbose = true

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
				// s.Verbose = true
				// s.VVerbose = true

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
				// s.Verbose = true
				// s.VVerbose = true

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
		{
			desc: "postgres 18",
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
				// s.Verbose = true
				// s.VVerbose = true

				dbs := &dbSettings{
					Settings: s,

					dataFilepath: "postgres",

					dockerImage: "postgres",
					version:     "18",
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

			db, purgeFn := setupDatabase(t, s)
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

			loadTestData(t, db.SQLDriver(), s)

			err := os.MkdirAll(s.Settings.OutputFilePath, 0755)
			if err != nil {
				t.Errorf("could not create output file path: %v", err)
				return
			}

			version, err := db.Version()
			if err != nil {
				t.Logf("could not get version: %v", err)
			} else {
				t.Logf("running tests against database %s\n", version)
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

func setupDatabase(t *testing.T, s *dbSettings) (database.Database, func() error) {
	t.Logf("spinning up database %s:%s ...\n", s.dockerImage, s.version)

	containerName := fmt.Sprintf("tables_to_go_%s_%s_integration", s.dockerImage, s.version)

	// Note: registering before the resource gets created because it happens that
	// the resource gets created but for some reason we cannot figure out if it's
	// ready or not. Using CTRL+C then would result in existing resource not
	// being cleaned up.
	done := registerCleanupSignalHandler(t, containerName)

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       containerName,
		Repository: s.dockerImage,
		Tag:        s.version,
		Env:        s.env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		t.Fatalf("could not start resource: %v", err)
	}
	_ = resource.Expire(resourceExpirationSeconds)

	purgeFn := func() error {
		if err := pool.Purge(resource); err != nil {
			return fmt.Errorf("could not purge MySQL: %w", err)
		}
		done <- struct{}{}
		return nil
	}

	var db database.Database

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
				t.Log(err.Error())
			}
			return err
		}
		return nil
	}); err != nil {
		t.Logf("could not connect to database: %v", err)
	}

	return db, purgeFn
}

func loadTestData(t *testing.T, db *sqlx.DB, s *dbSettings) {
	testDataPattern := filepath.Join(s.dataFilepath, testdataDirectoryName, "*.sql")
	files, err := filepath.Glob(testDataPattern)
	if err != nil {
		t.Errorf("could not find sql testdata: %v", err)
		return
	}

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Errorf("could not read %q: %v", f, err)
			return
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
				t.Errorf("could not create testdata %q: %v", f, err)
				return
			}
		}
	}
}
