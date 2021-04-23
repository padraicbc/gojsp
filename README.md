## Very early attempt at creating an [ANTLR](https://www.antlr.org/) visitor for javascript files using GO. Going to change like Irish weather so don't expect things to stay the same any time soon. 


The `base` is the result of running what is described [here](https://github.com/padraicbc/gojsp/tree/master/runantlr#readme)

The [vast](https://github.com/padraicbc/gojsp/tree/master/vast) folder is the implementation so far toward creating an ast which will allow easy manipulation/translation of the original source.
Thre are a few example functions in the exampls folder including how to parse and [manipulate](https://github.com/padraicbc/gojsp/blob/master/examples/singleexpress.go) a single expression. 


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

Quite verbose but more about getting it working than pretty to start...
