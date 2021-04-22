package vast

import "github.com/padraicbc/gojsp/base"

// declaration
//     : variableStatement
//     | classDeclaration
//     | functionDeclaration
//     ;
type Declaration struct {
	*SourceInfo
	Dec        VNode
	children   VNode
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
	return children(i.children)
}
func (v *Visitor) VisitDeclaration(ctx *base.DeclarationContext) interface{} {
	d := &Declaration{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if d.children == nil {
			d.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch

	}
	d.Dec = d.children
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
	children   VNode
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
	return children(i.children)
}
func (v *Visitor) VisitVariableDeclaration(ctx *base.VariableDeclarationContext) interface{} {
	// log.Println(ctx.SingleExpression().GetText(), ctx.Assignable().GetText())

	d := &VariableDeclaration{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if d.children == nil {
			d.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
		switch ch.(type) {
		case *ArrayLiteral: // todo ObjectLiteral":
			d.Assignable = ch
		case *ObjectLiteral: // todo ObjectLiteral":
			d.Assignable = ch
		case *LToken:
			t := ch.(Token)
			if t.RName("") == "identifier" {
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
	children             VNode
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
	return children(i.children)
}

func (v *Visitor) VisitVariableDeclarationList(ctx *base.VariableDeclarationListContext) interface{} {
	vdl := &VariableDeclarationList{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if vdl.children == nil {
			vdl.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
		switch ch.Type() {
		case "VariableDeclaration":
			vdl.VariableDeclarations = append(vdl.VariableDeclarations, ch.(*VariableDeclaration))
		case "LToken":
			t := ch.(Token)
			if t.SymbolName() == "Comma" {
				vdl.Commas = append(vdl.Commas, t)
				continue
			}
			if t.RName("") == "varModifier" {
				vdl.VarModifier = t
				continue
			}
			panic(ch.Type())

		default:
			panic(ch.Type())

		}

	}

	return vdl
}
func (v *Visitor) VisitVarModifier(ctx *base.VarModifierContext) interface{} {

	return v.VisitChildren(ctx)
}

// assignable
//     : identifier
//     | arrayLiteral
//     | objectLiteral
//     ;
// Just return children as it is one of 3 types so let others check what it is...
func (v *Visitor) VisitAssignable(ctx *base.AssignableContext) interface{} {
	// log.Println("VisitAssignable", ctx.GetText())
	return v.VisitChildren(ctx)
}
