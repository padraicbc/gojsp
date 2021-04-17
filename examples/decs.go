package main

import "github.com/padraicbc/gojsp/parser"

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;
type Declaration struct {
	*SourceInfo
	Node VNode
}

var _ VNode = (*Declaration)(nil)

func (i *Declaration) Type() string {
	return "Declaration"
}
func (i *Declaration) Code() string {
	return CodeDef(i)
}

func (i *Declaration) GetChildren() []VNode {
	return []VNode{
		i.Node,
	}
}
func (v *visitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	return &Declaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
		Node: v.VisitChildren(ctx).([]VNode)[0]}

}

// variableDeclaration
//     : assignable ('=' singleExpression)? // ECMAScript 6: Array & Object Matching
//     ;
type VariableDeclaration struct {
	*SourceInfo
	Assignable VNode
	Equals     Token
	Expression VNode
}

var _ VNode = (*VariableDeclaration)(nil)

func (i *VariableDeclaration) Type() string {
	return "VariableDeclaration"
}
func (i *VariableDeclaration) Code() string {
	return CodeDef(i)
}

func (i *VariableDeclaration) GetChildren() []VNode {
	return []VNode{
		i.Assignable,
		i.Equals,
		i.Expression,
	}
}
func (v *visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	// log.Println(ctx.SingleExpression().GetText(), ctx.Assignable().GetText())

	d := &VariableDeclaration{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	for _, ch := range v.VisitChildren(ctx).([]VNode) {

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
}

var _ VNode = (*VariableDeclarationList)(nil)

func (i *VariableDeclarationList) Type() string {
	return "VariableDeclarationList"
}
func (i *VariableDeclarationList) Code() string {
	return CodeDef(i)
}

func (i *VariableDeclarationList) GetChildren() []VNode {
	return []VNode{
		i.VarModifier,
		// i.Commas,
	}
}

func (v *visitor) VisitVariableDeclarationList(ctx *parser.VariableDeclarationListContext) interface{} {
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
func (v *visitor) VisitVarModifier(ctx *parser.VarModifierContext) interface{} {

	return v.VisitChildren(ctx)
}

// assignable
//     : identifier
//     | arrayLiteral
//     | objectLiteral
//     ;
type Assignable struct {
	*SourceInfo
	Node VNode
}

var _ VNode = (*ArrayElement)(nil)

func (i *Assignable) Type() string {
	return "Assignable"
}
func (i *Assignable) Code() string {
	return CodeDef(i)
}

func (i *Assignable) GetChildren() []VNode {
	return []VNode{
		i.Node,
	}
}

// Maybe just retrun children as it is not a concrete type and let others check what it is...
func (v *visitor) VisitAssignable(ctx *parser.AssignableContext) interface{} {
	// log.Println("VisitAssignable", ctx.GetText())
	return v.VisitChildren(ctx)
}
