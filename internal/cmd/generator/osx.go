package main

import (
	"io"
	"os"

	"github.com/bassosimone/apiclient/internal/cmd/internal/fatalx"
)

// file is an open file
type file interface {
	WriteString(s string)
	Close()
	io.Writer
}

// mustCreateFile creates a new File or calls log.Fatal. The returned
// file is such that any write error results in a failure.
func mustCreateFile(pathname string) file {
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
