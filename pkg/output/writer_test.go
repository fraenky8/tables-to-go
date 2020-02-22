package output

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileWriter_Write(t *testing.T) {
	tests := []struct {
		desc      string
		tableName string
		content   string
		isError   assert.ErrorAssertionFunc
	}{
		{
			desc:      "valid table name and valid content should write a file",
			tableName: "Bar",
			content:   "package dto\ntype Bar struct {\nID int `db:\"id\"`\n}",
			isError:   assert.NoError,
		},
		{
			desc:      "valid table name and invalid content should produce an error",
			tableName: "Bar",
			content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
			isError:   assert.Error,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			wd, err := os.Getwd()
			if err != nil {
				t.Fatalf("expected non error, got: %s", err)
			}

			file := path.Join(wd, test.tableName+FileWriterExtension)
			defer os.Remove(file)
			t.Logf("writing file: %s\n", file)

			fw := NewFileWriter(path.Dir(file))
			err = fw.Write(test.tableName, test.content)
			if err != nil {
				test.isError(t, err)
				return
			}

			fi, err := os.Stat(file)
			if err != nil {
				t.Fatalf("expected non error, got: %s", err)
			}
			assert.True(t, fi.Size() > 0)
		})
	}
}
