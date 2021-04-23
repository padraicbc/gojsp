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
	  }`)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)
	// all
	tree := p.Program()
	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	vp := visit(tree, v)
	log.Println(vp)

}
