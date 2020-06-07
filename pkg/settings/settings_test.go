package settings

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettings_Verify(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *Settings
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:     "default settings produce no error",
			settings: New,
			isError:  assert.NoError,
		},
		{
			desc: "wrong output file path produces error",
			settings: func() *Settings {
				s := New()
				s.OutputFilePath = ""
				return s
			},
			isError: assert.Error,
		},
		{
			desc: "output file path with file produces error",
			settings: func() *Settings {
				s := New()
				ex, err := os.Executable()
				assert.Nil(t, err)
				s.OutputFilePath = ex
				return s
			},
			isError: assert.Error,
		},
		{
			desc: "empty package name produces error",
			settings: func() *Settings {
				s := New()
				s.PackageName = ""
				return s
			},
			isError: assert.Error,
		},
		{
			desc: "set v-verbose mode activates verbose mode without error",
			settings: func() *Settings {
				s := New()
				s.VVerbose = true
				return s
			},
			isError: assert.NoError,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			settings := test.settings()
			err := settings.Verify()
			test.isError(t, err)
		})
	}
}

func TestSettings_IsNullTypeSQL(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *Settings
		expected bool
	}{
		{
			desc:     "in default settings sql NULL type is activated",
			settings: New,
			expected: true,
		},
		{
			desc: "explicit enabled sql NULL type ativates sql NULL type",
			settings: func() *Settings {
				s := New()
				s.Null = NullTypeSQL
				return s
			},
			expected: true,
		},
		{
			desc: "native NULL type deativates sql NULL type",
			settings: func() *Settings {
				s := New()
				s.Null = NullTypeNative
				return s
			},
			expected: false,
		},
		{
			desc: "primitve NULL type deativates sql NULL type",
			settings: func() *Settings {
				s := New()
				s.Null = NullTypePrimitive
				return s
			},
			expected: false,
		},
		{
			desc: "any other NULL type deativates sql NULL type",
			settings: func() *Settings {
				s := New()
				s.Null = NullType("any")
				return s
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			settings := test.settings()
			actual := settings.IsNullTypeSQL()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSettings_ShouldInitialism(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *Settings
		expected bool
	}{
		{
			desc:     "in default settings initialism is activated",
			settings: New,
			expected: true,
		},
		{
			desc: "explicit enabled initialism ativates initialism",
			settings: func() *Settings {
				s := New()
				s.NoInitialism = false
				return s
			},
			expected: true,
		},
		{
			desc: "disabled initialism deactivates initialism",
			settings: func() *Settings {
				s := New()
				s.NoInitialism = true
				return s
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			settings := test.settings()
			actual := settings.ShouldInitialism()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSettings_IsOutputFormatCamelCase(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *Settings
		expected bool
	}{
		{
			desc:     "in default settings camel case is activated",
			settings: New,
			expected: true,
		},
		{
			desc: "explicit enabled camel case ativates initialism",
			settings: func() *Settings {
				s := New()
				s.OutputFormat = OutputFormatCamelCase
				return s
			},
			expected: true,
		},
		{
			desc: "disabled camel case deactivates camel case",
			settings: func() *Settings {
				s := New()
				s.OutputFormat = OutputFormatOriginal
				return s
			},
			expected: false,
		},
		{
			desc: "any other output format deativates camel case",
			settings: func() *Settings {
				s := New()
				s.OutputFormat = OutputFormat("any")
				return s
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			settings := test.settings()
			actual := settings.IsOutputFormatCamelCase()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSettings_IsFileNameFormatSnakeCase(t *testing.T) {
	tests := []struct {
		desc     string
		settings func() *Settings
		expected bool
	}{
		{
			desc:     "in default settings camel case will be used",
			settings: New,
			expected: false,
		},
		{
			desc: "use snake case",
			settings: func() *Settings {
				s := New()
				s.FileNameFormat = FileNameFormatSnakeCase
				return s
			},
			expected: true,
		},
		{
			desc: "any other output format will converted to camel case",
			settings: func() *Settings {
				s := New()
				s.FileNameFormat = FileNameFormat("any")
				return s
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			settings := test.settings()
			actual := settings.IsFileNameFormatSnakeCase()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDbType_Set(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected DbType
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:     "typed supported db type produces no error and gets set",
			input:    string(DbTypePostgresql),
			expected: DbTypePostgresql,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed supported db type produces no error and gets set",
			input:    string("pg"),
			expected: DbTypePostgresql,
			isError:  assert.NoError,
		},
		{
			desc:     "empty db type produces no error and gets default",
			input:    "",
			expected: DbTypePostgresql,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed unsupported db type produces error and invalid db type",
			input:    string("invalid"),
			expected: DbType("invalid"),
			isError:  assert.Error,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := DbTypeMySQL
			err := actual.Set(test.input)
			test.isError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestNullType_Set(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected NullType
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:     "typed supported NULL type produces no error and gets set",
			input:    string(NullTypePrimitive),
			expected: NullTypePrimitive,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed supported NULL type produces no error and gets set",
			input:    string("native"),
			expected: NullTypeNative,
			isError:  assert.NoError,
		},
		{
			desc:     "empty NULL type produces no error and gets default",
			input:    "",
			expected: NullTypeSQL,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed unsupported NULL type produces error and invalid NULL type",
			input:    string("invalid"),
			expected: NullType("invalid"),
			isError:  assert.Error,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := NullTypeSQL
			err := actual.Set(test.input)
			test.isError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestOutputFormat_Set(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected OutputFormat
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:     "typed supported output type produces no error and gets set",
			input:    string(OutputFormatCamelCase),
			expected: OutputFormatCamelCase,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed supported output type produces no error and gets set",
			input:    string("o"),
			expected: OutputFormatOriginal,
			isError:  assert.NoError,
		},
		{
			desc:     "empty output type produces no error and gets default",
			input:    "",
			expected: OutputFormatCamelCase,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed unsupported output type produces error and invalid output type",
			input:    string("invalid"),
			expected: OutputFormat("invalid"),
			isError:  assert.Error,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := OutputFormatCamelCase
			err := actual.Set(test.input)
			test.isError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFileNameFormat_Set(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected FileNameFormat
		isError  assert.ErrorAssertionFunc
	}{
		{
			desc:     "typed supported filename type produces no error and gets set",
			input:    string(FileNameFormatCamelCase),
			expected: FileNameFormatCamelCase,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed supported filename type produces no error and gets set",
			input:    string("c"),
			expected: FileNameFormatCamelCase,
			isError:  assert.NoError,
		},
		{
			desc:     "empty output type produces no error and gets default",
			input:    "",
			expected: FileNameFormatCamelCase,
			isError:  assert.NoError,
		},
		{
			desc:     "string typed unsupported output type produces error and invalid output type",
			input:    string("invalid"),
			expected: FileNameFormat("invalid"),
			isError:  assert.Error,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := FileNameFormatCamelCase
			err := actual.Set(test.input)
			test.isError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSprintfSupportedDbTypes(t *testing.T) {
	tests := []struct {
		desc     string
		expected int
	}{
		{
			desc:     "print all supported DB types",
			expected: len(SupportedDbTypes),
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			printed := SprintfSupportedDbTypes()
			assert.NotEmpty(t, printed)

			actual := strings.Split(printed, " ")
			assert.Equal(t, test.expected, len(actual))
		})
	}
}

func TestSprintfSupportedNullTypes(t *testing.T) {
	tests := []struct {
		desc     string
		expected int
	}{
		{
			desc:     "print all supported NULL types",
			expected: len(supportedNullTypes),
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			printed := SprintfSupportedNullTypes()
			assert.NotEmpty(t, printed)

			actual := strings.Split(printed, " ")
			assert.Equal(t, test.expected, len(actual))
		})
	}
}
