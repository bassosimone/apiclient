// Command gendatamodel generates datamodel.go.
package main

import (
	"flag"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"sort"
	"strings"
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

	var decls []ast.Decl
	for _, fdata := range model.Files {
		for _, decl := range fdata.Decls {
			switch v := decl.(type) {
			case *ast.GenDecl:
				if len(v.Specs) == 1 {
					switch v.Specs[0].(type) {
					case *ast.TypeSpec:
						decls = append(decls, decl)
					}
				}
			}
		}
	}
	sort.SliceStable(decls, func(i, j int) bool {
		// we already exclude imports in the above loop
		left := decls[i].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name
		right := decls[j].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name
		return strings.Compare(left.String(), right.String()) <= 0
	})
	for _, decl := range decls {
		err = format.Node(filep, fset, decl)
		fatalx.OnError(err, "format.Node failed")
		filep.WriteString("\n\n")
	}
}
