package vast

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

// need pointer receiver for methods...
type Visitor struct {
	// any methods not implemented to satisfy JavaScriptParserVisitor checks in Accept...
	// JavaScriptParserVisitor embeds antlr.ParseTreeVisitor so we are also a "antlr.ParseTreeVisitor"
	// any visitor methods we don't add are called on BaseJavaScriptParserVisitor which are no-ops essentially -> nil
	base.BaseJavaScriptParserVisitor
	// todo:  syntax errors with line/col ...
	// errors []string

	ruleNames     []string
	symbolicNames []string
	Debug         bool
}

func NewVisitor(symbolicNames []string, ruleNames []string) *Visitor {
	return &Visitor{symbolicNames: symbolicNames, ruleNames: ruleNames}
}

func (v *Visitor) VisitArgument(ctx *base.ArgumentContext) interface{} {

	return v.VisitChildren(ctx)
}

// as per the docs but not sure if they will be used
func (v *Visitor) defaultResult() interface{} {
	return nil
}
func (v *Visitor) aggregateResult(aggregate interface{}, nextResult interface{}) interface{} {
	return nextResult
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
	log.Println(node)
	return nil
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
