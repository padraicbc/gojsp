package vast

import (
	"log"

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
		}
		prev = setSib(prev, ch)

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
		}
		prev = setSib(prev, ch)

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
		}
		prev = setSib(prev, ch)

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
	Async    Token
	Function Token
	Star     Token

	Identifier   Token
	OpenParen    Token
	PList        *FormalParameterList
	CloseParen   Token
	FunctionBody *FunctionBody

	firstChild VNode

	prev, next VNode
}

var _ VNode = (*FunctionDeclaration)(nil)

func (i *FunctionDeclaration) Type() string {
	return "FunctionDeclaration"
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

// functionDeclaration
// : Async? Function '*'? identifier '(' formalParameterList? ')' functionBody
// ;
func (v *Visitor) VisitFunctionDeclaration(ctx *base.FunctionDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionDeclaration", ctx.GetText(), ctx.GetChildCount())
	}
	fd := &FunctionDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fd.firstChild == nil {
			fd.firstChild = ch
		}
		prev = setSib(prev, ch)

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
			log.Printf("%+v %s\n", ch, ch.Type())

		}
	}
	return fd
}

func (v *Visitor) VisitFunctionDecl(ctx *base.FunctionDeclContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionDecl", ctx.GetText())
	}

	return v.VisitFunctionDeclaration(
		ctx.FunctionDeclaration().(*base.FunctionDeclarationContext)).(*FunctionDeclaration)
}

// same as any fun dec
// anoymousFunction
//     : functionDeclaration                                                       # FunctionDecl
//     | Async? Function '*'? '(' formalParameterList? ')' functionBody    # AnoymousFunctionDecl
//     | Async? arrowFunctionParameters '=>' arrowFunctionBody                     # ArrowFunction
//     ;
type AnoymousFunctionDecl struct {
	*SourceInfo
	Async    Token
	Function Token
	Star     Token

	OpenParen    Token
	PList        *FormalParameterList
	CloseParen   Token
	FunctionBody *FunctionBody

	firstChild VNode

	prev, next VNode
}

var _ VNode = (*AnoymousFunctionDecl)(nil)

func (i *AnoymousFunctionDecl) Type() string {
	return "AnoymousFunctionDecl"
}
func (i *AnoymousFunctionDecl) Code() string {
	return CodeDef(i)
}
func (i *AnoymousFunctionDecl) Next() VNode {

	return i.next
}
func (i *AnoymousFunctionDecl) SetNext(v VNode) {
	i.next = v
}
func (i *AnoymousFunctionDecl) Prev() VNode {

	return i.prev
}
func (i *AnoymousFunctionDecl) SetPrev(v VNode) {
	i.prev = v
}

func (i *AnoymousFunctionDecl) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitAnoymousFunctionDecl(ctx *base.AnoymousFunctionDeclContext) interface{} {
	if v.Debug {
		log.Println("VisitAnoymousFunctionDecl", ctx.GetText())
	}

	a := &AnoymousFunctionDecl{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if a.firstChild == nil {
			a.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {
		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {
			case "Multiply":
				a.Star = tk
			case "Function":
				a.Function = tk

			case "OpenParen":
				a.OpenParen = tk
			case "CloseParen":
				a.CloseParen = tk
			case "Async":
				a.Async = tk
			default:
				log.Panicf("%+v %s\n", ch, ch.Type())
			}

		case "FormalParameterList":
			a.PList = ch.(*FormalParameterList)
		case "FunctionBody":
			a.FunctionBody = ch.(*FunctionBody)

		default:
			log.Printf("%+v %s\n", ch, ch.Type())

		}
	}
	return a
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
		}
		prev = setSib(prev, ch)

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
		}
		prev = setSib(prev, ch)

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
	lf := &LastFormalParameterArg{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	lf.Ellipsis = ident(v, ctx.Ellipsis().GetSymbol())
	lf.firstChild = lf.Ellipsis
	lf.SingleExpression = v.Visit(ctx.SingleExpression()).(VNode)

	setAllSibs(lf.Ellipsis, lf.SingleExpression)
	return lf
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
		log.Println("VisitFunctionBody", ctx.GetText(), ctx.GetChildCount())
	}
	fb := &FunctionBody{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if fb.firstChild == nil {
			fb.firstChild = ch
		}
		prev = setSib(prev, ch)

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
			// todo: check a type?
			fb.SourceElements = append(fb.SourceElements, ch)

		}

	}
	return fb
}

// arguments
//     : '('(argument (',' argument)* ','?)?')'
//     ;
type Arguments struct {
	*SourceInfo
	OpenParen  Token
	Args       []*Argument
	Commas     []Token
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

	args := &Arguments{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if args.firstChild == nil {
			args.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {
		case "Argument":
			args.Args = append(args.Args, ch.(*Argument))

		case "LToken":

			switch tk := ch.(Token); tk.SymbolName() {
			case "Comma":
				args.Commas = append(args.Commas, tk)
			case "OpenParen":
				args.OpenParen = tk
			case "CloseParen":
				args.CloseParen = tk

			default:
				log.Panicf("%+v %s\n", ch, ch.Type())
			}
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}

	return args

}

// argument
//     : Ellipsis? (singleExpression | identifier)
//     ;
type Argument struct {
	*SourceInfo
	Ellipsis   Token
	SingleExp  VNode
	Identifier VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*Argument)(nil)

func (i *Argument) Type() string {
	return "Argument"
}
func (i *Argument) Code() string {
	return CodeDef(i)
}
func (i *Argument) Next() VNode {
	return i.next
}
func (i *Argument) SetNext(v VNode) {
	i.next = v
}
func (i *Argument) Prev() VNode {
	return i.prev
}
func (i *Argument) SetPrev(v VNode) {
	i.prev = v
}
func (i *Argument) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitArgument(ctx *base.ArgumentContext) interface{} {
	a := &Argument{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	if ctx.Ellipsis() != nil {
		a.Ellipsis = ident(v, ctx.Ellipsis().GetSymbol())
		a.firstChild = a.Ellipsis
	}
	if ctx.SingleExpression() != nil {
		a.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
		if a.firstChild == nil {
			a.firstChild = a.SingleExp
		}
	}
	if ctx.Identifier() != nil {
		a.Identifier = v.VisitIdentifier(ctx.Identifier().(*base.IdentifierContext)).(Token)
		if a.firstChild == nil {
			a.firstChild = a.Identifier
		}
	}
	setAllSibs(a.Ellipsis, a.SingleExp, a.Identifier)

	return a
}

// singleExpression '?'? '.' '#'? identifierName
type MemberDotExpression struct {
	*SourceInfo
	SingleExp      VNode
	QuestionMark   Token
	Dot            Token
	Hashtag        Token
	IdentifierName Token
	firstChild     VNode
	prev, next     VNode
}

var _ VNode = (*MemberDotExpression)(nil)

func (i *MemberDotExpression) Type() string {
	return "MemberDotExpression"
}
func (i *MemberDotExpression) Code() string {
	return CodeDef(i)
}
func (i *MemberDotExpression) Next() VNode {
	return i.next
}
func (i *MemberDotExpression) SetNext(v VNode) {
	i.next = v
}
func (i *MemberDotExpression) Prev() VNode {
	return i.prev
}
func (i *MemberDotExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *MemberDotExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitMemberDotExpression(ctx *base.MemberDotExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitMemberDotExpression", ctx.GetText())
	}
	md := &MemberDotExpression{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	md.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
	md.firstChild = md.SingleExp

	if ctx.QuestionMark() != nil {
		md.QuestionMark = ident(v, ctx.QuestionMark().GetSymbol())
	}
	if ctx.Dot() != nil {
		md.Dot = ident(v, ctx.Dot().GetSymbol())
	}
	if ctx.Hashtag() != nil {
		md.Hashtag = ident(v, ctx.Hashtag().GetSymbol())
	}
	md.IdentifierName = v.VisitIdentifierName(ctx.IdentifierName().(*base.IdentifierNameContext)).(Token)

	setAllSibs(md.SingleExp, md.QuestionMark, md.Dot, md.Hashtag, md.IdentifierName)

	return md

}

// singleExpression
//     : anoymousFunction                                                       # FunctionExpression
func (v *Visitor) VisitFunctionExpression(ctx *base.FunctionExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitFunctionExpression", ctx.GetText())
	}
	// todo: check what can be run
	return v.VisitChildren(ctx).([]VNode)[0]
}
