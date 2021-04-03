package main

import (
	"log"
	"strings"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp"
)

func main() {

	log.SetFlags(log.Lshortfile)

	antv()
}

func antv() {

	stream := antlr.NewInputStream(`
import defaultExport from "module-name";
import * as name from "module-name";
import { export1 } from "module-name";
import { export1 as alias1 } from "module-name";
import { export1 , export2 } from "module-name";
import { foo , bar } from "module-name/path/to/specific/un-exported/file";
import { export1 , export2 as alias} from "module-name";
import defaultExport, { export1 } from "module-name";
import defaultExport, * as name from "module-name";
import "module-name";
import {
	reallyReallyLongModuleExportName as shortName,
	anotherLongModuleName as short
  } from '/modules/my-module.js';
import { getUsefulContents } from '/modules2/file.js';
// todo
var promise = import("module-name");
let module = await import('/modules/my-module.js');

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

	// tree := p.Program()
	tree := p.SourceElements()
	// tree := p.ImportDefault() // to see a syntax error
	// p.Whatever
	visit(tree)

}
func walk(tree antlr.ParseTree) {
	pr := antlr.NewParseTreeWalker()
	pr.Walk(&listener{}, tree)

}
func visit(tree antlr.ParseTree) {
	vi := new(visitor).Visit(tree).([]string)
	log.Println(strings.Join(vi, "\n"))

}
