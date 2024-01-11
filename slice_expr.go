package astutil

import (
	"fmt"
	"go/ast"
)

type SliceExpr struct {
	*ast.SliceExpr
}

func NewSliceExpr(x *ast.SliceExpr) *SliceExpr {
	return &SliceExpr{SliceExpr: x}
}

func (s *SliceExpr) Low() Expr {
	return Expr{Expr: s.SliceExpr.Low}
}

func (s *SliceExpr) High() Expr {
	return Expr{Expr: s.SliceExpr.High}
}

func (s *SliceExpr) Expr() Expr {
	return Expr{Expr: s.SliceExpr.X}
}

// String returns a string representation of the SliceExpr.
//
// String concatenates the expression, low, and high values of the SliceExpr
// into a formatted string. If the Slice3 field is false, it uses the
// expression, low, and high values to construct the string. If the Slice3
// field is true, it returns the string "todo_slice_expr3_expr".
//
// Returns:
//
//	string: The string representation of the SliceExpr.
func (s *SliceExpr) String() string {
	if !s.Slice3 {
		return fmt.Sprintf("%s[%s:%s]", s.Expr(), s.Low(), s.High())
	}
	return "todo_slice_expr3_expr"
}
