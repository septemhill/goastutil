package astutil

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/samber/lo"
)

type ReturnStmt struct {
	*ast.ReturnStmt
}

func (r ReturnStmt) Results() []Expr {
	return lo.Map(r.ReturnStmt.Results, func(x ast.Expr, _ int) Expr {
		return Expr{Expr: x}
	})
}

func (r ReturnStmt) String() string {
	exprs := lo.Map(r.Results(), func(x Expr, _ int) string {
		return x.String()
	})
	return fmt.Sprintf("return %s", strings.Join(exprs, ", "))
}
