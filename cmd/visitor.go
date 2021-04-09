package main

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp"
)

type node struct {
	start, end, line int
	source           string
}

// need pointer receiver for methods...
type visitor struct {
	// any methods not implemented to satisfy JavaScriptParserVisitor checks in Accept...
	// JavaScriptParserVisitor embeds antlr.ParseTreeVisitor so we are also a "antlr.ParseTreeVisitor"
	// any visitor methods we don't add are called on BaseJavaScriptParserVisitor which are no-ops essentially -> nil
	gojsp.BaseJavaScriptParserVisitor
	// todo:  syntax errors with line/col ...
	errors  []string
	nodes   []node
	Imports []Import
	Expr    []Expression
}

func (v *visitor) getSourceInfo(ctx gojsp.BaseContext) node {
	return node{line: ctx.GetStart().GetLine(), start: ctx.GetStart().GetStart(), end: ctx.GetStop().GetStart(),
		source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: ctx.GetStart().GetStart(), Stop: ctx.GetStop().GetStop() + 1})}
}

func (v *visitor) VisitIdentifier(ctx *gojsp.IdentifierContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAssignable(ctx *gojsp.AssignableContext) interface{} {
	// log.Println("VisitAssignable", ctx.GetText())

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArgument(ctx *gojsp.ArgumentContext) interface{} {

	return v.VisitChildren(ctx)
}

// as per the docs but not sure if they will be used
func (v *visitor) defaultResult() interface{} {
	return nil
}
func (v *visitor) aggregateResult(aggregate interface{}, nextResult interface{}) interface{} {
	return nextResult
}
func (v *visitor) VisitStatement(ctx *gojsp.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitProgram(ctx *gojsp.ProgramContext) interface{} {

	// sourceElements as called when .Program() is used...
	return ctx.GetChild(0).(antlr.ParserRuleContext).Accept(v)

}

// Visit(tree ParseTree) interface{}
// VisitChildren(node RuleNode) interface{}
// VisitTerminal(node TerminalNode) interface{}
// VisitErrorNode(node ErrorNode) interface{}
func (v *visitor) VisitSourceElement(ctx *gojsp.SourceElementContext) interface{} {
	// log.Println("VisitSourceElement", ctx.GetText(), ctx.GetChildCount())
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
func (v *visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitSourceElements is called when production sourceElements is entered.
func (v *visitor) VisitSourceElements(ctx *gojsp.SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) shouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	return true
}

func (v *visitor) VisitChildren(node antlr.RuleNode) interface{} {

	// probably much better way...
	for _, ch := range node.GetChildren() {
		// todo: handle this EOF/;
		if ef, ok := ch.(*gojsp.EosContext); ok {
			// log.Println(ef)
			_ = ef
			continue
		}
		switch rr := ch.(antlr.ParseTree).Accept(v).(type) {

		case *gojsp.EosContext:
			log.Println(rr.GetText())
		case Expression:

			v.Expr = append(v.Expr, rr)
		case ImportDeclaration:

			log.Println(rr.ImportString())

			v.Imports = append(v.Imports, rr)

			// default:
			// 	panic(rr)

		}

	}

	return node

}

// not a token
func (v *visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}
func (v *visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	log.Println(node)
	return nil
}

func (v *visitor) VisitDeclaration(ctx *gojsp.DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitVariableStatement(ctx *gojsp.VariableStatementContext) interface{} {
	// log.Println("VisitVariableStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLabelledStatement(ctx *gojsp.LabelledStatementContext) interface{} {
	log.Println("VisitLabelledStatement")
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionDeclaration(ctx *gojsp.FunctionDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassDeclaration(ctx *gojsp.ClassDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassTail(ctx *gojsp.ClassTailContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassElement(ctx *gojsp.ClassElementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitMethodDefinition(ctx *gojsp.MethodDefinitionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFormalParameterList(ctx *gojsp.FormalParameterListContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFormalParameterArg(ctx *gojsp.FormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLastFormalParameterArg(ctx *gojsp.LastFormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionBody(ctx *gojsp.FunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrayLiteral(ctx *gojsp.ArrayLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitElementList(ctx *gojsp.ElementListContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrayElement(ctx *gojsp.ArrayElementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionProperty(ctx *gojsp.FunctionPropertyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyGetter(ctx *gojsp.PropertyGetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertySetter(ctx *gojsp.PropertySetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyShorthand(ctx *gojsp.PropertyShorthandContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyName(ctx *gojsp.PropertyNameContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArguments(ctx *gojsp.ArgumentsContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitMemberDotExpression(ctx *gojsp.MemberDotExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitFunctionExpression(ctx *gojsp.FunctionExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitIdentifierName(ctx *gojsp.IdentifierNameContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitParenthesizedExpression(ctx *gojsp.ParenthesizedExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitLiteral(ctx *gojsp.LiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitNumericLiteral(ctx *gojsp.NumericLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitObjectLiteralExpression(ctx *gojsp.ObjectLiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *visitor) VisitObjectLiteral(ctx *gojsp.ObjectLiteralContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitAwaitExpression(ctx *gojsp.AwaitExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitVariableDeclaration(ctx *gojsp.VariableDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitLet_(ctx *gojsp.Let_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitVariableDeclarationList(ctx *gojsp.VariableDeclarationListContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *visitor) VisitVarModifier(ctx *gojsp.VarModifierContext) interface{} {

	return v.VisitChildren(ctx)
}
