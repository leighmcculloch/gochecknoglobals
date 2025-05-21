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
