package settings

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsFlag_Set(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		args     []string
		expected StringsFlag
	}{
		{
			desc:     "no -table flag, no values",
			args:     []string{},
			expected: nil,
		},
		{
			desc:     "one -table flag",
			args:     []string{"-table", "test-table-1"},
			expected: StringsFlag{"test-table-1"},
		},
		{
			desc:     "multiple -table flags",
			args:     []string{"-table", "test-table-1", "-table", "test-table-2", "-table", "test-table-3"},
			expected: StringsFlag{"test-table-1", "test-table-2", "test-table-3"},
		},
		{
			desc:     "-table flag with comma separator",
			args:     []string{"-table", "test-table-1,test-table-2"},
			expected: StringsFlag{"test-table-1", "test-table-2"},
		},
		{
			desc:     "mixed -table flag with comma separator and standalone",
			args:     []string{"-table", "test-table-1,test-table-2", "-table", "test-table-3"},
			expected: StringsFlag{"test-table-1", "test-table-2", "test-table-3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var actual StringsFlag
			fs := flag.NewFlagSet("test", flag.ExitOnError)
			fs.Var(&actual, "table", "")
			err := fs.Parse(tt.args)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
