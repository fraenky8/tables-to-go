package tablestogo

import (
	"testing"
)

func TestCamelCaseString(t *testing.T) {

	var Cases = []struct {
		input    string
		expected string
	}{
		{"a_test_table", "ATestTable"},
		{"A_TeSt_tAbLe", "ATestTable"},
		{"id", "Id"},
		{"ATeSttAbLe", "ATeSttAbLe"},
		{"A_1test_2table", "A1test2table"},
	}

	for _, tt := range Cases {
		actual := CamelCaseString(tt.input)
		if actual != tt.expected {
			t.Errorf("camelCaseString(%q): expected %q, actual %q", tt.input, tt.expected, actual)
		}
	}

}
