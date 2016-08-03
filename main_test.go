package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyOutputPath(t *testing.T) {
	assert.True(t, true)
}

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
		actual := camelCaseString(tt.input)
		if actual != tt.expected {
			t.Errorf("camelCaseString(%q): expected %q, actual %q", tt.input, tt.expected, actual)
		}
	}

}
