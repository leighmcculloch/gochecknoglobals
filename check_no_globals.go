package gochecknoglobals

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"github.com/kyoh86/nolint"
	"golang.org/x/tools/go/analysis"
)

// allowedExpression is a struct representing packages and methods that will
// be an allowed combination to use as a global variable, f.ex. Name `regexp`
// and SelName `MustCompile`.
type allowedExpression struct {
	Name    string
	SelName string
}

// reports is the reports of found global variables that's not valid and will be
// reported to the analysis pass.
type report struct {
	name string
	pos  token.Pos
}

// Analyzer is the analasys analyzer for gochecknoglobals. Ironically enough,
// this is in fact a global variable.
var Analyzer = &analysis.Analyzer{ //nolint
	Name:             "gochecknoglobals",
	Doc:              "Don't allow global variables",
	Run:              run,
	Flags:            flags(),
	Requires:         []*analysis.Analyzer{nolint.Analyzer},
	RunDespiteErrors: true,
}

func flags() flag.FlagSet {
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", false, "Include tests")

	return *flags
}

func run(pass *analysis.Pass) (interface{}, error) {
	noLinter := pass.ResultOf[nolint.Analyzer].(*nolint.NoLinter)
	runTests, _ := strconv.ParseBool(pass.Analyzer.Flags.Lookup("t").Value.String())

	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Pos()).Filename
		if !runTests && strings.HasSuffix(filename, "_test.go") {
			continue
		}

		for _, decl := range file.Decls {
			if noLinter.IgnoreNode(decl, "global") {
				continue
			}

			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			if genDecl.Tok != token.VAR {
				continue
			}

			for _, issue := range checkNoGlobals(genDecl) {
				pass.Report(analysis.Diagnostic{
					Pos:      issue.pos,
					Category: "global",
					Message:  fmt.Sprintf("%s is a global variable", issue.name),
				})
			}
		}
	}

	return nil, nil
}

func isAllowed(v ast.Node) bool {
	switch i := v.(type) {
	case *ast.Ident:
		return i.Name == "_" || i.Name == "version" || looksLikeError(i)
	case *ast.CallExpr:
		if expr, ok := i.Fun.(*ast.SelectorExpr); ok {
			return isAllowedSelectorExpression(expr)
		}
	case *ast.CompositeLit:
		if expr, ok := i.Type.(*ast.SelectorExpr); ok {
			return isAllowedSelectorExpression(expr)
		}
	}

	return false
}

func isAllowedSelectorExpression(v *ast.SelectorExpr) bool {
	x, ok := v.X.(*ast.Ident)
	if !ok {
		return false
	}

	allowList := []allowedExpression{
		{Name: "regexp", SelName: "MustCompile"},
	}

	for _, i := range allowList {
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

func checkNoGlobals(genDecl *ast.GenDecl) []report {
	var globalVariables []report

	for _, spec := range genDecl.Specs {
		valueSpec := spec.(*ast.ValueSpec)
		onlyAllowedValues := false

		for _, vn := range valueSpec.Values {
			if isAllowed(vn) {
				onlyAllowedValues = true
				continue
			}

			onlyAllowedValues = false
			break
		}

		if onlyAllowedValues {
			continue
		}

		for _, vn := range valueSpec.Names {
			if isAllowed(vn) {
				continue
			}

			globalVariables = append(globalVariables, report{
				name: vn.Name,
				pos:  vn.Pos(),
			})
		}
	}

	return globalVariables
}
