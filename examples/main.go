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

}

// func walk(tree antlr.ParseTree) {
// 	pr := antlr.NewParseTreeWalker()
// 	pr.Walk(&listener{}, tree)

// }
func visit(tree antlr.ParseTree, v *Visitor) interface{} {

	return v.Visit(tree)

}

func impexp() {
	stream := antlr.NewInputStream(`
    import foo as name from "star-module-name";
 	import defaultExport from "default-module-name";
	import defaultname, { export1, export2 as alias} from "module-name";
 	import "module-name";

var promise = import("module-name");
import * as name from "star-module-name";
import { export1 } from "exp1-module-name";
export { name1, name2, nameN };
// let a = 123;
// $: {
	
// 	let foo = 123;
		
// }
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
	tree := p.Program()
	v := NewVisitor(lexer, p)
	v.lexer = lexer
	v.parser = p
	visit(tree, v) //.(*Program) // .Children()[0].Children()[0].Children()[0].Children()[1].(*ImportFromBlock).ImportDefault.Default)
	// sourceElement
	// : statement
	// ;
	// each  -> statement
	for se := range v.ParseTree.NextNodes() {
		// Statements ...
		for _, st := range se.Children() {

			for _, ch := range st.Children() {
				switch ch.Type() {
				case "ImportStatement":
					ims := st.Children()[0].(*ImportStatement)
					log.Println("Before code", st.Code())
					// "import 'whatever'"
					if ims.ImportFromBlock.StringLiteral != nil {
						ims.ImportFromBlock.StringLiteral.SetValue(`"a-new/path"`)
						// else has an ImportFrom
					} else {

						log.Println(ims.ImportFromBlock.ImportFrom.Path.Value())
						ims.ImportFromBlock.ImportFrom.Path.SetValue(`"some/new_path"`)
						log.Println(ims.ImportFromBlock.ImportFrom.Path.Value())
					}
					log.Println("After code", st.Code())

				}
				fmt.Println()

			}

			// vch(st)

		}
	}

}

func vch(v VNode) {

	for _, ch := range v.Children() {
		if ch == nil {
			return
		}

		if ch.Type() == "AliasName" {
			log.Println(ch.Code())
			log.Println("Alias before", ch.(*AliasName).IdentifierName.Value())
			ch.(*AliasName).IdentifierName.SetValue("changed")
			log.Println("Alias after", ch.(*AliasName).IdentifierName.Value())
			log.Println(ch.Code())
		}
		vch(ch)

	}
}
