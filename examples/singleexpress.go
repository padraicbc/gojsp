package main

import (
	"fmt"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
	"github.com/padraicbc/gojsp/vast"
)

func singleExp() {

	stream := antlr.NewInputStream(`i + j;`)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)
	// start at ExpressionSequence
	tree := p.ExpressionSequence()
	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	exp := visit(tree, v).(*vast.ExpressionSequence)

	expc := exp.FirstChild().(*vast.LRExpression)

	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())
	// change OP
	expc.OP().SetValue("/")
	expc.Right().(vast.Token).SetValue("1000")
	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

	// reuse lexer and parser
	stream.Seek(0)
	lexer.SetInputStream(stream)
	tokenStream = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p.SetInputStream(tokenStream)

	v = vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	// alternative using Body
	tree2 := p.Program()
	v = vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	exp2 := visit(tree2, v).(*vast.Program).Body[0].(*vast.ExpressionStatement).FirstChild()

	expc = exp2.FirstChild().(*vast.LRExpression)
	// can be any singleExpression so any VNode
	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

}
