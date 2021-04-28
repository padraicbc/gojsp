package main

import (
	"fmt"
	"log"

	"github.com/padraicbc/gojsp/vast"
)

func fs() {

	code := `function calcRectArea(width, height) {
		 width = width * 12 / 3 ** (function () {
			 a *= 12;
			return 10;
		})();
		return width * height;
	  }
	  
	  foo(a,b)=> a +b
	  `
	// could use this with .body way below
	// tree := p.Program()

	v := vast.NewVisitor(code)
	tree := v.Parser.FunctionDeclaration()
	// other way to go
	// tree := p.Program()
	// fd := visit(tree, v).(*vast.Program).Body[0].(*vast.FunctionDeclaration)
	fd := visit(tree, v).(*vast.FunctionDeclaration)
	// makes no logical sense but shows how to change
	fd.PList.FormalParameterArgs[0].Assignable.(vast.Token).SetValue("new")
	fmt.Println(fd.Code())

}

func arrow() {
	code := `

(a,b) => a + b;

(c, d) => {
	return c + d;
}
let a = null;
a = 13;
`

	v := vast.NewVisitor(code)
	v.Debug = true
	// do whatever with errors
	go func() {
		e := <-v.Errors
		log.Fatal(e)
	}()
	tree := v.Parser.Program()

	go func() {
		e := <-v.Errors
		log.Fatal(e)
	}()
	rfs := visit(tree, v).(*vast.Program).Body

	// ExpressionStatement -> ExpressionSequence -> ArrowFunction
	// either way below can access
	exp1, exp2 := rfs[0], rfs[1]
	ar1, ar2 := exp1.(*vast.ExpressionStatement).ExpSequence.FirstChild().(*vast.ArrowFunction),
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
	log.Println(rfs[2].(*vast.VariableDeclarationList).VarModifier)
}

func toes5() {
	code := `
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
 `

	v := vast.NewVisitor(code)
	tree := v.Parser.Program()
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
