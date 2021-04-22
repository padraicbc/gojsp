package vast

// import (
// 	"fmt"
// 	"log"

// 	antlr "github.com/padraicbc/antlr4"
// 	"github.com/padraicbc/gojsp"

// 	"github.com/padraicbc/gojsp/base"
// )

// type listener struct {
// 	gojsp.BaseJavaScriptParserListener
// 	imports map[string]*ImportFromBlock
// }

// func show(fname, text string) {

// 	// log.Printf("%s %s\n", fname, text)
// }
// func (l *listener) EnterBlock(ctx *base.BlockContext) {
// 	show("EnterBlock", ctx.GetText())
// }

// func (l *listener) EnterFunctionDeclaration(ctx *base.FunctionDeclarationContext) {
// 	show("EnterFunctionDeclaration", ctx.FunctionBody().GetText())
// }

// // VisitTerminal is called when a terminal node is visited.
// func (s *listener) VisitTerminal(node antlr.Identifier) {
// 	// log.Println(node)
// }

// // importStatement
// //     : Import importFromBlock
// //     ;

// // importFromBlock
// //     : importDefault? (importNamespace | moduleItems) importF
// // EnterImportStatement is called when production importStatement is entered.
// func (s *listener) EnterImportStatement(ctx *base.ImportStatementContext) {
// 	impf := ctx.ImportFromBlock()
// 	log.Println(impf.GetChildCount())

// 	line, start := impf.GetStart().GetLine(), ctx.GetStart().GetStart()

// 	if impf != nil {
// 		for _, ch := range impf.Children() {
// 			switch vv := ch.(type) {
// 			case *base.ImportDefaultContext:
// 				fmt.Println("\tImportDefaultContext", vv.GetText())

// 			case *base.ImportNamespaceContext:
// 				fmt.Println("\tImportNamespaceContext", vv.GetText())

// 			case *base.ModuleItemsContext:
// 				fmt.Println("\tModuleItemsContext", vv.GetText())

// 			case *base.ImportFromContext:
// 				fmt.Println("\tPath", vv.GetText())

// 			case *base.EosContext:
// 				fmt.Println("\t", vv.GetText())
// 			case *base.LiteralContext:
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
// func (s *listener) ExitImportStatement(ctx *base.ImportStatementContext) {
// 	show("ExitImportStatement", ctx.GetText())
// 	// log.Println(ctx.ToStringTree(nil, nil))

// }

// // EnterImportFromBlock is called when production importFromBlock is entered.
// func (s *listener) EnterImportFromBlock(ctx *base.ImportFromBlockContext) {

// 	show("EnterImportFromBlock", ctx.GetText())
// }

// // ExitImportFromBlock is called when production importFromBlock is exited.
// func (s *listener) ExitImportFromBlock(ctx *base.ImportFromBlockContext) {
// 	show("ExitImportFromBlock", ctx.GetText())

// }

// // EnterModuleItems is called when production moduleItems is entered.
// func (s *listener) EnterModuleItems(ctx *base.ModuleItemsContext) {
// 	show("EnterModuleItems", ctx.GetText())

// }

// // ExitModuleItems is called when production moduleItems is exited.
// func (s *listener) ExitModuleItems(ctx *base.ModuleItemsContext) {
// 	show("ExitModuleItems", ctx.GetText())

// }

// // EnterImportDefault is called when production importDefault is entered.
// func (s *listener) EnterImportDefault(ctx *base.ImportDefaultContext) {
// 	show("EnterImportDefault", ctx.GetText())

// }

// // ExitImportDefault is called when production importDefault is exited.
// func (s *listener) ExitImportDefault(ctx *base.ImportDefaultContext) {
// 	show("ExitImportDefault", ctx.GetText())
// }

// // EnterImportNamespace is called when production importNamespace is entered.
// func (s *listener) EnterImportNamespace(ctx *base.ImportNamespaceContext) {
// 	show("EnterImportNamespace", ctx.GetText())
// }

// // ExitImportNamespace is called when production importNamespace is exited.
// func (s *listener) ExitImportNamespace(ctx *base.ImportNamespaceContext) {
// 	show("ExitImportNamespace", ctx.GetText())
// }

// // EnterImportFrom is called when production importFrom is entered.
// func (s *listener) EnterImportFrom(ctx *base.ImportFromContext) {
// 	show("EnterImportFrom", ctx.GetText())
// }

// // ExitImportFrom is called when production importFrom is exited.
// func (s *listener) ExitImportFrom(ctx *base.ImportFromContext) {
// 	show("ExitImportFrom", ctx.GetText())
// }
