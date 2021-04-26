package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// labelledStatement
//     : identifier ':' statement
//     ;
type LabeledStatement struct {
	*SourceInfo
	Statement  VNode
	Label      Token
	Colon      Token
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*LabeledStatement)(nil)

func (i *LabeledStatement) Next() VNode {

	return i.next
}
func (i *LabeledStatement) SetNext(v VNode) {
	i.next = v
}
func (i *LabeledStatement) Prev() VNode {

	return i.prev
}
func (i *LabeledStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *LabeledStatement) Type() string {
	return "LabeledStatement"
}
func (i *LabeledStatement) Code() string {
	return CodeDef(i)
}
func (i *LabeledStatement) FirstChild() VNode {

	return i.firstChild

}

// special case for $: ... todo: a type
func (v *Visitor) VisitLabelledStatement(ctx *base.LabelledStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitLabelledStatement", ctx.GetText())
	}

	if ctx.Identifier().GetText() == "$" {
		// log.Println("Reactive?")
	}
	lst := &LabeledStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if lst.firstChild == nil {
			lst.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch
		switch ch.Type() {
		case "LToken":
			t := ch.(Token)
			if t.SymbolName() == "Identifier" {
				lst.Label = t
				continue
			}
			if t.SymbolName() == "Colon" {
				lst.Colon = t
				continue
			}
			//  Statement can be idenifier also?
			panic(ch)

		case "Block":
			lst.Statement = ch
		default:
			panic(ch.Type())

		}

	}
	return lst
}

// block
//     : '{' statementList? '}'
//     ;
type Block struct {
	*SourceInfo
	// *StatementList
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*Block)(nil)

func (i *Block) Next() VNode {

	return i.next
}
func (i *Block) SetNext(v VNode) {
	i.next = v
}
func (i *Block) Prev() VNode {

	return i.prev
}
func (i *Block) SetPrev(v VNode) {
	i.prev = v
}

func (i *Block) Type() string {
	return "Block"
}
func (i *Block) Code() string {
	return CodeDef(i)
}
func (i *Block) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitBlock(ctx *base.BlockContext) interface{} {
	if v.Debug {
		log.Println("VisitBlock", ctx.GetText())
	}
	b := &Block{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if b.firstChild == nil {
			b.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch
	}
	return b
}

// expressionStatement
//     : {p.notOpenBraceAndNotFunction()}? expressionSequence eos
//     ;
type ExpressionStatement struct {
	*SourceInfo
	// singlexpression(s)
	firstChild VNode
	Eos        Token

	prev, next VNode
}

func (i *ExpressionStatement) Next() VNode {

	return i.next
}
func (i *ExpressionStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ExpressionStatement) Prev() VNode {

	return i.prev
}
func (i *ExpressionStatement) SetPrev(v VNode) {
	i.prev = v
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
func (i *ExpressionStatement) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitExpressionStatement(ctx *base.ExpressionStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitExpressionStatement", ctx.GetText(), ctx.GetChildCount())
	}
	exp := &ExpressionStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if exp.firstChild == nil {
			exp.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch

		if tk, ok := ch.(Token); ok {
			exp.Eos = tk

		}

	}
	return exp
}

// expressionSequence
//     : singleExpression (',' singleExpression)*
//     ;
type ExpressionSequence struct {
	*SourceInfo
	// 1-n singleexpressions
	Commas     []Token
	firstChild VNode

	prev, next VNode
}

func (i *ExpressionSequence) Next() VNode {

	return i.next
}
func (i *ExpressionSequence) SetNext(v VNode) {
	i.next = v
}
func (i *ExpressionSequence) Prev() VNode {

	return i.prev
}
func (i *ExpressionSequence) SetPrev(v VNode) {
	i.prev = v
}

func (e *ExpressionSequence) GetInfo() *SourceInfo {
	return e.SourceInfo
}
func (e *ExpressionSequence) Type() string {
	return "ExpressionSequence"
}
func (i *ExpressionSequence) Code() string {

	return CodeDef(i)
}
func (i *ExpressionSequence) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitExpressionSequence(ctx *base.ExpressionSequenceContext) interface{} {
	if v.Debug {
		log.Println("VisitExpressionSequence", ctx.GetText())
	}
	e := &ExpressionSequence{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		if e.firstChild == nil {
			e.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch
		switch ch.Type() {
		case "LToken":
			// always this?
			e.Commas = append(e.Commas, ch.(Token))

			// todo: leave as is or add specific SinglesExpressions?
			// also maybe type all singleExpressions.
		default:
			// log.Println(ch.Type())
			// ArrowFunction
			// PlusExpression

		}

	}
	return e

}

// emptyStatement_
//     : SemiColon
//     ;
func (v *Visitor) VisitEmptyStatement_(ctx *base.EmptyStatement_Context) interface{} {
	return ident(v, ctx.SemiColon().GetSymbol())
}

// ifStatement
//     : If '(' expressionSequence ')' statement (Else statement)?
//     ;
func (v *Visitor) VisitIfStatement(ctx *base.IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

// iterationStatement
//     : Do statement While '(' expressionSequence ')' eos                                                                       # DoStatement
//    ....
type DoStatement struct {
	*SourceInfo
	Do                 Token
	Statement          VNode
	While              Token
	OpenParen          Token
	CloseParen         Token
	ExpressionSequence *ExpressionSequence

	firstChild VNode

	prev, next VNode
}

func (i *DoStatement) Next() VNode {

	return i.next
}
func (i *DoStatement) SetNext(v VNode) {
	i.next = v
}
func (i *DoStatement) Prev() VNode {

	return i.prev
}
func (i *DoStatement) SetPrev(v VNode) {
	i.prev = v
}

func (e *DoStatement) GetInfo() *SourceInfo {
	return e.SourceInfo
}
func (e *DoStatement) Type() string {
	return "DoStatement"
}
func (i *DoStatement) Code() string {

	return CodeDef(i)
}
func (i *DoStatement) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitDoStatement(ctx *base.DoStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitDoStatement", ctx.GetText())
	}
	d := &DoStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	_ = d
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWhileStatement(ctx *base.WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForStatement(ctx *base.ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForInStatement(ctx *base.ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForOfStatement(ctx *base.ForOfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitContinueStatement(ctx *base.ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBreakStatement(ctx *base.BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

// returnStatement
//     : Return ({p.notLineTerminator()}? expressionSequence)? eos
//     ;
type ReturnStatement struct {
	*SourceInfo
	Return Token

	ExpSeq *ExpressionSequence
	Eos    Token

	firstChild VNode

	prev, next VNode
}

func (i *ReturnStatement) Next() VNode {

	return i.next
}
func (i *ReturnStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ReturnStatement) Prev() VNode {

	return i.prev
}
func (i *ReturnStatement) SetPrev(v VNode) {
	i.prev = v
}

func (e *ReturnStatement) GetInfo() *SourceInfo {
	return e.SourceInfo
}
func (e *ReturnStatement) Type() string {
	return "ReturnStatement"
}
func (i *ReturnStatement) Code() string {

	return CodeDef(i)
}
func (i *ReturnStatement) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitReturnStatement(ctx *base.ReturnStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitReturnStatement", ctx.GetText())
	}
	r := &ReturnStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if r.firstChild == nil {
			r.firstChild = ch
		} else {
			prev.SetNext(ch)

		}

		ch.SetPrev(prev)

		prev = ch

		switch ch.Type() {
		case "LToken":
			tk := ch.(Token)
			switch tk.SymbolName() {
			case "Return":
				r.Return = tk

			case "SemiColon":

				r.Eos = tk

			default:
				log.Panicf("%+v\n", ch)
			}
		case "ExpressionSequence":
			r.ExpSeq = ch.(*ExpressionSequence)

		default:

			log.Panicf("%+v %s\n", ch, ch.Type())

		}

	}
	return r
}

func (v *Visitor) VisitYieldStatement(ctx *base.YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWithStatement(ctx *base.WithStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSwitchStatement(ctx *base.SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}
