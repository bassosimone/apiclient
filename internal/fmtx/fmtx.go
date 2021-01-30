// Package fmtx extends the fmt package
package fmtx

import (
	"fmt"
	"io"
	"log"
)

// Fprintf is like fmt.Fprintf but calls log.Fatal on failure.
func Fprintf(w io.Writer, format string, v ...interface{}) {
	if _, err := fmt.Fprintf(w, format, v...); err != nil {
		log.Fatal(err)
	}
}

// Fprint is like fmt.Fprint but calls log.Fatal on failure.
func Fprint(w io.Writer, v ...interface{}) {
	if _, err := fmt.Fprint(w, v...); err != nil {
		log.Fatal(err)
	}
}
