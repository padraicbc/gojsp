package main

// import (
// 	"fmt"
// 	"log"

// 	antlr "github.com/padraicbc/antlr4"
// 	"github.com/padraicbc/gojsp"

// 	"github.com/padraicbc/gojsp/parser"
// )

// type listener struct {
// 	gojsp.BaseJavaScriptParserListener
// 	imports map[string]*ImportDeclaration
// }

// func show(fname, text string) {

// 	// log.Printf("%s %s\n", fname, text)
// }
// func (l *listener) EnterBlock(ctx *parser.BlockContext) {
// 	show("EnterBlock", ctx.GetText())
// }

// func (l *listener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
// 	show("EnterFunctionDeclaration", ctx.FunctionBody().GetText())
// }

// // VisitTerminal is called when a terminal node is visited.
// func (s *listener) VisitTerminal(node antlr.TerminalNode) {
// 	// log.Println(node)
// }

// // importStatement
// //     : Import importFromBlock
// //     ;

// // importFromBlock
// //     : importDefault? (importNamespace | moduleItems) importF
// // EnterImportStatement is called when production importStatement is entered.
// func (s *listener) EnterImportStatement(ctx *parser.ImportStatementContext) {
// 	impf := ctx.ImportFromBlock()
// 	log.Println(impf.GetChildCount())

// 	line, start := impf.GetStart().GetLine(), ctx.GetStart().GetStart()

// 	if impf != nil {
// 		for _, ch := range impf.GetChildren() {
// 			switch vv := ch.(type) {
// 			case *parser.ImportDefaultContext:
// 				fmt.Println("\tImportDefaultContext", vv.GetText())

// 			case *parser.ImportNamespaceContext:
// 				fmt.Println("\tImportNamespaceContext", vv.GetText())

// 			case *parser.ModuleItemsContext:
// 				fmt.Println("\tModuleItemsContext", vv.GetText())

// 			case *parser.ImportFromContext:
// 				fmt.Println("\tPath", vv.GetText())

// 			case *parser.EosContext:
// 				fmt.Println("\t", vv.GetText())
// 			case *parser.LiteralContext:
// 				fmt.Println("\tSimple", vv.GetText())

// 			default:
// 				log.Printf("%t\n", vv.GetPayload())
// 				// log.Panic(vv)

// 			}
// 		}

// 		log.Printf("Import On Line: %d from index -> %d to -> %d\n", line, start, impf.GetStop().GetStop()+1)
// 		orig := ctx.GetStart().GetInputStream().GetText(start, impf.GetStop().GetStop()+1)
// 		log.Println(orig)

// 	}

// }

// // ExitImportStatement is called when production importStatement is exited.
// func (s *listener) ExitImportStatement(ctx *parser.ImportStatementContext) {
// 	show("ExitImportStatement", ctx.GetText())
// 	// log.Println(ctx.ToStringTree(nil, nil))

// }

// // EnterImportFromBlock is called when production importFromBlock is entered.
// func (s *listener) EnterImportFromBlock(ctx *parser.ImportFromBlockContext) {

// 	show("EnterImportFromBlock", ctx.GetText())
// }

// // ExitImportFromBlock is called when production importFromBlock is exited.
// func (s *listener) ExitImportFromBlock(ctx *parser.ImportFromBlockContext) {
// 	show("ExitImportFromBlock", ctx.GetText())

// }

// // EnterModuleItems is called when production moduleItems is entered.
// func (s *listener) EnterModuleItems(ctx *parser.ModuleItemsContext) {
// 	show("EnterModuleItems", ctx.GetText())

// }

// // ExitModuleItems is called when production moduleItems is exited.
// func (s *listener) ExitModuleItems(ctx *parser.ModuleItemsContext) {
// 	show("ExitModuleItems", ctx.GetText())

// }

// // EnterImportDefault is called when production importDefault is entered.
// func (s *listener) EnterImportDefault(ctx *parser.ImportDefaultContext) {
// 	show("EnterImportDefault", ctx.GetText())

// }

// // ExitImportDefault is called when production importDefault is exited.
// func (s *listener) ExitImportDefault(ctx *parser.ImportDefaultContext) {
// 	show("ExitImportDefault", ctx.GetText())
// }

// // EnterImportNamespace is called when production importNamespace is entered.
// func (s *listener) EnterImportNamespace(ctx *parser.ImportNamespaceContext) {
// 	show("EnterImportNamespace", ctx.GetText())
// }

// // ExitImportNamespace is called when production importNamespace is exited.
// func (s *listener) ExitImportNamespace(ctx *parser.ImportNamespaceContext) {
// 	show("ExitImportNamespace", ctx.GetText())
// }

// // EnterImportFrom is called when production importFrom is entered.
// func (s *listener) EnterImportFrom(ctx *parser.ImportFromContext) {
// 	show("EnterImportFrom", ctx.GetText())
// }

// // ExitImportFrom is called when production importFrom is exited.
// func (s *listener) ExitImportFrom(ctx *parser.ImportFromContext) {
// 	show("ExitImportFrom", ctx.GetText())
// }
