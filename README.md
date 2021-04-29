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

	expc = exp2.FirstChild().(*vast.LRExpression)
	// can be any singleExpression so any VNode
	fmt.Println(expc.Left().(vast.Token).Value(), expc.OP().Value(), expc.Right().(vast.Token).Value())

And a very incomplete conversion from arrow to es5 functions but it shows the general idea:

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
	go v.DefaultError()
	// v.Debug = true

	rfs := visit(tree, v).(*vast.Program).Body
	// Statement ->ExpressionStatememts -> ExpressionSequence
	for _, fn := range rfs {

		var trans string
		// all with one child
		af := fn.FirstChild().FirstChild().FirstChild().(*vast.ArrowFunction)

		// either has a function body with {} of a single expression.
		if af.Body.SingleExpression != nil {
			// can be there or not
			var open, close string
			if af.FunctionParameters.OpenParen == nil {
				open, close = "(", ")"
			}

			trans = fmt.Sprintf("function%s%s%s {\n\treturn %s;\n}",
				open, af.FunctionParameters.Source, close, af.Body.SingleExpression.Code())

		} else {

			trans = fmt.Sprintf("function%s %s", af.FunctionParameters.Source, af.Body.Source)

		}
		fmt.Println("After ->", trans)

	}
	
Quite verbose but more about getting it working than pretty to start...
