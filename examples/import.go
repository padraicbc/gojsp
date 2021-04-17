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
}

var _ VNode = (*ImportStatement)(nil)

func (i *ImportStatement) Type() string {
	return "ImportStatement"
}
func (i *ImportStatement) Code() string {
	return CodeDef(i)
}

func (i *ImportStatement) GetChildren() []VNode {

	return []VNode{i.Import, i.ImportFromBlock}
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
}

var _ VNode = (*ImportFromBlock)(nil)

func (i *ImportFromBlock) Type() string {
	return "ImportFromBlock"
}
func (i *ImportFromBlock) Code() string {
	return CodeDef(i)
}

func (i *ImportFromBlock) GetChildren() []VNode {

	return []VNode{
		i.Default,
		i.ModuleItems,
		i.ImportNamespace,
		i.StringLiteral,
		i.ImportFrom,
		i.Eos,
	}
}

type ImportFrom struct {
	*SourceInfo

	From Token
	Path Token
}

var _ VNode = (*ImportFrom)(nil)

func f(v VNode) {

}
func (i *ImportFrom) Type() string {
	return "ImportFrom"
}

func (i *ImportFrom) Code() string {
	return CodeDef(i)
}

func (i *ImportFrom) GetChildren() []VNode {

	return []VNode{i.From, i.Path}
}

// import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	*SourceInfo
	Import     Token
	Module     Token
	OpenParen  Token
	CloseParen Token
}

var _ VNode = (*ImportExpression)(nil)

func (i *ImportExpression) Code() string {
	return CodeDef(i)
}

func (i *ImportExpression) Type() string {
	return "ImportExpression"
}

func (i *ImportExpression) GetChildren() []VNode {

	return []VNode{i.Import, i.OpenParen, i.Module, i.CloseParen}
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
}

var _ VNode = (*ModuleItems)(nil)

func (m *ModuleItems) Code() string {

	return CodeDef(m)
}
func (m *ModuleItems) Type() string {
	return "ModuleItems"
}

/// todo: AliasNames
func (m *ModuleItems) GetChildren() []VNode {
	if m == nil {
		return nil
	}
	nd := []VNode{m.OpenBrace}
	for i, al := range m.AliasNames {
		nd = append(nd, al)
		if i > 0 {
			nd = append(nd, m.Commas[i-1])
		}

	}
	return append(nd, m.CloseBrace)
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
}

var _ VNode = (*ImportNamespace)(nil)

func (in *ImportNamespace) Code() string {
	return CodeDef(in)
}
func (i *ImportNamespace) Type() string {
	return "ImportNamespace"
}

func (i *ImportNamespace) GetChildren() []VNode {

	return []VNode{i.Star, i.IdentifierName, i.AliasName, i.As}
}

// importDefault
// : aliasName ','
type ImportDefault struct {
	*SourceInfo
	Default *AliasName
	Comma   Token
}

var _ VNode = (*ImportDefault)(nil)

func (i *ImportDefault) Type() string {
	return "ImportDefault"
}

func (i *ImportDefault) Code() string {
	return CodeDef(i)
}
func (i *ImportDefault) GetChildren() []VNode {
	// if i == nil {
	// 	return nil
	// }
	return []VNode{i.Default, i.Comma}
}

func (v *visitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {
	// log.Println("VisitImportStatement", ctx.GetText())
	// could vists and switch check but this is the same thing.
	im := &ImportStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		Import:     ident(v, ctx.Import().GetSymbol()),
		ImportFromBlock: v.VisitImportFromBlock(
			ctx.ImportFromBlock().(*parser.ImportFromBlockContext)).(*ImportFromBlock)}
	return im

}

func (v *visitor) VisitImportFromBlock(ctx *parser.ImportFromBlockContext) interface{} {
	imf := &ImportFromBlock{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// iterate here as some are there, some not.
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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
func (v *visitor) VisitImportExpression(ctx *parser.ImportExpressionContext) interface{} {
	// log.Println("VisitImportExpression", ctx.GetText())
	ime := &ImportExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// alwyas tokens
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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

func (v *visitor) VisitModuleItems(ctx *parser.ModuleItemsContext) interface{} {

	mit := &ModuleItems{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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
func (v *visitor) VisitImportDefault(ctx *parser.ImportDefaultContext) interface{} {
	// could iterate over children but this is the same thing.
	// todo: check types?
	return &ImportDefault{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		Default:    v.VisitAliasName(ctx.AliasName().(*parser.AliasNameContext)).(*AliasName),
		Comma:      ident(v, ctx.Comma().GetSymbol())}
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *visitor) VisitImportNamespace(ctx *parser.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())

	imn := &ImportNamespace{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for i, ch := range v.VisitChildren(ctx).([]VNode) {
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
func (v *visitor) VisitImportFrom(ctx *parser.ImportFromContext) interface{} {
	imfr := &ImportFrom{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
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
}

var _ VNode = (*AliasName)(nil)

func (a *AliasName) Code() string {
	return CodeDef(a)
}
func (i *AliasName) Type() string {
	return "AliasName"
}

func (i *AliasName) GetChildren() []VNode {

	return []VNode{i.IdentifierName, i.As, i.Alias}
}

func (v *visitor) VisitAliasName(ctx *parser.AliasNameContext) interface{} {
	al := &AliasName{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	for i, ch := range v.VisitChildren(ctx).([]VNode) {
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
