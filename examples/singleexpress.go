package main

import (
	"fmt"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/vast"
)

func singleExp() {

	code := `i + j;`
	v := vast.NewVisitor(code)
	// do whatever with errors
	go v.DefaultError()
	// start at ExpressionSequence
	tree := v.Parser.ExpressionSequence()
	exp := visit(tree, v).(*vast.ExpressionSequence)

	expc := exp.FirstChild().(*vast.LRExpression)

	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())
	// change OP
	expc.OP().SetValue("/")
	expc.Right().(vast.Token).SetValue("1000")
	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

	// reuse lexer and parser
	v.Stream.Seek(0)
	v.Lexer.SetInputStream(v.Stream)
	tokenStream := antlr.NewCommonTokenStream(v.Lexer, antlr.TokenDefaultChannel)
	v.Parser.SetInputStream(tokenStream)

	tree2 := v.Parser.ExpressionStatement()
	exp2 := visit(tree2, v).(*vast.ExpressionStatement).FirstChild()

	expc = exp2.FirstChild().(*vast.LRExpression)
	// can be any singleExpression so any VNode
	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

}
