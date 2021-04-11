package main

import "github.com/padraicbc/gojsp/parser"

func (v *visitor) VisitArrowFunction(ctx *parser.ArrowFunctionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrowFunctionParameters(ctx *parser.ArrowFunctionParametersContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrowFunctionBody(ctx *parser.ArrowFunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitAnoymousFunctionDecl(ctx *parser.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return v.VisitChildren(ctx)
}
