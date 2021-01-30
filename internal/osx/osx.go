// Package osx extends the os package.
package osx

import (
	"io"
	"os"

	"github.com/bassosimone/apiclient/internal/fatalx"
)

// File is an open file
type File interface {
	WriteString(s string)
	Close()
	io.Writer
}

// MustCreate creates a new File or calls log.Fatal
func MustCreate(pathname string) File {
	filep, err := os.Create(pathname)
	fatalx.OnError(err, "os.Create failed")
	return &fileWrapper{File: filep}
}

type fileWrapper struct {
	*os.File
}

func (fw *fileWrapper) WriteString(s string) {
	_, err := fw.File.WriteString(s)
	fatalx.OnError(err, "WriteString failed")
}

func (fw *fileWrapper) Close() {
	fatalx.OnError(fw.File.Close(), "Close failed")
}
