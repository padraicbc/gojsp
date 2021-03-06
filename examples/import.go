package main

import (
	"fmt"

	"github.com/padraicbc/gojsp/vast"
)

func impexp() {
	code := `
	import foo as name from "star-module-name";
import defaultExport from "default-module-name";
import defaultname, { export1, export2 as alias} from "module-name";
import "module-name";

var promise = import("module-name");
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
 var promise = import("promise-module-name");
 let module = await import('/await/modules/my-module.js');
 export default foo;
 export default function name1() { } // also class, function*
export { name1 as default };
 `

	// all

	v := vast.NewVisitor(code)
	// v.Debug = true
	// do whatever with errors
	go v.DefaultError()

	tree := v.Parser.Program()
	vp := visit(tree, v)

	for _, ch := range vp.(*vast.Program).Body {
		ch = ch.FirstChild()
		switch ch.Type() {
		case "ImportStatement":

			ims := ch.(*vast.ImportStatement)
			fmt.Println("Before code", ch.Code())
			// "import 'whatever'"
			if ims.ImportFromBlock.StringLiteral != nil {
				// change path StringLiteral...
				ims.ImportFromBlock.StringLiteral.SetValue(`"a-new/path"`)
				// else has an ImportFrom
			} else {

				fmt.Println(ims.ImportFromBlock.ImportFrom.Path.Value())
				// Change ImportFrom path
				ims.ImportFromBlock.ImportFrom.Path.SetValue(`"some/new_path"`)
				fmt.Println(ims.ImportFromBlock.ImportFrom.Path.Value())
			}
			fmt.Println("After code", ch.Code())
			fmt.Println()
		case "VariableDeclarationList":
			vl := ch.(*vast.VariableDeclarationList)
			fmt.Println(vl.VarModifier.Value())
			fmt.Println()
		default:
			// fmt.Printf("%+v %s\n", ch, ch.Type())

		}

	}

}
