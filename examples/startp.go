package main

import (
	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

func (v *visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	// sourceElements as called when .Program() is used...
	return ctx.GetChild(0).(antlr.ParserRuleContext).Accept(v)

}

// Visit(tree ParseTree) interface{}
// VisitChildren(node RuleNode) interface{}
// VisitTerminal(node Identifier) interface{}
// VisitErrorNode(node ErrorNode) interface{}
func (v *visitor) VisitSourceElement(ctx *parser.SourceElementContext) interface{} {
	// log.Println("VisitSourceElement", ctx.GetText(), ctx.GetChildCount())
	s := &SourceElement{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	s.FirstChild = s.Children[0]
	if v.ParseTree != nil {
		lst := v.ParseTree.LastChild
		lst.Next = s
		s.Prev = lst
		v.ParseTree.LastChild = s
		return s
	}
	v.ParseTree = new(PTree)

	v.ParseTree = &PTree{Root: s, LastChild: s}

	return s

}
