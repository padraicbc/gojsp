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
}

var _ VNode = (*ArrayLiteral)(nil)

func (i *ArrayLiteral) Type() string {
	return "ArrayLiteral"
}
func (i *ArrayLiteral) Code() string {
	return CodeDef(i)
}

func (i *ArrayLiteral) GetChildren() []VNode {
	// todo: flatten
	return []VNode{
		i.OpenBracket,
		// i.ElementList,
		i.CloseBracket,
	}
}
func (v *visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	arl := &ArrayLiteral{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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
}

var _ VNode = (*ElementList)(nil)

func (i *ElementList) Type() string {
	return "ElementList"
}
func (i *ElementList) Code() string {
	return CodeDef(i)
}

func (i *ElementList) GetChildren() []VNode {
	// todo: flatteb
	return []VNode{
		//  i.ArrayElements
	}
}

func (v *visitor) VisitElementList(ctx *parser.ElementListContext) interface{} {
	el := &ElementList{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
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
}

var _ VNode = (*ArrayElement)(nil)

func (i *ArrayElement) Type() string {
	return "ArrayElement"
}
func (i *ArrayElement) Code() string {
	return CodeDef(i)
}

func (i *ArrayElement) GetChildren() []VNode {
	return []VNode{
		i.Ellipsis,
		i.SingleExpression,
	}
}
func (v *visitor) VisitArrayElement(ctx *parser.ArrayElementContext) interface{} {
	ae := ArrayElement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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
