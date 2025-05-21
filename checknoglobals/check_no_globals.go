package checknoglobals

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// allowedExpression is a struct representing packages and methods that will
// be an allowed combination to use as a global variable, f.ex. Name `regexp`
// and SelName `MustCompile`.
type allowedExpression struct {
	Name    string
	SelName string
}

const Doc = `detects direct mutations of global variables.

This analyzer identifies package-level variables (globals) and reports any direct
assignments to them within function scopes. The goal is to help developers
track and manage state changes that can have wide-ranging side effects.

A global variable is a variable declared in package scope. While their declaration
is not flagged, their mutation is, as this can lead to complex state management
and debugging challenges.

Mutation Detection Scope:
- Detected Mutations: The linter currently detects direct assignments to global
  variables. This includes:
  - Simple assignment: globalVar = newValue
  - Compound assignment: globalVar += someValue
  - Reassignment of global pointers, slices, maps: globalSlice = newSlice

- Undetected Mutations (Known Limitations): The linter currently does NOT detect
  indirect mutations, such as:
  - Modifying a field of a global struct: globalStruct.Field = value
  - Modifying an element of a global array, slice, or map: globalSlice[0] = value
  - Mutations via function/method calls: modifyGlobal(&globalStruct)

Reason/Future: Detecting these indirect mutations is more complex and may be
considered for future enhancements.`

// Analyzer provides an Analyzer that checks that there are no global
// variables, except for errors and variables containing regular
// expressions.
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:             "gochecknoglobals",
		Doc:              Doc,
		Run:              checkNoGlobals,
		RunDespiteErrors: true,
	}
}

func isAllowed(cm ast.CommentMap, v ast.Node, ti *types.Info) bool {
	switch i := v.(type) {
	case *ast.GenDecl:
		return hasEmbedComment(cm, i)
	case *ast.Ident:
		return i.Name == "_" || i.Name == "version" || isError(i, ti) || identHasEmbedComment(cm, i)
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

// isError reports whether the AST identifier looks like
// an error and implements the error interface.
func isError(i *ast.Ident, ti *types.Info) bool {
	return looksLikeError(i) && implementsError(i, ti)
}

// looksLikeError returns true if the AST identifier starts
// with 'err' or 'Err', or false otherwise.
func looksLikeError(i *ast.Ident) bool {
	prefix := "err"
	if i.IsExported() {
		prefix = "Err"
	}
	return strings.HasPrefix(i.Name, prefix)
}

// implementsError reports whether the AST identifier
// implements the error interface.
func implementsError(i *ast.Ident, ti *types.Info) bool {
	t := ti.TypeOf(i)
	et := types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
	return types.Implements(t, et)
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
	// Phase 1: Collect all package-level 'var' declarations.
	// The 'isAllowed' logic is intentionally not used here, as the goal is to
	// identify all globals first, and then separately check for mutations.
	globalVars := make(map[types.Object]bool)

	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Pos()).Filename
		if !strings.HasSuffix(filename, ".go") {
			continue
		}

		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			if genDecl.Tok != token.VAR {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec, okSpec := spec.(*ast.ValueSpec)
				if !okSpec {
					continue
				}

				for _, name := range valueSpec.Names {
					// Use pass.TypesInfo.Defs to get the object defined by the name.
					if obj := pass.TypesInfo.Defs[name]; obj != nil {
						globalVars[obj] = true
					}
				}
			}
		}
	}

	// Phase 2: Detect and report direct mutations of these global variables
	// within function scopes. The 'isAllowed' logic is not used to filter mutations.
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Body == nil {
				continue
			}

			ast.Inspect(fn.Body, func(n ast.Node) bool {
				assignStmt, isAssign := n.(*ast.AssignStmt)
				if !isAssign {
					return true // Continue traversal
				}

				for _, lhsExpr := range assignStmt.Lhs {
					ident, isIdent := lhsExpr.(*ast.Ident)
					if !isIdent {
						continue
					}

					obj := pass.TypesInfo.Uses[ident]
					if obj == nil {
						continue
					}

					if _, isGlobal := globalVars[obj]; isGlobal {
						pass.Report(analysis.Diagnostic{
							Pos:      ident.Pos(),
							Category: "mutation",
							Message:  fmt.Sprintf("mutation of global variable '%s'", ident.Name),
						})
					}
				}
				return true
			})
		}
	}

	return nil, nil
}
