package main

import (
	"log"

	"github.com/padraicbc/gojsp"
)

type Expression struct {
	OP          string
	Left, Right string
}

func (v *visitor) VisitPropertyExpressionAssignment(ctx *gojsp.PropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitComputedPropertyExpressionAssignment(ctx *gojsp.ComputedPropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExpressionStatement(ctx *gojsp.ExpressionStatementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitAdditiveExpression(ctx *gojsp.AdditiveExpressionContext) interface{} {
	log.Println(ctx.GetText())
	return Expression{OP: "+", Left: ctx.SingleExpression(0).GetText(), Right: ctx.SingleExpression(1).GetText()}

}
func (v *visitor) VisitExpressionSequence(ctx *gojsp.ExpressionSequenceContext) interface{} {
	log.Println("VisitExpressionSequence", ctx.GetText())

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAssignmentExpression(ctx *gojsp.AssignmentExpressionContext) interface{} {
	log.Println("VisitAssignmentExpression", ctx.GetText())
	return v.VisitChildren(ctx)
}
func (v *visitor) VisitIdentifierExpression(ctx *gojsp.IdentifierExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArgumentsExpression(ctx *gojsp.ArgumentsExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLiteralExpression(ctx *gojsp.LiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
