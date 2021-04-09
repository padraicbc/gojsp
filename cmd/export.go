package main

import (
	"github.com/padraicbc/gojsp"
)

func (v *visitor) VisitExportDeclaration(ctx *gojsp.ExportDeclarationContext) interface{} {
	v.nodes = append(v.nodes, v.getSourceInfo(ctx))

	// log.Println("VisitExportDeclaration", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExportDefaultDeclaration(ctx *gojsp.ExportDefaultDeclarationContext) interface{} {
	v.nodes = append(v.nodes, v.getSourceInfo(ctx))
	// log.Println("VisitExportDefaultDeclaration", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExportFromBlock(ctx *gojsp.ExportFromBlockContext) interface{} {
	// log.Println("VisitExportFromBlock", ctx.GetText())
	return v.VisitChildren(ctx)
}
