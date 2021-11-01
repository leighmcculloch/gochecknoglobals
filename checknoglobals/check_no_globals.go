package checknoglobals

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var (
	errInvalidName = errors.New("invalid name of var")
	errInvalidType = errors.New("var can be const")
)

// allowedExpression is a struct representing packages and methods that will
// be an allowed combination to use as a global variable, f.ex. Name `regexp`
// and SelName `MustCompile`.
type allowedExpression struct {
	Name    string
	SelName string
}

const Doc = `check that no global variables exist

This analyzer checks for global variables and errors on any found.

A global variable is a variable declared in package scope and that can be read
and written to by any function within the package. Global variables can cause
side effects which are difficult to keep track of. A code in one function may
change the variables state while another unrelated chunk of code may be
effected by it.`

// Analyzer provides an Analyzer that checks that there are no global
// variables, except for errors and variables containing regular
// expressions.
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:             "gochecknoglobals",
		Doc:              Doc,
		Run:              checkNoGlobals,
		Flags:            flags(),
		RunDespiteErrors: true,
	}
}

func flags() flag.FlagSet {
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", false, "Include tests")

	return *flags
}

func isAllowed(cm ast.CommentMap, name ast.Node, value ast.Node) error {
	i, ok := name.(*ast.Ident)
	if !ok {
		return errInvalidName
	}

	// Allow specific expressions or things that we cannot guarantee doesn't
	// implement the error interface. This means we don't allow basic literals
	// since they can be converted to a const.
	switch v := value.(type) {
	case *ast.CallExpr:
		if expr, ok := v.Fun.(*ast.SelectorExpr); ok {
			if isAllowedSelectorExpression(expr) {
				return nil
			}
		}
	case *ast.CompositeLit:
		if expr, ok := v.Type.(*ast.SelectorExpr); ok {
			if isAllowedSelectorExpression(expr) {
				return nil
			}
		}
	case *ast.BasicLit:
		if i.Name != "_" && i.Name != "version" {
			return errInvalidType
		}
	}

	// Allow embed vars.
	if identHasEmbedComment(cm, i) {
		return nil
	}

	// Independent of type, the name of the variable must be allowed.
	if i.Name != "_" && i.Name != "version" && !looksLikeError(i) {
		return errInvalidName
	}

	return nil
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

func identHasEmbedComment(cm ast.CommentMap, i *ast.Ident) bool {
	if i.Obj == nil {
		return false
	}

	spec, ok := i.Obj.Decl.(*ast.ValueSpec)
	if !ok {
		return false
	}

	return hasEmbedComment(cm, spec)
}

// hasEmbedComment returns true if the AST node has
// a '//go:embed ' comment, or false otherwise.
func hasEmbedComment(cm ast.CommentMap, n ast.Node) bool {
	for _, g := range cm[n] {
		for _, c := range g.List {
			if strings.HasPrefix(c.Text, "//go:embed ") {
				return true
			}
		}
	}
	return false
}

func checkNoGlobals(pass *analysis.Pass) (interface{}, error) {
	includeTests := pass.Analyzer.Flags.Lookup("t").Value.(flag.Getter).Get().(bool)

	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Pos()).Filename
		if !strings.HasSuffix(filename, ".go") {
			continue
		}
		if !includeTests && strings.HasSuffix(filename, "_test.go") {
			continue
		}

		fileCommentMap := ast.NewCommentMap(pass.Fset, file, file.Comments)

		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			if genDecl.Tok != token.VAR {
				continue
			}

			if hasEmbedComment(fileCommentMap, genDecl) {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec := spec.(*ast.ValueSpec)

				for i := range valueSpec.Names {
					var (
						name  = valueSpec.Names[i]
						value ast.Node
					)

					if len(valueSpec.Values) > i {
						value = valueSpec.Values[i]
					}

					if err := isAllowed(fileCommentMap, name, value); err != nil {
						message := fmt.Sprintf("%s is a global variable", name.Name)
						if errors.Is(errInvalidType, err) {
							message += ", should be a const"
						}

						pass.Report(analysis.Diagnostic{
							Pos:      name.Pos(),
							Category: "global",
							Message:  message,
						})
					}
				}
			}
		}
	}

	return nil, nil
}
