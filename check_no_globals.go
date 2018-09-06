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

func isWhitelisted(i *ast.Ident) bool {
	return i.Name == "_" || looksLikeError(i)
}

// looksLikeError checks if the AST identifier starts with 'err' or 'Err'.
//
// See the following issue on how to further improve this func:
// https://github.com/leighmcculloch/gochecknoglobals/issues/5
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
			line := fset.Position(genDecl.TokPos).Line
			valueSpec := genDecl.Specs[0].(*ast.ValueSpec)
			for i := 0; i < len(valueSpec.Names); i++ {
				vn := valueSpec.Names[i]
				if isWhitelisted(vn) {
					continue
				}
				message := fmt.Sprintf("%s:%d %s is a global variable", filename, line, vn.Name)
				messages = append(messages, message)
			}
		}
		return nil
	})

	return messages, err
}
