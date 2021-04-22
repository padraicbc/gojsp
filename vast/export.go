package vast

import "github.com/padraicbc/gojsp/base"

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;

func (v *Visitor) VisitExportDefaultDeclaration(ctx *base.ExportDefaultDeclarationContext) interface{} {
	// log.Println("VisitExportDefaultDeclaration", ctx.GetText())
	return v.VisitChildren(ctx)
}

// exportFromBlock
//     : importNamespace importFrom eos
//     | moduleItems importFrom? eos
//     ;
func (v *Visitor) VisitExportFromBlock(ctx *base.ExportFromBlockContext) interface{} {
	// log.Println("VisitExportFromBlock", ctx.GetText())

	return v.VisitChildren(ctx)
}

// exportStatement
//     : Export (exportFromBlock | declaration) eos    # ExportDeclaration
//     | Export Default singleExpression eos           # ExportDefaultDeclaration
//     ;
func (v *Visitor) VisitExportDeclaration(ctx *base.ExportDeclarationContext) interface{} {

	return v.VisitChildren(ctx)

}
