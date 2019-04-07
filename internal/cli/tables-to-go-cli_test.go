package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toInitialisms(t *testing.T) {
	tests := []struct {
		desc     string
		intput   string
		expected string
	}{
		{
			desc:     "id at the end of string should be upper case",
			intput:   "userId",
			expected: "userID",
		},
		{
			desc:     "id at the beginning of string should be upper case",
			intput:   "Iduser",
			expected: "IDuser",
		},
		{
			desc:     "id in the middle of string should be upper case",
			intput:   "userIdprim",
			expected: "userIDprim",
		},
		{
			desc:     "multiple occurences should be upper case",
			intput:   "userIdasJsonWithUrl",
			expected: "userIDasJSONWithURL",
		},
		{
			desc:     "multiple id in the string should be upper case",
			intput:   "IduserId",
			expected: "IDuserID",
		},
		{
			desc:     "non replacement in the string should be return original string",
			intput:   "name",
			expected: "name",
		},
		{
			desc:     "replacements only in the string should be return original string",
			intput:   "IdjsonuRlHtTp",
			expected: "IDJSONURLHTTP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := toInitialisms(tt.intput)
			assert.Equal(t, tt.expected, actual, "test case input: "+tt.intput)
		})
	}
}
