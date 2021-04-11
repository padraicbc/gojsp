package main

import "github.com/padraicbc/gojsp/parser"

func (v *visitor) VisitExportDeclaration(ctx *parser.ExportDeclarationContext) interface{} {
	// v.nodes = append(v.nodes, getSourceInfo(*ctx.BaseParserRuleContext))

	// log.Println("VisitExportDeclaration", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExportDefaultDeclaration(ctx *parser.ExportDefaultDeclarationContext) interface{} {
	// v.nodes = append(v.nodes, getSourceInfo(ctx))
	// log.Println("VisitExportDefaultDeclaration", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExportFromBlock(ctx *parser.ExportFromBlockContext) interface{} {
	// log.Println("VisitExportFromBlock", ctx.GetText())
	return v.VisitChildren(ctx)
}
