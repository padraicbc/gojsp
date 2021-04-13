package main

import (
	"log"

	"github.com/padraicbc/gojsp/parser"
)

type Statement struct {
	*SourceInfo
	From       string
	prev, next VNode
	Children   []VNode //
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

// statementList
//     : statement+
//     ;
// 	statementList
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
	return nil
}
func (v *visitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	log.Println("VisitStatementList", ctx.GetText())
	return &StatementList{Children: v.VisitChildren(ctx).([]VNode)}
}

// block
//     : '{' statementList? '}'
//     ;
type Block struct {
	*SourceInfo
	// *StatementList
	Children []VNode
}

var _ VNode = (*Block)(nil)

func (i *Block) Type() string {
	return "Block"
}
func (i *Block) Code() string {
	return CodeDef(i)
}
func (i *Block) GetChildren() []VNode {
	return nil
}
func (v *visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	log.Println("VisitBlock", ctx.GetText())
	return &Block{Children: v.VisitChildren(ctx).([]VNode)}
}
