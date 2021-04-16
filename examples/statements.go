package main

import (
	"log"

	"github.com/padraicbc/gojsp/parser"
)

func (v *visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	// log.Println("VisitStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

// statementList
//     : statement+
//     ;
// 	statementList
//     : statement+
//     ;
func (v *visitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	// log.Println("VisitStatementList", ctx.GetText())
	return v.VisitChildren(ctx)
}

// Block
//     A block statement is used to group zero or more statements. The block is delimited by a pair of curly brackets.
// block
//     : '{' statementList? '}'
//     ;

func (v *visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// log.Println("VisitBlock", ctx.GetText())
	return v.VisitChildren(ctx)
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/break
// break [label];
// label Optional
//     Identifier associated with the label of the statement. If the statement is not a loop or switch, this is required.

// Description
// The break statement includes an optional label that allows the program to break out of a labeled statement. The break statement needs to be nested within the referenced label. The labeled statement can be any block statement; it does not have to be preceded by a loop statement.
// A break statement, with or without a following label, cannot be used within the body of a function that is itself nested within the current loop, switch, or label statement that the break statement is intended to break out of.

// breakStatement
//     : Break ({p.notLineTerminator()}? identifier)? eos
//     ;
// don't think we care about Literal/Termianl node?
func (v *visitor) VisitBreak(ctx *parser.BreakStatementContext) interface{} {
	log.Println("VisitBreak", ctx.GetText())

	return v.VisitChildren(ctx)
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/continue
// continue [label];
// label
//     Identifier associated with the label of the statement.

// Description
// In contrast to the break statement, continue does not terminate the execution of the loop entirely: instead,
//     In a while loop, it jumps back to the condition.
//     In a for loop, it jumps to the update expression.

// The continue statement can include an optional label that allows the program to jump to the next iteration of a labeled loop statement instead of the current loop. In this case, the continue statement needs to be nested within this labeled statement.

// continueStatement
//     : Continue ({p.notLineTerminator()}? identifier)? eos
//     ;
func (v *visitor) VisitContinue(ctx *parser.ContinueStatementContext) interface{} {
	log.Println("VisitContinue", ctx.GetText())

	return v.VisitChildren(ctx)
}

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/Empty
// Syntax
// ;

// Description
// The empty statement is a semicolon (;) indicating that no statement will be executed, even if JavaScript syntax requires one.
// The opposite behavior, where you want multiple statements, but JavaScript only allows a single one, is possible using a block statement, which combines several statements into a single one.
type EmptyStatement struct {
	*SourceInfo
	Children []VNode
}

var _ VNode = (*EmptyStatement)(nil)

func (i *EmptyStatement) Type() string {
	return "EmptyStatement"
}
func (i *EmptyStatement) Code() string {
	return CodeDef(i)
}
func (i *EmptyStatement) GetChildren() []VNode {
	return i.Children
}

// Not sure what we should do with this
func (v *visitor) VisitEmpty(ctx *parser.EmptyStatement_Context) interface{} {
	log.Println("VisitEmpty", ctx.GetText())
	return &EmptyStatement{
		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

}

type IfStatement struct {
	*SourceInfo
	// Test       *ExpressionStatement
	// Consequent *Statement
	// Alternate  *Statement
	Children []VNode
}

var _ VNode = (*IfStatement)(nil)

func (i *IfStatement) Type() string {
	return "IfStatement"
}
func (i *IfStatement) Code() string {
	return CodeDef(i)
}
func (i *IfStatement) GetChildren() []VNode {
	return i.Children
}

// ifStatement
//     : If '(' expressionSequence ')' statement (Else statement)?
//     ;
//     expressionSequence
//         : singleExpression (',' singleExpression)*
//         ;
func (v *visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	log.Println("VisitIfStatement", ctx.GetText())

	ifs := &IfStatement{

		Children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	return ifs
}
