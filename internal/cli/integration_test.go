// +build integration

package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/pkg/database"
	"github.com/fraenky8/tables-to-go/pkg/output"
	"github.com/fraenky8/tables-to-go/pkg/settings"
)

const (
	expectedFilePath = "testdata/expected"
	outputFilePath   = "testdata/output"
)

type cliWriter struct{}

func (c cliWriter) Write(tableName string, content string) error {
	_, err := fmt.Println(content)
	return err
}

type dbSettings struct {
	*settings.Settings

	dockerImage string
	version     string
	env         []string
}

func TestIntegration(t *testing.T) {
	log.Println("running Tables-to-Go integration tests")

	tests := []struct {
		desc     string
		settings func() dbSettings
	}{
		{
			desc: "mysql 8",
			settings: func() dbSettings {
				s := settings.New()
				s.DbType = settings.DbTypeMySQL
				s.User = "root"
				s.Pswd = "mysecretpassword"
				s.DbName = "public"
				s.Host = "localhost"
				s.Port = "3306"
				s.OutputFilePath = outputFilePath
				//s.Verbose = true
				//s.VVerbose = true

				return dbSettings{
					dockerImage: "mysql",
					version:     "8",
					env:         []string{"MYSQL_DATABASE=public", "MYSQL_ROOT_PASSWORD=mysecretpassword"},
					Settings:    s,
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			s := test.settings()

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

				_ = os.RemoveAll(outputFilePath)
				_ = os.MkdirAll(outputFilePath, 0755)
			}()

			writer := output.NewFileWriter(outputFilePath)
			//writer := cliWriter{}

			s.Settings.OutputFilePath = outputFilePath
			prefix := strings.Title(s.dockerImage + s.version)
			s.Settings.Prefix = prefix + "_"

			err = Run(s.Settings, db, writer)
			assert.NoError(t, err)

			checkFiles(t, prefix)
		})
	}
}

func checkFiles(t *testing.T, prefix string) {
	expectedPattern := filepath.Join(expectedFilePath, prefix+"*")
	expected, err := filepath.Glob(expectedPattern)
	assert.NoError(t, err)

	actualPattern := filepath.Join(outputFilePath, prefix+"*")
	actual, err := filepath.Glob(actualPattern)
	assert.NoError(t, err)

	if len(expected) != len(actual) {
		t.Fatalf("expected and actual files differ in length: %v (%d) vs. %v (%d)",
			expected, len(expected), actual, len(actual))
	}

	sort.Strings(expected)
	sort.Strings(actual)

	for i, ef := range expected {
		af := actual[i]
		f1, err := ioutil.ReadFile(ef)
		assert.NoError(t, err)
		f2, err := ioutil.ReadFile(af)
		assert.NoError(t, err)
		assert.True(t, bytes.Equal(f1, f2))
	}
}

func setupDatabase(settings dbSettings) (database.Database, func() error, error) {
	log.Printf("spinning up Database %s:%s ...\n", settings.dockerImage, settings.version)
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to Docker: %v", err)
	}
	pool.MaxWait = 2 * time.Minute

	resource, err := pool.Run(settings.dockerImage, settings.version, settings.env)
	if err != nil {
		return nil, nil, fmt.Errorf("could not start resource: %s", err)
	}
	_ = resource.Expire(60)

	purgeFn := func() error {
		if err := pool.Purge(resource); err != nil {
			return fmt.Errorf("could not purge resource: %v", err)
		}
		return nil
	}

	var db database.Database

	// give docker some time to spin up the database
	// also reduce unnecessary output of packets.go:36: unexpected EOF errors when spinning up mysql
	//time.Sleep(25 * time.Second)

	if err = pool.Retry(func() error {
		s := settings.Settings
		port := resource.GetPort(s.Port + "/tcp")
		if port != "" {
			s.Port = port
		}
		db = database.New(s)
		return db.Connect()
	}); err != nil {
		return nil, purgeFn, fmt.Errorf("could not connect to Docker: %v", err)
	}

	err = createTestData(db.SQLDriver(), settings.dockerImage+settings.version)
	if err != nil {
		return nil, purgeFn, err
	}

	return db, purgeFn, nil
}

func createTestData(db *sqlx.DB, dockerImage string) error {
	data, err := ioutil.ReadFile("testdata/" + dockerImage + ".sql")
	if err != nil {
		return fmt.Errorf("could not read sql testdata: %v", err)
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
			return fmt.Errorf("could not insert testdata %q: %v", q, err)
		}
	}

	return nil
}
