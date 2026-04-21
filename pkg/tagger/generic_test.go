package tagger

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraenky8/tables-to-go/v2/pkg/database"
)

func TestGeneric_GenerateTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		tagger   Generic
		column   database.Column
		expected string
	}{
		{
			desc:     "empty tag name generates no tag",
			tagger:   NewGeneric(""),
			column:   database.Column{Name: "column_name"},
			expected: "",
		},
		{
			desc:     "custom tag name generates tag with column name",
			tagger:   NewGeneric("json"),
			column:   database.Column{Name: "column_name"},
			expected: `json:"column_name"`,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := test.tagger.GenerateTag(nil, test.column)
			assert.Equal(t, test.expected, actual)
		})
	}
}
