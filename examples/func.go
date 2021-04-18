package main

import "github.com/padraicbc/gojsp/parser"

// interface Function <: Node {
//     id: Identifier | null;
//     params: [ Pattern ];
//     body: FunctionBody;
// }
type Function struct {
}

func (v *Visitor) VisitArrowFunction(ctx *parser.ArrowFunctionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowFunctionParameters(ctx *parser.ArrowFunctionParametersContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowFunctionBody(ctx *parser.ArrowFunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnoymousFunctionDecl(ctx *parser.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return v.VisitChildren(ctx)
}
