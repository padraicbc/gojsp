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
		}
		prev = setSib(prev, ch)

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
		}
		prev = setSib(prev, ch)

	}
	return b
}

// expressionStatement
//     : {p.notOpenBraceAndNotFunction()}? expressionSequence eos
//     ;
type ExpressionStatement struct {
	*SourceInfo
	// singlexpression(s)
	ExpressionSequence *ExpressionSequence
	firstChild         VNode
	Eos                Token

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
	exp := &ExpressionStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	exp.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	exp.firstChild = exp.ExpressionSequence

	if tk, ok := v.VisitEos(ctx.Eos().(*base.EosContext)).(Token); ok {
		exp.Eos = tk
		exp.ExpressionSequence.SetNext(tk)
		tk.SetPrev(exp.ExpressionSequence)

	}

	return exp
}

// expressionSequence
//     : singleExpression (',' singleExpression)*
//     ;
type ExpressionSequence struct {
	*SourceInfo
	// 1-n singleexpressions
	Commas      []Token
	Expressions []VNode
	firstChild  VNode

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
		}

		prev = setSib(prev, ch)

		switch rr := ch.(type) {
		case *LToken:
			// always this?
			e.Commas = append(e.Commas, rr)
			// todo: leave as is or add specific SinglesExpressions?
			// also maybe type all singleExpressions.
		case VNode:
			e.Expressions = append(e.Expressions, ch)
		default:
			log.Panic(ch.Type())
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
type IfStatement struct {
	*SourceInfo
	If         Token
	OpenParen  Token
	Test       *ExpressionSequence
	CloseParen Token
	Consequent VNode
	Else       Token
	Alternate  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*IfStatement)(nil)

func (i *IfStatement) Type() string {
	return "IfStatement"
}
func (i *IfStatement) Code() string {
	return CodeDef(i)
}
func (i *IfStatement) Next() VNode {
	return i.next
}
func (i *IfStatement) SetNext(v VNode) {
	i.next = v
}
func (i *IfStatement) Prev() VNode {
	return i.prev
}
func (i *IfStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *IfStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitIfStatement(ctx *base.IfStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitIfStatement", ctx.GetText())
	}
	ifs := &IfStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	ifs.If = ident(v, ctx.If().GetSymbol())
	ifs.firstChild = ifs.If
	ifs.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	ifs.Test = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	ifs.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	ifs.Consequent = v.Visit(ctx.Statement(0)).(VNode)
	if ctx.Else() != nil {
		ifs.Else = ident(v, ctx.Else().GetSymbol())
	}
	if st2 := ctx.Statement(1); st2 != nil {
		ifs.Alternate = v.Visit(st2).(VNode)
	}

	setAllSibs(ifs.If, ifs.OpenParen, ifs.Test, ifs.CloseParen, ifs.Consequent, ifs.Else, ifs.Alternate)
	return ifs
}

//  Do statement While '(' expressionSequence ')' eos                                                                       # DoStatement
type DoStatement struct {
	*SourceInfo
	Do                 Token
	Statement          VNode
	While              Token
	OpenParen          Token
	ExpressionSequence *ExpressionSequence
	CloseParen         Token
	Eos                Token

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
	d.Do = ident(v, ctx.Do().GetSymbol())
	d.firstChild = d.Do
	d.Statement = v.Visit(ctx.Statement()).(VNode)
	d.While = ident(v, ctx.While().GetSymbol())
	d.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	d.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	d.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	if ctx.Eos() != nil {
		if tk, ok := v.VisitEos(ctx.Eos().(*base.EosContext)).(Token); ok {
			d.Eos = tk
		}
	}
	setAllSibs(d.Do, d.Statement, d.While, d.OpenParen, d.ExpressionSequence, d.CloseParen, d.Eos)

	return d
}

// continueStatement
//     : Continue ({p.notLineTerminator()}? identifier)? eos
//     ;
type ContinueStatement struct {
	*SourceInfo
	Continue   Token
	Identifier Token
	Eos        Token

	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ContinueStatement)(nil)

func (i *ContinueStatement) Type() string {
	return "ContinueStatement"
}
func (i *ContinueStatement) Code() string {
	return CodeDef(i)
}
func (i *ContinueStatement) Next() VNode {
	return i.next
}
func (i *ContinueStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ContinueStatement) Prev() VNode {
	return i.prev
}
func (i *ContinueStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ContinueStatement) FirstChild() VNode {
	return i.firstChild
}

func (v *Visitor) VisitContinueStatement(ctx *base.ContinueStatementContext) interface{} {
	c := &ContinueStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		Continue:   ident(v, ctx.Continue().GetSymbol())}

	c.firstChild = c.Continue
	if ctx.Identifier() != nil {
		c.Identifier = v.VisitIdentifier(ctx.Identifier().(*base.IdentifierContext)).(Token)

	}
	if ctx.Eos() != nil {
		if tk, ok := v.VisitEos(ctx.Eos().(*base.EosContext)).(Token); ok {
			c.Eos = tk

		}

	}
	setAllSibs(c.Continue, c.Identifier, c.Eos)

	return c
}

// While '(' expressionSequence ')' statement
type WhileStatement struct {
	*SourceInfo
	While              Token
	OpenParen          Token
	ExpressionSequence *ExpressionSequence
	Statement          VNode
	CloseParen         Token
	firstChild         VNode
	prev, next         VNode
}

var _ VNode = (*WhileStatement)(nil)

func (i *WhileStatement) Type() string {
	return "WhileStatement"
}
func (i *WhileStatement) Code() string {
	return CodeDef(i)
}
func (i *WhileStatement) Next() VNode {
	return i.next
}
func (i *WhileStatement) SetNext(v VNode) {
	i.next = v
}
func (i *WhileStatement) Prev() VNode {
	return i.prev
}
func (i *WhileStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *WhileStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitWhileStatement(ctx *base.WhileStatementContext) interface{} {
	wh := &WhileStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		While:      ident(v, ctx.While().GetSymbol()),
	}
	wh.firstChild = wh.While

	wh.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	wh.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	wh.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	wh.Statement = v.VisitStatement(ctx.Statement().(*base.StatementContext)).(VNode)
	setAllSibs(wh.While, wh.OpenParen, wh.ExpressionSequence, wh.CloseParen, wh.Statement)
	return wh
}

// For '(' (expressionSequence | variableDeclarationList)? ';' expressionSequence? ';' expressionSequence? ')' statement
type ForStatement struct {
	*SourceInfo
	For       Token
	OpenParen Token
	// Init VNode  for both?
	VariableDeclarationList *VariableDeclarationList
	ExpressionSequence      *ExpressionSequence

	SemiCol1 Token
	SemiCol2 Token

	CloseParen Token

	Test       *ExpressionSequence
	Update     *ExpressionSequence
	Body       *Statement
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ForStatement)(nil)

func (i *ForStatement) Type() string {
	return "ForStatement"
}
func (i *ForStatement) Code() string {
	return CodeDef(i)
}
func (i *ForStatement) Next() VNode {
	return i.next
}
func (i *ForStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ForStatement) Prev() VNode {
	return i.prev
}
func (i *ForStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ForStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitForStatement(ctx *base.ForStatementContext) interface{} {
	f := &ForStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	f.For = ident(v, ctx.For().GetSymbol())
	f.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	if ctx.VariableDeclarationList() != nil {
		f.VariableDeclarationList = v.VisitVariableDeclarationList(
			ctx.VariableDeclarationList().(*base.VariableDeclarationListContext)).(*VariableDeclarationList)
		// use test as way to differentiate expressions
	} else if ctx.Test != nil {
		f.ExpressionSequence = v.VisitExpressionSequence(
			ctx.Test.(*base.ExpressionSequenceContext)).(*ExpressionSequence)

	}
	if ctx.Update != nil {
		f.Update = v.VisitExpressionSequence(
			ctx.Update.(*base.ExpressionSequenceContext)).(*ExpressionSequence)

	}

	f.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	// should neevr be nil?
	if s1 := ctx.SemiColon(0); s1 != nil {
		f.SemiCol1 = ident(v, s1.GetSymbol())
	}
	if s1 := ctx.SemiColon(1); s1 != nil {
		f.SemiCol1 = ident(v, s1.GetSymbol())
	}

	f.Body = v.VisitStatement(ctx.Body.(*base.StatementContext)).(*Statement)

	setAllSibs(
		f.For,
		f.OpenParen,
		f.ExpressionSequence,
		f.VariableDeclarationList,
		f.SemiCol1,
		f.Test,
		f.SemiCol2,
		f.Update,
		f.CloseParen,
		f.Body)
	return f
}

// For '(' (singleExpression | variableDeclarationList) In expressionSequence ')' statement
type ForInStatement struct {
	*SourceInfo
	For       Token
	OpenParen Token

	// Left for both?
	SingleExp               VNode
	VariableDeclarationList *VariableDeclarationList

	In Token
	// Right?
	ExpressionSequence *ExpressionSequence
	CloseParen         Token

	Body       *Statement
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ForInStatement)(nil)

func (i *ForInStatement) Type() string {
	return "ForInStatement"
}
func (i *ForInStatement) Code() string {
	return CodeDef(i)
}
func (i *ForInStatement) Next() VNode {
	return i.next
}
func (i *ForInStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ForInStatement) Prev() VNode {
	return i.prev
}
func (i *ForInStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ForInStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitForInStatement(ctx *base.ForInStatementContext) interface{} {
	f := &ForInStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	f.For = ident(v, ctx.For().GetSymbol())
	f.firstChild = f.For
	f.OpenParen = ident(v, ctx.OpenParen().GetSymbol())

	if ctx.SingleExpression() != nil {
		f.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
	} else if ctx.VariableDeclarationList() != nil {
		f.VariableDeclarationList = v.VisitVariableDeclarationList(
			ctx.VariableDeclarationList().(*base.VariableDeclarationListContext)).(*VariableDeclarationList)
	}
	f.In = ident(v, ctx.In().GetSymbol())
	f.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	f.CloseParen = ident(v, ctx.CloseParen().GetSymbol())

	f.Body = v.VisitStatement(ctx.Statement().(*base.StatementContext)).(*Statement)
	setAllSibs(f.For,
		f.OpenParen,
		f.SingleExp, f.VariableDeclarationList, f.In, f.ExpressionSequence,
		f.CloseParen,
		f.Body)
	return f
}

// todo: check this works as "strange, 'of' is an identifier. and p.p("of") not work in sometime." notr in grammar
// For Await? '(' (singleExpression | variableDeclarationList) identifier{p.p("of")}? expressionSequence ')' statement
type ForOfStatement struct {
	*SourceInfo
	For                     Token
	Await                   Token
	OpenParen               Token
	SingleExp               VNode
	VariableDeclarationList *VariableDeclarationList
	Of                      Token
	ExpressionSequence      *ExpressionSequence
	CloseParen              Token
	Body                    *Statement
	firstChild              VNode
	prev, next              VNode
}

var _ VNode = (*ForOfStatement)(nil)

func (i *ForOfStatement) Type() string {
	return "ForOfStatement"
}
func (i *ForOfStatement) Code() string {
	return CodeDef(i)
}
func (i *ForOfStatement) Next() VNode {
	return i.next
}
func (i *ForOfStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ForOfStatement) Prev() VNode {
	return i.prev
}
func (i *ForOfStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ForOfStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitForOfStatement(ctx *base.ForOfStatementContext) interface{} {
	f := &ForOfStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	f.For = ident(v, ctx.For().GetSymbol())
	f.firstChild = f.For
	if ctx.Await() != nil {
		f.Await = ident(v, ctx.Await().GetSymbol())
	}

	f.OpenParen = ident(v, ctx.OpenParen().GetSymbol())

	if ctx.SingleExpression() != nil {
		f.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
	} else if ctx.VariableDeclarationList() != nil {
		f.VariableDeclarationList = v.VisitVariableDeclarationList(
			ctx.VariableDeclarationList().(*base.VariableDeclarationListContext)).(*VariableDeclarationList)
	}

	f.Of = v.VisitIdentifier(ctx.Identifier().(*base.IdentifierContext)).(Token)
	f.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	f.CloseParen = ident(v, ctx.CloseParen().GetSymbol())

	f.Body = v.VisitStatement(ctx.Statement().(*base.StatementContext)).(*Statement)
	setAllSibs(f.For,
		f.Await,
		f.OpenParen,
		f.SingleExp, f.VariableDeclarationList, f.Of, f.ExpressionSequence,
		f.CloseParen,
		f.Body)
	return f
}

// breakStatement
//     : Break ({p.notLineTerminator()}? identifier)? eos
//     ;
type BreakStatement struct {
	*SourceInfo
	Break      Token
	Identifier Token
	Eos        Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*BreakStatement)(nil)

func (i *BreakStatement) Type() string {
	return "BreakStatement"
}
func (i *BreakStatement) Code() string {
	return CodeDef(i)
}
func (i *BreakStatement) Next() VNode {
	return i.next
}
func (i *BreakStatement) SetNext(v VNode) {
	i.next = v
}
func (i *BreakStatement) Prev() VNode {
	return i.prev
}
func (i *BreakStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *BreakStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitBreakStatement(ctx *base.BreakStatementContext) interface{} {
	b := &BreakStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	b.Break = ident(v, ctx.Break().GetSymbol())
	b.firstChild = b.Break
	if ctx.Identifier() != nil {
		b.Identifier = v.VisitIdentifier(ctx.Identifier().(*base.IdentifierContext)).(Token)
	}
	if ctx.Eos() != nil {
		if tk, ok := v.VisitEos(ctx.Eos().(*base.EosContext)).(Token); ok {
			b.Eos = tk
		}
	}
	setAllSibs(b.Break, b.Identifier, b.Eos)
	return b
}

// returnStatement
//     : Return ({p.notLineTerminator()}? expressionSequence)? eos
//     ;
type ReturnStatement struct {
	*SourceInfo
	Return Token

	ExpressionSequence *ExpressionSequence
	Eos                Token

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

// returnStatement
//     : Return ({p.notLineTerminator()}? expressionSequence)? eos
//     ;
func (v *Visitor) VisitReturnStatement(ctx *base.ReturnStatementContext) interface{} {
	if v.Debug {
		log.Println("VisitReturnStatement", ctx.GetText())
	}
	r := &ReturnStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	r.Return = ident(v, ctx.Return().GetSymbol())
	r.firstChild = r.Return
	if ctx.ExpressionSequence() != nil {
		r.ExpressionSequence = v.VisitExpressionSequence(
			ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)

	}
	if ctx.Eos() != nil {
		if tk, ok := v.VisitEos(ctx.Eos().(*base.EosContext)).(Token); ok {
			r.Eos = tk
		}
	}
	setAllSibs(r.Return, r.ExpressionSequence, r.Eos)

	return r
}

// yieldStatement
// : Yield ({p.notLineTerminator()}? expressionSequence)? eos
// ;
type YieldStatement struct {
	*SourceInfo
	Yield              Token
	ExpressionSequence *ExpressionSequence
	Eos                Token
	firstChild         VNode
	prev, next         VNode
}

var _ VNode = (*YieldStatement)(nil)

func (i *YieldStatement) Type() string {
	return "YieldStatement"
}
func (i *YieldStatement) Code() string {
	return CodeDef(i)
}
func (i *YieldStatement) Next() VNode {
	return i.next
}
func (i *YieldStatement) SetNext(v VNode) {
	i.next = v
}
func (i *YieldStatement) Prev() VNode {
	return i.prev
}
func (i *YieldStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *YieldStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitYieldStatement(ctx *base.YieldStatementContext) interface{} {
	r := &YieldStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	r.Yield = ident(v, ctx.Yield().GetSymbol())
	r.firstChild = r.Yield
	if ctx.ExpressionSequence() != nil {
		r.ExpressionSequence = v.VisitExpressionSequence(
			ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)

	}
	if ctx.Eos() != nil {
		if tk, ok := v.VisitEos(ctx.Eos().(*base.EosContext)).(Token); ok {
			r.Eos = tk
		}
	}
	setAllSibs(r.Yield, r.ExpressionSequence, r.Eos)

	return r
}

// withStatement
//     : With '(' expressionSequence ')' statement
//     ;
type WithStatement struct {
	*SourceInfo
	With               Token
	OpenParen          Token
	ExpressionSequence *ExpressionSequence
	CloseParen         Token
	Body               *Statement
	firstChild         VNode
	prev, next         VNode
}

var _ VNode = (*WithStatement)(nil)

func (i *WithStatement) Type() string {
	return "WithStatement"
}
func (i *WithStatement) Code() string {
	return CodeDef(i)
}
func (i *WithStatement) Next() VNode {
	return i.next
}
func (i *WithStatement) SetNext(v VNode) {
	i.next = v
}
func (i *WithStatement) Prev() VNode {
	return i.prev
}
func (i *WithStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *WithStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitWithStatement(ctx *base.WithStatementContext) interface{} {
	r := &WithStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	r.With = ident(v, ctx.With().GetSymbol())
	r.firstChild = r.With
	r.OpenParen = ident(v, ctx.OpenParen().GetSymbol())

	r.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	r.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	r.Body = v.VisitStatement(ctx.Statement().(*base.StatementContext)).(*Statement)
	setAllSibs(r.With, r.OpenParen, r.ExpressionSequence, r.CloseParen, r.Body)

	return r
}
