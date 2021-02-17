// Command generator generates code and tests.
package main

func main() {
	filep := mustCreateFile("generated.go")
	defer filep.Close()
	genAPIModel(filep)
	genSwagger(filep)
}
