package vast

import (
	"io"
	"log"

	"github.com/padraicbc/gojsp/base"
)

type IdentifierName struct {
	*SourceInfo
	Identifier Token
	prev, next VNode
}

var _ VNode = (*IdentifierName)(nil)

func (i *IdentifierName) Next() VNode {

	return i.next
}
func (i *IdentifierName) SetNext(v VNode) {
	i.next = v
}
func (i *IdentifierName) Prev() VNode {

	return i.prev
}
func (i *IdentifierName) SetPrev(v VNode) {
	i.prev = v
}
func (i *IdentifierName) Code() string {
	return CodeDef(i)
}

func (i *IdentifierName) Type() string {
	return "IdentifierName"
}

func (i *IdentifierName) FirstChild() VNode {

	return i.Identifier
}

// identifierName
//     : identifier
//     | reservedWord
//     ;

func (v *Visitor) VisitIdentifierName(ctx *base.IdentifierNameContext) interface{} {
	if v.Debug {
		log.Println("VisitIdentifierName", ctx.GetText())
	}

	return v.VisitChildren(ctx)
	// Maybe just return &IdentifierName ctx.Identifier()...

}

func (v *Visitor) VisitKeyword(ctx *base.KeywordContext) interface{} {
	if v.Debug {
		log.Println("VisitKeyword", ctx.GetText())
	}
	return v.VisitChildren(ctx)

}

// reservedWord
//     : keyword
//     | NullLiteral
//     | BooleanLiteral
//     ;
func (v *Visitor) VisitReservedWord(ctx *base.ReservedWordContext) interface{} {
	if v.Debug {
		log.Println("VisitReservedWord", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEos(ctx *base.EosContext) interface{} {
	if v.Debug {
		log.Println("VisitEos", ctx.GetText())
	}
	if ctx.GetChildCount() == 0 || ctx.EOF() != nil {
		return io.EOF
	}
	if v.Debug {
		log.Println("VisitEos", ctx.GetText(), v.VisitChildren(ctx).([]VNode)[0].Type())
	}
	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitIdentifierExpression(ctx *base.IdentifierExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitIdentifierExpression", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}

// identifier
//     : Identifier
//     | Async
//     ;
func (v *Visitor) VisitIdentifier(ctx *base.IdentifierContext) interface{} {
	// log.Println("VisitIdentifier", ctx.GetText(), ctx.GetChildCount())
	if v.Debug {
		log.Println("VisitIdentifier", ctx.GetText())
	}
	// VisitChildren would return the same inside a list but we don't need it
	return v.VisitChildren(ctx)
}
