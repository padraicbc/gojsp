package main

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp"
)

func main() {

	log.SetFlags(log.Lshortfile)

	antv()
}

func antv() {

	stream := antlr.NewInputStream(`import defaultExport from "default-module-name";
import * as name from "star-module-name";
import { export1 } from "exp1-module-name";
import { export1 as alias1 } from "module-name";
import { export1 , export2 } from "module-name";
import { foo , bar } from "module-name/path/to/specific/un-exported/file";
import { export1 , export2 as alias} from "module-name";
import defaultExport, { export1 } from "module-name";
import defaultExport, * as name from "module-name";
import "module-name-singleExpression";
import {
	reallyReallyLongModuleExportName as shortName,
	anotherLongModuleName as short
  } from '/modules/my-module.js';
import { getUsefulContents } from '/modules2/file.js';
// todo
var promise = import("promise-module-name");
let module = await import('/await/modules/my-module.js');
import('/then/modules/my-module.js')
  .then((module) => {
    // Do something with the module.
  });// Export list
export { name1, name2, nameN };

// Renaming exports
export { variable1 as name1, variable2 as name2, nameN };
export * from 'whatever'; // does not set the default export
// Exporting destructured assignments with renaming
export const { name1, name2: bar } = o;

export default expression;
export default function () { } // also class, function*
// odds  = evens.map(v => v + 1);
// pairs = evens.map(v => ({ even: v, odd: v + 1 }));
// nums  = evens.map((v, i) => v + i)
// let top2 = allContent
// 	  .filter((content) => content.type == "article")
// 	  .sort((a, b) => a.date > b.date)
// 	  .slice(0, 2);
	  	  `)

	// Create the js Lexer
	lexer := gojsp.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := gojsp.NewJavaScriptParser(tokenStream)
	// all
	tree := p.Program()

	v := new(visitor)
	visit(tree, v)
	// log.Println(v.Imports)
	// Create the js Lexer

	// two exp
	stream = antlr.NewInputStream(`i + j;f + k`)
	lexer = gojsp.NewJavaScriptLexer(stream)

	tokenStream = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p = gojsp.NewJavaScriptParser(tokenStream)

	tree2 := p.Program()
	v = new(visitor)

	visit(tree2, v)
	for _, exp := range v.Expr {
		log.Println(exp.Left, exp.OP, exp.Right)
	}

	// single exp
	stream = antlr.NewInputStream(`i + j;`)
	lexer = gojsp.NewJavaScriptLexer(stream)

	tokenStream = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p = gojsp.NewJavaScriptParser(tokenStream)

	tree3 := p.ExpressionStatement()
	v = new(visitor)

	visit(tree3, v)
	exp := v.Expr[0]
	log.Println(exp.Left, exp.OP, exp.Right)

}
func walk(tree antlr.ParseTree) {
	pr := antlr.NewParseTreeWalker()
	pr.Walk(&listener{}, tree)

}
func visit(tree antlr.ParseTree, v *visitor) interface{} {

	return v.Visit(tree)

}
