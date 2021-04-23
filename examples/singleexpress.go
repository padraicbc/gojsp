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
	exp := visit(tree, v).(*vast.ExpressionStatement)

	chi := exp.Children()
	expc := chi[0].(*vast.LRExpression)

	log.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())
	expc.OP().SetValue("/")
	expc.Right().(vast.Token).SetValue("1000")
	log.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

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
	expc = exp.Children()[0].(*vast.LRExpression)

	log.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

}
