package main

import (
	"fmt"
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
	"github.com/padraicbc/gojsp/vast"
)

func fs() {

	stream := antlr.NewInputStream(`function calcRectArea(width, height) {
		 width = width * 12 / 3 ** (function () {
			 a *= 12;
			return 10;
		})();
		return width * height;
	  }
	  
	  foo(a,b)=> a +b
	  `)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)

	tree := p.FunctionDeclaration()
	// could use this with .body way below
	// tree := p.Program()

	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())

	// other way to go
	// tree := p.Program()
	// fd := visit(tree, v).(*vast.Program).Body[0].(*vast.FunctionDeclaration)
	fd := visit(tree, v).(*vast.FunctionDeclaration)
	fd.FunctionBody.Children()[0].Type()
	// makes no logival sense but shows how to change
	fd.FormalParameterList.FormalParameterArgs[0].Assignable.(vast.Token).SetValue("new")
	log.Println(fd.Code())

}

func arrow() {
	stream := antlr.NewInputStream(`
	 
// Arrow Function Break Down

// 1. Remove the word "function" and place arrow between the argument and opening body bracket
(a) => {
  return a + 100;
}

// 2. Remove the body brackets and word "return" -- the return is implied.
(b) => b + 100;

// 3. Remove the argument parentheses
c => c 

// Arrow Function
(a, b) => {
  let chuck = 42;
  return a + b + chuck;
}
	 `)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)

	tree := p.Program()

	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())

	rfs := visit(tree, v).(*vast.Program).Body
	// *ExpressionStatememts
	for _, fn := range rfs {
		var trans string
		// all with one child
		af := fn.Children()[0].(*vast.ArrowFunction)
		fmt.Println("Before ->", af.Code())
		// either has a fucntion body with {} of a signle expression.
		if af.FunctionBody.SingleExpression != nil {
			// can be there or not
			var open, close string
			if af.FunctionParameters.OpenParen == nil {
				open, close = "(", ")"
			}
			trans = fmt.Sprintf("function%s%s%s {\n\treturn %s\n}",
				open, af.FunctionParameters.Source, close, af.FunctionBody.SingleExpression.Code())

		}
		if af.FunctionBody.FunctionBody != nil {

			trans = fmt.Sprintf("function%s %s", af.FunctionParameters.Source, af.FunctionBody.FunctionBody.Source)

		}
		fmt.Println("After ->", trans)

	}
}
