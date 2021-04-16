package main

import (
	"log"

	"github.com/padraicbc/gojsp/parser"
)

func (v *visitor) VisitPropertyExpressionAssignment(ctx *parser.PropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitComputedPropertyExpressionAssignment(ctx *parser.ComputedPropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLeftRightExpression(ctx *parser.LeftRightExpressionContext) interface{} {
	return v.VisitChildren(ctx)

}

func (v *visitor) VisitExpressionSequence(ctx *parser.ExpressionSequenceContext) interface{} {
	// log.Println("VisitExpressionSequence", ctx.GetText())
	return v.VisitChildren(ctx)

}
func (v *visitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	log.Println("VisitAssignmentExpression", ctx.OP.GetText())

	return v.VisitChildren(ctx)

}

func (v *visitor) VisitArgumentsExpression(ctx *parser.ArgumentsExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitTernaryExpression(ctx *parser.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPowerExpression(ctx *parser.PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
