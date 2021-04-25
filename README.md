## Very early attempt at creating an [ANTLR](https://www.antlr.org/) visitor for javascript files using GO. Going to change like Irish weather so don't expect things to stay the same any time soon. 


The `base` is the result of running what is described [here](https://github.com/padraicbc/gojsp/tree/master/runantlr#readme)

The [vast](https://github.com/padraicbc/gojsp/tree/master/vast) folder is the implementation so far toward creating an ast which will allow easy manipulation/translation of the original source.
Thre are a few example functions in the exampls folder including how to parse and [manipulate](https://github.com/padraicbc/gojsp/blob/master/examples/singleexpress.go) a single expression. 


    import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
	"github.com/padraicbc/gojsp/vast"
    )



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
	log.Println(expc.Code())

And a very incomplete conversion from arrow to es5 fucntions but it shows the general idea:

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
	
Quite verbose but more about getting it working than pretty to start...
