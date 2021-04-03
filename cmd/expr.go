package main

import (
	"log"

	"github.com/padraicbc/gojsp"
)

func (v *visitor) VisitImportExpression(ctx *gojsp.ImportExpressionContext) interface{} {

	log.Println("VisitImportExpression", ctx.GetText())
	return v.VisitChildren(ctx)
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
func (v *visitor) VisitExpressionSequence(ctx *gojsp.ExpressionSequenceContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAssignmentExpression(ctx *gojsp.AssignmentExpressionContext) interface{} {

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
