package vast

import (
	"fmt"

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
	children        VNode
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
	return children(i.children)
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
	children        VNode
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

	return children(i.children)
}

type ImportFrom struct {
	*SourceInfo
	children   VNode
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

	return children(i.children)
}

// import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	*SourceInfo
	Import     Token
	Module     Token
	OpenParen  Token
	CloseParen Token
	children   VNode
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

	return children(i.children)
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
	children   VNode
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

	return children(m.children)
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
	children       VNode
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

	return children(i.children)
}

// importDefault
// : aliasName ','
type ImportDefault struct {
	*SourceInfo
	Default    *AliasName
	Comma      Token
	children   VNode
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

	return children(i.children)
}

func (v *Visitor) VisitImportStatement(ctx *base.ImportStatementContext) interface{} {
	// log.Println("VisitImportStatement", ctx.GetText())
	// could vists and switch check but this is the same thing.

	im := &ImportStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if im.children == nil {
			im.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch

		if ch.Type() == "LToken" {
			im.Import = ch.(Token)

		} else {
			im.ImportFromBlock = ch.(*ImportFromBlock)
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
		if imf.children == nil {
			imf.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)

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
	// log.Println("VisitImportExpression", ctx.GetText())
	ime := &ImportExpression{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	// alwyas tokens
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ime.children == nil {
			ime.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
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

func (v *Visitor) VisitModuleItems(ctx *base.ModuleItemsContext) interface{} {

	mit := &ModuleItems{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if mit.children == nil {
			mit.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
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
func (v *Visitor) VisitImportDefault(ctx *base.ImportDefaultContext) interface{} {
	// could iterate over children but this is the same thing.
	// todo: check types?
	ind := &ImportDefault{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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
		if ind.children == nil {
			ind.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
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
		if imn.children == nil {
			imn.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
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
		if imfr.children == nil {
			imfr.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)

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
