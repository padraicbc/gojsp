package vast

import "github.com/padraicbc/gojsp/base"

// objectLiteral
//     : '{' (propertyAssignment (',' propertyAssignment)*)? ','? '}'
//     ;
type ObjectLiteral struct {
	*SourceInfo
	OpenBracket  Token
	ElementList  *ElementList
	CloseBracket Token
	children     VNode
	next, prev   VNode
}

var _ VNode = (*ObjectLiteral)(nil)

func (i *ObjectLiteral) Next(v VNode) VNode {

	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *ObjectLiteral) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *ObjectLiteral) Type() string {
	return "ObjectLiteral"
}
func (i *ObjectLiteral) Code() string {
	return CodeDef(i)
}

func (i *ObjectLiteral) Children() []VNode {
	// todo: flatten
	return children(i.children)
}

func (v *Visitor) VisitObjectLiteralExpression(ctx *base.ObjectLiteralExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitObjectLiteral(ctx *base.ObjectLiteralContext) interface{} {

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

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertySetter(ctx *base.PropertySetterContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyShorthand(ctx *base.PropertyShorthandContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPropertyName(ctx *base.PropertyNameContext) interface{} {

	return v.VisitChildren(ctx)
}
