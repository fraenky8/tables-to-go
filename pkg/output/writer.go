package output

import (
	"io/ioutil"
	"path"
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
	path       string
	decorators []Decorator
}

// NewFileWriter constructs a new FileWriter.
func NewFileWriter(path string) *FileWriter {
	return &FileWriter{
		path: path,
		decorators: []Decorator{
			FormatDecorator{},
			ImportDecorator{},
		},
	}
}

// Write is the implementation of the Writer interface. The FilerWriter writes
// decorated content to the file specified by the given path and table name.
func (w FileWriter) Write(tableName string, content string) error {
	fileName := path.Join(w.path, tableName+FileWriterExtension)

	decorated, err := w.decorate(content)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(decorated), 0666)
}

// decorate applies some decorations like formatting and empty import removal.
func (w FileWriter) decorate(content string) (decorated string, err error) {
	for _, decorator := range w.decorators {
		content, err = decorator.Decorate(content)
		if err != nil {
			return content, err
		}
	}

	return content, nil
}
