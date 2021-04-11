package main

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

func main() {

	log.SetFlags(log.Lshortfile)

	antv()
}

func antv() {

	impexp()
	multAddExp()
	singleExp()
	labeledStatementOK()
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

	visit(tree, v)
	exp := v.Expr[0].(*Expression)
	log.Println("Single -> ", exp.Left, exp.OP, exp.Right)
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

	for _, v := range v.Expr {
		log.Println(v.Type())
		log.Println(v.GetInfo().Source)
		log.Println(v.Code())
		v.(*Expression).Left = "100"
		log.Println(v.Code())

	}

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
	import defaultExport from "default-module-name";
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
todo
var promise = import("promise-module-name");
let module = await import('/await/modules/my-module.js');
import('/then/modules/my-module.js')
  .then((module) => {
    // Do something with the module.
  });// Export list
export { name1, name2, nameN };

//Renaming exports
export { variable1 as name1, variable2 as name2, nameN };
export * from 'whatever'; // does not set the default export
// Exporting destructured assignments with renaming
export const { name1, name2: bar } = o;

export default expression;
export default function () { } // also class, function*
`)
	lexer := parser.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(tokenStream)
	// all
	tree := p.Program()

	v := new(visitor)
	visit(tree, v)

	for _, v := range v.Imports {

		log.Println(v.Code())
		if v.Type() == "ImportDeclaration" {
			id := v.(*ImportDeclaration)
			id.ImportFrom = `"/whatever/new/path"`

			if id.ModulesItems != nil {
				id.ModulesItems.AliasNames[0].IdentifierName = "WhateverName"
			}
		}

		log.Println(v.Code())

	}

}
