package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;
type Declaration struct {
	*SourceInfo
	Dec        VNode
	firstChild VNode

	prev, next VNode
}

var _ VNode = (*Declaration)(nil)

func (i *Declaration) Next() VNode {

	return i.next
}
func (i *Declaration) SetNext(v VNode) {
	i.next = v
}
func (i *Declaration) Prev() VNode {

	return i.prev
}
func (i *Declaration) SetPrev(v VNode) {
	i.prev = v
}
func (i *Declaration) Type() string {
	return "Declaration"
}
func (i *Declaration) Code() string {
	return CodeDef(i)
}

func (i *Declaration) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitDeclaration(ctx *base.DeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitDeclaration", ctx.GetText())
	}
	d := &Declaration{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if d.firstChild == nil {
			d.firstChild = ch
		}
		prev = setSib(prev, ch)

	}
	d.Dec = d.firstChild

	return d

}

// variableDeclaration
//     : assignable ('=' singleExpression)? // ECMAScript 6: Array & Object Matching
//     ;
type VariableDeclaration struct {
	*SourceInfo
	Assignable VNode
	Equals     Token
	Expression VNode
	firstChild VNode

	next, prev VNode
}

var _ VNode = (*VariableDeclaration)(nil)

func (i *VariableDeclaration) Next() VNode {

	return i.next
}
func (i *VariableDeclaration) SetNext(v VNode) {
	i.next = v
}
func (i *VariableDeclaration) Prev() VNode {

	return i.prev
}
func (i *VariableDeclaration) SetPrev(v VNode) {
	i.prev = v
}
func (i *VariableDeclaration) Type() string {
	return "VariableDeclaration"
}
func (i *VariableDeclaration) Code() string {
	return CodeDef(i)
}

func (i *VariableDeclaration) FirstChild() VNode {

	return i.firstChild

}
func (v *Visitor) VisitVariableDeclaration(ctx *base.VariableDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitVariableDeclaration", ctx.GetText(), ctx.GetChildCount())
	}

	d := &VariableDeclaration{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if d.firstChild == nil {
			d.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.(type) {
		case *ArrayLiteral: // todo ArrayLiteral":
			d.Assignable = ch
		case *ObjectLiteral: // todo ObjectLiteral":
			d.Assignable = ch
		case *LToken:
			t := ch.(Token)
			if t.rname("") == "identifier" {
				d.Assignable = t
				continue
			}
			d.Equals = ch.(Token)
		case VNode:
			d.Expression = ch
		default:
			panic(ch.Type())

		}

	}
	return d
}

// variableDeclarationList
//     : varModifier variableDeclaration (',' variableDeclaration)*
//     ;
type VariableDeclarationList struct {
	*SourceInfo
	VarModifier          Token // var, let, const
	VariableDeclarations []*VariableDeclaration
	Commas               []Token
	firstChild           VNode

	prev, next VNode
}

var _ VNode = (*VariableDeclarationList)(nil)

func (i *VariableDeclarationList) Next() VNode {

	return i.next
}
func (i *VariableDeclarationList) SetNext(v VNode) {
	i.next = v
}
func (i *VariableDeclarationList) Prev() VNode {

	return i.prev
}
func (i *VariableDeclarationList) SetPrev(v VNode) {
	i.prev = v
}
func (i *VariableDeclarationList) Type() string {
	return "VariableDeclarationList"
}
func (i *VariableDeclarationList) Code() string {
	return CodeDef(i)
}

func (i *VariableDeclarationList) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitVariableDeclarationList(ctx *base.VariableDeclarationListContext) interface{} {
	if v.Debug {
		log.Println("VisitVariableDeclarationList", ctx.GetText())
	}

	vdl := &VariableDeclarationList{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if vdl.firstChild == nil {
			vdl.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {
		case "VariableDeclaration":
			vdl.VariableDeclarations = append(vdl.VariableDeclarations, ch.(*VariableDeclaration))
		case "LToken":
			t := ch.(Token)
			switch t.SymbolName() {

			case "Comma":
				vdl.Commas = append(vdl.Commas, t)

			case "Var", "Const", "Let":
				vdl.VarModifier = t
			default:
				panic(ch.Type() + t.rname("") + t.SymbolName())
			}

		default:
			panic(ch.Type())

		}

	}

	return vdl
}

// varModifier  // let, const - ECMAScript 6
//     : Var
//     | Let
//     | Const
//     ;
func (v *Visitor) VisitVarModifier(ctx *base.VarModifierContext) interface{} {
	if v.Debug {
		log.Println("VisitVarModifier", ctx.GetText())
	}
	if ctx.Let() != nil {
		return ident(v, ctx.Let().GetSymbol())
	}
	if ctx.Var() != nil {
		return ident(v, ctx.Var().GetSymbol())
	}
	if ctx.Const() != nil {
		return ident(v, ctx.Const().GetSymbol())
	}
	panic("VisitVarModifier")

}

// assignable
//     : identifier
//     | arrayLiteral
//     | objectLiteral
//     ;
// Just return children as it is one of 3 types so let others check what it is...
func (v *Visitor) VisitAssignable(ctx *base.AssignableContext) interface{} {
	if v.Debug {
		log.Println("VisitAssignable", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}
