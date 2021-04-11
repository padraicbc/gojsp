package main

import (
	"fmt"
	"log"

	antlr "github.com/padraicbc/antlr4"

	"github.com/padraicbc/gojsp/parser"
)

type Expression struct {
	*SourceInfo
	expr        string
	OP          string
	Left, Right string
}

func (e *Expression) Type() string {
	return e.expr
}
func (i *Expression) Code() string {
	if i == nil {
		return ""
	}

	return fmt.Sprintf("%s %s %s", i.Left, i.OP, i.Right)
}
func (v *visitor) VisitPropertyExpressionAssignment(ctx *parser.PropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitComputedPropertyExpressionAssignment(ctx *parser.ComputedPropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	return &Expression{OP: ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText(), Left: ctx.SingleExpression(0).GetText(),
		Right: ctx.SingleExpression(1).GetText(), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext), expr: "AdditiveExpression"}

}
func (v *visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return &Expression{OP: ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText(), Left: ctx.SingleExpression(0).GetText(),
		Right: ctx.SingleExpression(1).GetText(), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext), expr: "MultiplicativeExpression"}
}

func (v *visitor) VisitExpressionSequence(ctx *parser.ExpressionSequenceContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	log.Println("VisitAssignmentExpression", ctx.GetText())
	return v.VisitChildren(ctx)
}
func (v *visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArgumentsExpression(ctx *parser.ArgumentsExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
