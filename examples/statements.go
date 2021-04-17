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
	Children []VNode
}

var _ VNode = (*StatementList)(nil)

func (i *StatementList) Type() string {
	return "StatementList"
}
func (i *StatementList) Code() string {
	return CodeDef(i)
}
func (i *StatementList) GetChildren() []VNode {

	return i.Children
}
func (v *visitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	// log.Println("VisitStatementList", ctx.GetText())
	return &StatementList{Children: v.VisitChildren(ctx).([]VNode)}
}

type Statement struct {
	*SourceInfo
	From     string
	Children []VNode //
}

var _ VNode = (*Statement)(nil)

func (i *Statement) Type() string {
	return "Statement"
}

func (i *Statement) Code() string {
	return CodeDef(i)
}
func (i *Statement) GetChildren() []VNode {
	return i.Children
}

func (v *visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	// log.Println("VisitStatement", ctx.GetText())
	return &Statement{Children: v.VisitChildren(ctx).([]VNode)}
}

// labelledStatement
//     : identifier ':' statement
//     ;
type LabeledStatement struct {
	*SourceInfo
	Statement *Statement
	Label     Token
	Colon     Token
}

var _ VNode = (*LabeledStatement)(nil)

func (i *LabeledStatement) Type() string {
	return "LabeledStatement"
}
func (i *LabeledStatement) Code() string {
	return CodeDef(i)
}
func (i *LabeledStatement) GetChildren() []VNode {
	if i == nil {
		return nil
	}
	return []VNode{i.Label, i.Colon, i.Statement}
}

// special case for $: ... todo: a type
func (v *visitor) VisitLabelledStatement(ctx *parser.LabelledStatementContext) interface{} {
	// log.Println("VisitLabelledStatement", ctx.GetText())

	if ctx.Identifier().GetText() == "$" {
		// log.Println("Reactive?")

	}
	lst := &LabeledStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
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
}

var _ VNode = (*Block)(nil)

func (i *Block) Type() string {
	return "Block"
}
func (i *Block) Code() string {
	return CodeDef(i)
}
func (i *Block) GetChildren() []VNode {
	return i.StatementList
}
func (v *visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// log.Println("VisitBlock", ctx.GetText())
	return &Block{StatementList: v.VisitChildren(ctx).([]VNode), SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}
