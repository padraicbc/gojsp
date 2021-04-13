package main

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;
type Declaration struct {
	*SourceInfo
	// always one of 3 above
	Children []VNode
}

var _ VNode = (*Declaration)(nil)

func (i *Declaration) Type() string {
	return "Declaration"
}
func (i *Declaration) Code() string {
	return CodeDef(i)
}
func (i *Declaration) GetChildren() []VNode {
	return i.Children
}

func (v *visitor) VisitExportDefaultDeclaration(ctx *parser.ExportDefaultDeclarationContext) interface{} {
	// log.Println("VisitExportDefaultDeclaration", ctx.GetText())

	return &Declaration{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
}

type ExportFromBlock struct {
	*SourceInfo
	Children []VNode
	// todo: use thse or children?
	// one of thse 2
	ImportNamespace *ImportNamespace
	ModulesItems    *ModulesItems
	// always with ImportNamespace optional wirh moduleitems
	ImportFrom *ImportFrom
	// always this
	Path *ImportFrom
}

var _ VNode = (*ExportFromBlock)(nil)

func (i *ExportFromBlock) Type() string {
	return "ExportFromBlock"
}
func (i *ExportFromBlock) Code() string {
	return CodeDef(i)
}
func (i *ExportFromBlock) GetChildren() []VNode {
	return nil
}

// exportFromBlock
//     : importNamespace importFrom eos
//     | moduleItems importFrom? eos
//     ;
func (v *visitor) VisitExportFromBlock(ctx *parser.ExportFromBlockContext) interface{} {
	log.Println("VisitExportFromBlock", ctx.GetText())
	ef := &ExportFromBlock{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		switch rr := ch.(type) {

		case *ImportNamespace:
			ef.ImportNamespace = rr
		case *ModulesItems:

			ef.ModulesItems = rr
		case *ImportFrom:
			ef.Path = rr
		}
	}

	return ef
}

type ExportStatement struct {
	*SourceInfo
	Children []VNode
	Export   string

	// todo: use thse or children?
	ExportFromBlock  *ExportFromBlock
	Declaration      *Declaration
	ModulesItems     *ModulesItems
	Default          string
	SingleExpression string
}

var _ VNode = (*ExportStatement)(nil)

func (i *ExportStatement) Type() string {
	return "ExportStatement"
}
func (i *ExportStatement) Code() string {
	return CodeDef(i)
}
func (i *ExportStatement) GetChildren() []VNode {
	return i.Children
}

// exportStatement
//     : Export (exportFromBlock | declaration) eos    # ExportDeclaration
//     | Export Default singleExpression eos           # ExportDefaultDeclaration
//     ;
func (v *visitor) VisitExportDeclaration(ctx *parser.ExportDeclarationContext) interface{} {

	es := &ExportStatement{
		Export:     ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText(),
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	if se, ok := ctx.GetChild(1).(*antlr.TerminalNodeImpl); ok {
		es.SingleExpression = se.GetText()
	}

	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		switch rr := ch.(type) {

		case *ExportFromBlock:
			es.ExportFromBlock = rr
		case *Declaration:
			es.Declaration = rr
		case *ModulesItems:
			es.ModulesItems = rr

		}
	}
	return es

}
