// This script generates datamodel.go
package main

import (
	"flag"
	"go/format"
	"go/parser"
	"go/token"
	"time"

	"github.com/bassosimone/apiclient/internal/fatalx"
	"github.com/bassosimone/apiclient/internal/fmtx"
	"github.com/bassosimone/apiclient/internal/osx"
)

func main() {
	outfile := flag.String("outfile", "datamodel.go", "Output file")
	pkgdir := flag.String("pkg", "./internal/datamodel", "Package directory path")
	flag.Parse()

	filep := osx.MustCreate(*outfile)
	defer filep.Close()

	filep.WriteString("// Code generated by go generate; DO NOT EDIT.\n")
	fmtx.Fprintf(filep, "// %+v\n\n", time.Now())
	filep.WriteString("package apiclient\n\n")

	fmtx.Fprint(filep, "//go:generate go run ./internal/gendatamodel/...\n\n")

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, *pkgdir, nil, parser.ParseComments)
	fatalx.OnError(err, "parser.ParseDir failed")

	model := pkgs["datamodel"]
	fatalx.IfNil(model, "cannot find the datamodel package")

	for _, fdata := range model.Files {
		for _, decl := range fdata.Decls {
			err = format.Node(filep, fset, decl)
			fatalx.OnError(err, "format.Node failed")
			filep.WriteString("\n\n")
		}
	}
}
