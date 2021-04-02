package main

import (
	"fmt"
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/parserv"
)

type listener struct {
	parserv.BaseJavaScriptParserListener
}

func show(fname, text string) {

	// log.Printf("%s %s\n", fname, text)
}
func (l *listener) EnterBlock(ctx *parserv.BlockContext) {
	show("EnterBlock", ctx.GetText())
}

func (l *listener) EnterFunctionDeclaration(ctx *parserv.FunctionDeclarationContext) {
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
func (s *listener) EnterImportStatement(ctx *parserv.ImportStatementContext) {
	impf := ctx.ImportFromBlock()
	log.Println(impf.GetChildCount())

	line, start := impf.GetStart().GetLine(), ctx.GetStart().GetStart()

	if impf != nil {
		for _, ch := range impf.GetChildren() {
			switch vv := ch.(type) {
			case *parserv.ImportDefaultContext:
				fmt.Println("\tImportDefaultContext", vv.GetText())

			case *parserv.ImportNamespaceContext:
				fmt.Println("\tImportNamespaceContext", vv.GetText())

			case *parserv.ModuleItemsContext:
				fmt.Println("\tModuleItemsContext", vv.GetText())

			case *parserv.ImportFromContext:
				fmt.Println("\tPath", vv.GetText())

			case *parserv.EosContext:
				fmt.Println("\t", vv.GetText())
			case *parserv.LiteralContext:
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
func (s *listener) ExitImportStatement(ctx *parserv.ImportStatementContext) {
	show("ExitImportStatement", ctx.GetText())
	// log.Println(ctx.ToStringTree(nil, nil))

}

// EnterImportFromBlock is called when production importFromBlock is entered.
func (s *listener) EnterImportFromBlock(ctx *parserv.ImportFromBlockContext) {

	show("EnterImportFromBlock", ctx.GetText())
}

// ExitImportFromBlock is called when production importFromBlock is exited.
func (s *listener) ExitImportFromBlock(ctx *parserv.ImportFromBlockContext) {
	show("ExitImportFromBlock", ctx.GetText())

}

// EnterModuleItems is called when production moduleItems is entered.
func (s *listener) EnterModuleItems(ctx *parserv.ModuleItemsContext) {
	show("EnterModuleItems", ctx.GetText())

}

// ExitModuleItems is called when production moduleItems is exited.
func (s *listener) ExitModuleItems(ctx *parserv.ModuleItemsContext) {
	show("ExitModuleItems", ctx.GetText())

}

// EnterImportDefault is called when production importDefault is entered.
func (s *listener) EnterImportDefault(ctx *parserv.ImportDefaultContext) {
	show("EnterImportDefault", ctx.GetText())

}

// ExitImportDefault is called when production importDefault is exited.
func (s *listener) ExitImportDefault(ctx *parserv.ImportDefaultContext) {
	show("ExitImportDefault", ctx.GetText())
}

// EnterImportNamespace is called when production importNamespace is entered.
func (s *listener) EnterImportNamespace(ctx *parserv.ImportNamespaceContext) {
	show("EnterImportNamespace", ctx.GetText())
}

// ExitImportNamespace is called when production importNamespace is exited.
func (s *listener) ExitImportNamespace(ctx *parserv.ImportNamespaceContext) {
	show("ExitImportNamespace", ctx.GetText())
}

// EnterImportFrom is called when production importFrom is entered.
func (s *listener) EnterImportFrom(ctx *parserv.ImportFromContext) {
	show("EnterImportFrom", ctx.GetText())
}

// ExitImportFrom is called when production importFrom is exited.
func (s *listener) ExitImportFrom(ctx *parserv.ImportFromContext) {
	show("ExitImportFrom", ctx.GetText())
}
