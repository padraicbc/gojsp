package main

import (
	"fmt"
	"log"
	"strings"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/parserv"
)

type visitor struct {
	*parserv.BaseJavaScriptParserVisitor
	// todo:  syntax errors with line/col ...
	errors []string
}

// as per the docs but not sure if they will be used
func (v *visitor) defaultResult() interface{} {
	return nil
}
func (v *visitor) aggregateResult(aggregate interface{}, nextResult interface{}) interface{} {
	return nextResult
}
func (v *visitor) VisitStatement(ctx *parserv.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitProgram(ctx *parserv.ProgramContext) interface{} {

	// sourceElements as called when .Program() is used...
	return ctx.GetChild(0).(antlr.ParserRuleContext).Accept(v)

}

// Visit(tree ParseTree) interface{}
// VisitChildren(node RuleNode) interface{}
// VisitTerminal(node TerminalNode) interface{}
// VisitErrorNode(node ErrorNode) interface{}
func (v *visitor) VisitSourceElement(ctx *parserv.SourceElementContext) interface{} {
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
func (v *visitor) VisitSourceElements(ctx *parserv.SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) shouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	return true
}

func (v *visitor) VisitChildren(node antlr.RuleNode) interface{} {

	var result = []string{}
	// probably much better way...
	for _, ch := range node.GetChildren() {
		if !v.shouldVisitNextChild(node, result) {
			return result
		}
		if ef, ok := ch.(*parserv.EosContext); ok {
			result = append(result, ef.GetText())
			continue
		}

		switch rr := ch.(antlr.ParseTree).Accept(v).(type) {
		case string:
			result = append(result, rr)
		case []string:
			result = append(result, rr...)
		case *parserv.EosContext:
			result = append(result, rr.GetText())
		case nil:

		default:
			panic(rr)

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

func (v *visitor) VisitImportExpression(ctx *parserv.ImportExpressionContext) interface{} {

	log.Println("VisitImportExpression", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitImportStatement(ctx *parserv.ImportStatementContext) interface{} {
	var st string
	// we could format based on some spec. Could be done here, at each step or at the very end...
	for _, nd := range v.VisitChildren(ctx).([]string) {
		if nd == "," || nd == ";" {
			st += nd
			continue
		}

		st += " " + nd

	}
	return st
}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
//     ;
func (v *visitor) VisitImportFromBlock(ctx *parserv.ImportFromBlockContext) interface{} {

	// todo: error as can't have both?
	// if 	ctx.ImportNamespace() != nil && ctx.ModuleItems() != nil{

	// }

	return v.VisitChildren(ctx)
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
// just pass on the wor if nothign to change
func (v *visitor) VisitModuleItems1(ctx *parserv.ModuleItemsContext) interface{} {

	return v.VisitChildren(ctx)
}

// call other visit method directly...
func (v *visitor) VisitModuleItems2(ctx *parserv.ModuleItemsContext) interface{} {

	var out string
	for i, mc := range ctx.AllAliasName() {
		out += v.VisitAliasName(mc.(*parserv.AliasNameContext)).(string)
		// add comma if more than one...
		if c := ctx.Comma(i); c != nil {
			out += ", "

		}

	}
	//todo: validate syntax... ctx.OpenBrace().GetText() ctx.CloseBrace().GetText()
	return fmt.Sprintf("{%s}", strings.TrimSpace(out))

}

//  alternative version where we do the AliasName work ourselves so we can change...
func (v *visitor) VisitModuleItems(ctx *parserv.ModuleItemsContext) interface{} {

	var out string
	for i, mc := range ctx.AllAliasName() {
		tmp := []string{}
		// always this
		actx := (mc.(*parserv.AliasNameContext))
		if ident := actx.IdentifierName(0); ident != nil {
			tmp = append(tmp, "changed"+fmt.Sprint(i)) // ident.GetText())

		}
		if as := actx.As(); as != nil {
			// can just use "as" .. vs actx.As().GetText(
			tmp = append(tmp, "as")
			// todo: again syntax error if nil as . is a syntax error
			if ident := actx.IdentifierName(1); ident != nil {
				tmp = append(tmp, "a new alias") //ident.GetText())

			}

		}
		// each moduleitems..
		out += strings.Join(tmp, " ")

		// add comma if more than one...
		if c := ctx.Comma(i); c != nil {
			out += ", "

		}

	}
	// . ctx.OpenBrace().GetText() ctx.CloseBrace().GetText()?
	return fmt.Sprintf("{%s}", out)
}
func (v *visitor) VisitImportDefault(ctx *parserv.ImportDefaultContext) interface{} {
	// log.Println("VisitImportDefault", ctx.GetText())
	return v.VisitChildren(ctx)
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
func (v *visitor) VisitImportNamespace(ctx *parserv.ImportNamespaceContext) interface{} {
	// log.Println("VisitImportNamespace", ctx.GetText())
	// todo: add to errors
	if ctx.GetChildCount() != 1 && ctx.GetChildCount() != 3 {
		panic("VisitImportNamespace shoudl have 1 or 3 children but had " + fmt.Sprint(ctx.GetChildCount()))
	}
	var out = []string{}
	// todo: validate syntax
	// either this or *parserv.IdentifierNameContext i.e * or any identifier
	if vv, ok := ctx.GetChild(0).(*antlr.TerminalNodeImpl); ok {
		out = append(out, vv.GetText())

	}

	if as := ctx.As(); as != nil {
		// can just use "as" ..
		out = append(out, "as")
		// todo: again syntax error if nil as . is a syntax error
		if ident := ctx.IdentifierName(1); ident != nil {
			out = append(out, ident.GetText())
		}

	}
	// out should be "as " here so check later
	if ident := ctx.IdentifierName(0); ident != nil {
		out = append(out, ident.GetText())
	}
	return out

}

// importFrom
//     : From StringLiteral
//     ;
func (v *visitor) VisitImportFrom(ctx *parserv.ImportFromContext) interface{} {

	if ctx.GetChildCount() != 2 {
		// todo: error
		// ctx.From().GetSymbol().GetLine()
		panic("wrong child count for importfrom")
	}

	// could do work here also ctx.From() and ctx.GetChild(1) -> path
	return v.VisitChildren(ctx)

}

// aliasName
//     : identifierName (As identifierName)?
//     ;s
func (v *visitor) VisitAliasName(ctx *parserv.AliasNameContext) interface{} {
	// log.Println("VisitAliasName", ctx.GetChildCount(), ctx.GetText())
	var out = []string{}
	// todo: syntax error if nil
	if ident := ctx.IdentifierName(0); ident != nil {
		out = append(out, ident.GetText())

	}
	if as := ctx.As(); as != nil {
		// can just use "as" .. vs ctx.As().GetText(
		out = append(out, "as")
		// todo: again syntax error if nil as . is a syntax error
		if ident := ctx.IdentifierName(1); ident != nil {
			out = append(out, ident.GetText())

		}

	}

	return strings.Join(out, " ")
}

func (v *visitor) VisitExportDeclaration(ctx *parserv.ExportDeclarationContext) interface{} {
	// log.Println("VisitExportDeclaration", ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExportDefaultDeclaration(ctx *parserv.ExportDefaultDeclarationContext) interface{} {
	log.Println("VisitExportDefaultDeclaration", ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitExportFromBlock(ctx *parserv.ExportFromBlockContext) interface{} {
	log.Println("VisitExportFromBlock", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitDeclaration(ctx *parserv.DeclarationContext) interface{} {
	// log.Println("VisitDeclaration", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitVariableStatement(ctx *parserv.VariableStatementContext) interface{} {
	// log.Println("VisitVariableStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLabelledStatement(ctx *parserv.LabelledStatementContext) interface{} {
	log.Println("VisitLabelledStatement")
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionDeclaration(ctx *parserv.FunctionDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassDeclaration(ctx *parserv.ClassDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassTail(ctx *parserv.ClassTailContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitClassElement(ctx *parserv.ClassElementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitMethodDefinition(ctx *parserv.MethodDefinitionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFormalParameterList(ctx *parserv.FormalParameterListContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFormalParameterArg(ctx *parserv.FormalParameterArgContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitLastFormalParameterArg(ctx *parserv.LastFormalParameterArgContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionBody(ctx *parserv.FunctionBodyContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrayLiteral(ctx *parserv.ArrayLiteralContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitElementList(ctx *parserv.ElementListContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArrayElement(ctx *parserv.ArrayElementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyExpressionAssignment(ctx *parserv.PropertyExpressionAssignmentContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitComputedPropertyExpressionAssignment(ctx *parserv.ComputedPropertyExpressionAssignmentContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitFunctionProperty(ctx *parserv.FunctionPropertyContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyGetter(ctx *parserv.PropertyGetterContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertySetter(ctx *parserv.PropertySetterContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyShorthand(ctx *parserv.PropertyShorthandContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitPropertyName(ctx *parserv.PropertyNameContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitArguments(ctx *parserv.ArgumentsContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}
