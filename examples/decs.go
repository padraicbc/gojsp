package main

import "github.com/padraicbc/gojsp/parser"

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;
type Declaration struct {
	*SourceInfo
	Node       VNode
	children   []VNode
	prev, next VNode
}

var _ VNode = (*Declaration)(nil)

func (i *Declaration) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *Declaration) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *Declaration) Type() string {
	return "Declaration"
}
func (i *Declaration) Code() string {
	return CodeDef(i)
}

func (i *Declaration) Children() []VNode {
	return i.children
}
func (v *Visitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	d := &Declaration{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	d.Node = d.children[0]
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
	children   []VNode
	next, prev VNode
}

var _ VNode = (*VariableDeclaration)(nil)

func (i *VariableDeclaration) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *VariableDeclaration) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *VariableDeclaration) Type() string {
	return "VariableDeclaration"
}
func (i *VariableDeclaration) Code() string {
	return CodeDef(i)
}

func (i *VariableDeclaration) Children() []VNode {
	return i.children
}
func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	// log.Println(ctx.SingleExpression().GetText(), ctx.Assignable().GetText())

	d := &VariableDeclaration{
		children:   v.VisitChildren(ctx).([]VNode),
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	for _, ch := range d.children {

		switch ch.(type) {
		case *ArrayLiteral: // todo ObjectLiteral":
			d.Assignable = ch
		case *LToken:
			t := ch.(Token)
			if t.RName() == "identifier" {
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
	children             []VNode
	prev, next           VNode
}

var _ VNode = (*VariableDeclarationList)(nil)

func (i *VariableDeclarationList) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *VariableDeclarationList) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *VariableDeclarationList) Type() string {
	return "VariableDeclarationList"
}
func (i *VariableDeclarationList) Code() string {
	return CodeDef(i)
}

func (i *VariableDeclarationList) Children() []VNode {
	return i.children
}

func (v *Visitor) VisitVariableDeclarationList(ctx *parser.VariableDeclarationListContext) interface{} {
	vdl := &VariableDeclarationList{}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

		switch ch.Type() {
		case "VariableDeclaration":
			vdl.VariableDeclarations = append(vdl.VariableDeclarations, ch.(*VariableDeclaration))
		case "LToken":
			t := ch.(Token)
			if t.SymbolName() == "Comma" {
				vdl.Commas = append(vdl.Commas, t)
				continue
			}
			vdl.VarModifier = t

		default:
			panic(ch.Type())

		}
	}

	return v.VisitChildren(ctx)
}
func (v *Visitor) VisitVarModifier(ctx *parser.VarModifierContext) interface{} {

	return v.VisitChildren(ctx)
}

// assignable
//     : identifier
//     | arrayLiteral
//     | objectLiteral
//     ;
type Assignable struct {
	*SourceInfo
	Node       VNode
	children   []VNode
	prev, next VNode
}

var _ VNode = (*Assignable)(nil)

func (i *Assignable) Type() string {
	return "Assignable"
}
func (i *Assignable) Code() string {
	return CodeDef(i)
}
func (i *Assignable) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *Assignable) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *Assignable) Children() []VNode {
	return i.children
}

// Maybe just retrun children as it is not a concrete type and let others check what it is...
func (v *Visitor) VisitAssignable(ctx *parser.AssignableContext) interface{} {
	// log.Println("VisitAssignable", ctx.GetText())
	return v.VisitChildren(ctx)
}
