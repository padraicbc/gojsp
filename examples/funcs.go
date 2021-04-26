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
	// makes no logical sense but shows how to change
	fd.PList.FormalParameterArgs[0].Assignable.(vast.Token).SetValue("new")
	fmt.Println(fd.Code())

}

func arrow() {
	stream := antlr.NewInputStream(`

(a,b) => a + b;

(a, b) => {
	return a + b;
}
	 `)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := base.NewJavaScriptParser(tokenStream)

	tree := p.Program()

	v := vast.NewVisitor(lexer.SymbolicNames, p.GetRuleNames())
	// v.Debug = true

	rfs := visit(tree, v).(*vast.Program).Body

	// ExpressionStatement -> ExpressionSequence -> ArrowFunction
	exp1, exp2 := rfs[0], rfs[1]
	ar1, ar2 := exp1.FirstChild().FirstChild().(*vast.ArrowFunction),
		exp2.FirstChild().FirstChild().(*vast.ArrowFunction)

	fmt.Println(exp1.Code())
	lr := ar1.Body.SingleExpression.(*vast.LRExpression)
	lr.OP().SetValue("*")
	lft := lr.Left().(vast.Token)
	lft.SetValue(lft.Value() + " * 100")
	fmt.Println(exp1.Code() + "\n")

	// change operator
	fmt.Println(exp2.Code())
	bdy := ar2.Body.FBody

	// FirstChild() -> .Next() = brace then return
	ret := bdy.FirstChild().Next().(*vast.ReturnStatement)
	// is a left/right with single token expressions
	ret.ExpSeq.FirstChild().(*vast.LRExpression).OP().SetValue("/")
	fmt.Println(exp2.Code())
	// source stays the same
	fmt.Println(exp2.GetInfo().Source)
}

func toes5() {
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
	// v.Debug = true

	rfs := visit(tree, v).(*vast.Program).Body
	// *ExpressionStatememts -> ExpressionSequence, iterate and check types
	for _, fn := range rfs {
		log.Println(fn.Type())
		var trans string
		// all with one child
		af := fn.FirstChild().FirstChild().(*vast.ArrowFunction)
		fmt.Println("Before ->", af.Code())
		// either has a fucntion body with {} of a single expression.
		if af.Body.FirstChild() != nil {
			// can be there or not
			var open, close string
			if af.FunctionParameters.OpenParen == nil {
				open, close = "(", ")"
			}
			log.Println(af.Body.FirstChild().Type(), af.FunctionParameters.Source)
			trans = fmt.Sprintf("function%s%s%s {\n\treturn %s\n}",
				open, af.FunctionParameters.Source, close, af.Body.FirstChild().GetInfo().Source)

		}
		if af.Body != nil {

			trans = fmt.Sprintf("function%s %s", af.FunctionParameters.Source, af.Body.Source)

		}
		fmt.Println("After ->", trans)

	}
}
