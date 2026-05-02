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

func TestCommentsMode_Set(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		input    string
		expected CommentsMode
		isErr    assert.ErrorAssertionFunc
	}{
		{
			desc:     "unsupported comments mode returns error",
			input:    "invalid",
			expected: "invalid",
			isErr:    assert.Error,
		},
		{
			desc:     "comments mode off is accepted",
			input:    "off",
			expected: CommentsModeOff,
			isErr:    assert.NoError,
		},
		{
			desc:     "comments mode line is accepted",
			input:    "line",
			expected: CommentsModeLine,
			isErr:    assert.NoError,
		},
		{
			desc:     "comments mode inline is accepted",
			input:    "inline",
			expected: CommentsModeInline,
			isErr:    assert.NoError,
		},
		{
			desc:     "empty value maps to line",
			input:    "",
			expected: CommentsModeLine,
			isErr:    assert.NoError,
		},
		{
			desc:     "true value maps to line",
			input:    "true",
			expected: CommentsModeLine,
			isErr:    assert.NoError,
		},
		{
			desc:     "false value maps to off",
			input:    "false",
			expected: CommentsModeOff,
			isErr:    assert.NoError,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := CommentsModeOff
			err := actual.Set(test.input)
			test.isErr(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCommentsMode_IsBoolFlag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		input    CommentsMode
		expected bool
	}{
		{
			desc:     "comments mode is bool-like flag",
			input:    CommentsModeOff,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := test.input.IsBoolFlag()
			assert.Equal(t, test.expected, actual)
		})
	}
}
