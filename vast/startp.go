package vast

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

type Program struct {
	Body []VNode
}

func (v *Visitor) VisitProgram(ctx *base.ProgramContext) interface{} {
	// sourceElements as called when .Program() is used...
	pg := &Program{}
	var prev VNode
	// this just adds prev/next nodes to traverse the tree. Might just remove it and return ctx.GetChild(0).(antlr.ParserRuleContext).Accept(v)
	for _, ch := range ctx.GetChild(0).(antlr.ParserRuleContext).Accept(v).([]VNode) {
		if prev != nil {
			prev.SetNext(ch)

		}

		pg.Body = append(pg.Body, ch)

	}
	return pg

}

// Visit(tree ParseTree) interface{}
// VisitChildren(node RuleNode) interface{}
// VisitTerminal(node Identifier) interface{}
// VisitErrorNode(node ErrorNode) interface{}
func (v *Visitor) VisitSourceElement(ctx *base.SourceElementContext) interface{} {
	if v.Debug {

		log.Println("VisitSourceElement", ctx.GetText(), ctx.GetChildCount())
	}

	return v.Visit(ctx.GetChild(0).(antlr.ParseTree)).(VNode)

}

// statementList
//     : statement+
//     ;
type StatementList struct {
	*SourceInfo
	Statements []VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*StatementList)(nil)

func (i *StatementList) Type() string {
	return "StatementList"
}
func (i *StatementList) Code() string {
	return CodeDef(i)
}
func (i *StatementList) Next() VNode {
	return i.next
}
func (i *StatementList) SetNext(v VNode) {
	i.next = v
}
func (i *StatementList) Prev() VNode {
	return i.prev
}
func (i *StatementList) SetPrev(v VNode) {
	i.prev = v
}
func (i *StatementList) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitStatementList(ctx *base.StatementListContext) interface{} {
	if v.Debug {
		log.Println("VisitStatementList", ctx.GetText())
	}
	sl := &StatementList{}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if sl.firstChild == nil {
			sl.firstChild = ch
		}
		prev = setSib(prev, ch)
		sl.Statements = append(sl.Statements, ch)

	}
	return sl
}

type Statement struct {
	*SourceInfo
	firstChild VNode
	Eos        Token
	prev, next VNode
}

var _ VNode = (*Statement)(nil)

func (i *Statement) Type() string {
	return "Statement"
}
func (i *Statement) Code() string {
	return CodeDef(i)
}
func (i *Statement) Next() VNode {
	return i.next
}
func (i *Statement) SetNext(v VNode) {
	i.next = v
}
func (i *Statement) Prev() VNode {
	return i.prev
}
func (i *Statement) SetPrev(v VNode) {
	i.prev = v
}
func (i *Statement) FirstChild() VNode {
	return i.firstChild
}

// not sure if just returning .Children() would be enough..
func (v *Visitor) VisitStatement(ctx *base.StatementContext) interface{} {
	if v.Debug {
		log.Println("VisitStatement", ctx.GetText(), ctx.GetChildCount())
	}
	st := &Statement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if st.firstChild == nil {
			st.firstChild = ch
		}
		prev = setSib(prev, ch)
		// not sure what can be here if grammar is correct
		if tk, ok := ch.(Token); ok && tk.SymbolName() == "SemiColon" {
			st.Eos = tk
		}

	}

	return st
}
