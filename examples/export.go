package main

import (
	"log"

	"github.com/padraicbc/gojsp/parser"
)

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;

func (v *visitor) VisitExportDefaultDeclaration(ctx *parser.ExportDefaultDeclarationContext) interface{} {
	log.Println("VisitExportDefaultDeclaration", ctx.GetText())

	return v.VisitChildren(ctx)
}

// exportFromBlock
//     : importNamespace importFrom eos
//     | moduleItems importFrom? eos
//     ;
func (v *visitor) VisitExportFromBlock(ctx *parser.ExportFromBlockContext) interface{} {
	// log.Println("VisitExportFromBlock", ctx.GetText())

	return v.VisitChildren(ctx)
}

// exportStatement
//     : Export (exportFromBlock | declaration) eos    # ExportDeclaration
//     | Export Default singleExpression eos           # ExportDefaultDeclaration
//     ;
func (v *visitor) VisitExportDeclaration(ctx *parser.ExportDeclarationContext) interface{} {

	return v.VisitChildren(ctx)

}
