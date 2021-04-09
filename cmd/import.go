package main

import (
	"fmt"
	"strings"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp"
)

type Import interface {
	ImportString() string

	// AddInfo()
}

type SourceInfo struct {
}

// func (v *visitor) getSourceInfo(ctx gojsp.BaseContext) node {
// 	return node{line: ctx.GetStart().GetLine(), start: ctx.GetStart().GetStart(), end: ctx.GetStop().GetStart(),
// 		source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
// 			Start: ctx.GetStart().GetStart(), Stop: ctx.GetStop().GetStop() + 1})}
// }

// ImportDeclaration

// interface ImportDeclaration <: ModuleDeclaration {
//     type: "ImportDeclaration";
//     specifiers: [ ImportSpecifier |  | ImportNamespaceSpecifier ];
//     source: Literal;
// }
// Import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	SingleExpression string
}

func (i ImportExpression) ImportString() string {
	return fmt.Sprintf("import (%s)", i.SingleExpression)
}

// importStatement
//     : Import importFromBlock
//     ;

// aliasName
//     : identifierName (As identifierName)?
//     ;
type AliasName struct {
	IdentifierName string
	Alias          string
}

func (a AliasName) ImportString() string {
	if a.Alias != "" {
		return "as " + a.IdentifierName
	}
	return a.IdentifierName
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
type ModulesItems struct {
	AliasNames []AliasName
}

func (m *ModulesItems) ImportString() string {
	out := []string{}
	for _, a := range m.AliasNames {
		if a.Alias != "" {
			out = append(out, "as "+a.IdentifierName)
		} else {
			out = append(out, a.IdentifierName)
		}
	}
	if len(out) == 0 {
		return ""
	}

	return "{" + strings.Join(out, ", ") + "}"
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
type ImportNamespace struct {
	// Star           string
	IdentifierName string
	AliasName      string
}

func (in *ImportNamespace) ImportString() string {
	if in.AliasName != "" {
		return strings.TrimSpace(in.IdentifierName + " as " + in.AliasName)
	}
	return in.IdentifierName
}

// importDefault
//     : aliasName ','
//     ;
type ImportDefault struct {
	AliasName AliasName
}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
// importDefault
// : aliasName ','
type ImportDeclaration struct {
	From            string
	Default         ImportDefault
	ModulesItems    ModulesItems
	ImportNamespace ImportNamespace
	StringLiteral   string
	ImportFrom      string
}

func (i ImportDeclaration) ImportString() string {
	var def, from string
	if td := i.Default.AliasName.ImportString(); td != "" {
		def = " " + td
	}
	if i.From != "" {
		from = " from "
	}

	return fmt.Sprintf("import%s %s%s%s", def, i.ImportNamespace.ImportString()+
		i.ModulesItems.ImportString(), from, i.StringLiteral+i.ImportFrom)
}

func (v *visitor) VisitImportStatement(ctx *gojsp.ImportStatementContext) interface{} {

	// we could format based on some spec. Could be done here, at each step or at the very end...
	return v.VisitChildren(ctx)

}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
//     ;
func (v *visitor) VisitImportFromBlock(ctx *gojsp.ImportFromBlockContext) interface{} {

	imp := ImportDeclaration{}
	if ct := ctx.ImportDefault(); ct != nil {
		imp.Default = v.VisitImportDefault(ct.(*gojsp.ImportDefaultContext)).(ImportDefault)
	}

	if ctx.ModuleItems() != nil {
		imp.ModulesItems = v.VisitModuleItems(ctx.ModuleItems().(*gojsp.ModuleItemsContext)).(ModulesItems)
	}

	if ct := ctx.ImportFrom(); ct != nil {
		imp.From = "from"
		imp.ImportFrom = (ct.GetChild(1).(*antlr.TerminalNodeImpl).GetText())
	}
	if ct := ctx.ImportNamespace(); ct != nil {
		imp.ImportNamespace = v.VisitImportNamespace(ct.(*gojsp.ImportNamespaceContext)).(ImportNamespace)
	}
	if ct := ctx.StringLiteral(); ct != nil {
		imp.StringLiteral = ct.GetText()
	}

	return imp
}

// Import '(' singleExpression ')'
func (v *visitor) VisitImportExpression(ctx *gojsp.ImportExpressionContext) interface{} {

	// log.Println("VisitImportExpression", ctx.GetText())

	return ImportExpression{SingleExpression: ctx.SingleExpression().GetText()}
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
// just pass on the wor if nothign to change
func (v *visitor) VisitModuleItems1(ctx *gojsp.ModuleItemsContext) interface{} {

	return v.VisitChildren(ctx)
}

//  alternative version where we do the AliasName work ourselves so we can change...
func (v *visitor) VisitModuleItems(ctx *gojsp.ModuleItemsContext) interface{} {
	var m ModulesItems

	for _, mc := range ctx.AllAliasName() {

		m.AliasNames = append(m.AliasNames, v.VisitAliasName((mc.(*gojsp.AliasNameContext))).(AliasName))

	}

	// . ctx.OpenBrace().GetText() ctx.CloseBrace().GetText()?
	return m
}
func (v *visitor) VisitImportDefault(ctx *gojsp.ImportDefaultContext) interface{} {
	return ImportDefault{AliasName: v.VisitAliasName(ctx.AliasName().(*gojsp.AliasNameContext)).(AliasName)}
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *visitor) VisitImportNamespace(ctx *gojsp.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())
	// todo: add to errors

	var imp ImportNamespace
	// either this or *gojsp.IdentifierNameContext i.e * or any identifier
	if _, ok := ctx.GetChild(0).(*antlr.TerminalNodeImpl); ok {
		imp.IdentifierName = "*"

	} else if ident := ctx.IdentifierName(0); ident != nil {
		// log.Println(ident.GetText())
		imp.IdentifierName = ident.GetText()
	}

	return imp

}

// importFrom
//     : From StringLiteral
//     ;
func (v *visitor) VisitImportFrom(ctx *gojsp.ImportFromContext) interface{} {

	if ctx.GetChildCount() != 2 {
		// todo: error
		// ctx.From().GetSymbol().GetLine()
		panic("wrong child count for importfrom")
	}

	// could do work here also ctx.From() and ctx.GetChild(1) -> path
	return v.VisitChildren(ctx)

}

// aliasName
//     : identifierName (As identifierName)?
//     ;
func (v *visitor) VisitAliasName(ctx *gojsp.AliasNameContext) interface{} {
	// log.Println("VisitAliasName", ctx.GetChildCount(), ctx.GetText())
	a := AliasName{}
	// todo: syntax error if nil
	if ident := ctx.IdentifierName(0); ident != nil {
		a.IdentifierName = ident.GetText()

	}

	// todo: again syntax error if nil as . is a syntax error
	if ident := ctx.IdentifierName(1); ident != nil {
		a.Alias = ident.GetText()

	}

	return a
}
