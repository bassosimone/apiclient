// Package fatalx simplifies checking for errors.
package fatalx

import "log"

// OnError calls log.Fatal if err is not nil
func OnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}
