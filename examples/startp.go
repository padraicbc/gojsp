package main

import (
	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

type Program struct {
	prev, next VNode
	tree       PTree
	children   []VNode
}

var _ VNode = (*Program)(nil)

func (i *Program) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *Program) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *Program) Code() string {
	return CodeDef(i)
}

func (i *Program) Type() string {
	return "Program"
}

func (i *Program) GetInfo() *SourceInfo {

	return nil
}
func (i *Program) Children() []VNode {

	return i.children
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	// sourceElements as called when .Program() is used...
	children := ctx.GetChild(0).(antlr.ParserRuleContext).Accept(v).([]VNode)

	return &Program{children: children}

}

// Visit(tree ParseTree) interface{}
// VisitChildren(node RuleNode) interface{}
// VisitTerminal(node Identifier) interface{}
// VisitErrorNode(node ErrorNode) interface{}
type SourceElement struct {
	*SourceInfo
	// VNodes have their own next/prev. Can visit all children from here
	children []VNode

	prev, next VNode
}

var _ VNode = (*SourceElement)(nil)

func (i *SourceElement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *SourceElement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *SourceElement) Code() string {
	return CodeDef(i)
}

func (i *SourceElement) Type() string {
	return "SourceElement"
}

func (i *SourceElement) Children() []VNode {

	return i.children
}

func (v *Visitor) VisitSourceElement(ctx *parser.SourceElementContext) interface{} {
	// log.Println("VisitSourceElement", ctx.GetText(), ctx.GetChildCount())
	s := &SourceElement{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	if v.ParseTree.Root == nil {
		v.ParseTree.Root = s

	}
	v.ParseTree.LastChild = s

	return s

}
