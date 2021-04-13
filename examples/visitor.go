package main

import (
	"log"

	antlr "github.com/padraicbc/antlr4"

	"github.com/padraicbc/gojsp/parser"
)

// need pointer receiver for methods...
type visitor struct {
	// any methods not implemented to satisfy JavaScriptParserVisitor checks in Accept...
	// JavaScriptParserVisitor embeds antlr.ParseTreeVisitor so we are also a "antlr.ParseTreeVisitor"
	// any visitor methods we don't add are called on BaseJavaScriptParserVisitor which are no-ops essentially -> nil
	parser.BaseJavaScriptParserVisitor
	// todo:  syntax errors with line/col ...
	// errors []string

	ParseTree *PTree
}

func (v *visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAssignable(ctx *parser.AssignableContext) interface{} {
	// log.Println("VisitAssignable", ctx.GetText())

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {

	return v.VisitChildren(ctx)
}

// as per the docs but not sure if they will be used
func (v *visitor) defaultResult() interface{} {
	return nil
}
func (v *visitor) aggregateResult(aggregate interface{}, nextResult interface{}) interface{} {
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
func (v *visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitSourceElements is called when production sourceElements is entered.
func (v *visitor) VisitSourceElements(ctx *parser.SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) shouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	return true
}

// todo: forget this and implement method to return []VNode specifically
func (v *visitor) VisitChildren(node antlr.RuleNode) interface{} {

	var result []VNode

	for _, ch := range node.GetChildren() {
		// todo: handle this EOF/;
		if ef, ok := ch.(*parser.EosContext); ok {
			// log.Println(ef)
			_ = ef
			// result.Eos = ef.GetText()
			break
		}

		res := ch.(antlr.ParseTree).Accept(v)
		switch rr := res.(type) {
		// case *Expression:
		// 	result = append(result.([]*Expression), rr)
		// case *ImportFromBlock:
		// 	result = append(result.([]*ImportFromBlock), rr)
		case VNode:

			result = append(result, rr)
			// v.ParseTree.LastChild.Children = append(v.ParseTree.LastChild.Children, rr)

			// default:
			// 	log.Println(rr)

		}

	}

	return result

}

// not a token
func (v *visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}
func (v *visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	log.Println(node)
	return nil
}

func (v *visitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitVariableStatement(ctx *parser.VariableStatementContext) interface{} {
	// log.Println("VisitVariableStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

// special case for $: ... todo: a type
func (v *visitor) VisitLabelledStatement(ctx *parser.LabelledStatementContext) interface{} {
	log.Println("VisitLabelledStatement", ctx.GetText())

	if ctx.Identifier().GetText() == "$" {
		log.Println("Reactive?")

	}
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassTail(ctx *parser.ClassTailContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassElement(ctx *parser.ClassElementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitMethodDefinition(ctx *parser.MethodDefinitionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFormalParameterArg(ctx *parser.FormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLastFormalParameterArg(ctx *parser.LastFormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionBody(ctx *parser.FunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitElementList(ctx *parser.ElementListContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrayElement(ctx *parser.ArrayElementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionProperty(ctx *parser.FunctionPropertyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyGetter(ctx *parser.PropertyGetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertySetter(ctx *parser.PropertySetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyShorthand(ctx *parser.PropertyShorthandContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyName(ctx *parser.PropertyNameContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitMemberDotExpression(ctx *parser.MemberDotExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitFunctionExpression(ctx *parser.FunctionExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitIdentifierName(ctx *parser.IdentifierNameContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitParenthesizedExpression(ctx *parser.ParenthesizedExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitNumericLiteral(ctx *parser.NumericLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitObjectLiteralExpression(ctx *parser.ObjectLiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitObjectLiteral(ctx *parser.ObjectLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAwaitExpression(ctx *parser.AwaitExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitLet_(ctx *parser.Let_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitVariableDeclarationList(ctx *parser.VariableDeclarationListContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitVarModifier(ctx *parser.VarModifierContext) interface{} {

	return v.VisitChildren(ctx)
}
