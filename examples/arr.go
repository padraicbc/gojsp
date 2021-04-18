package main

import "github.com/padraicbc/gojsp/parser"

// arrayLiteral
//     : ('[' elementList ']')
//     ;
type ArrayLiteral struct {
	*SourceInfo
	OpenBracket  Token
	ElementList  *ElementList
	CloseBracket Token
	children     []VNode
	next, prev   VNode
}

var _ VNode = (*ArrayLiteral)(nil)

func (i *ArrayLiteral) Next(v VNode) VNode {

	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ArrayLiteral) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ArrayLiteral) Type() string {
	return "ArrayLiteral"
}
func (i *ArrayLiteral) Code() string {
	return CodeDef(i)
}

func (i *ArrayLiteral) Children() []VNode {
	// todo: flatten
	return i.children
}
func (v *Visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	arl := &ArrayLiteral{children: v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range arl.children {
		arl.children = append(arl.children, ch)
		switch ch.Type() {
		case "OpenBracket":
			arl.OpenBracket = ch.(Token)
		case "CloseBracket":
			arl.CloseBracket = ch.(Token)
		case "ElementList":
			arl.ElementList = ch.(*ElementList)

		}
	}
	return arl
}

// elementList
//     : ','* arrayElement? (','+ arrayElement)* ','* // Yes, everything is optional
//     ;
type ElementList struct {
	*SourceInfo
	ArrayElements []*ArrayElement
	Commas        []Token
	children      []VNode
	prev, next    VNode
}

var _ VNode = (*ElementList)(nil)

func (i *ElementList) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ElementList) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ElementList) Type() string {
	return "ElementList"
}
func (i *ElementList) Code() string {
	return CodeDef(i)
}

func (i *ElementList) Children() []VNode {
	// todo: flatteb
	return i.children
}

func (v *Visitor) VisitElementList(ctx *parser.ElementListContext) interface{} {
	el := &ElementList{children: v.VisitChildren(ctx).([]VNode), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		switch ch.Type() {
		case "ArrayElement":
			el.ArrayElements = append(el.ArrayElements, ch.(*ArrayElement))
		case "Comma":
			el.Commas = append(el.Commas, ch.(Token))
		default:
			panic(ch.Type())

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
	children         []VNode
	prev, next       VNode
}

var _ VNode = (*ArrayElement)(nil)

func (i *ArrayElement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ArrayElement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ArrayElement) Type() string {
	return "ArrayElement"
}
func (i *ArrayElement) Code() string {
	return CodeDef(i)
}

func (i *ArrayElement) Children() []VNode {
	return i.children
}
func (v *Visitor) VisitArrayElement(ctx *parser.ArrayElementContext) interface{} {
	ae := ArrayElement{children: v.VisitChildren(ctx).([]VNode), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range ae.children {

		switch ch.Type() {
		case "Ellipsis":
			ae.Ellipsis = ch.(Token)
		case "SingleExpression":
			ae.SingleExpression = ch
		default:
			panic(ch.Type())
		}
	}
	return ae
}
