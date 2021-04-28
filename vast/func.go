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
	Body               *ArrowFunctionBody
	firstChild         VNode

	prev, next VNode
}

var _ VNode = (*ArrowFunction)(nil)

func (i *ArrowFunction) Type() string {
	return "ArrowFunction"
}
func (i *ArrowFunction) Code() string {
	return CodeDef(i)
}
func (i *ArrowFunction) Next() VNode {

	return i.next
}
func (i *ArrowFunction) SetNext(v VNode) {
	i.next = v
}
func (i *ArrowFunction) Prev() VNode {

	return i.prev
}
func (i *ArrowFunction) SetPrev(v VNode) {
	i.prev = v
}

func (i *ArrowFunction) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitArrowFunction(ctx *base.ArrowFunctionContext) interface{} {
	if v.Debug {
		log.Println("VisitArrowFunction", ctx.GetText())
	}
	af := &ArrowFunction{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode

	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if af.firstChild == nil {
			af.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch
		// log.Printf("%+v\n", ch)

		switch ch.Type() {
		case "ArrowFunctionBody":
			af.Body = ch.(*ArrowFunctionBody)
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
			// log.Printf("%s\n", ch.Type())
			af.Body = ch.(*ArrowFunctionBody)

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
	firstChild          VNode

	prev, next VNode
}

var _ VNode = (*ArrowFunctionParameters)(nil)

func (i *ArrowFunctionParameters) Type() string {
	return "ArrowFunctionParameters"
}
func (i *ArrowFunctionParameters) Code() string {
	return CodeDef(i)
}
func (i *ArrowFunctionParameters) Next() VNode {

	return i.next
}
func (i *ArrowFunctionParameters) SetNext(v VNode) {
	i.next = v
}
func (i *ArrowFunctionParameters) Prev() VNode {

	return i.prev
}
func (i *ArrowFunctionParameters) SetPrev(v VNode) {
	i.prev = v
}

func (i *ArrowFunctionParameters) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitArrowFunctionParameters(ctx *base.ArrowFunctionParametersContext) interface{} {
	if v.Debug {
		log.Println("VisitArrowFunctionParameters", ctx.GetText())
	}
	ar := &ArrowFunctionParameters{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ar.firstChild == nil {
			ar.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

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
	FBody            *FunctionBody
	CloseParen       Token
	firstChild       VNode

	prev, next VNode
}

var _ VNode = (*ArrowFunctionBody)(nil)

func (i *ArrowFunctionBody) Type() string {
	return "ArrowFunctionBody"
}
func (i *ArrowFunctionBody) Code() string {
	return CodeDef(i)
}
func (i *ArrowFunctionBody) Next() VNode {

	return i.next
}
func (i *ArrowFunctionBody) SetNext(v VNode) {
	i.next = v
}
func (i *ArrowFunctionBody) Prev() VNode {

	return i.prev
}
func (i *ArrowFunctionBody) SetPrev(v VNode) {
	i.prev = v
}

func (i *ArrowFunctionBody) FirstChild() VNode {

	return i.firstChild
}

func (v *Visitor) VisitArrowFunctionBody(ctx *base.ArrowFunctionBodyContext) interface{} {
	if v.Debug {
		log.Println("VisitArrowFunctionBody", ctx.GetText())
	}
	afb := &ArrowFunctionBody{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if afb.firstChild == nil {
			afb.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch
		switch ch.Type() {
		case "FunctionBody":
			afb.FBody = ch.(*FunctionBody)
		default:
			afb.SingleExpression = ch
			// log.Printf("%+v %+v", ch, ch.Type())
		}

	}
	return afb
}

// todo: maybe create types for anon ec.. separate
// functionDeclaration
//     : Async? Function '*'? identifier '(' formalParameterList? ')' functionBody
//     ;
type FunctionDeclaration struct {
	*SourceInfo
	Async        Token
	Function     Token
	Star         Token
	FunctionBody *FunctionBody
	typeOf       string

	Identifier Token
	OpenParen  Token
	PList      *FormalParameterList
	CloseParen Token
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*FunctionDeclaration)(nil)

func (i *FunctionDeclaration) Type() string {
	return i.typeOf
}
func (i *FunctionDeclaration) Code() string {
	return CodeDef(i)
}
func (i *FunctionDeclaration) Next() VNode {

	return i.next
}
func (i *FunctionDeclaration) SetNext(v VNode) {
	i.next = v
}
func (i *FunctionDeclaration) Prev() VNode {

	return i.prev
}
func (i *FunctionDeclaration) SetPrev(v VNode) {
	i.prev = v
}

func (i *FunctionDeclaration) FirstChild() VNode {

	return i.firstChild

}

func fdecl(fd *FunctionDeclaration, ctx antlr.RuleNode, v *Visitor) interface{} {

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fd.firstChild == nil {
			fd.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

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
			fd.PList = ch.(*FormalParameterList)
		case "FunctionBody":
			fd.FunctionBody = ch.(*FunctionBody)
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}
	return fd
}

// functionDeclaration
// : Async? Function '*'? identifier '(' formalParameterList? ')' functionBody
// ;
func (v *Visitor) VisitFunctionDeclaration(ctx *base.FunctionDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionDeclaration", ctx.GetText())
	}

	return fdecl(&FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext), typeOf: "FunctionDeclaration"}, ctx, v)

}

func (v *Visitor) VisitFunctionDecl(ctx *base.FunctionDeclContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionDecl", ctx.GetText())
	}
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
	if v.Debug {
		log.Println("VisitAnoymousFunctionDecl", ctx.GetText())
	}

	return fdecl(&FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		typeOf: "AnoymousFunctionDecl"}, ctx, v)
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

	firstChild VNode

	prev, next VNode
}

var _ VNode = (*FormalParameterList)(nil)

func (i *FormalParameterList) Type() string {
	return "FormalParameterList"
}
func (i *FormalParameterList) Code() string {
	return CodeDef(i)
}
func (i *FormalParameterList) Next() VNode {

	return i.next
}
func (i *FormalParameterList) SetNext(v VNode) {
	i.next = v
}
func (i *FormalParameterList) Prev() VNode {

	return i.prev
}
func (i *FormalParameterList) SetPrev(v VNode) {
	i.prev = v
}

func (i *FormalParameterList) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitFormalParameterList(ctx *base.FormalParameterListContext) interface{} {
	if v.Debug {
		log.Println("VisitFormalParameterList", ctx.GetText())
	}
	fp := &FormalParameterList{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fp.firstChild == nil {
			fp.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

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
	firstChild       VNode

	prev, next VNode
}

var _ VNode = (*FormalParameterArg)(nil)

func (i *FormalParameterArg) Type() string {
	return "FormalParameterArg"
}
func (i *FormalParameterArg) Code() string {
	return CodeDef(i)
}
func (i *FormalParameterArg) Next() VNode {

	return i.next
}
func (i *FormalParameterArg) SetNext(v VNode) {
	i.next = v
}
func (i *FormalParameterArg) Prev() VNode {

	return i.prev
}
func (i *FormalParameterArg) SetPrev(v VNode) {
	i.prev = v
}

func (i *FormalParameterArg) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitFormalParameterArg(ctx *base.FormalParameterArgContext) interface{} {
	if v.Debug {
		log.Println("VisitFormalParameterArg", ctx.GetText())
	}
	fa := &FormalParameterArg{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode

	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fa.firstChild == nil {
			fa.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

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
	firstChild       VNode

	prev, next VNode
}

var _ VNode = (*LastFormalParameterArg)(nil)

func (i *LastFormalParameterArg) Type() string {
	return "LastFormalParameterArg"
}
func (i *LastFormalParameterArg) Code() string {
	return CodeDef(i)
}
func (i *LastFormalParameterArg) Next() VNode {

	return i.next
}
func (i *LastFormalParameterArg) SetNext(v VNode) {
	i.next = v
}
func (i *LastFormalParameterArg) Prev() VNode {

	return i.prev
}
func (i *LastFormalParameterArg) SetPrev(v VNode) {
	i.prev = v
}

func (i *LastFormalParameterArg) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitLastFormalParameterArg(ctx *base.LastFormalParameterArgContext) interface{} {

	if v.Debug {
		log.Println("VisitFormalParameterArg", ctx.GetText())
	}

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
	firstChild     VNode
	Eos            VNode

	prev, next VNode
}

var _ VNode = (*FunctionBody)(nil)

func (i *FunctionBody) Type() string {
	return "FunctionBody"
}
func (i *FunctionBody) Code() string {
	return CodeDef(i)
}
func (i *FunctionBody) Next() VNode {

	return i.next
}
func (i *FunctionBody) SetNext(v VNode) {
	i.next = v
}
func (i *FunctionBody) Prev() VNode {

	return i.prev
}
func (i *FunctionBody) SetPrev(v VNode) {
	i.prev = v
}

func (i *FunctionBody) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitFunctionBody(ctx *base.FunctionBodyContext) interface{} {
	if v.Debug {
		log.Println("VisitArrowFunctionBody", ctx.GetText())
	}
	if v.Debug {
		log.Println("VisitFunctionBody", ctx.GetText(), ctx.GetChildCount())
	}
	fb := &FunctionBody{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fb.firstChild == nil {
			fb.firstChild = ch
		} else {
			prev.SetNext(ch)

		}

		ch.SetPrev(prev)

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

			fb.SourceElements = append(fb.SourceElements, ch)

		}

	}
	return fb
}

// Async? '*'? propertyName '(' formalParameterList?  ')'  functionBody
func (v *Visitor) VisitFunctionProperty(ctx *base.FunctionPropertyContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionProperty", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

// arguments
//     : '('(argument (',' argument)* ','?)?')'
//     ;
type Arguments struct {
	*SourceInfo
	OpenBrace  Token
	Args       *Argument
	CloseParen Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*Arguments)(nil)

func (i *Arguments) Type() string {
	return "Arguments"
}
func (i *Arguments) Code() string {
	return CodeDef(i)
}
func (i *Arguments) Next() VNode {
	return i.next
}
func (i *Arguments) SetNext(v VNode) {
	i.next = v
}
func (i *Arguments) Prev() VNode {
	return i.prev
}
func (i *Arguments) SetPrev(v VNode) {
	i.prev = v
}
func (i *Arguments) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitArguments(ctx *base.ArgumentsContext) interface{} {
	if v.Debug {
		log.Println("VisitArguments", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

// singleExpression '?'? '.' '#'? identifierName
func (v *Visitor) VisitMemberDotExpression(ctx *base.MemberDotExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitMemberDotExpression", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

// singleExpression
//     : anoymousFunction                                                      # FunctionExpression
func (v *Visitor) VisitFunctionExpression(ctx *base.FunctionExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionExpression", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}
