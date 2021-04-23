package vast

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

// Async? arrowFunctionParameters '=>' arrowFunctionBody
type ArrowFunction struct {
	*SourceInfo
	Async              Token
	Arrow              Token
	FunctionParameters *ArrowFunctionParameters
	FunctionBody       VNode
	children           VNode
	prev, next         VNode
}

var _ VNode = (*ArrowFunction)(nil)

func (i *ArrowFunction) Type() string {
	return "ArrowFunction"
}
func (i *ArrowFunction) Code() string {
	return CodeDef(i)
}
func (i *ArrowFunction) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ArrowFunction) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *ArrowFunction) Children() []VNode {
	return children(i.children)
}

func (v *Visitor) VisitArrowFunction(ctx *base.ArrowFunctionContext) interface{} {
	af := &ArrowFunction{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if af.children == nil {
			af.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch

		switch ch.Type() {
		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {

			case "Async":
				af.Async = tk
			case "ARROW":
				af.Arrow = tk

			default:
				log.Panicf("%+v %s\n", ch, ch.Type())
			}

		case "ArrowFunctionParameters":
			af.FunctionParameters = ch.(*ArrowFunctionParameters)
			// todo: check this better
		default:
			log.Printf("%s\n", ch.Type())
			af.FunctionBody = ch

		}
	}
	return af
}

// arrowFunctionParameters
//     : identifier
//     | '(' formalParameterList? ')'
//     ;
type ArrowFunctionParameters struct {
	*SourceInfo
	Identifier Token
	OpenParen  Token

	FormalParameterList *FormalParameterList
	CloseParen          Token
	children            VNode
	prev, next          VNode
}

var _ VNode = (*ArrowFunctionParameters)(nil)

func (i *ArrowFunctionParameters) Type() string {
	return "ArrowFunctionParameters"
}
func (i *ArrowFunctionParameters) Code() string {
	return CodeDef(i)
}
func (i *ArrowFunctionParameters) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ArrowFunctionParameters) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *ArrowFunctionParameters) Children() []VNode {
	return children(i.children)
}

func (v *Visitor) VisitArrowFunctionParameters(ctx *base.ArrowFunctionParametersContext) interface{} {
	ar := &ArrowFunctionParameters{}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ar.children == nil {
			ar.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch

		switch ch.Type() {
		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {

			case "Identifier":
				ar.Identifier = tk
			case "OpenParen":
				ar.OpenParen = tk
			case "CloseParen":
				ar.CloseParen = tk

			default:
				log.Panicf("%+v %s\n", ch, ch.Type())
			}

		case "FormalParameterList":
			ar.FormalParameterList = ch.(*FormalParameterList)

		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}

	return ar
}

// arrowFunctionBody
//     : singleExpression
//     | functionBody
//     ;
func (v *Visitor) VisitArrowFunctionBody(ctx *base.ArrowFunctionBodyContext) interface{} {
	if ctx.FunctionBody() != nil {
		return v.Visit(ctx.FunctionBody())
	}
	return v.Visit(ctx.SingleExpression())
}

// function name([param[, param[, ... param]]]) {
// 	statements
//  }
// functionDeclaration
//     : Async? Function '*'? identifier '(' formalParameterList? ')' functionBody
//     ;
type FunctionDeclaration struct {
	*SourceInfo
	Async        Token
	Function     Token
	Star         Token
	FunctionBody *FunctionBody

	Identifier          Token
	OpenParen           Token
	FormalParameterList *FormalParameterList
	CloseParen          Token
	children            VNode
	prev, next          VNode
}

var _ VNode = (*FunctionDeclaration)(nil)

func (i *FunctionDeclaration) Type() string {
	return "FunctionDeclaration"
}
func (i *FunctionDeclaration) Code() string {
	return CodeDef(i)
}
func (i *FunctionDeclaration) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *FunctionDeclaration) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *FunctionDeclaration) Children() []VNode {
	return children(i.children)
}

func fdecl(fd *FunctionDeclaration, ctx antlr.RuleNode, v *Visitor) interface{} {
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fd.children == nil {
			fd.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch

		switch ch.Type() {
		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {
			case "Multiply":
				fd.Star = tk
			case "Function":
				fd.Function = tk
			case "Identifier":
				fd.Identifier = tk
			case "OpenParen":
				fd.OpenParen = tk
			case "CloseParen":
				fd.CloseParen = tk
			case "Async":
				fd.Async = tk
			default:
				log.Panicf("%+v %s\n", ch, ch.Type())
			}

		case "FormalParameterList":
			fd.FormalParameterList = ch.(*FormalParameterList)
		case "FunctionBody":
			fd.FunctionBody = ch.(*FunctionBody)
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}
	return fd
}

// Visit a parse tree produced by JavaScriptParser#functionDeclaration.
func (v *Visitor) VisitFunctionDeclaration(ctx *base.FunctionDeclarationContext) interface{} {

	return fdecl(&FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}, ctx, v)

}

func (v *Visitor) VisitFunctionDecl(ctx *base.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return fdecl(&FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}, ctx, v)
}

// same as any fun dec
// anoymousFunction
//     : functionDeclaration                                                       # FunctionDecl
//     | Async? Function '*'? '(' formalParameterList? ')' functionBody    # AnoymousFunctionDecl
//     | Async? arrowFunctionParameters '=>' arrowFunctionBody                     # ArrowFunction
//     ;
func (v *Visitor) VisitAnoymousFunctionDecl(ctx *base.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return fdecl(&FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}, ctx, v)
}

// formalParameterList
//     : formalParameterArg (',' formalParameterArg)* (',' lastFormalParameterArg)?
//     | lastFormalParameterArg
//     ;
type FormalParameterList struct {
	*SourceInfo
	FormalParameterArgs    []*FormalParameterArg
	LastFormalParameterArg *LastFormalParameterArg
	Commas                 []Token

	children   VNode
	prev, next VNode
}

var _ VNode = (*FormalParameterList)(nil)

func (i *FormalParameterList) Type() string {
	return "FormalParameterList"
}
func (i *FormalParameterList) Code() string {
	return CodeDef(i)
}
func (i *FormalParameterList) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *FormalParameterList) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *FormalParameterList) Children() []VNode {
	return children(i.children)
}

func (v *Visitor) VisitFormalParameterList(ctx *base.FormalParameterListContext) interface{} {
	fp := &FormalParameterList{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fp.children == nil {
			fp.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch

		switch ch.Type() {
		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {
			case "Comma":
				fp.Commas = append(fp.Commas, tk)

			default:
				log.Panicf("%+v\n", ch)
			}
		case "FormalParameterArg":
			fp.FormalParameterArgs = append(fp.FormalParameterArgs, ch.(*FormalParameterArg))
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}
	return fp
}

// formalParameterArg
//     : assignable ('=' singleExpression)?      // ECMAScript 6: Initialization
//     ;
type FormalParameterArg struct {
	*SourceInfo
	Assignable       VNode
	Equals           Token
	SingleExpression VNode
	children         VNode
	prev, next       VNode
}

var _ VNode = (*FormalParameterArg)(nil)

func (i *FormalParameterArg) Type() string {
	return "FormalParameterArg"
}
func (i *FormalParameterArg) Code() string {
	return CodeDef(i)
}
func (i *FormalParameterArg) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *FormalParameterArg) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *FormalParameterArg) Children() []VNode {
	return children(i.children)
}
func (v *Visitor) VisitFormalParameterArg(ctx *base.FormalParameterArgContext) interface{} {
	fa := &FormalParameterArg{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode

	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fa.children == nil {
			fa.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch
		switch tk := ch.(Token); tk.SymbolName() {
		case "Identifier":
			fa.Assignable = tk
		case "SingleExpression":
			fa.SingleExpression = tk
		case "Assign":
			fa.Equals = tk
		default:
			log.Panicf("%+v\n", tk)
		}

	}
	return fa
}

// lastFormalParameterArg                        // ECMAScript 6: Rest Parameter
//     : Ellipsis singleExpression
//     ;
type LastFormalParameterArg struct {
	*SourceInfo
	Ellipsis         Token
	SingleExpression VNode
	children         VNode
	prev, next       VNode
}

var _ VNode = (*LastFormalParameterArg)(nil)

func (i *LastFormalParameterArg) Type() string {
	return "LastFormalParameterArg"
}
func (i *LastFormalParameterArg) Code() string {
	return CodeDef(i)
}
func (i *LastFormalParameterArg) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *LastFormalParameterArg) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *LastFormalParameterArg) Children() []VNode {
	return children(i.children)
}
func (v *Visitor) VisitLastFormalParameterArg(ctx *base.LastFormalParameterArgContext) interface{} {

	return v.VisitChildren(ctx)
}

// functionBody
//     : '{' sourceElements? '}'
//     ;
type FunctionBody struct {
	*SourceInfo
	OpenBrace Token
	// statements
	SourceElements []VNode
	CloseBrace     Token
	Eos            Token //
	children       VNode
	prev, next     VNode
}

var _ VNode = (*FunctionBody)(nil)

func (i *FunctionBody) Type() string {
	return "FunctionBody"
}
func (i *FunctionBody) Code() string {
	return CodeDef(i)
}
func (i *FunctionBody) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *FunctionBody) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *FunctionBody) Children() []VNode {
	return children(i.children)
}
func (v *Visitor) VisitFunctionBody(ctx *base.FunctionBodyContext) interface{} {
	fb := &FunctionBody{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fb.children == nil {
			fb.children = ch
		} else {
			prev.Next(ch)

		}

		ch.Prev(prev)
		prev = ch

		switch ch.Type() {
		case "LToken":
			tk := ch.(Token)
			switch tk.SymbolName() {
			case "OpenBrace":
				fb.OpenBrace = tk
			case "CloseBrace":
				fb.CloseBrace = tk
			case "SemiColon":
				fb.Eos = tk

			default:
				log.Panicf("%+v\n", ch)
			}

		default:
			// ExpressionStatement ...
			// log.Printf("%+v %s\n", ch, ch.Type())
			fb.SourceElements = append(fb.SourceElements, ch)

		}

	}
	return fb
}

func (v *Visitor) VisitFunctionProperty(ctx *base.FunctionPropertyContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArguments(ctx *base.ArgumentsContext) interface{} {

	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitMemberDotExpression(ctx *base.MemberDotExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

// singleExpression
//     : anoymousFunction                                                      # FunctionExpression
func (v *Visitor) VisitFunctionExpression(ctx *base.FunctionExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
