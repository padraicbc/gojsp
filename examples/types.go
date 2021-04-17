package main

import (
	"github.com/padraicbc/gojsp/parser"
)

type IdentifierName struct {
	*SourceInfo
	Identifier Token
}

var _ VNode = (*IdentifierName)(nil)

func (i *IdentifierName) Code() string {
	return CodeDef(i)
}

func (i *IdentifierName) Type() string {
	return "IdentifierName"
}

func (i *IdentifierName) GetChildren() []VNode {

	return []VNode{i.Identifier}
}

// identifierName
//     : identifier
//     | reservedWord
//     ;

func (v *visitor) VisitIdentifierName(ctx *parser.IdentifierNameContext) interface{} {
	// log.Println("VisitIdentifierName", ctx.GetText())
	//

	return v.VisitChildren(ctx)
	// Maybe just return &IdentifierName ctx.Identifier()...

}

func (v *visitor) VisitKeyword(ctx *parser.KeywordContext) interface{} {
	return v.VisitChildren(ctx)

}

// reservedWord
//     : keyword
//     | NullLiteral
//     | BooleanLiteral
//     ;
func (v *visitor) VisitReservedWord(ctx *parser.ReservedWordContext) interface{} {
	// log.Println("VisitReservedWord", ctx.Keyword().GetText())

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitEos(ctx *parser.EosContext) interface{} {
	if ctx.GetChildCount() == 0 || ctx.EOF() != nil {
		return nil
	}
	return v.VisitChildren(ctx)
}
func (v *visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	// log.Println("VisitIdentifierExpression", ctx.GetText())
	return v.VisitChildren(ctx)
}

// identifier
//     : Identifier
//     | Async
//     ;
func (v *visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	// log.Println("VisitIdentifier", ctx.GetText(), ctx.GetChildCount())
	// VisitChildren would return the same inside a list but we don't need it
	return v.VisitChildren(ctx)
}
