// Package fmtx extends the fmt package
package fmtx

import (
	"fmt"
	"io"

	"github.com/bassosimone/apiclient/internal/fatalx"
)

// Fprintf is like fmt.Fprintf but calls log.Fatal on failure.
func Fprintf(w io.Writer, format string, v ...interface{}) {
	_, err := fmt.Fprintf(w, format, v...)
	fatalx.OnError(err, "fmt.Fprintf failed")
}

// Fprint is like fmt.Fprint but calls log.Fatal on failure.
func Fprint(w io.Writer, v ...interface{}) {
	_, err := fmt.Fprint(w, v...)
	fatalx.OnError(err, "fmt.Fprint failed")
}
