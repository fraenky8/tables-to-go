package output

import (
	"io/ioutil"
)

const (
	// FileWriterExtension is the extension to write files of.
	FileWriterExtension = ".go"
)

// Writer represents an interface to write the produced struct content.
type Writer interface {
	Write(tableName string, content string) error
}

// FileWriter is a writer that writes to a file given by the path and the table name.
type FileWriter struct {
	path string
}

// NewFileWriter constructs a new FileWriter.
func NewFileWriter(path string) *FileWriter {
	return &FileWriter{path: path}
}

// Write is the implementation of the Writer interface. The FilerWriter writes
// decorated content to the file specified by the given path and table name.
func (w FileWriter) Write(tableName string, content string) error {
	fileName := w.path + tableName + FileWriterExtension

	decorated, err := decorate(content)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(decorated), 0666)
}

// decorate applies some decorations like formatting and empty import removal.
func decorate(content string) (decorated string, err error) {
	decorators := []Decorator{
		FormatDecorator{},
		ImportDecorator{},
	}
	for _, decorator := range decorators {
		content, err = decorator.Decorate(content)
		if err != nil {
			return content, err
		}
	}

	return content, nil
}
