package vast

import "github.com/padraicbc/gojsp/base"

// arrayLiteral
//     : ('[' elementList ']')
//     ;
type ArrayLiteral struct {
	*SourceInfo
	OpenBracket  Token
	ElementList  *ElementList
	CloseBracket Token
	children     VNode
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
	return children(i.children)
}
func (v *Visitor) VisitArrayLiteral(ctx *base.ArrayLiteralContext) interface{} {
	arl := &ArrayLiteral{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		switch ch.Type() {
		case "OpenBracket":
			arl.OpenBracket = ch.(Token)
		case "CloseBracket":
			arl.CloseBracket = ch.(Token)
		case "ElementList":
			arl.ElementList = ch.(*ElementList)

		}
		if arl.children == nil {
			arl.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
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
	children      VNode
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
	return children(i.children)
}

func (v *Visitor) VisitElementList(ctx *base.ElementListContext) interface{} {
	el := &ElementList{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if el.children == nil {
			el.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
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
	children         VNode
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
	return children(i.children)
}
func (v *Visitor) VisitArrayElement(ctx *base.ArrayElementContext) interface{} {
	ae := ArrayElement{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		switch ch.Type() {
		case "Ellipsis":
			ae.Ellipsis = ch.(Token)
		case "SingleExpression":
			ae.SingleExpression = ch
		default:
			panic(ch.Type())
		}
		if ae.children == nil {
			ae.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
	}
	return ae
}
