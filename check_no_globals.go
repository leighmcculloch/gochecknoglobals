package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// allowedExpression is a struct representing packages and methods that will
// be an allowed combination to use as a global variable, f.ex. Name `regexp`
// and SelName `MustCompile`.
type allowedExpression struct {
	Name    string
	SelName string
	IsErr   bool
}

func isAllowed(ident *ast.Ident, v ast.Node) bool {
	switch i := v.(type) {
	case *ast.CallExpr:
		if expr, ok := i.Fun.(*ast.SelectorExpr); ok {
			return isAllowedSelectorExpression(ident, expr)
		}
	case *ast.CompositeLit:
		if expr, ok := i.Type.(*ast.SelectorExpr); ok {
			return isAllowedSelectorExpression(ident, expr)
		}
	}

	return ident.Name == "version" || ident.Name == "_"
}

func isAllowedSelectorExpression(ident *ast.Ident, v *ast.SelectorExpr) bool {
	x, ok := v.X.(*ast.Ident)
	if !ok {
		return false
	}

	allowList := []allowedExpression{
		{Name: "regexp", SelName: "MustCompile", IsErr: false},
		{Name: "fmt", SelName: "Errorf", IsErr: true},
		{Name: "errors", SelName: "New", IsErr: true},
	}

	for _, i := range allowList {
		if x.Name == i.Name && v.Sel.Name == i.SelName {
			if i.IsErr {
				return looksLikeError(ident)
			}

			return true
		}
	}

	return false
}

// looksLikeError returns true if the AST identifier starts
// with 'err' or 'Err', or false otherwise.
//
// TODO: https://github.com/leighmcculloch/gochecknoglobals/issues/5
func looksLikeError(i *ast.Ident) bool {
	prefix := "err"
	if i.IsExported() {
		prefix = "Err"
	}
	return strings.HasPrefix(i.Name, prefix)
}

func checkNoGlobals(rootPath string, includeTests bool) ([]string, error) {
	const recursiveSuffix = string(filepath.Separator) + "..."
	recursive := false
	if strings.HasSuffix(rootPath, recursiveSuffix) {
		recursive = true
		rootPath = rootPath[:len(rootPath)-len(recursiveSuffix)]
	}

	messages := []string{}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if !recursive && path != rootPath {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if !includeTests && strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}

		addError := func(filename string, line int, variable string) {
			message := fmt.Sprintf("%s:%d %s is a global variable", filename, line, variable)
			messages = append(messages, message)
		}

		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			if genDecl.Tok != token.VAR {
				continue
			}

			filename := fset.Position(genDecl.TokPos).Filename
			for _, spec := range genDecl.Specs {
				valueSpec := spec.(*ast.ValueSpec)

				if len(valueSpec.Values) != len(valueSpec.Names) {
					addError(filename, fset.Position(valueSpec.Names[0].Pos()).Line, valueSpec.Names[0].Name)
					continue
				}

				for i := 0; i < len(valueSpec.Names); i++ {
					if !isAllowed(valueSpec.Names[i], valueSpec.Values[i]) {
						addError(filename, fset.Position(valueSpec.Names[i].Pos()).Line, valueSpec.Names[i].Name)
					}
				}
			}
		}
		return nil
	})

	return messages, err
}
