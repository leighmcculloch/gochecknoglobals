package checknoglobals

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/astutil"
)

// allowedExpression is a struct representing packages and methods that will
// be an allowed combination to use as a global variable, f.ex. Name `regexp`
// and SelName `MustCompile`.
type allowedExpression struct {
	Name    string
	SelName string
}

const Doc = `check that global variables are not mutated

This analyzer checks for mutations of global variables and errors on any found.

A global variable is a variable declared in package scope and that can be read
and written to by any function within the package. Global variables can cause
side effects which are difficult to keep track of. A code in one function may
change the variables state while another unrelated chunk of code may be
affected by it.

This analyzer allows global variables but disallows mutation of them, 
encouraging them to be used like constants.`

// Analyzer provides an Analyzer that checks that global variables
// are not mutated. Global variables themselves are allowed, but their values 
// should not be changed after initialization.
func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:             "gochecknoglobals",
		Doc:              Doc,
		Run:              checkNoGlobals,
		RunDespiteErrors: true,
		Requires:         []*analysis.Analyzer{inspect.Analyzer},
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
	// Map to store global variables
	globals := make(map[string]token.Pos)

	// First pass: collect all global variables
	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Pos()).Filename
		if !strings.HasSuffix(filename, ".go") {
			continue
		}

		// Find all global variables
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			if genDecl.Tok != token.VAR {
				continue
			}
			
			for _, spec := range genDecl.Specs {
				valueSpec := spec.(*ast.ValueSpec)
				
				for _, vn := range valueSpec.Names {
					if vn.Name == "_" {
						continue // Skip blank identifier
					}
					
					// Store the global variable
					globals[vn.Name] = vn.Pos()
				}
			}
		}
	}

	// Second pass: check for mutations of global variables
	for _, file := range pass.Files {
		filename := pass.Fset.Position(file.Pos()).Filename
		if !strings.HasSuffix(filename, ".go") {
			continue
		}

		// Visit all nodes in the file
		ast.Inspect(file, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.AssignStmt:
				// Skip declarations (first assignment)
				if n.Tok == token.DEFINE {
					return true
				}
				
				// Check if this is a function or method
				if isFunctionScope(pass, file, n) {
					// Check if this is an assignment to a global variable
					for _, lhs := range n.Lhs {
						ident, ok := lhs.(*ast.Ident)
						if !ok {
							continue
						}
						
						// Check if it's a global variable
						if _, exists := globals[ident.Name]; exists {
							// Check if this identifier refers to the global variable
							if obj := pass.TypesInfo.Uses[ident]; obj != nil {
								if _, ok := obj.(*types.Var); ok && obj.Pos() == globals[ident.Name] {
									message := fmt.Sprintf("global variable %s is being mutated", ident.Name)
									pass.Report(analysis.Diagnostic{
										Pos:      ident.Pos(),
										Category: "global-mutation",
										Message:  message,
									})
								}
							}
						}
					}
				}
			case *ast.IncDecStmt:
				// Check if this is a function or method
				if isFunctionScope(pass, file, n) {
					// Check if this is incrementing or decrementing a global variable
					if ident, ok := n.X.(*ast.Ident); ok {
						// Check if it's a global variable
						if _, exists := globals[ident.Name]; exists {
							// Check if this identifier refers to the global variable
							if obj := pass.TypesInfo.Uses[ident]; obj != nil {
								if _, ok := obj.(*types.Var); ok && obj.Pos() == globals[ident.Name] {
									message := fmt.Sprintf("global variable %s is being mutated", ident.Name)
									pass.Report(analysis.Diagnostic{
										Pos:      ident.Pos(),
										Category: "global-mutation",
										Message:  message,
									})
								}
							}
						}
					}
				}
			}
			return true
		})
	}

	return nil, nil
}

// isFunctionScope checks if a node is within a function body
func isFunctionScope(pass *analysis.Pass, file *ast.File, node ast.Node) bool {
	// Find the enclosing function for this node
	path, _ := astutil.PathEnclosingInterval(file, node.Pos(), node.End())
	
	for _, n := range path {
		switch n.(type) {
		case *ast.FuncDecl, *ast.FuncLit:
			return true
		}
	}
	return false
}
