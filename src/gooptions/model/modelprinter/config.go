package modelprinter

import (
	"io"
	"os"
)

type Config struct {
	FileWriteCloserFactory func(path string) (io.WriteCloser, error)
}

var DefaultFileWriteCloserFactory = func(path string) (io.WriteCloser, error) {
	file, err := os.Create(path)
	return file, err
}
