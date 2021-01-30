package fatalx

import "log"

// OnError calls log.Fatal if err is not nil
func OnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

// IfNil calls log.Fatal if ptr is nil
func IfNil(ptr interface{}, msg string) {
	if ptr == nil {
		log.Fatal(msg)
	}
}
