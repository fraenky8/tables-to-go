package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralDatabase_andInClause(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc         string
		field        string
		params       []string
		args         []any
		expected     string
		expectedArgs []any
	}{
		{
			desc:         "empty field returns early",
			params:       nil,
			args:         nil,
			expected:     "",
			expectedArgs: nil,
		},
		{
			desc:         "nil params returns early",
			field:        "table_name",
			params:       nil,
			args:         nil,
			expected:     "",
			expectedArgs: nil,
		},
		{
			desc:         "zero params returns early",
			field:        "table_name",
			params:       []string{},
			args:         nil,
			expected:     "",
			expectedArgs: nil,
		},
		{
			desc:         "one param returns AND IN clause",
			field:        "table_name",
			params:       []string{"foo"},
			args:         []any{},
			expected:     "AND table_name IN (?)",
			expectedArgs: []any{"foo"},
		},
		{
			desc:         "multiple params returns AND IN clause",
			field:        "table_name",
			params:       []string{"foo", "bar", "baz", "qux"},
			args:         []any{},
			expected:     "AND table_name IN (?,?,?,?)",
			expectedArgs: []any{"foo", "bar", "baz", "qux"},
		},
		{
			desc:         "multiple params and existing args returns AND IN clause",
			field:        "table_name",
			params:       []string{"baz", "qux", "quux", "corge"},
			args:         []any{"foo", "bar"},
			expected:     "AND table_name IN (?,?,?,?)",
			expectedArgs: []any{"foo", "bar", "baz", "qux", "quux", "corge"},
		},
	}
	for i, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := (&GeneralDatabase{}).andInClause(tt.field, tt.params, &tests[i].args)
			assert.Equal(t, tt.expected, actual)
			assert.Equal(t, tt.expectedArgs, tests[i].args)
		})
	}
}
