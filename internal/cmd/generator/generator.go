// Command generator generates code and tests.
package main

func main() {
	{
		filep := mustCreateFile("generated.go")
		defer filep.Close()
		genAPIModel(filep)
		genSwagger(filep)
	}
	{
		filep := mustCreateFile("generated_test.go")
		defer filep.Close()
		genTests(filep)
	}
}
