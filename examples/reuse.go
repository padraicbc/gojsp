package main

import (
	"log"
	"time"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
	"github.com/padraicbc/gojsp/vast"
)

func coldHot() {
	s := time.Now()
	stream := antlr.NewInputStream(`
import foo as name from "star-module-name";
import defaultExport from "default-module-name";
import defaultname, { export1, export2 as alias} from "module-name";
import "module-name";

var promise = import("module-name");
import * as name from "star-module-name";
import { export1 } from "exp1-module-name";
export { name1, name2, nameN };

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
Exporting destructured assignments with renaming
// export const { name1, name2: bar } = o;

// export default expression;
export default function () { } // also class, function*
`)

	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)
	// all
	tree := p.Program()
	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	visit(tree, v)
	crun := time.Now().Sub(s)
	s = time.Now()
	// Reuse lexer and parser. Can pass new stream altogether. I was too lazy..
	stream.Seek(0)
	lexer.SetInputStream(stream)
	tokenStream = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(tokenStream)

	v = vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	tree2 := p.Program()
	visit(tree2, v)
	log.Println("Cold run -> ", crun)
	log.Println("Hot run -> ", time.Now().Sub(s))
}
