// Command generator generates code and tests.
package main

import (
	"github.com/bassosimone/apiclient/internal/cmd/internal/fatalx"
	"golang.org/x/sys/execabs"
)

func gofmt(filename string) {
	cmd := execabs.Command("go", "fmt", filename)
	fatalx.OnError(cmd.Run(), "cmd.Run failed")
}

func main() {
	{
		filep := mustCreateFile("generated.go")
		defer filep.Close()
		genAPIModel(filep)
		genSwagger(filep)
	}
	gofmt("generated.go")
	{
		filep := mustCreateFile("generated_test.go")
		defer filep.Close()
		genTests(filep)
	}
	gofmt("generated_test.go")
}
