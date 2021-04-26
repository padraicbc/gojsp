package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// objectLiteral
//     : '{' (propertyAssignment (',' propertyAssignment)*)? ','? '}'
//     ;
type ObjectLiteral struct {
	*SourceInfo
	OpenBracket  Token
	ElementList  *ElementList
	CloseBracket Token
	firstChild   VNode

	next, prev VNode
}

var _ VNode = (*ObjectLiteral)(nil)

func (i *ObjectLiteral) Next() VNode {

	return i.next
}
func (i *ObjectLiteral) SetNext(v VNode) {
	i.next = v
}
func (i *ObjectLiteral) Prev() VNode {

	return i.prev
}
func (i *ObjectLiteral) SetPrev(v VNode) {
	i.prev = v
}
func (i *ObjectLiteral) Type() string {
	return "ObjectLiteral"
}
func (i *ObjectLiteral) Code() string {
	return CodeDef(i)
}

func (i *ObjectLiteral) FirstChild() VNode {

	// todo: flatten
	return i.firstChild

}

func (v *Visitor) VisitObjectLiteralExpression(ctx *base.ObjectLiteralExpressionContext) interface{} {
	if v.Debug {
		log.Println("VisitObjectLiteralExpression", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitObjectLiteral(ctx *base.ObjectLiteralContext) interface{} {
	if v.Debug {
		log.Println("VisitObjectLiteral", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

// propertyAssignment
//     : propertyName ':' singleExpression                                             # PropertyExpressionAssignment
//     | '[' singleExpression ']' ':' singleExpression                                 # ComputedPropertyExpressionAssignment
//     | Async? '*'? propertyName '(' formalParameterList?  ')'  functionBody  # FunctionProperty
//     | getter '(' ')' functionBody                                           # PropertyGetter
//     | setter '(' formalParameterArg ')' functionBody                        # PropertySetter
//     | Ellipsis? singleExpression                                                    # PropertyShorthand
//     ;

func (v *Visitor) VisitPropertyGetter(ctx *base.PropertyGetterContext) interface{} {

	if v.Debug {
		log.Println("VisitPropertyGetter", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertySetter(ctx *base.PropertySetterContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertySetter", ctx.GetText())
	}
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyShorthand(ctx *base.PropertyShorthandContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertyShorthand", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyName(ctx *base.PropertyNameContext) interface{} {
	if v.Debug {
		log.Println("VisitPropertyName", ctx.GetText())
	}

	return v.VisitChildren(ctx)
}
