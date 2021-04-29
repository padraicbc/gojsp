package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// arrayLiteral
func (v *Visitor) VisitArrayLiteralExpression(ctx *base.ArrayLiteralExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitArrayLiteralExpression", ctx.GetText())
	}

	return v.VisitArrayLiteral(ctx.ArrayLiteral().(*base.ArrayLiteralContext))
}

// arrayLiteral
//     : ('[' elementList ']')
//     ;
type ArrayLiteral struct {
	*SourceInfo
	OpenBracket  Token
	Elems        *ElementList
	CloseBracket Token
	firstChild   VNode

	next, prev VNode
}

var _ VNode = (*ArrayLiteral)(nil)

func (i *ArrayLiteral) Next() VNode {

	return i.next
}
func (i *ArrayLiteral) SetNext(v VNode) {
	i.next = v
}
func (i *ArrayLiteral) Prev() VNode {

	return i.prev
}
func (i *ArrayLiteral) SetPrev(v VNode) {
	i.prev = v
}
func (i *ArrayLiteral) Type() string {
	return "ArrayLiteral"
}
func (i *ArrayLiteral) Code() string {
	return CodeDef(i)
}

func (i *ArrayLiteral) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitArrayLiteral(ctx *base.ArrayLiteralContext) interface{} {
	if v.Debug {
		log.Println("VisitArrayLiteral", ctx.GetText())
	}
	arl := &ArrayLiteral{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	arl.OpenBracket = ident(v, ctx.OpenBracket().GetSymbol())
	arl.firstChild = arl.OpenBracket
	arl.Elems = v.VisitElementList(
		ctx.ElementList().(*base.ElementListContext)).(*ElementList)
	arl.CloseBracket = ident(v, ctx.CloseBracket().GetSymbol())
	setAllSibs(arl.OpenBracket, arl.Elems, arl.CloseBracket)
	return arl
}

// elementList
//     : ','* arrayElement? (','+ arrayElement)* ','* // Yes, everything is optional
//     ;
type ElementList struct {
	*SourceInfo
	ArrayElements []*ArrayElement
	Commas        []Token
	firstChild    VNode

	prev, next VNode
}

var _ VNode = (*ElementList)(nil)

func (i *ElementList) Next() VNode {

	return i.next
}
func (i *ElementList) SetNext(v VNode) {
	i.next = v
}
func (i *ElementList) Prev() VNode {

	return i.prev
}
func (i *ElementList) SetPrev(v VNode) {
	i.prev = v
}
func (i *ElementList) Type() string {
	return "ElementList"
}
func (i *ElementList) Code() string {
	return CodeDef(i)
}

func (i *ElementList) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitElementList(ctx *base.ElementListContext) interface{} {
	if v.Debug {
		log.Println("VisitElementList", ctx.GetText())
	}
	el := &ElementList{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if el.firstChild == nil {
			el.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {
		case "ArrayElement":
			el.ArrayElements = append(el.ArrayElements, ch.(*ArrayElement))
		case "LToken":
			tk := ch.(Token)
			if tk.SymbolName() != "Comma" {
				log.Panicf("%+v %s\n", tk, tk.SymbolName())
			}
			el.Commas = append(el.Commas, ch.(Token))
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}

	}
	return el
}

// arrayElement
//     : Ellipsis? singleExpression
//     ;
type ArrayElement struct {
	*SourceInfo
	Ellipsis         Token
	SingleExpression VNode
	firstChild       VNode

	prev, next VNode
}

var _ VNode = (*ArrayElement)(nil)

func (i *ArrayElement) Next() VNode {

	return i.next
}
func (i *ArrayElement) SetNext(v VNode) {
	i.next = v
}
func (i *ArrayElement) Prev() VNode {

	return i.prev
}
func (i *ArrayElement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ArrayElement) Type() string {
	return "ArrayElement"
}
func (i *ArrayElement) Code() string {
	return CodeDef(i)
}

func (i *ArrayElement) FirstChild() VNode {
	return i.firstChild

}
func (v *Visitor) VisitArrayElement(ctx *base.ArrayElementContext) interface{} {
	if v.Debug {
		log.Println("VisitArrayElement", ctx.GetText())
	}
	ae := &ArrayElement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	if ctx.Ellipsis() != nil {

		ae.Ellipsis = ident(v, ctx.Ellipsis().GetSymbol())
		ae.firstChild = ae.Ellipsis
	}
	ae.SingleExpression = v.Visit(ctx.SingleExpression()).(VNode)
	if ae.firstChild == nil {
		ae.firstChild = ae.SingleExpression
	}
	setAllSibs(ae.Ellipsis, ae.SingleExpression)

	return ae
}
