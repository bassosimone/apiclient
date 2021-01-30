// +build ignore

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"time"
)

func fatalOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func fatalIfNil(pkg *ast.Package, msg string) {
	if pkg == nil {
		log.Fatal(msg)
	}
}

func isRequestOrResponse(decl ast.Decl) bool {
	gendecl, good := decl.(*ast.GenDecl)
	if !good {
		return false
	}
	switch gendecl.Tok {
	case token.TYPE, token.VAR:
		return true
	default:
		return false
	}
}

func main() {
	outfile := flag.String("outfile", "datamodel.go", "Output file")
	pkgdir := flag.String("pkg", "", "Package directory path")
	flag.Parse()

	filep, err := os.Create(*outfile)
	fatalOnError(err, "os.Create failed")

	_, err = filep.Write([]byte(
		"// Code generated by go generate; DO NOT EDIT.\n"))
	fatalOnError(err, "filep.Write failed")

	_, err = fmt.Fprintf(filep, "// %+v\n\n", time.Now())
	fatalOnError(err, "filep.Write failed")

	_, err = filep.Write([]byte("package apiclient\n\n"))
	fatalOnError(err, "filep.Write failed")

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, *pkgdir, nil, parser.ParseComments)
	fatalOnError(err, "parser.ParseDir failed")

	model := pkgs["datamodel"]
	fatalIfNil(model, "cannot find the datamodel package")

	for _, fdata := range model.Files {
		for _, decl := range fdata.Decls {
			if !isRequestOrResponse(decl) {
				continue
			}

			err = format.Node(filep, fset, decl)
			fatalOnError(err, "format.Node failed")

			_, err = filep.Write([]byte("\n\n"))
			fatalOnError(err, "filep.Write failed")
		}
	}

	err = filep.Close()
	fatalOnError(err, "filep.Close failed")
}
