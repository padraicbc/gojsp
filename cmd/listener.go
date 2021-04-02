package main

import (
	"fmt"
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp"
)

type listener struct {
	gojsp.BaseJavaScriptParserListener
}

func show(fname, text string) {

	// log.Printf("%s %s\n", fname, text)
}
func (l *listener) EnterBlock(ctx *gojsp.BlockContext) {
	show("EnterBlock", ctx.GetText())
}

func (l *listener) EnterFunctionDeclaration(ctx *gojsp.FunctionDeclarationContext) {
	show("EnterFunctionDeclaration", ctx.FunctionBody().GetText())
}

// VisitTerminal is called when a terminal node is visited.
func (s *listener) VisitTerminal(node antlr.TerminalNode) {
	// log.Println(node)
}

// importStatement
//     : Import importFromBlock
//     ;

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importF
// EnterImportStatement is called when production importStatement is entered.
func (s *listener) EnterImportStatement(ctx *gojsp.ImportStatementContext) {
	impf := ctx.ImportFromBlock()
	log.Println(impf.GetChildCount())

	line, start := impf.GetStart().GetLine(), ctx.GetStart().GetStart()

	if impf != nil {
		for _, ch := range impf.GetChildren() {
			switch vv := ch.(type) {
			case *gojsp.ImportDefaultContext:
				fmt.Println("\tImportDefaultContext", vv.GetText())

			case *gojsp.ImportNamespaceContext:
				fmt.Println("\tImportNamespaceContext", vv.GetText())

			case *gojsp.ModuleItemsContext:
				fmt.Println("\tModuleItemsContext", vv.GetText())

			case *gojsp.ImportFromContext:
				fmt.Println("\tPath", vv.GetText())

			case *gojsp.EosContext:
				fmt.Println("\t", vv.GetText())
			case *gojsp.LiteralContext:
				fmt.Println("\tSimple", vv.GetText())

			default:
				log.Printf("%t\n", vv.GetPayload())
				// log.Panic(vv)

			}
		}

		log.Printf("Import On Line: %d from index -> %d to -> %d\n", line, start, impf.GetStop().GetStop()+1)
		orig := ctx.GetStart().GetInputStream().GetText(start, impf.GetStop().GetStop()+1)
		log.Println(orig)

	}

}

// ExitImportStatement is called when production importStatement is exited.
func (s *listener) ExitImportStatement(ctx *gojsp.ImportStatementContext) {
	show("ExitImportStatement", ctx.GetText())
	// log.Println(ctx.ToStringTree(nil, nil))

}

// EnterImportFromBlock is called when production importFromBlock is entered.
func (s *listener) EnterImportFromBlock(ctx *gojsp.ImportFromBlockContext) {

	show("EnterImportFromBlock", ctx.GetText())
}

// ExitImportFromBlock is called when production importFromBlock is exited.
func (s *listener) ExitImportFromBlock(ctx *gojsp.ImportFromBlockContext) {
	show("ExitImportFromBlock", ctx.GetText())

}

// EnterModuleItems is called when production moduleItems is entered.
func (s *listener) EnterModuleItems(ctx *gojsp.ModuleItemsContext) {
	show("EnterModuleItems", ctx.GetText())

}

// ExitModuleItems is called when production moduleItems is exited.
func (s *listener) ExitModuleItems(ctx *gojsp.ModuleItemsContext) {
	show("ExitModuleItems", ctx.GetText())

}

// EnterImportDefault is called when production importDefault is entered.
func (s *listener) EnterImportDefault(ctx *gojsp.ImportDefaultContext) {
	show("EnterImportDefault", ctx.GetText())

}

// ExitImportDefault is called when production importDefault is exited.
func (s *listener) ExitImportDefault(ctx *gojsp.ImportDefaultContext) {
	show("ExitImportDefault", ctx.GetText())
}

// EnterImportNamespace is called when production importNamespace is entered.
func (s *listener) EnterImportNamespace(ctx *gojsp.ImportNamespaceContext) {
	show("EnterImportNamespace", ctx.GetText())
}

// ExitImportNamespace is called when production importNamespace is exited.
func (s *listener) ExitImportNamespace(ctx *gojsp.ImportNamespaceContext) {
	show("ExitImportNamespace", ctx.GetText())
}

// EnterImportFrom is called when production importFrom is entered.
func (s *listener) EnterImportFrom(ctx *gojsp.ImportFromContext) {
	show("EnterImportFrom", ctx.GetText())
}

// ExitImportFrom is called when production importFrom is exited.
func (s *listener) ExitImportFrom(ctx *gojsp.ImportFromContext) {
	show("ExitImportFrom", ctx.GetText())
}
