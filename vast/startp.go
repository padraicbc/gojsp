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
		ch.SetPrev(prev)

		prev = ch
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
		log.Println("VisitSourceElement", ctx.GetText())
	}
	return v.VisitChildren(ctx).([]VNode)

}

func (v *Visitor) VisitStatementList(ctx *base.StatementListContext) interface{} {
	if v.Debug {
		log.Println("VisitStatementList", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatement(ctx *base.StatementContext) interface{} {
	if v.Debug {
		log.Println("VisitStatement", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}
