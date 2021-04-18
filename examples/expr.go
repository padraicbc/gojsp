package main

import (
	"log"

	"github.com/padraicbc/gojsp/parser"
)

func (v *Visitor) VisitPropertyExpressionAssignment(ctx *parser.PropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitComputedPropertyExpressionAssignment(ctx *parser.ComputedPropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLeftRightExpression(ctx *parser.LeftRightExpressionContext) interface{} {
	return v.VisitChildren(ctx)

}

func (v *Visitor) VisitExpressionSequence(ctx *parser.ExpressionSequenceContext) interface{} {
	// log.Println("VisitExpressionSequence", ctx.GetText())
	return v.VisitChildren(ctx)

}
func (v *Visitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	log.Println("VisitAssignmentExpression", ctx.OP.GetText())

	return v.VisitChildren(ctx)

}

func (v *Visitor) VisitArgumentsExpression(ctx *parser.ArgumentsExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTernaryExpression(ctx *parser.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPowerExpression(ctx *parser.PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
