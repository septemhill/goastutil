package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type CallExpr struct {
	*ast.CallExpr
}

func NewCallExpr(c *ast.CallExpr) *CallExpr {
	return &CallExpr{CallExpr: c}
}

func (c *CallExpr) Func() Expr {
	return Expr{Expr: c.CallExpr.Fun}
}

func (c *CallExpr) Args() []Expr {
	return lo.Map(c.CallExpr.Args, func(x ast.Expr, _ int) Expr {
		return Expr{Expr: x}
	})
}

// String returns a string representation of the CallExpr.
//
// It concatenates the name of the function and its arguments into a single string.
// It returns the resulting string.
func (c *CallExpr) String() string {
	args := lo.Map(c.Args(), func(x Expr, _ int) string {
		return x.String()
	})

	if c.Ellipsis.IsValid() {
		return fmt.Sprintf("%s(%s...)", c.Func(), strings.Join(args, ", "))
	}
	return fmt.Sprintf("%s(%s)", c.Func(), strings.Join(args, ", "))
}
