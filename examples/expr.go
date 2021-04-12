package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/padraicbc/gojsp/parser"
)

type ExpressionSequence struct {
	typeOf   string
	Children []*Expression
	VNode
}

func (e *ExpressionSequence) Type() string {
	return e.typeOf
}
func (e *ExpressionSequence) Code() string {
	var str strings.Builder
	for _, ee := range e.Children {
		str.WriteString(ee.Code())
	}
	return str.String()
}

type Expression struct {
	*SourceInfo
	OP          string
	Left, Right string
	typeOf      string
	VNode
}

func (e *Expression) GetInfo() *SourceInfo {
	return e.SourceInfo
}
func (e *Expression) Type() string {
	return e.typeOf
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

func (v *visitor) VisitLeftRightExpression(ctx *parser.LeftRightExpressionContext) interface{} {
	var typeOf string

	switch ctx.OP.GetText() {
	case "==", "!=", "===", "!==", "<", "<=", ">", ">=", "<<", ">>", ">>>", "+", "-", "*", "/", "%", ",", "^", "&", "in", "instanceof":
		typeOf = "BinaryExpression"

	case "||", "&&":
		typeOf = "LogicalExpression"
	default:
		panic(ctx.OP.GetText())

	}
	return &Expression{
		OP:         ctx.OP.GetText(),
		Left:       ctx.Left.GetText(),
		Right:      ctx.SingleExpression(1).GetText(),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext), typeOf: typeOf}

}

func (v *visitor) VisitExpressionSequence(ctx *parser.ExpressionSequenceContext) interface{} {
	log.Println("VisitExpressionSequence", ctx.GetText())
	v.Nodes = append(v.Nodes, &ExpressionSequence{typeOf: "ExpressionSequence", Children: v.VisitChildren(ctx).([]*Expression)})
	return nil

}
func (v *visitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	log.Println("VisitAssignmentExpression", ctx.OP.GetText())

	return &Expression{
		OP:         ctx.OP.GetText(),
		Left:       ctx.Left.GetText(),
		Right:      ctx.SingleExpression(1).GetText(),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext), typeOf: "AssignmentExpression"}

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

func (v *visitor) VisitTernaryExpression(ctx *parser.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPowerExpression(ctx *parser.PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
