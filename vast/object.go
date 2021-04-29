package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// objectLiteral
//     : '{' (propertyAssignment (',' propertyAssignment)*)? ','? '}'
//     ;
type ObjectLiteral struct {
	*SourceInfo
	OpenBrace Token
	// propertyAssignment rule
	PropAssignments []VNode
	Commas          []Token
	CloseBrace      Token
	firstChild      VNode

	next, prev VNode
}

var _ VNode = (*ObjectLiteral)(nil)

func (i *ObjectLiteral) Next() VNode {

	return i.next
}
func (i *ObjectLiteral) SetNext(v VNode) {
	i.next = v
}
func (i *ObjectLiteral) Prev() VNode {

	return i.prev
}
func (i *ObjectLiteral) SetPrev(v VNode) {
	i.prev = v
}
func (i *ObjectLiteral) Type() string {
	return "ObjectLiteral"
}
func (i *ObjectLiteral) Code() string {
	return CodeDef(i)
}

func (i *ObjectLiteral) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitObjectLiteral(ctx *base.ObjectLiteralContext) interface{} {
	if v.Debug {
		log.Println("VisitObjectLiteral", ctx.GetText())
	}
	ol := &ObjectLiteral{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ol.firstChild == nil {
			ol.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {
		case "LToken":
			tk := ch.(Token)
			switch tk.SymbolName() {
			case "OpenBrace":
				ol.OpenBrace = tk
			case "CloseBrace":
				ol.OpenBrace = tk
			case "Comma":
				ol.Commas = append(ol.Commas, tk)

			default:
				log.Panicf("%+v %s\n", tk, tk.SymbolName())
			}
		case "PropertyShorthand", "PropertyExpressionAssignment",
			"ComputedPropertyExpressionAssignment",
			"FunctionProperty", "PropertyGetter", "PropertySetter":
			ol.PropAssignments = append(ol.PropAssignments, ch)

		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}

	return ol
}

func (v *Visitor) VisitObjectLiteralExpression(ctx *base.ObjectLiteralExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitObjectLiteralExpression", ctx.GetText())
	}
	return v.VisitObjectLiteral(
		ctx.ObjectLiteral().(*base.ObjectLiteralContext)).(*ObjectLiteral)
}

// Async? '*'? propertyName '(' formalParameterList?  ')'  functionBody
type FunctionProperty struct {
	*SourceInfo
	Async               Token
	Multiply            Token
	PropertyName        *PropertyName
	OpenParen           Token
	FormalParameterList *FormalParameterList
	CloseParen          Token
	FunctionBody        *FunctionBody

	firstChild VNode
	prev, next VNode
}

var _ VNode = (*FunctionProperty)(nil)

func (i *FunctionProperty) Type() string {
	return "FunctionProperty"
}
func (i *FunctionProperty) Code() string {
	return CodeDef(i)
}
func (i *FunctionProperty) Next() VNode {
	return i.next
}
func (i *FunctionProperty) SetNext(v VNode) {
	i.next = v
}
func (i *FunctionProperty) Prev() VNode {
	return i.prev
}
func (i *FunctionProperty) SetPrev(v VNode) {
	i.prev = v
}
func (i *FunctionProperty) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitFunctionProperty(ctx *base.FunctionPropertyContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionProperty", ctx.GetText())
	}
	fp := &FunctionProperty{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fp.firstChild == nil {
			fp.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {
		case "LToken":
			tk := ch.(Token)
			switch tk.SymbolName() {
			case "Async":
				fp.Async = tk
			case "Multiply":
				fp.Multiply = tk
			case "OpenParen":
				fp.OpenParen = tk
			case "CloseParen":
				fp.CloseParen = tk

			default:
				log.Panicf("%+v\n", ch)
			}
		case "PropertyName":
			fp.PropertyName = ch.(*PropertyName)
		case "FunctionProperty":
			fp.FunctionBody = ch.(*FunctionBody)
		case "FormalParameterList":
			fp.FormalParameterList = ch.(*FormalParameterList)
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}

	}
	return fp
}

// getter '(' ')' functionBody
type PropertyGetter struct {
	*SourceInfo
	Getter       *Getter
	OpenParen    Token
	CloseParen   Token
	FunctionBody *FunctionBody
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*PropertyGetter)(nil)

func (i *PropertyGetter) Type() string {
	return "PropertyGetter"
}
func (i *PropertyGetter) Code() string {
	return CodeDef(i)
}
func (i *PropertyGetter) Next() VNode {
	return i.next
}
func (i *PropertyGetter) SetNext(v VNode) {
	i.next = v
}
func (i *PropertyGetter) Prev() VNode {
	return i.prev
}
func (i *PropertyGetter) SetPrev(v VNode) {
	i.prev = v
}
func (i *PropertyGetter) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPropertyGetter(ctx *base.PropertyGetterContext) interface{} {

	if v.Debug {
		log.Println("VisitPropertyGetter", ctx.GetText())
	}
	pg := &PropertyGetter{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	pg.Getter = v.VisitGetter(ctx.Getter().(*base.GetterContext)).(*Getter)
	pg.firstChild = pg.Getter
	pg.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	pg.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	pg.FunctionBody = v.VisitFunctionBody(
		ctx.FunctionBody().(*base.FunctionBodyContext)).(*FunctionBody)
	setAllSibs(pg.Getter, pg.OpenParen, pg.CloseParen, pg.FunctionBody)

	return pg
}

// setter '(' formalParameterArg ')' functionBody
type PropertySetter struct {
	*SourceInfo
	Setter       *Setter
	OpenParen    Token
	Arg          *FormalParameterArg
	CloseParen   Token
	FunctionBody *FunctionBody
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*PropertySetter)(nil)

func (i *PropertySetter) Type() string {
	return "PropertySetter"
}
func (i *PropertySetter) Code() string {
	return CodeDef(i)
}
func (i *PropertySetter) Next() VNode {
	return i.next
}
func (i *PropertySetter) SetNext(v VNode) {
	i.next = v
}
func (i *PropertySetter) Prev() VNode {
	return i.prev
}
func (i *PropertySetter) SetPrev(v VNode) {
	i.prev = v
}
func (i *PropertySetter) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPropertySetter(ctx *base.PropertySetterContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertySetter", ctx.GetText())
	}
	vp := &PropertySetter{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	vp.Setter = v.VisitSetter(ctx.Setter().(*base.SetterContext)).(*Setter)
	vp.firstChild = vp.Setter

	vp.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	vp.Arg = v.VisitFormalParameterArg(
		ctx.FormalParameterArg().(*base.FormalParameterArgContext)).(*FormalParameterArg)
	vp.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	vp.FunctionBody = v.VisitFunctionBody(
		ctx.FunctionBody().(*base.FunctionBodyContext)).(*FunctionBody)
	setAllSibs(vp.Setter, vp.OpenParen, vp.Arg, vp.CloseParen, vp.FunctionBody)

	return vp
}

// Ellipsis? singleExpression
type PropertyShorthand struct {
	*SourceInfo
	Ellipsis   Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*PropertyShorthand)(nil)

func (i *PropertyShorthand) Type() string {
	return "PropertyShorthand"
}
func (i *PropertyShorthand) Code() string {
	return CodeDef(i)
}
func (i *PropertyShorthand) Next() VNode {
	return i.next
}
func (i *PropertyShorthand) SetNext(v VNode) {
	i.next = v
}
func (i *PropertyShorthand) Prev() VNode {
	return i.prev
}
func (i *PropertyShorthand) SetPrev(v VNode) {
	i.prev = v
}
func (i *PropertyShorthand) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPropertyShorthand(ctx *base.PropertyShorthandContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertyShorthand", ctx.GetText())
	}
	ps := &PropertyShorthand{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	if ctx.Ellipsis() != nil {
		ps.Ellipsis = ident(v, ctx.Ellipsis().GetSymbol())
		ps.firstChild = ps.Ellipsis
	}
	ps.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
	if ps.firstChild != nil {
		ps.firstChild = ps.SingleExp
	}

	return ps
}

// propertyName
//     : identifierName
//     | StringLiteral
//     | numericLiteral
//     | '[' singleExpression ']'
//     ;
type PropertyName struct {
	*SourceInfo
	IdentifierName Token
	StringLiteral  Token
	NumericLiteral Token
	OpenBracket    Token
	SingleExp      VNode
	CloseBracket   Token
	firstChild     VNode
	prev, next     VNode
}

var _ VNode = (*PropertyName)(nil)

func (i *PropertyName) Type() string {
	return "PropertyName"
}
func (i *PropertyName) Code() string {
	return CodeDef(i)
}
func (i *PropertyName) Next() VNode {
	return i.next
}
func (i *PropertyName) SetNext(v VNode) {
	i.next = v
}
func (i *PropertyName) Prev() VNode {
	return i.prev
}
func (i *PropertyName) SetPrev(v VNode) {
	i.prev = v
}
func (i *PropertyName) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPropertyName(ctx *base.PropertyNameContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertyName", ctx.GetText())
	}

	pn := &PropertyName{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	if ctx.OpenBracket() != nil {
		pn.OpenBracket = ident(v, ctx.OpenBracket().GetSymbol())
		pn.firstChild = pn.OpenBracket

		if ctx.SingleExpression() != nil {
			pn.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)

		}
		pn.CloseBracket = ident(v, ctx.CloseBracket().GetSymbol())
		setAllSibs(pn.OpenBracket, pn.SingleExp, pn.CloseBracket)
		return pn
	}
	// just tokens, maybe best to use one field to cover all 3 and use getchild(0)
	if ctx.IdentifierName() != nil {
		pn.IdentifierName = v.VisitIdentifierName(ctx.IdentifierName().(*base.IdentifierNameContext)).(Token)
		pn.firstChild = pn.IdentifierName

	} else if ctx.StringLiteral() != nil {
		pn.StringLiteral = ident(v, ctx.StringLiteral().GetSymbol())
		pn.firstChild = pn.StringLiteral
	} else if ctx.NumericLiteral() != nil {
		pn.NumericLiteral = v.VisitNumericLiteral(ctx.NumericLiteral().(*base.NumericLiteralContext)).(Token)
		pn.firstChild = pn.NumericLiteral

	}

	return pn
}
