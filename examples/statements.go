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
	children   []VNode
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

	return i.children
}
func (v *Visitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	// log.Println("VisitStatementList", ctx.GetText())
	return &StatementList{children: v.VisitChildren(ctx).([]VNode), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}

type Statement struct {
	*SourceInfo
	From       string
	children   []VNode
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
	return i.children
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	// log.Println("VisitStatement", ctx.GetText())
	return &Statement{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}

// labelledStatement
//     : identifier ':' statement
//     ;
type LabeledStatement struct {
	*SourceInfo
	Statement  *Statement
	Label      Token
	Colon      Token
	children   []VNode
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

	return i.children
}

// special case for $: ... todo: a type
func (v *Visitor) VisitLabelledStatement(ctx *parser.LabelledStatementContext) interface{} {
	// log.Println("VisitLabelledStatement", ctx.GetText())

	if ctx.Identifier().GetText() == "$" {
		// log.Println("Reactive?")

	}
	lst := &LabeledStatement{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	for _, ch := range lst.children {
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
	}
	return lst
}

// block
//     : '{' statementList? '}'
//     ;
type Block struct {
	*SourceInfo
	// *StatementList
	StatementList []VNode
	prev, next    VNode
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
	return i.StatementList
}
func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// log.Println("VisitBlock", ctx.GetText())
	return &Block{
		StatementList: v.VisitChildren(ctx).([]VNode),
		SourceInfo:    getSourceInfo(*ctx.BaseParserRuleContext)}
}
