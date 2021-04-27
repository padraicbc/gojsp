package vast

import (
	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

type SynError struct {
	Line, Column int
	Msg          string
	Exc          antlr.RecognitionException
}

// need pointer receiver for methods...
type Visitor struct {
	// any methods not implemented to satisfy JavaScriptParserVisitor checks in Accept...
	// JavaScriptParserVisitor embeds antlr.ParseTreeVisitor so we are also a "antlr.ParseTreeVisitor"
	// any visitor methods we don't add are called on BaseJavaScriptParserVisitor which are no-ops essentially -> nil
	base.BaseJavaScriptParserVisitor
	// todo:  syntax errors with line/col ...
	// errors []string

	ruleNames       []string
	symbolicNames   []string
	Debug           bool
	Errors          chan SynError
	Lexer           *base.JavaScriptLexer
	Parser          *base.JavaScriptParser
	Stream          *antlr.InputStream
	SyntaxErrorFunc func(errors chan SynError, recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int,
		msg string, e antlr.RecognitionException)
}

// TODO:  Should all be on a  differnt "Error" type but easier for now. Plus add fields similar to SyntaxErrorFunc to Visitor so easy to implement

// This method is called by the parser when a full-context prediction results in an ambiguity.
func (d *Visitor) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {

}

// This method is called when an SLL conflict occurs and the parser is about to use the full context information to make an LL decision.
func (d *Visitor) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// log.Println(recognizer.GetInputStream().Index(), conflictingAlts)
}

// This method is called by the parser when a full-context prediction has a unique result.
func (d *Visitor) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	// log.Println(recognizer.GetInputStream().Index(), prediction)

}

func DefaultSyntaxError(errors chan SynError, recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int,
	msg string, e antlr.RecognitionException) {
	errors <- SynError{Line: line, Column: column, Msg: msg, Exc: e}
}

// Upon syntax error, notify any interested parties.
func (c *Visitor) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

	c.SyntaxErrorFunc(c.Errors, recognizer, offendingSymbol, line, column, msg, e)

}
func NewVisitor(code string) *Visitor {
	stream := antlr.NewInputStream(code)
	lexer := base.NewJavaScriptLexer(stream)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := base.NewJavaScriptParser(tokenStream)

	symbolicNames, ruleNames := lexer.SymbolicNames, parser.GetRuleNames()
	vis := &Visitor{symbolicNames: symbolicNames,
		Lexer:     lexer,
		ruleNames: ruleNames, Parser: parser,
		Errors: make(chan SynError, 1), Stream: stream,
		SyntaxErrorFunc: DefaultSyntaxError}

	parser.RemoveErrorListeners()
	parser.AddErrorListener(vis)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(vis)

	return vis
}

func (v *Visitor) VisitArgument(ctx *base.ArgumentContext) interface{} {

	return v.VisitChildren(ctx)
}

// public T visit(ParseTree tree)
// Visit a parse tree, and return a user-defined result of the operation.
// The default implementation calls ParseTree.accept(org.antlr.v4.runtime.tree.ParseTreeVisitor<? extends T>) on the specified tree.

// Specified by:
// visit in interface ParseTreeVisitor<T>
// Parameters:
// tree - The ParseTree to visit.
// Returns:
//     the result of visiting the parse tree.
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitSourceElements is called when production sourceElements is entered.
// Program -> this -> SourceElement
func (v *Visitor) VisitSourceElements(ctx *base.SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

// maybe make an a files on Visitor and non nil implementation gets called...
func (v *Visitor) ShouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	return true
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return node
}

// variableStatement
//     : variableDeclarationList eos
//     ;
type VariableStatement struct {
	*SourceInfo
}

func (v *Visitor) VisitVariableStatement(ctx *base.VariableStatementContext) interface{} {
	// log.Println("VisitVariableStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

// classDeclaration
//     : Class identifier classTail
//     ;
func (v *Visitor) VisitClassDeclaration(ctx *base.ClassDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

// classTail
//     : (Extends singleExpression)? '{' classElement* '}'
//     ;
func (v *Visitor) VisitClassTail(ctx *base.ClassTailContext) interface{} {

	return v.VisitChildren(ctx)
}

// classElement
//     : (Static | {p.n("static")}? identifier | Async)* (methodDefinition | assignable '=' objectLiteral ';')
//     | emptyStatement_
//     | '#'? propertyName '=' singleExpression
//     ;
func (v *Visitor) VisitClassElement(ctx *base.ClassElementContext) interface{} {

	return v.VisitChildren(ctx)
}

// methodDefinition
//     : '*'? '#'? propertyName '(' formalParameterList? ')' functionBody
//     | '*'? '#'? getter '(' ')' functionBody
//     | '*'? '#'? setter '(' formalParameterList? ')' functionBody
//     ;
func (v *Visitor) VisitMethodDefinition(ctx *base.MethodDefinitionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParenthesizedExpression(ctx *base.ParenthesizedExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAwaitExpression(ctx *base.AwaitExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

// not a token
func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return ident(v, node.GetSymbol())

}

func ident(v *Visitor, token antlr.Token) *LToken {

	start, end := token.GetStart(), token.GetStop()+1

	return &LToken{
		sn:    v.symbolicNames[token.GetTokenType()],
		value: token.GetText(),
		SourceInfo: &SourceInfo{
			Line:   token.GetLine(),
			Column: token.GetColumn(),
			Start:  start,
			End:    end,
			Source: token.GetInputStream().GetTextFromInterval(
				&antlr.Interval{
					Start: start,
					Stop:  end,
				}),
		},
	}
}
