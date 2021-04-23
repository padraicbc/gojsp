package main

import (
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
	 
	 (a,b) => a + b;

	(a, b) => {
		 return a+b;
	 }
	 `)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)

	tree := p.Program()

	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())

	// other way to go
	rfs := visit(tree, v).(*vast.Program).Body
	a, b := rfs[0].Children()[0].(*vast.ArrowFunction), rfs[1].Children()[0].(*vast.ArrowFunction)

	// makes no logical sense but shows how to change
	log.Println(a.Code())

	log.Println(b.Code())
}
