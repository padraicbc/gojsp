package main

import (
	"github.com/padraicbc/gojsp/parser"
)

func (v *visitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {

	im := &ImportStatement{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	im.ImportFromBlock = im.Children[0].(*ImportFromBlock)
	return im

}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
//     ;
func (v *visitor) VisitImportFromBlock(ctx *parser.ImportFromBlockContext) interface{} {
	im := &ImportFromBlock{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		switch rr := ch.(type) {
		case *ImportDefault:
			im.Default = rr
		case *ImportNamespace:
			im.ImportNamespace = rr
		case *ModulesItems:
			im.ModulesItems = rr
		case *ImportFrom:
			im.ImportFrom = rr
		}
	}
	return im

}

// *SourceInfo '(' singleExpression ')'
func (v *visitor) VisitImportExpression(ctx *parser.ImportExpressionContext) interface{} {

	return &ImportExpression{
		Children: v.VisitChildren(ctx).([]VNode),
		// SingleExpression: ctx.SingleExpression().GetText(),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}

//  alternative version where we do the AliasName work ourselves so we can change...
func (v *visitor) VisitModuleItems(ctx *parser.ModuleItemsContext) interface{} {

	// . ctx.OpenBrace().GetText() ctx.CloseBrace().GetText()?
	m := &ModulesItems{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	for _, ch := range m.Children {
		m.AliasNames = append(m.AliasNames, ch.(*AliasName))
	}
	return m
}
func (v *visitor) VisitImportDefault(ctx *parser.ImportDefaultContext) interface{} {

	return &ImportDefault{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *visitor) VisitImportNamespace(ctx *parser.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())

	return &ImportNamespace{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

}

// importFrom
//     : From StringLiteral
//     ;
func (v *visitor) VisitImportFrom(ctx *parser.ImportFromContext) interface{} {

	// if ctx.GetChildCount() != 2 {
	// 	// todo: error
	// 	// ctx.From().GetSymbol().GetLine()
	// 	panic("wrong child count for importfrom")
	// }
	a := &ImportFrom{
		From: ctx.From().GetText(), Path: ctx.StringLiteral().GetText(),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	return a

}

// aliasName
//     : identifierName (As identifierName)?
//     ;
func (v *visitor) VisitAliasName(ctx *parser.AliasNameContext) interface{} {

	a := &AliasName{
		IdentifierName: ctx.IdentifierName(0).GetText(),
		SourceInfo:     getSourceInfo(*ctx.BaseParserRuleContext)}
	if as := ctx.As(); as != nil {
		a.Alias = ctx.IdentifierName(1).GetText()
		a.IdentifierName = ctx.IdentifierName(0).GetText()
	}
	return a

}
