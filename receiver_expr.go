package goastutil

import (
	"fmt"
	"go/ast"
)

type ReceiverExpr struct {
	*ast.FieldList
}

func NewReceiverExpr(recv *ast.FieldList) *ReceiverExpr {
	return &ReceiverExpr{FieldList: recv}
}

func (r *ReceiverExpr) Name() *Ident {
	return &Ident{Ident: r.List[0].Names[0]}
}

func (r *ReceiverExpr) Type() Expr {
	return NewExpr(r.List[0].Type)
}

// String returns a string representation of the ReceiverExpr.
//
// It returns a formatted string that includes the name and type of the ReceiverExpr.
func (r *ReceiverExpr) String() string {
	return fmt.Sprintf("(%s %s)", r.Name(), r.Type())
}
