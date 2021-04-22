package main

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
	"github.com/padraicbc/gojsp/vast"
)

func singleExp() {

	stream := antlr.NewInputStream(`i + j;`)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)
	// one
	tree := p.ExpressionStatement()
	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	exp := visit(tree, v).([]vast.VNode)[0].(*vast.ExpressionStatement)

	log.Println(exp.Left.Value(), exp.OP.Value(), exp.Right.Value())
	exp.OP.SetValue("/")
	log.Println(exp.Left.Value(), exp.OP.Value(), exp.Right.Value())

	// reuse lexer and parser
	stream.Seek(0)
	lexer.SetInputStream(stream)
	tokenStream = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(tokenStream)

	v = vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	// alterntive using Body
	tree2 := p.Program()
	v = vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	exp = visit(tree2, v).(*vast.Program).Body[0].(*vast.ExpressionStatement)

	log.Println(exp.Left.Value(), exp.OP.Value(), exp.Right.Value())

}
