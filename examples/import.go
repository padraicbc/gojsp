package main

import (
	"fmt"
	"strings"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

// ** Not sure if pointer receivers are better or not. Means checkign for nil a lot more
// but does allow easily manipulating struct values...**
// ImportDeclaration

// interface ImportDeclaration <: ModuleDeclaration {
//     type: "ImportDeclaration";
//     specifiers: [ ImportSpecifier |  | ImportNamespaceSpecifier ];
//     source: Literal;
// }
// import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	*SourceInfo
	SingleExpression string
}

func (i *ImportExpression) Code() string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("import (%s)", i.SingleExpression)
}
func (i *ImportExpression) Type() string {
	return "ImportExpression"
}

// importStatement
//     : *SourceInfo importFromBlock
//     ;

// aliasName
//     : identifierName (As identifierName)?
//     ;
type AliasName struct {
	*SourceInfo
	IdentifierName string
	Alias          string
}

func (a AliasName) Code() string {
	if a.Alias != "" {
		return "as " + a.IdentifierName
	}
	return a.IdentifierName
}
func (i AliasName) Type() string {
	return "AliasName"
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
type ModulesItems struct {
	*SourceInfo
	AliasNames []AliasName
}

func (m *ModulesItems) Code() string {
	if m == nil {
		return ""
	}
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
func (i *ModulesItems) Type() string {
	return "ModulesItems"
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
type ImportNamespace struct {
	*SourceInfo
	// Star           string
	IdentifierName string
	AliasName      string
}

func (in *ImportNamespace) Code() string {
	if in == nil {
		return ""
	}
	if in.AliasName != "" {
		return strings.TrimSpace(in.IdentifierName + " as " + in.AliasName)
	}
	return in.IdentifierName
}
func (i *ImportNamespace) Type() string {
	return "ImportNamespace"
}

// importDefault
//     : aliasName ','
//     ;
type ImportDefault struct {
	*SourceInfo
	AliasName AliasName
}

func (i *ImportDefault) Type() string {
	return "ImportDefault"
}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
// importDefault
// : aliasName ','
type ImportDeclaration struct {
	*SourceInfo
	From            string
	Default         *ImportDefault
	ModulesItems    *ModulesItems
	ImportNamespace *ImportNamespace
	StringLiteral   string
	ImportFrom      string
}

func (i *ImportDeclaration) Type() string {
	return "ImportDeclaration"
}
func (i *ImportDeclaration) Code() string {
	var def, from string
	if df := i.Default; df != nil && i.Default.AliasName.Code() != "" {
		def = " " + i.Default.AliasName.Code()
	}
	if i.From != "" {
		from = " from "
	}

	return fmt.Sprintf("import%s %s%s%s", def, i.ImportNamespace.Code()+
		i.ModulesItems.Code(), from, i.StringLiteral+i.ImportFrom)
}

func (v *visitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {

	// we could format based on some spec. Could be done here, at each step or at the very end...
	return v.VisitChildren(ctx)

}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
//     ;
func (v *visitor) VisitImportFromBlock(ctx *parser.ImportFromBlockContext) interface{} {

	var imp = &ImportDeclaration{}
	imp.SourceInfo = getSourceInfo(*ctx.BaseParserRuleContext)
	if ct := ctx.ImportDefault(); ct != nil {
		imp.Default = v.VisitImportDefault(ct.(*parser.ImportDefaultContext)).(*ImportDefault)
	}

	if ctx.ModuleItems() != nil {
		imp.ModulesItems = v.VisitModuleItems(ctx.ModuleItems().(*parser.ModuleItemsContext)).(*ModulesItems)
	}

	if ct := ctx.ImportFrom(); ct != nil {
		imp.From = "from"
		imp.ImportFrom = (ct.GetChild(1).(*antlr.TerminalNodeImpl).GetText())
	}
	if ct := ctx.ImportNamespace(); ct != nil {
		imp.ImportNamespace = v.VisitImportNamespace(ct.(*parser.ImportNamespaceContext)).(*ImportNamespace)
	}
	if ct := ctx.StringLiteral(); ct != nil {
		imp.StringLiteral = ct.GetText()
	}

	return imp
}

// *SourceInfo '(' singleExpression ')'
func (v *visitor) VisitImportExpression(ctx *parser.ImportExpressionContext) interface{} {

	return ImportExpression{
		SingleExpression: ctx.SingleExpression().GetText(),
		SourceInfo:       getSourceInfo(*ctx.BaseParserRuleContext)}
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
// just pass on the wor if nothign to change
func (v *visitor) VisitModuleItems1(ctx *parser.ModuleItemsContext) interface{} {

	return v.VisitChildren(ctx)
}

//  alternative version where we do the AliasName work ourselves so we can change...
func (v *visitor) VisitModuleItems(ctx *parser.ModuleItemsContext) interface{} {
	var m = &ModulesItems{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, mc := range ctx.AllAliasName() {
		m.AliasNames = append(m.AliasNames, v.VisitAliasName((mc.(*parser.AliasNameContext))).(AliasName))
	}

	// . ctx.OpenBrace().GetText() ctx.CloseBrace().GetText()?
	return m
}
func (v *visitor) VisitImportDefault(ctx *parser.ImportDefaultContext) interface{} {
	return &ImportDefault{AliasName: v.VisitAliasName(ctx.AliasName().(*parser.AliasNameContext)).(AliasName),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *visitor) VisitImportNamespace(ctx *parser.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())
	// todo: add to errors

	var imp = &ImportNamespace{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// either this or *parser.IdentifierNameContext i.e * or any identifier
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
func (v *visitor) VisitImportFrom(ctx *parser.ImportFromContext) interface{} {

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
func (v *visitor) VisitAliasName(ctx *parser.AliasNameContext) interface{} {
	// log.Println("VisitAliasName", ctx.GetChildCount(), ctx.GetText())
	a := AliasName{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
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
