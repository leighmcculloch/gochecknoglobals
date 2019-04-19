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

type whitelistedExpression struct {
	Name    string
	SelName string
}

func isWhitelisted(v ast.Node) bool {
	switch i := v.(type) {
	case *ast.Ident:
		return i.Name == "_" || i.Name == "version" || looksLikeError(i)
	case *ast.CallExpr:
		if expr, ok := i.Fun.(*ast.SelectorExpr); ok {
			return isWhitelistedSelectorExpression(expr)
		}
	case *ast.CompositeLit:
		if expr, ok := i.Type.(*ast.SelectorExpr); ok {
			return isWhitelistedSelectorExpression(expr)
		}
	}

	return false
}

func isWhitelistedSelectorExpression(v *ast.SelectorExpr) bool {
	x, ok := v.X.(*ast.Ident)
	if !ok {
		return false
	}

	whitelist := []whitelistedExpression{
		{
			Name:    "errors",
			SelName: "New",
		},
		{
			Name:    "fmt",
			SelName: "Errorf",
		},
		{
			Name:    "regexp",
			SelName: "MustCompile",
		},
	}

	for _, i := range whitelist {
		if x.Name == i.Name && v.Sel.Name == i.SelName {
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
				onlyWhitelistedValues := false

				for _, vn := range valueSpec.Values {
					if isWhitelisted(vn) {
						onlyWhitelistedValues = true
						continue
					}

					onlyWhitelistedValues = false
					break
				}

				if onlyWhitelistedValues {
					continue
				}

				for _, vn := range valueSpec.Names {
					if isWhitelisted(vn) {
						continue
					}
					line := fset.Position(vn.Pos()).Line
					message := fmt.Sprintf("%s:%d %s is a global variable", filename, line, vn.Name)
					messages = append(messages, message)
				}
			}
		}
		return nil
	})

	return messages, err
}
