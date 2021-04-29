package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// simple helper to add tree pointers when not iterating
func setLR(v *Visitor, l, r base.ISingleExpressionContext, par ExpI) {
	left, right := flt(v.Visit(l)), flt(v.Visit(r))
	par.SetChild(left)
	left.SetNext(par.OP())
	par.OP().SetNext(right)
	par.OP().SetPrev(left)
	right.SetPrev(par.OP())
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
	firstChild  VNode
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
	i.firstChild = vn
}
func (i *LRExpression) Next() VNode {

	return i.next
}
func (i *LRExpression) SetNext(v VNode) {
	i.next = v
}
func (i *LRExpression) Prev() VNode {

	return i.prev
}
func (i *LRExpression) SetPrev(v VNode) {
	i.prev = v
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
func (i *LRExpression) FirstChild() VNode {

	return i.firstChild

}

// <assoc=right> singleExpression '=' singleExpression
func (v *Visitor) VisitAssignmentExpression(ctx *base.AssignmentExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitAssignmentExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	exp.op = ident(v, ctx.Assign().GetSymbol())
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp

}

//     | <assoc=right> singleExpression '**' singleExpression
func (v *Visitor) VisitPowerExpression(ctx *base.PowerExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitPowerExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.Power().GetSymbol())
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)

	return exp
}

// Visit a parse tree produced by JavaScriptParser#MultiplicativeExpression.
// singleExpression ('*' | '/' | '%') singleExpression
func (v *Visitor) VisitMultiplicativeExpression(ctx *base.MultiplicativeExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitMultiplicativeExpression", ctx.GetText())
	}
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

// singleExpression ('+' | '-') singleExpression
func (v *Visitor) VisitAdditiveExpression(ctx *base.AdditiveExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitAdditiveExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	if ctx.Plus() != nil {
		exp.op = ident(v, ctx.Plus().GetSymbol())

	} else if ctx.Minus() != nil {
		exp.op = ident(v, ctx.Minus().GetSymbol())

	}

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp

}

// singleExpression '??' singleExpression
func (v *Visitor) VisitCoalesceExpression(ctx *base.CoalesceExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitCoalesceExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.NullCoalesce().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// singleExpression ('<<' | '>>' | '>>>') singleExpression
func (v *Visitor) VisitBitShiftExpression(ctx *base.BitShiftExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitBitShiftExpression", ctx.GetText())
	}
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

// singleExpression ('<' | '>' | '<=' | '>=') singleExpression
func (v *Visitor) VisitRelationalExpression(ctx *base.RelationalExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitRelationalExpression", ctx.GetText())
	}
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

// singleExpression Instanceof singleExpression
func (v *Visitor) VisitInstanceofExpression(ctx *base.InstanceofExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitInstanceofExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.Instanceof().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// singleExpression In singleExpression
func (v *Visitor) VisitInExpression(ctx *base.InExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitInExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.In().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// singleExpression ('==' | '!=' | '===' | '!==') singleExpression
func (v *Visitor) VisitEqualityExpression(ctx *base.EqualityExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitEqualityExpression", ctx.GetText())
	}
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

//  singleExpression '&' singleExpression
func (v *Visitor) VisitBitAndExpression(ctx *base.BitAndExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitBitAndExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.BitAnd().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// singleExpression '^' singleExpression
func (v *Visitor) VisitBitXOrExpression(ctx *base.BitXOrExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitBitXOrExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.BitXOr().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// singleExpression '|' singleExpression
func (v *Visitor) VisitBitOrExpression(ctx *base.BitOrExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitBitOrExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.BitOr().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// singleExpression '&&' singleExpression
func (v *Visitor) VisitLogicalAndExpression(ctx *base.LogicalAndExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitLogicalAndExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.And().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

//  singleExpression '||' singleExpression
func (v *Visitor) VisitLogicalOrExpression(ctx *base.LogicalOrExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitLogicalOrExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	exp.op = ident(v, ctx.Or().GetSymbol())

	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

// <assoc=right> singleExpression assignmentOperator singleExpression
func (v *Visitor) VisitAssignmentOperatorExpression(ctx *base.AssignmentOperatorExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitAssignmentOperatorExpression", ctx.GetText())
	}
	exp := &LRExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	// ident...
	exp.op = v.VisitAssignmentOperator(ctx.AssignmentOperator().(*base.AssignmentOperatorContext)).(Token)
	setLR(v, ctx.SingleExpression(0), ctx.SingleExpression(1), exp)
	return exp
}

//  singleExpression '?' singleExpression ':' singleExpression
type TernaryExpression struct {
	*SourceInfo
	Test         VNode
	Colon        Token
	Alternate    VNode
	QuestionMark Token
	Consequent   VNode
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*TernaryExpression)(nil)

func (i *TernaryExpression) Type() string {
	return "TernaryExpression"
}
func (i *TernaryExpression) Code() string {
	return CodeDef(i)
}
func (i *TernaryExpression) Next() VNode {
	return i.next
}
func (i *TernaryExpression) SetNext(v VNode) {
	i.next = v
}
func (i *TernaryExpression) Prev() VNode {
	return i.prev
}
func (i *TernaryExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *TernaryExpression) FirstChild() VNode {
	return i.firstChild
}

//  singleExpression '?' singleExpression ':' singleExpression
func (v *Visitor) VisitTernaryExpression(ctx *base.TernaryExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitTernaryExpression", ctx.GetText())
	}
	te := &TernaryExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	te.Test = v.Visit(ctx.SingleExpression(0)).(VNode)
	te.firstChild = te.Test
	te.Consequent = v.Visit(ctx.SingleExpression(1)).(VNode)
	te.Alternate = v.Visit(ctx.SingleExpression(2)).(VNode)
	te.Colon = ident(v, ctx.Colon().GetSymbol())
	te.QuestionMark = ident(v, ctx.QuestionMark().GetSymbol())
	setAllSibs(te.Test, te.QuestionMark, te.Consequent, te.Colon, te.Alternate)

	return te

}

// singleExpression arguments
type ArgumentsExpression struct {
	*SourceInfo
	SingleExp  VNode
	Arguments  *Arguments
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ArgumentsExpression)(nil)

func (i *ArgumentsExpression) Type() string {
	return "ArgumentsExpression"
}
func (i *ArgumentsExpression) Code() string {
	return CodeDef(i)
}
func (i *ArgumentsExpression) Next() VNode {
	return i.next
}
func (i *ArgumentsExpression) SetNext(v VNode) {
	i.next = v
}
func (i *ArgumentsExpression) Prev() VNode {
	return i.prev
}
func (i *ArgumentsExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *ArgumentsExpression) FirstChild() VNode {
	return i.firstChild
}

//  singleExpression arguments
func (v *Visitor) VisitArgumentsExpression(ctx *base.ArgumentsExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitArgumentsExpression", ctx.GetText())
	}

	a := &ArgumentsExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	a.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
	a.firstChild = a.SingleExp
	a.Arguments = v.VisitArguments(ctx.Arguments().(*base.ArgumentsContext)).(*Arguments)
	setAllSibs(a.SingleExp, a.Arguments)

	return a
}

// propertyAssignment
//     : propertyName ':' singleExpression
type PropertyExpressionAssignment struct {
	*SourceInfo
	PropertyName *PropertyName
	SingleExp    VNode
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*PropertyExpressionAssignment)(nil)

func (i *PropertyExpressionAssignment) Type() string {
	return "PropertyExpressionAssignment"
}
func (i *PropertyExpressionAssignment) Code() string {
	return CodeDef(i)
}
func (i *PropertyExpressionAssignment) Next() VNode {
	return i.next
}
func (i *PropertyExpressionAssignment) SetNext(v VNode) {
	i.next = v
}
func (i *PropertyExpressionAssignment) Prev() VNode {
	return i.prev
}
func (i *PropertyExpressionAssignment) SetPrev(v VNode) {
	i.prev = v
}
func (i *PropertyExpressionAssignment) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPropertyExpressionAssignment(ctx *base.PropertyExpressionAssignmentContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertyExpressionAssignment", ctx.GetText())
	}
	pa := &PropertyExpressionAssignment{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	pa.PropertyName = v.VisitPropertyName(ctx.PropertyName().(*base.PropertyNameContext)).(*PropertyName)
	pa.firstChild = pa.PropertyName
	pa.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)

	setAllSibs(pa.PropertyName, pa.SingleExp)

	return pa
}

// [' singleExpression ']' ':' singleExpression
type ComputedPropertyExpressionAssignment struct {
	*SourceInfo
	OpenBracket  Token
	Key          VNode
	Colon        Token
	Val          VNode
	CloseBracket Token
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*ComputedPropertyExpressionAssignment)(nil)

func (i *ComputedPropertyExpressionAssignment) Type() string {
	return "ComputedPropertyExpressionAssignment"
}
func (i *ComputedPropertyExpressionAssignment) Code() string {
	return CodeDef(i)
}
func (i *ComputedPropertyExpressionAssignment) Next() VNode {
	return i.next
}
func (i *ComputedPropertyExpressionAssignment) SetNext(v VNode) {
	i.next = v
}
func (i *ComputedPropertyExpressionAssignment) Prev() VNode {
	return i.prev
}
func (i *ComputedPropertyExpressionAssignment) SetPrev(v VNode) {
	i.prev = v
}
func (i *ComputedPropertyExpressionAssignment) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitComputedPropertyExpressionAssignment(ctx *base.ComputedPropertyExpressionAssignmentContext) interface{} {
	if v.Debug {
		log.Println("VisitComputedPropertyExpressionAssignment", ctx.GetText())
	}
	cp := &ComputedPropertyExpressionAssignment{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	cp.OpenBracket = ident(v, ctx.OpenBracket().GetSymbol())
	cp.Key = v.Visit(ctx.SingleExpression(0)).(VNode)
	cp.CloseBracket = ident(v, ctx.CloseBracket().GetSymbol())
	cp.Colon = ident(v, ctx.Colon().GetSymbol())
	cp.Val = v.Visit(ctx.SingleExpression(1)).(VNode)
	cp.firstChild = cp.OpenBracket
	setAllSibs(cp.OpenBracket, cp.Key, cp.CloseBracket, cp.Colon, cp.Val)

	return cp
}

// Single node
func (v *Visitor) VisitLiteralExpression(ctx *base.LiteralExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitLiteralExpression", ctx.GetText())
	}
	return v.VisitLiteral(ctx.Literal().(*base.LiteralContext))
}

// identifierExpression
// 	identifier
//     : Identifier
//     | Async
//     ;
func (v *Visitor) VisitIdentifierExpression(ctx *base.IdentifierExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitIdentifierExpression", ctx.GetText())
	}
	// Token
	return v.VisitIdentifier(ctx.Identifier().(*base.IdentifierContext))

}
