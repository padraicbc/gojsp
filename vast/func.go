package vast

import "github.com/padraicbc/gojsp/base"

// interface Function <: Node {
//     id: Identifier | null;
//     params: [ Pattern ];
//     body: FunctionBody;
// }
type Function struct {
}

func (v *Visitor) VisitArrowFunction(ctx *base.ArrowFunctionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowFunctionParameters(ctx *base.ArrowFunctionParametersContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowFunctionBody(ctx *base.ArrowFunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionDecl(ctx *base.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnoymousFunctionDecl(ctx *base.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return v.VisitChildren(ctx)
}
