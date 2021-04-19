package main

import (
	"github.com/padraicbc/gojsp/parser"
)

// statementList
//     : statement+
//     ;
type StatementList struct {
	*SourceInfo
	// *Statements
	children   VNode
	prev, next VNode
}

var _ VNode = (*StatementList)(nil)

func (i *StatementList) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *StatementList) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *StatementList) Type() string {
	return "StatementList"
}
func (i *StatementList) Code() string {
	return CodeDef(i)
}
func (i *StatementList) Children() []VNode {

	return children(i.children)
}
func (v *Visitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	// log.Println("VisitStatementList", ctx.GetText())
	stl := &StatementList{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		if stl.children == nil {
			stl.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
	}
	return stl
}

type Statement struct {
	*SourceInfo
	From       string
	children   VNode
	prev, next VNode //
}

var _ VNode = (*Statement)(nil)

func (i *Statement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *Statement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *Statement) Type() string {
	return "Statement"
}

func (i *Statement) Code() string {
	return CodeDef(i)
}
func (i *Statement) Children() []VNode {
	return children(i.children)
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	// log.Println("VisitStatement", ctx.GetText())
	st := &Statement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if st.children == nil {
			st.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
	}
	return st
}

// labelledStatement
//     : identifier ':' statement
//     ;
type LabeledStatement struct {
	*SourceInfo
	Statement  *Statement
	Label      Token
	Colon      Token
	children   VNode
	prev, next VNode
}

var _ VNode = (*LabeledStatement)(nil)

func (i *LabeledStatement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *LabeledStatement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *LabeledStatement) Type() string {
	return "LabeledStatement"
}
func (i *LabeledStatement) Code() string {
	return CodeDef(i)
}
func (i *LabeledStatement) Children() []VNode {

	return children(i.children)
}

// special case for $: ... todo: a type
func (v *Visitor) VisitLabelledStatement(ctx *parser.LabelledStatementContext) interface{} {
	// log.Println("VisitLabelledStatement", ctx.GetText())

	if ctx.Identifier().GetText() == "$" {
		// log.Println("Reactive?")
	}
	lst := &LabeledStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		switch ch.Type() {
		case "LToken":
			t := ch.(Token)
			if t.RName() == "identifier" {
				lst.Label = t
				continue
			}
			lst.Colon = t
		case "Statement":
			lst.Statement = ch.(*Statement)
		default:
			panic(ch.Type())

		}
		if lst.children == nil {
			lst.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
	}
	return lst
}

// block
//     : '{' statementList? '}'
//     ;
type Block struct {
	*SourceInfo
	// *StatementList
	children   VNode
	prev, next VNode
}

var _ VNode = (*Block)(nil)

func (i *Block) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *Block) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *Block) Type() string {
	return "Block"
}
func (i *Block) Code() string {
	return CodeDef(i)
}
func (i *Block) Children() []VNode {
	return children(i.children)
}
func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// log.Println("VisitBlock", ctx.GetText())
	b := &Block{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if b.children == nil {
			b.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
	}
	return b
}
