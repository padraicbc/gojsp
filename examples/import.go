package main

import (
	"github.com/padraicbc/gojsp/parser"
)

func (v *visitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {

	return v.VisitChildren(ctx)

}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
//     ;
func (v *visitor) VisitImportFromBlock(ctx *parser.ImportFromBlockContext) interface{} {
	// log.Printf("%+v\n", v.VisitChildren(ctx)[0].Code())

	return v.VisitChildren(ctx)

}

// *SourceInfo '(' singleExpression ')'
func (v *visitor) VisitImportExpression(ctx *parser.ImportExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

//  alternative version where we do the AliasName work ourselves so we can change...
func (v *visitor) VisitModuleItems(ctx *parser.ModuleItemsContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitImportDefault(ctx *parser.ImportDefaultContext) interface{} {

	return v.VisitChildren(ctx)
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *visitor) VisitImportNamespace(ctx *parser.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())

	return v.VisitChildren(ctx)

}

// importFrom
//     : From StringLiteral
//     ;
func (v *visitor) VisitImportFrom(ctx *parser.ImportFromContext) interface{} {

	return v.VisitChildren(ctx)
}

// aliasName
//     : identifierName (As identifierName)?
//     ;
func (v *visitor) VisitAliasName(ctx *parser.AliasNameContext) interface{} {

	return v.VisitChildren(ctx)
}
