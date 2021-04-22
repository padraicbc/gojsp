package vast

import (
	"fmt"
	"log"

	"github.com/padraicbc/gojsp/base"
)

// interface Function <: Node {
//     id: Identifier | null;
//     params: [ Pattern ];
//     body: FunctionBody;
// }
type Function struct {
}

func (v *Visitor) VisitArrowFunction(ctx *base.ArrowFunctionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowFunctionParameters(ctx *base.ArrowFunctionParametersContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrowFunctionBody(ctx *base.ArrowFunctionBodyContext) interface{} {

	return v.VisitChildren(ctx)
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

// Visit a parse tree produced by JavaScriptParser#functionDeclaration.
func (v *Visitor) VisitFunctionDeclaration(ctx *base.FunctionDeclarationContext) interface{} {
	fd := &FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		switch ch.Type() {
		case "LToken":
			tk := ch.(Token)
			switch tk.SymbolName() {
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
				log.Panic(fmt.Sprintf("%+v\n", ch))
			}
		// case ""
		case "FormalParameterList":
			fd.FormalParameterList = ch.(*FormalParameterList)
		default:
			// log.Printf("%+v\n", ch)

		}
	}
	return fd
}

func (v *Visitor) VisitFunctionDecl(ctx *base.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnoymousFunctionDecl(ctx *base.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return v.VisitChildren(ctx)
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
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		switch ch.Type() {
		case "LToken":
			tk := ch.(Token)
			switch tk.SymbolName() {
			case "Comma":
				fp.Commas = append(fp.Commas, tk)

			default:
				log.Println(fmt.Sprintf("%+v\n", ch))
			}
		case "FormalParameterList":
			fp.FormalParameterArgs = append(fp.FormalParameterArgs, ch.(*FormalParameterArg))
		default:
			// log.Println(ch)

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
		// log.Printf("%+vz\n", ch)

		// if ch.Type() == "LToken" {
		// 	fa.faport = ch.(Token)

		// } else {
		// 	im.ImportFromBlock = ch.(*ImportFromBlock)
		// }

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
	OpenBrace      Token
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
				log.Println(fmt.Sprintf("%+v\n", ch))
			}

		default:
			log.Println(fmt.Sprintf("%+v\n", ch))
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
func (v *Visitor) VisitFunctionExpression(ctx *base.FunctionExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}
