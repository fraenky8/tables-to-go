package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatDecorator_Decorate(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:  "well formed golang code should get decorated",
			input: "package dto\ntype Bar struct {\nID int `db:\"id\"`\n}",
			expected: `package dto

type Bar struct {
	ID int ` + "`db:\"id\"" + "`" + `
}
`,
			isError: assert.NoError,
		},
		{
			desc:     "arbitrary text throws error",
			input:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
			expected: "",
			isError:  assert.Error,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			decorator := FormatDecorator{}
			actual, err := decorator.Decorate(test.input)
			if err != nil {
				test.isError(t, err)
				return
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestImportDecorator_Decorate(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:     "well formed golang code with inport-statement should get decorated",
			input:    "package dto\n\nimport ()\n\ntype Bar struct {\nID int `db:\"id\"`\n}",
			expected: "package dto\n\ntype Bar struct {\nID int `db:\"id\"`\n}",
			isError:  assert.NoError,
		},
		{
			desc:     "well formed golang code without inport-statement should stay unchanged",
			input:    "package dto\n\ntype Bar struct {\nID int `db:\"id\"`\n}",
			expected: "package dto\n\ntype Bar struct {\nID int `db:\"id\"`\n}",
			isError:  assert.NoError,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			decorator := ImportDecorator{}
			actual, err := decorator.Decorate(test.input)
			if err != nil {
				test.isError(t, err)
				return
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
