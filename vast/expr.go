package vast

import (
	"fmt"
	"log"
	"strings"

	"github.com/padraicbc/gojsp/base"
)

type ExpressionSequence struct {
	typeOf   string
	Children []*ExpressionStatement
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

type ThisExpression struct {
	typeOf string
}

// interface ExpressionStatement <: Statement {
//     type: "ExpressionStatement";
//     expression: Expression;
// }

// An expression statement, i.e., a statement consisting of a single expression.

type ExpressionStatement struct {
	*SourceInfo
	OP          Token
	typeOf      string
	Left, Right VNode
	children    VNode
	prev, next  VNode
}

func (i *ExpressionStatement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ExpressionStatement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (e *ExpressionStatement) GetInfo() *SourceInfo {
	return e.SourceInfo
}
func (e *ExpressionStatement) Type() string {
	return "ExpressionStatement"
}
func (i *ExpressionStatement) Code() string {

	return CodeDef(i)
}
func (i *ExpressionStatement) Children() []VNode {

	return children(i.children)
}

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
func (v *Visitor) VisitPropertyExpressionAssignment(ctx *base.PropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitComputedPropertyExpressionAssignment(ctx *base.ComputedPropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpressionStatement(ctx *base.ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

// func (v *Visitor) VisitLeftRightExpression(ctx *base.LeftRightExpressionContext) interface{} {

// 	exp := &ExpressionStatement{
// 		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
// 	}
// 	var prev VNode
// 	for _, ch := range v.VisitChildren(ctx).([]VNode) {
// 		if exp.children == nil {
// 			exp.children = ch
// 		} else {
// 			prev.Next(ch)
// 		}
// 		ch.Prev(prev)
// 		prev = ch
// 		switch ch.(Token).SymbolName() {
// 		case "DecimalLiteral":

// 		case "Identifier":
// 			if exp.Left == nil {
// 				exp.Left = ch.(Token)
// 				continue
// 			}
// 			exp.Right = ch.(Token)
// 		default:
// 			if ch.(Token).RName("") != "singleExpression" {
// 				log.Panicf("%+v\n", ch)

// 			}
// 			switch ch.(Token).Value() {

// 			case "==", "!=", "===", "!==", "<", "<=", ">", ">=", "<<", ">>", ">>>", "+", "-", "*", "/", "%", ",", "^", "&", "in", "instanceof":
// 				exp.typeOf = "BinaryExpression"
// 				exp.OP = ch.(Token)
// 			case "||", "&&":
// 				exp.typeOf = "LogicalExpression"
// 				exp.OP = ch.(Token)
// 			default:
// 				panic(ch)
// 			}

// 		}
// 	}
// 	return exp

// }

func (v *Visitor) VisitExpressionSequence(ctx *base.ExpressionSequenceContext) interface{} {
	// log.Println("VisitExpressionSequence", ctx.GetText())
	return v.VisitChildren(ctx)

}

type AssignmentExpression struct {
	*SourceInfo
	Left       VNode
	OP         Token
	Right      VNode
	children   VNode
	prev, next VNode
}

var _ VNode = (*AssignmentExpression)(nil)

func (i *AssignmentExpression) Type() string {
	return "AssignmentExpression"
}
func (i *AssignmentExpression) Code() string {
	return CodeDef(i)
}
func (i *AssignmentExpression) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *AssignmentExpression) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *AssignmentExpression) Children() []VNode {
	return children(i.children)
}
func (v *Visitor) VisitAssignmentExpression(ctx *base.AssignmentExpressionContext) interface{} {
	// log.Println("VisitAssignmentExpression", ctx.OP.GetText())
	ae := &AssignmentExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode

	for i, ch := range v.VisitChildren(ctx).([]VNode) {
		if ae.children == nil {
			ae.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch
		switch i {
		case 0:
			ae.Left = ch
		case 1:
			ae.OP = ch.(Token)
		case 2:
			ae.Right = ch
		}

	}
	log.Println()
	return ae

}

// Visit a parse tree produced by JavaScriptParser#MultiplicativeExpression.
func (v *Visitor) VisitMultiplicativeExpression(ctx *base.MultiplicativeExpressionContext) interface{} {
	log.Println("VisitMultiplicativeExpression", ctx.GetText())
	exp := &ExpressionStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		typeOf: "MultiplicativeExpression"}
	var prev VNode

	for i, ch := range v.VisitChildren(ctx).([]VNode) {
		if exp.children == nil {
			exp.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch
		switch i {
		case 0:
			exp.Left = ch
		case 1:
			exp.OP = ch.(Token)
			exp.typeOf = exp.OP.Value()
		case 2:
			exp.Right = ch
		}

	}
	return exp

}
func (v *Visitor) VisitArgumentsExpression(ctx *base.ArgumentsExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLiteralExpression(ctx *base.LiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTernaryExpression(ctx *base.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPowerExpression(ctx *base.PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
