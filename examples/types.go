package main

import (
	"github.com/padraicbc/gojsp/parser"
)

type IdentifierName struct {
	*SourceInfo
	Identifier Token
	prev, next VNode
}

var _ VNode = (*IdentifierName)(nil)

func (i *IdentifierName) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *IdentifierName) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *IdentifierName) Code() string {
	return CodeDef(i)
}

func (i *IdentifierName) Type() string {
	return "IdentifierName"
}

func (i *IdentifierName) Children() []VNode {

	return []VNode{i.Identifier}
}

// identifierName
//     : identifier
//     | reservedWord
//     ;

func (v *Visitor) VisitIdentifierName(ctx *parser.IdentifierNameContext) interface{} {
	// log.Println("VisitIdentifierName", ctx.GetText())
	//

	return v.VisitChildren(ctx)
	// Maybe just return &IdentifierName ctx.Identifier()...

}

func (v *Visitor) VisitKeyword(ctx *parser.KeywordContext) interface{} {
	return v.VisitChildren(ctx)

}

// reservedWord
//     : keyword
//     | NullLiteral
//     | BooleanLiteral
//     ;
func (v *Visitor) VisitReservedWord(ctx *parser.ReservedWordContext) interface{} {
	// log.Println("VisitReservedWord", ctx.Keyword().GetText())

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEos(ctx *parser.EosContext) interface{} {
	if ctx.GetChildCount() == 0 || ctx.EOF() != nil {
		return nil
	}
	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	// log.Println("VisitIdentifierExpression", ctx.GetText())
	return v.VisitChildren(ctx)
}

// identifier
//     : Identifier
//     | Async
//     ;
func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	// log.Println("VisitIdentifier", ctx.GetText(), ctx.GetChildCount())
	// VisitChildren would return the same inside a list but we don't need it
	return v.VisitChildren(ctx)
}
