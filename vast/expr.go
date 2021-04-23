package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// one type to catch all left = right expressions
type ExpI interface {
	Left() VNode
	OP() Token
	SetLeft(VNode)
	SetOP(Token)
	SetRight(VNode)
	SetChild(VNode)
	SetType(string)
	Type() string
}
type LRExpression struct {
	*SourceInfo
	left, right VNode
	op          Token
	typeOf      string
	children    VNode
	prev, next  VNode
}

func (i *LRExpression) Left() VNode {
	return i.left
}
func (i *LRExpression) OP() Token {
	return i.op
}
func (i *LRExpression) Right() VNode {
	return i.right
}
func (i *LRExpression) SetLeft(vn VNode) {
	i.left = vn
}
func (i *LRExpression) SetType(s string) {
	i.typeOf = s
}
func (i *LRExpression) SetOP(tk Token) {
	i.op = tk
}
func (i *LRExpression) SetRight(vn VNode) {
	i.right = vn
}
func (i *LRExpression) SetChild(vn VNode) {
	i.children = vn
}
func (i *LRExpression) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *LRExpression) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (e *LRExpression) GetInfo() *SourceInfo {
	return e.SourceInfo
}
func (e *LRExpression) Type() string {
	return e.typeOf
}
func (i *LRExpression) Code() string {

	return CodeDef(i)
}
func (i *LRExpression) Children() []VNode {

	return children(i.children)
}

// simple helper to add tree pointers when not iterating
func setLR(v *Visitor, l, r base.ISingleExpressionContext, par ExpI) {
	left, right := flt(v.Visit(l)), flt(v.Visit(r))
	par.SetChild(left)
	left.Next(par.OP())
	par.OP().Next(right)
	par.OP().Prev(left)
	right.Prev(par.OP())
	par.SetLeft(left)
	par.SetRight(right)
	// not sure we need +"Expression" as it cannot but be one.
	// also maybe best to set general types in each visitor
	par.SetType(par.OP().SymbolName() + "Expression")

}

// temp measure I think as should be concrete single types when all implemented.
func flt(i interface{}) VNode {
	switch i.(type) {
	case []VNode:
		return i.([]VNode)[0]
	case VNode:
		return i.(VNode)
	default:
		log.Panicf("%+v\n", i)

	}
	return nil

}

type ExpressionStatement struct {
	*SourceInfo
	// singlerexpressions
	children   VNode
	prev, next VNode
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

func (v *Visitor) VisitExpressionStatement(ctx *base.ExpressionStatementContext) interface{} {
	exp := &ExpressionStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if exp.children == nil {
			exp.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch

	}
	return exp
}

// expressionSequence
//     : singleExpression (',' singleExpression)*
//     ;
func (v *Visitor) VisitExpressionSequence(ctx *base.ExpressionSequenceContext) interface{} {
	// log.Println("VisitExpressionSequence", ctx.GetText())
	return v.VisitChildren(ctx)

}

// <assoc=right> singleExpression '=' singleExpression

func (v *Visitor) VisitAssignmentExpression(ctx *base.AssignmentExpressionContext) interface{} {
	// log.Println("VisitAssignmentExpression", ctx.OP.GetText())
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	exp.op = ident(v, ctx.Assign().GetSymbol())
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp

}

//     | <assoc=right> singleExpression '**' singleExpression
func (v *Visitor) VisitPowerExpression(ctx *base.PowerExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.Power().GetSymbol())
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)

	return exp
}

// Visit a parse tree produced by JavaScriptParser#MultiplicativeExpression.
func (v *Visitor) VisitMultiplicativeExpression(ctx *base.MultiplicativeExpressionContext) interface{} {
	// log.Println("VisitMultiplicativeExpression", ctx.GetText())
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	if ctx.Multiply() != nil {
		exp.op = ident(v, ctx.Multiply().GetSymbol())

	} else if ctx.Divide() != nil {
		exp.op = ident(v, ctx.Divide().GetSymbol())

	} else if ctx.Modulus() != nil {
		exp.op = ident(v, ctx.Modulus().GetSymbol())

	}
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp

}
func (v *Visitor) VisitAdditiveExpression(ctx *base.AdditiveExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	if ctx.Plus() != nil {
		exp.op = ident(v, ctx.Plus().GetSymbol())

	} else if ctx.Minus() != nil {
		exp.op = ident(v, ctx.Minus().GetSymbol())

	}

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp

}
func (v *Visitor) VisitCoalesceExpression(ctx *base.CoalesceExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.NullCoalesce().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

func (v *Visitor) VisitBitShiftExpression(ctx *base.BitShiftExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	if ctx.LeftShiftArithmetic() != nil {
		exp.op = ident(v, ctx.LeftShiftArithmetic().GetSymbol())
	} else if ctx.RightShiftArithmetic() != nil {
		exp.op = ident(v, ctx.RightShiftArithmetic().GetSymbol())
	} else if ctx.RightShiftLogical() != nil {
		exp.op = ident(v, ctx.RightShiftLogical().GetSymbol())
	}

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

func (v *Visitor) VisitRelationalExpression(ctx *base.RelationalExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	if ctx.LessThan() != nil {
		exp.op = ident(v, ctx.LessThan().GetSymbol())
	} else if ctx.LessThanEquals() != nil {
		exp.op = ident(v, ctx.LessThan().GetSymbol())
	} else if ctx.MoreThan() != nil {
		exp.op = ident(v, ctx.MoreThan().GetSymbol())
	} else if ctx.GreaterThanEquals() != nil {
		exp.op = ident(v, ctx.GreaterThanEquals().GetSymbol())
	}

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp

}

func (v *Visitor) VisitInstanceofExpression(ctx *base.InstanceofExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.Instanceof().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

func (v *Visitor) VisitInExpression(ctx *base.InExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.In().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}
func (v *Visitor) VisitEqualityExpression(ctx *base.EqualityExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	if ctx.NotEquals() != nil {
		exp.op = ident(v, ctx.NotEquals().GetSymbol())
	} else if ctx.Equals_() != nil {
		exp.op = ident(v, ctx.Equals_().GetSymbol())
	} else if ctx.IdentityEquals() != nil {
		exp.op = ident(v, ctx.IdentityEquals().GetSymbol())
	} else if ctx.IdentityNotEquals() != nil {
		exp.op = ident(v, ctx.IdentityNotEquals().GetSymbol())
	}

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}
func (v *Visitor) VisitBitAndExpression(ctx *base.BitAndExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.BitAnd().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

func (v *Visitor) VisitBitXOrExpression(ctx *base.BitXOrExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.BitXOr().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}
func (v *Visitor) VisitBitOrExpression(ctx *base.BitOrExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.BitOr().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

func (v *Visitor) VisitLogicalAndExpression(ctx *base.LogicalAndExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.And().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}
func (v *Visitor) VisitLogicalOrExpression(ctx *base.LogicalOrExpressionContext) interface{} {
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.Or().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}
func (v *Visitor) VisitAssignmentOperator(ctx *base.AssignmentOperatorContext) interface{} {
	// ctx.GetStart().GetLine() etc.. could be  used directly but still have to switch on a lot of types so as easy do this.
	return v.VisitChildren(ctx).([]VNode)[0].(Token)
}
func (v *Visitor) VisitAssignmentOperatorExpression(ctx *base.AssignmentOperatorExpressionContext) interface{} {
	// log.Println("VisitAssignmentOperatorExpression")
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// ctx.AssignmentOperator().GetText() but we want source info?
	exp.op = v.Visit(ctx.AssignmentOperator()).(Token)
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}
func (v *Visitor) VisitTernaryExpression(ctx *base.TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgumentsExpression(ctx *base.ArgumentsExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLiteralExpression(ctx *base.LiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyExpressionAssignment(ctx *base.PropertyExpressionAssignmentContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitComputedPropertyExpressionAssignment(ctx *base.ComputedPropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
} // An expression statement, i.e., a statement consisting of a single expression.
