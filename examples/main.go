package main

import (
	"fmt"
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

func main() {

	log.SetFlags(log.Llongfile)

	antv()
}

func antv() {

	impexp()
	// multAddExp()
	// singleExp()
	// seqExp()
	// labeledStatementOK()
	// labeledStatementRecurLoop()

}

// func walk(tree antlr.ParseTree) {
// 	pr := antlr.NewParseTreeWalker()
// 	pr.Walk(&listener{}, tree)

// }
func visit(tree antlr.ParseTree, v *visitor) interface{} {

	return v.Visit(tree)

}
func singleExp() {

	// single exp
	stream := antlr.NewInputStream(`i + j;`)
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)

	tree := p.ExpressionStatement()
	v := new(visitor)
	v.Visit(tree)
	// visit(tree, v)

	// exp := v.Nodes[0].(*ExpressionSequence).Children[0]

	// log.Println("Single -> ", exp.Left, exp.OP, exp.Right)
}
func seqExp() {

	// single exp
	stream := antlr.NewInputStream(`i + j, h + o;`)
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)

	tree := p.ExpressionStatement()
	v := new(visitor)

	visit(tree, v)
	// for _, exp := range v.Expr {
	// 	log.Println("Sequence -> ", exp.Left, exp.OP, exp.Right, exp.Source)
	// }
}
func multAddExp() {

	stream := antlr.NewInputStream(`4 / 8
4 % 8
4 -8
4 =8
4 * 8`)
	// Create the js L exer
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)
	// all
	tree := p.Program()

	v := new(visitor)
	visit(tree, v)

	// for _, v := range v.Expr {
	// 	log.Println(v.Type())
	// 	log.Println(v.GetInfo().Source)
	// 	log.Println(v.Code())
	// 	v.Left = "100"
	// 	log.Println(v.Code())

	// }

}
func labeledStatementOK() {
	stream := antlr.NewInputStream(`
	export let title;


	$: document.title = title;

	$: {
		console.log("multiple statements can be combined");
		console.log("the current title is ${title}"");
	}
	`)
	// Create the js L exer
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)
	// all
	v := new(visitor)
	tree := p.Program()
	visit(tree, v)
}

// {} blows the stack. Think it came from mixing listeners and visitors... maybe ..
func labeledStatementRecurLoop() {

	stream := antlr.NewInputStream(`
	$: {
		
	 }
	`)
	// Create the js L exer
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)
	// all
	v := new(visitor)
	tree := p.Program()
	visit(tree, v)
}
func impexp() {
	stream := antlr.NewInputStream(`
// 	// import foo as name from "star-module-name";
	import defaultExport from "default-module-name";
	import defaultname, { export1 , export2 as alias} from "module-name";
// 	import "module-name";

// var promise = import("module-name");
// import * as name from "star-module-name";
// import { export1 } from "exp1-module-name";
// export { name1, name2, nameN };
// let a = 123;
$: {
	
	let foo = 123;
		
}
// import { export1 as alias1 } from "module-name";
// import { export1 , export2 } from "module-name";
// import { foo , bar } from "module-name/path/to/specific/un-exported/file";
// import { export1 , export2 as alias} from "module-name";
// import defaultExport, { export1 } from "module-name";
// import defaultExport, * as name from "module-name";
// import "module-name-singleExpression";
// import {
// 	reallyReallyLongModuleExportName as shortName,
// 	anotherLongModuleName as short
//   } from '/modules/my-module.js';
// import { getUsefulContents } from '/modules2/file.js';
// var promise = import("promise-module-name");
// let module = await import('/await/modules/my-module.js');
// import('/then/modules/my-module.js')
//   .then((module) => {
//     // Do something with the module.
//   });// Export list
// export { name1, name2, nameN };

// //Renaming exports
// export { variable1 as name1, variable2 as name2, nameN };
// export * from 'whatever'; // does not set the default export
// // Exporting destructured assignments with renaming
// export const { name1, name2: bar } = o;

// export default expression;
// export default function () { } // also class, function*
`)
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)
	// all
	tree := p.SourceElements()

	v := new(visitor)
	v.BaseJavaScriptParserVisitor.VisitChildren = v.VisitChildren
	v.lexer = lexer
	v.parser = p
	log.Println(visit(tree, v))
	// sourceElement
	// : statement
	// ;
	// each  -> statement
	for se := range v.ParseTree.NextNodes() {
		for _, st := range se.Children {
			fmt.Println(st.(*Statement).Children[0].Code())
			fmt.Println(st.(*Statement).Children[0].GetInfo().Source)

		}
	}

}
