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
