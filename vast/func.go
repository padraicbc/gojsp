package vast

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

// Async? arrowFunctionParameters '=>' arrowFunctionBody
// arrowFunctionParameters
//     : identifier
//     | '(' formalParameterList? ')'
//     ;
// arrowFunctionBody
//     : singleExpression
//     | functionBody
//     ;
type ArrowFunction struct {
	*SourceInfo
	Async              Token
	Arrow              Token
	FunctionParameters *ArrowFunctionParameters
	FunctionBody       *ArrowFunctionBody
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
	// if ctx.Async() != nil {
	// 	af.Async = ident(v, ctx.Async().GetSymbol())
	// 	af.children = af.Async
	// 	prev = af.Async
	// }
	// f := v.Visit(ctx.ArrowFunctionParameters()).(*ArrowFunctionParameters)
	// if prev != nil {
	// 	prev.Next(f)
	// } else {
	// 	af.children = f
	// }
	// f.prev = prev
	// prev = f
	// af.FunctionParameters = f

	// af.Arrow = ident(v, ctx.ARROW().GetSymbol())
	// af.FunctionBody = v.Visit(ctx.ArrowFunctionBody()).(*Arr)
	// log.Println("VisitArrowFunction", v.Visit(ctx.ArrowFunctionParameters()))
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		if af.children == nil {
			af.children = ch
		} else {
			prev.Next(ch)

		}
		ch.Prev(prev)
		prev = ch
		// log.Printf("%+v\n", ch)

		switch ch.Type() {
		case "ArrowFunctionBody":
			af.FunctionBody = ch.(*ArrowFunctionBody)
		case "ArrowFunctionParameters":
			af.FunctionParameters = ch.(*ArrowFunctionParameters)
		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {
			case "Async":
				af.Async = tk
			case "ARROW":
				af.Arrow = tk

			default:
				log.Panicf("%+v %s\n", ch, ch.Type())
			}

			// todo: check this better
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

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

	ar := &ArrowFunctionParameters{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

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

			case "OpenParen":
				ar.OpenParen = tk
			case "CloseParen":
				ar.CloseParen = tk
			case "Identifier":
				ar.Identifier = tk

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
type ArrowFunctionBody struct {
	*SourceInfo
	SingleExpression VNode

	FunctionBody *FunctionBody

	children   VNode
	prev, next VNode
}

var _ VNode = (*ArrowFunctionBody)(nil)

func (i *ArrowFunctionBody) Type() string {
	return "ArrowFunctionBody"
}
func (i *ArrowFunctionBody) Code() string {
	return CodeDef(i)
}
func (i *ArrowFunctionBody) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ArrowFunctionBody) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *ArrowFunctionBody) Children() []VNode {
	return children(i.children)
}

func (v *Visitor) VisitArrowFunctionBody(ctx *base.ArrowFunctionBodyContext) interface{} {
	af := &ArrowFunctionBody{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	if ctx.FunctionBody() != nil {

		ch := v.Visit(ctx.FunctionBody()).(*FunctionBody)
		af.children = ch
		af.FunctionBody = ch
		return af
	}
	// todo: validate and check if always should have flat element?
	ch := v.Visit(ctx.SingleExpression())
	if v, ok := ch.([]VNode); ok {
		af.children = v[0]
		af.SingleExpression = v[0]
		return af

	}
	af.children = ch.(VNode)
	af.SingleExpression = af.children
	return af

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
	typeOf              string
	prev, next          VNode
}

var _ VNode = (*FunctionDeclaration)(nil)

func (i *FunctionDeclaration) Type() string {
	return i.typeOf
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

	return fdecl(&FunctionDeclaration{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		typeOf:     "FunctionDeclaration"}, ctx, v)

}

func (v *Visitor) VisitFunctionDecl(ctx *base.FunctionDeclContext) interface{} {
	// log.Println(ctx)
	return fdecl(&FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		typeOf: "AnonymousFunctionDeclaration"}, ctx, v)
}

// same as any fun dec
// anoymousFunction
//     : functionDeclaration                                                       # FunctionDecl
//     | Async? Function '*'? '(' formalParameterList? ')' functionBody    # AnoymousFunctionDecl
//     | Async? arrowFunctionParameters '=>' arrowFunctionBody                     # ArrowFunction
//     ;
func (v *Visitor) VisitAnoymousFunctionDecl(ctx *base.AnoymousFunctionDeclContext) interface{} {
	// log.Println("VisitAnoymousFunctionDecl", ctx.FunctionBody().GetText())
	return fdecl(&FunctionDeclaration{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		typeOf:     "AnonymousFunctionDeclaration"}, ctx, v)
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
