package astutil

import (
	"fmt"
	"go/ast"
)

type IndexExpr struct {
	*ast.IndexExpr
}

func (i IndexExpr) Expr() Expr {
	return Expr{Expr: i.IndexExpr.X}
}

func (i IndexExpr) Index() Expr {
	return Expr{Expr: i.IndexExpr.Index}
}

// String returns a string representation of the IndexExpr.
//
// It returns a formatted string that represents the IndexExpr as `Expr[Index]`.
// The returned string can be used for debugging or displaying the IndexExpr.
//
// Returns:
// - a string representation of the IndexExpr.
func (i IndexExpr) String() string {
	return fmt.Sprintf("%s[%s]", i.Expr(), i.Index())
}
