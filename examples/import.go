package main

import (
	"fmt"

	"github.com/padraicbc/gojsp/parser"
)

// ** Not sure if pointer receivers are better or not. Means checkign for nil a lot more
// but does allow easily manipulating struct values...**

// importStatement
//     : Import importFromBlock
//     ;
type ImportStatement struct {
	*SourceInfo
	Import Token
	// one child *ImportFromBlock
	ImportFromBlock *ImportFromBlock
	children        []VNode
	prev, next      VNode
}

var _ VNode = (*ImportStatement)(nil)

func (i *ImportStatement) Type() string {
	return "ImportStatement"
}
func (i *ImportStatement) Code() string {
	return CodeDef(i)
}
func (i *ImportStatement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ImportStatement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *ImportStatement) Children() []VNode {

	return i.children
}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
type ImportFromBlock struct {
	*SourceInfo

	Default         *ImportDefault
	ModuleItems     *ModuleItems
	ImportNamespace *ImportNamespace
	StringLiteral   Token
	ImportFrom      *ImportFrom
	Eos             Token
	children        []VNode
	prev, next      VNode
}

var _ VNode = (*ImportFromBlock)(nil)

func (i *ImportFromBlock) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ImportFromBlock) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ImportFromBlock) Type() string {
	return "ImportFromBlock"
}
func (i *ImportFromBlock) Code() string {
	return CodeDef(i)
}

func (i *ImportFromBlock) Children() []VNode {

	return i.children
}

type ImportFrom struct {
	*SourceInfo
	children   []VNode
	prev, next VNode

	From Token
	Path Token
}

var _ VNode = (*ImportFrom)(nil)

func (i *ImportFrom) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ImportFrom) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ImportFrom) Type() string {
	return "ImportFrom"
}

func (i *ImportFrom) Code() string {
	return CodeDef(i)
}

func (i *ImportFrom) Children() []VNode {

	return i.children
}

// import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	*SourceInfo
	Import     Token
	Module     Token
	OpenParen  Token
	CloseParen Token
	children   []VNode
	prev, next VNode
}

var _ VNode = (*ImportExpression)(nil)

func (i *ImportExpression) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ImportExpression) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ImportExpression) Code() string {
	return CodeDef(i)
}

func (i *ImportExpression) Type() string {
	return "ImportExpression"
}

func (i *ImportExpression) Children() []VNode {

	return i.children
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
type ModuleItems struct {
	*SourceInfo
	// always AliasName(s) not sure of any nice way dynamically add so 1-n so juyst using a slice of them
	AliasNames []*AliasName
	Commas     []Token
	OpenBrace  Token
	CloseBrace Token
	children   []VNode
	prev, next VNode
}

var _ VNode = (*ModuleItems)(nil)

func (i *ModuleItems) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ModuleItems) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (m *ModuleItems) Code() string {

	return CodeDef(m)
}
func (m *ModuleItems) Type() string {
	return "ModuleItems"
}

/// todo: AliasNames
func (m *ModuleItems) Children() []VNode {

	return m.children
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
	children       []VNode
	prev, next     VNode
}

var _ VNode = (*ImportNamespace)(nil)

func (i *ImportNamespace) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ImportNamespace) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (in *ImportNamespace) Code() string {
	return CodeDef(in)
}
func (i *ImportNamespace) Type() string {
	return "ImportNamespace"
}

func (i *ImportNamespace) Children() []VNode {

	return i.children
}

// importDefault
// : aliasName ','
type ImportDefault struct {
	*SourceInfo
	Default    *AliasName
	Comma      Token
	children   []VNode
	prev, next VNode
}

var _ VNode = (*ImportDefault)(nil)

func (i *ImportDefault) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ImportDefault) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ImportDefault) Type() string {
	return "ImportDefault"
}

func (i *ImportDefault) Code() string {
	return CodeDef(i)
}
func (i *ImportDefault) Children() []VNode {

	return i.children
}

func (v *Visitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {
	// log.Println("VisitImportStatement", ctx.GetText())
	// could vists and switch check but this is the same thing.
	im := &ImportStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		children:   v.VisitChildren(ctx).([]VNode)}
	return im

}

func (v *Visitor) VisitImportFromBlock(ctx *parser.ImportFromBlockContext) interface{} {
	imf := &ImportFromBlock{
		children: v.VisitChildren(ctx).([]VNode), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// iterate here as some are there, some not.
	for _, ch := range imf.children {
		switch ch.Type() {
		case "ImportDefault":
			imf.Default = ch.(*ImportDefault)
		case "ImportNamespace":
			imf.ImportNamespace = ch.(*ImportNamespace)
		case "ModuleItems":
			imf.ModuleItems = ch.(*ModuleItems)
		case "ImportFrom":
			imf.ImportFrom = ch.(*ImportFrom)

		default:
			if t, ok := ch.(Token); ok {
				// sn:SemiColon
				imf.Eos = t
				continue

			}
			panic(fmt.Sprintf("%+v %s\n", ch, ch.Type()))
		}
	}
	return imf

}

//  Import '(' singleExpression ')' || Import "whatever"
func (v *Visitor) VisitImportExpression(ctx *parser.ImportExpressionContext) interface{} {
	// log.Println("VisitImportExpression", ctx.GetText())
	ime := &ImportExpression{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// alwyas tokens
	for _, ch := range ime.children {
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

func (v *Visitor) VisitModuleItems(ctx *parser.ModuleItemsContext) interface{} {

	mit := &ModuleItems{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range mit.children {
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
			panic(t.SymbolName())

		default:

			panic(fmt.Sprintf("%+v %s\n", ch, ch.Type()))
		}
	}
	return mit

}

// importDefault
//     : aliasName ','
//     ;
func (v *Visitor) VisitImportDefault(ctx *parser.ImportDefaultContext) interface{} {
	// could iterate over children but this is the same thing.
	// todo: check types?
	ind := &ImportDefault{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	for _, ch := range ind.children {
		switch ch.Type() {
		case "AliasName":
			ind.Default = ch.(*AliasName)
		case "LToken":
			t := ch.(Token)

			if t.SymbolName() == "Comma" {
				ind.Comma = t
				continue
			}
			panic(t.SymbolName())

		default:

			panic(fmt.Sprintf("%+v %s\n", ch, ch.Type()))
		}
	}
	return ind
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *Visitor) VisitImportNamespace(ctx *parser.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())

	imn := &ImportNamespace{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	for i, ch := range imn.children {
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
func (v *Visitor) VisitImportFrom(ctx *parser.ImportFromContext) interface{} {
	imfr := &ImportFrom{children: v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range imfr.children {
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

// aliasName
//     : identifierName (As identifierName)?
//     ;
type AliasName struct {
	*SourceInfo
	IdentifierName Token
	Alias          Token
	As             Token
	Comma          Token
	children       []VNode
	prev, next     VNode
}

var _ VNode = (*AliasName)(nil)

func (i *AliasName) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *AliasName) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (a *AliasName) Code() string {
	return CodeDef(a)
}
func (i *AliasName) Type() string {
	return "AliasName"
}

func (i *AliasName) Children() []VNode {

	return i.children
}

func (v *Visitor) VisitAliasName(ctx *parser.AliasNameContext) interface{} {
	al := &AliasName{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	for i, ch := range al.children {
		t := ch.(Token)
		switch t.SymbolName() {
		case "Identifier":
			// always there
			if i == 0 {
				al.IdentifierName = t
				// > 0 means alias
			} else {
				al.Alias = t
			}
		case "As":
			al.As = t
		default:
			panic(t.SymbolName())

		}
	}
	return al
}
