package vast

import (
	"fmt"
	"log"

	"github.com/padraicbc/gojsp/base"
)

// importStatement
//     : Import importFromBlock
//     ;
type ImportStatement struct {
	*SourceInfo
	Import Token
	// one child *ImportFromBlock
	ImportFromBlock *ImportFromBlock
	firstChild      VNode

	prev, next VNode
}

var _ VNode = (*ImportStatement)(nil)

func (i *ImportStatement) Type() string {
	return "ImportStatement"
}
func (i *ImportStatement) Code() string {
	return CodeDef(i)
}
func (i *ImportStatement) Next() VNode {

	return i.next
}
func (i *ImportStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ImportStatement) Prev() VNode {

	return i.prev
}
func (i *ImportStatement) SetPrev(v VNode) {
	i.prev = v
}

func (i *ImportStatement) FirstChild() VNode {

	return i.firstChild

}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
type ImportFromBlock struct {
	*SourceInfo

	ImportDefault   *ImportDefault
	ModuleItems     *ModuleItems
	ImportNamespace *ImportNamespace
	StringLiteral   Token
	ImportFrom      *ImportFrom
	Eos             Token
	firstChild      VNode

	prev, next VNode
}

var _ VNode = (*ImportFromBlock)(nil)

func (i *ImportFromBlock) Next() VNode {

	return i.next
}
func (i *ImportFromBlock) SetNext(v VNode) {
	i.next = v
}
func (i *ImportFromBlock) Prev() VNode {

	return i.prev
}
func (i *ImportFromBlock) SetPrev(v VNode) {
	i.prev = v
}
func (i *ImportFromBlock) Type() string {
	return "ImportFromBlock"
}
func (i *ImportFromBlock) Code() string {
	return CodeDef(i)
}

func (i *ImportFromBlock) FirstChild() VNode {

	return i.firstChild

}

type ImportFrom struct {
	*SourceInfo
	firstChild VNode

	prev, next VNode

	From Token
	Path Token
}

var _ VNode = (*ImportFrom)(nil)

func (i *ImportFrom) Next() VNode {

	return i.next
}
func (i *ImportFrom) SetNext(v VNode) {
	i.next = v
}
func (i *ImportFrom) Prev() VNode {

	return i.prev
}
func (i *ImportFrom) SetPrev(v VNode) {
	i.prev = v
}
func (i *ImportFrom) Type() string {
	return "ImportFrom"
}

func (i *ImportFrom) Code() string {
	return CodeDef(i)
}

func (i *ImportFrom) FirstChild() VNode {

	return i.firstChild

}

// import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	*SourceInfo
	Import     Token
	Module     Token
	OpenParen  Token
	CloseParen Token
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*ImportExpression)(nil)

func (i *ImportExpression) Next() VNode {

	return i.next
}
func (i *ImportExpression) SetNext(v VNode) {
	i.next = v
}
func (i *ImportExpression) Prev() VNode {

	return i.prev
}
func (i *ImportExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *ImportExpression) Code() string {
	return CodeDef(i)
}

func (i *ImportExpression) Type() string {
	return "ImportExpression"
}

func (i *ImportExpression) FirstChild() VNode {

	return i.firstChild

}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
type ImportNamespace struct {
	*SourceInfo
	Star           Token
	IdentifierName Token
	AliasName      Token
	As             Token
	firstChild     VNode

	prev, next VNode
}

var _ VNode = (*ImportNamespace)(nil)

func (i *ImportNamespace) Next() VNode {

	return i.next
}
func (i *ImportNamespace) SetNext(v VNode) {
	i.next = v
}
func (i *ImportNamespace) Prev() VNode {

	return i.prev
}
func (i *ImportNamespace) SetPrev(v VNode) {
	i.prev = v
}

func (in *ImportNamespace) Code() string {
	return CodeDef(in)
}
func (i *ImportNamespace) Type() string {
	return "ImportNamespace"
}

func (i *ImportNamespace) FirstChild() VNode {

	return i.firstChild

}

// importDefault
// : aliasName ','
type ImportDefault struct {
	*SourceInfo
	Default    *AliasName
	Comma      Token
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*ImportDefault)(nil)

func (i *ImportDefault) Next() VNode {

	return i.next
}
func (i *ImportDefault) SetNext(v VNode) {
	i.next = v
}
func (i *ImportDefault) Prev() VNode {

	return i.prev
}
func (i *ImportDefault) SetPrev(v VNode) {
	i.prev = v
}
func (i *ImportDefault) Type() string {
	return "ImportDefault"
}

func (i *ImportDefault) Code() string {
	return CodeDef(i)
}
func (i *ImportDefault) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitImportStatement(ctx *base.ImportStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitImportStatement", ctx.GetText())
	}
	// could vists and switch check but this is the same thing.

	im := &ImportStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if im.firstChild == nil {
			im.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch
		switch ch.Type() {
		case "LToken":
			im.Import = ch.(Token)

		case "ImportFromBlock":
			im.ImportFromBlock = ch.(*ImportFromBlock)
		default:
			log.Panicf("%+v\n", ch)
		}

	}

	return im

}

func (v *Visitor) VisitImportFromBlock(ctx *base.ImportFromBlockContext) interface{} {
	imf := &ImportFromBlock{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	// iterate here as some are there, some not.
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if imf.firstChild == nil {
			imf.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch
		switch ch.Type() {

		case "ImportDefault":
			imf.ImportDefault = ch.(*ImportDefault)
		case "ImportNamespace":
			imf.ImportNamespace = ch.(*ImportNamespace)
		case "ModuleItems":
			imf.ModuleItems = ch.(*ModuleItems)
		case "ImportFrom":
			imf.ImportFrom = ch.(*ImportFrom)
		case "LToken":
			tk := ch.(Token)
			if tk.SymbolName() == "StringLiteral" {
				imf.StringLiteral = tk
				continue
			}
			if tk.SymbolName() == "SemiColon" {
				imf.Eos = tk
				continue
			}

			panic(fmt.Sprintf("%+v %s\n", ch, ch.Type()))

		}

	}

	return imf

}

//  Import '(' singleExpression ')' || Import "whatever"
func (v *Visitor) VisitImportExpression(ctx *base.ImportExpressionContext) interface{} {

	if v.Debug {
		log.Println("VisitImportExpression", ctx.GetText(), ctx.GetChildCount())
	}
	ime := &ImportExpression{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	// alwyas tokens
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ime.firstChild == nil {
			ime.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch
		t := ch.(Token)
		switch t.SymbolName() {
		case "Import":
			ime.Import = t
		case "OpenParen":
			ime.OpenParen = t
		case "StringLiteral":
			ime.Module = t
		case "CloseParen":
			ime.CloseParen = t
		default:
			panic(t.SymbolName())
		}

	}

	return ime
}

// importDefault
//     : aliasName ','
//     ;
func (v *Visitor) VisitImportDefault(ctx *base.ImportDefaultContext) interface{} {
	// could iterate over children but this is the same thing.
	// todo: check types?
	ind := &ImportDefault{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ind.firstChild == nil {
			ind.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch

		switch ch.Type() {
		case "AliasName":
			ind.Default = ch.(*AliasName)
		case "LToken":
			t := ch.(Token)

			if t.SymbolName() == "Comma" {
				ind.Comma = t
				continue
			}
			log.Panic(t.SymbolName())

		default:

			log.Panicf("%+v %s\n", ch, ch.Type())
		}

	}
	return ind
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *Visitor) VisitImportNamespace(ctx *base.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())

	imn := &ImportNamespace{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for i, ch := range v.VisitChildren(ctx).([]VNode) {
		if imn.firstChild == nil {
			imn.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch
		t := ch.(Token)
		switch t.SymbolName() {
		case "Identifier":
			if i == 0 {
				imn.IdentifierName = t
			} else {
				imn.AliasName = t
			}
		case "As":
			imn.As = t
		case "Multiply":
			imn.Star = t

		default:
			panic(fmt.Sprintf("%+v\n", ch))
		}

	}
	return imn

}

// importFrom
//     : From StringLiteral
//     ;
func (v *Visitor) VisitImportFrom(ctx *base.ImportFromContext) interface{} {
	imfr := &ImportFrom{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if imfr.firstChild == nil {
			imfr.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch
		t := ch.(Token)
		switch t.SymbolName() {
		case "From":
			imfr.From = t
		case "StringLiteral":
			imfr.Path = t
		default:
			panic(t.SymbolName())
		}

	}

	return imfr
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
type ModuleItems struct {
	*SourceInfo
	// always AliasName(s) not sure of any nice way dynamically add so 1-n so just using a slice of them
	AliasNames []*AliasName
	Commas     []Token
	OpenBrace  Token
	CloseBrace Token
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*ModuleItems)(nil)

func (i *ModuleItems) Next() VNode {

	return i.next
}
func (i *ModuleItems) SetNext(v VNode) {
	i.next = v
}
func (i *ModuleItems) Prev() VNode {

	return i.prev
}
func (i *ModuleItems) SetPrev(v VNode) {
	i.prev = v
}
func (m *ModuleItems) Code() string {

	return CodeDef(m)
}
func (m *ModuleItems) Type() string {
	return "ModuleItems"
}

func (m *ModuleItems) FirstChild() VNode {

	return m.firstChild

}

func (v *Visitor) VisitModuleItems(ctx *base.ModuleItemsContext) interface{} {

	mit := &ModuleItems{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if mit.firstChild == nil {
			mit.firstChild = ch
		}
		prev = setSib(prev, ch)

		prev = ch

		switch ch.Type() {
		case "AliasName":
			mit.AliasNames = append(mit.AliasNames, ch.(*AliasName))
		case "LToken":
			t := ch.(Token)
			if t.SymbolName() == "OpenBrace" {
				mit.OpenBrace = t
				continue
			}
			if t.SymbolName() == "CloseBrace" {
				mit.CloseBrace = t
				continue
			}
			if t.SymbolName() == "Comma" {
				mit.Commas = append(mit.Commas, t)
				continue
			}
			log.Panic(t.SymbolName())

		default:
			log.Panicf("%+v %s\n", ch, ch.Type())
		}

	}

	return mit

}
