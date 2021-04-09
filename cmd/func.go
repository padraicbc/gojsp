package main

import (
	"github.com/padraicbc/gojsp"
)

func (v *visitor) VisitArrowFunction(ctx *gojsp.ArrowFunctionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrowFunctionParameters(ctx *gojsp.ArrowFunctionParametersContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrowFunctionBody(ctx *gojsp.ArrowFunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionDecl(ctx *gojsp.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitAnoymousFunctionDecl(ctx *gojsp.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return v.VisitChildren(ctx)
}
