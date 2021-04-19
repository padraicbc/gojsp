package main

import (
	"log"
	"reflect"

	antlr "github.com/padraicbc/antlr4"

	"github.com/padraicbc/gojsp/parser"
)

// need pointer receiver for methods...
type Visitor struct {
	// any methods not implemented to satisfy JavaScriptParserVisitor checks in Accept...
	// JavaScriptParserVisitor embeds antlr.ParseTreeVisitor so we are also a "antlr.ParseTreeVisitor"
	// any visitor methods we don't add are called on BaseJavaScriptParserVisitor which are no-ops essentially -> nil
	parser.BaseJavaScriptParserVisitor
	// todo:  syntax errors with line/col ...
	// errors []string

	ParseTree *PTree
	lexer     *parser.JavaScriptLexer
	parser    *parser.JavaScriptParser
}

func NewVisitor(lexer *parser.JavaScriptLexer, parser *parser.JavaScriptParser) *Visitor {
	return &Visitor{lexer: lexer, parser: parser, ParseTree: &PTree{}}
}

func (v *Visitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {

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
func (v *Visitor) VisitSourceElements(ctx *parser.SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) shouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	return true
}

// todo: build tree from result nodes not each iteration and remove dupes
func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {

	var result []VNode
	prev := v.ParseTree.LastChild
	for _, ch := range node.GetChildren() {

		res := ch.(antlr.ParseTree).Accept(v)

		// first node
		switch rr := res.(type) {

		case *LToken:
			rr.rn = v.parser.GetRuleNames()[node.GetRuleContext().GetRuleIndex()]
			result = append(result, rr)

		case VNode:
			result = append(result, rr)
			v.ParseTree.LastChild = rr
			if prev != nil {
				prev.Next(rr)
				rr.Prev(prev)
			}
			prev = rr
		case []VNode:
			result = append(result, rr...)
		case nil:
			// panic(rr)
		default:

			panic(reflect.TypeOf(rr))

		}

	}

	return result

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

func (v *Visitor) VisitVariableStatement(ctx *parser.VariableStatementContext) interface{} {
	// log.Println("VisitVariableStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

// functionDeclaration
//     : Async? Function '*'? identifier '(' formalParameterList? ')' functionBody
//     ;
func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

// classDeclaration
//     : Class identifier classTail
//     ;
func (v *Visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

// classTail
//     : (Extends singleExpression)? '{' classElement* '}'
//     ;
func (v *Visitor) VisitClassTail(ctx *parser.ClassTailContext) interface{} {

	return v.VisitChildren(ctx)
}

// classElement
//     : (Static | {p.n("static")}? identifier | Async)* (methodDefinition | assignable '=' objectLiteral ';')
//     | emptyStatement_
//     | '#'? propertyName '=' singleExpression
//     ;
func (v *Visitor) VisitClassElement(ctx *parser.ClassElementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodDefinition(ctx *parser.MethodDefinitionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFormalParameterArg(ctx *parser.FormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLastFormalParameterArg(ctx *parser.LastFormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionBody(ctx *parser.FunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionProperty(ctx *parser.FunctionPropertyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyGetter(ctx *parser.PropertyGetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertySetter(ctx *parser.PropertySetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyShorthand(ctx *parser.PropertyShorthandContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyName(ctx *parser.PropertyNameContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitMemberDotExpression(ctx *parser.MemberDotExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitFunctionExpression(ctx *parser.FunctionExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitParenthesizedExpression(ctx *parser.ParenthesizedExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNumericLiteral(ctx *parser.NumericLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitObjectLiteralExpression(ctx *parser.ObjectLiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitObjectLiteral(ctx *parser.ObjectLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitAwaitExpression(ctx *parser.AwaitExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

// not a token
func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {

	return ident(v, node.GetSymbol())

}

func ident(v *Visitor, token antlr.Token) *LToken {

	start, end := token.GetStart(), token.GetStop()+1

	return &LToken{
		sn:    v.lexer.SymbolicNames[token.GetTokenType()],
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

// func (v *BaseJavaScriptParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
// 	return v.VisitChildren(ctx)
// }
