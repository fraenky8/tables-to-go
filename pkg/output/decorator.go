package output

import (
	"fmt"
	"go/format"
	"strings"
)

// Decorator represents an interface to decorate the given content.
type Decorator interface {
	Decorate(content string) (string, error)
}

// FormatDecorator applies a formatting decoration to the given content.
type FormatDecorator struct{}

// Decorate is the implementation of the Decorator interface.
func (FormatDecorator) Decorate(content string) (string, error) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		return content, fmt.Errorf("could not format content: %v", err)
	}
	return string(formatted), nil
}

// ImportDecorator removes empty import statements from the given content.
type ImportDecorator struct{}

// Decorate is the implementation of the Decorator interface.
func (ImportDecorator) Decorate(content string) (string, error) {
	// fight the symptom instead of the cause - if we didnt imported anything, remove it
	decorated := strings.ReplaceAll(content, "\nimport ()\n", "")
	return decorated, nil
}
