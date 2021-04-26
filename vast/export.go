package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;

func (v *Visitor) VisitExportDefaultDeclaration(ctx *base.ExportDefaultDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitExportDefaultDeclaration", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}

// exportFromBlock
//     : importNamespace importFrom eos
//     | moduleItems importFrom? eos
//     ;
func (v *Visitor) VisitExportFromBlock(ctx *base.ExportFromBlockContext) interface{} {
	if v.Debug {
		log.Println("VisitExportFromBlock", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}

// exportStatement
//     : Export (exportFromBlock | declaration) eos    # ExportDeclaration
//     | Export Default singleExpression eos           # ExportDefaultDeclaration
//     ;
func (v *Visitor) VisitExportDeclaration(ctx *base.ExportDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitExportDeclaration", ctx.GetText())
	}

	return v.VisitChildren(ctx)

}
