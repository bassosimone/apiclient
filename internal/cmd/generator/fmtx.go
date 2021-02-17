package main

import (
	"fmt"
	"io"

	"github.com/bassosimone/apiclient/internal/cmd/internal/fatalx"
)

// fprintf is like fmt.fprintf but calls log.Fatal on failure.
func fprintf(w io.Writer, format string, v ...interface{}) {
	_, err := fmt.Fprintf(w, format, v...)
	fatalx.OnError(err, "fmt.Fprintf failed")
}

// fprint is like fmt.fprint but calls log.Fatal on failure.
func fprint(w io.Writer, v ...interface{}) {
	_, err := fmt.Fprint(w, v...)
	fatalx.OnError(err, "fmt.Fprint failed")
}
