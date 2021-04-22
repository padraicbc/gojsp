package vast

import "github.com/padraicbc/gojsp/base"

// labelledStatement
//     : identifier ':' statement
//     ;
type LabeledStatement struct {
	*SourceInfo
	Statement  VNode
	Label      Token
	Colon      Token
	children   VNode
	prev, next VNode
}

var _ VNode = (*LabeledStatement)(nil)

func (i *LabeledStatement) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *LabeledStatement) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (i *LabeledStatement) Type() string {
	return "LabeledStatement"
}
func (i *LabeledStatement) Code() string {
	return CodeDef(i)
}
func (i *LabeledStatement) Children() []VNode {

	return children(i.children)
}

// special case for $: ... todo: a type
func (v *Visitor) VisitLabelledStatement(ctx *base.LabelledStatementContext) interface{} {
	// log.Println("VisitLabelledStatement", ctx.GetText())

	if ctx.Identifier().GetText() == "$" {
		// log.Println("Reactive?")
	}
	lst := &LabeledStatement{
		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext),
	}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if lst.children == nil {
			lst.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
		switch ch.Type() {
		case "LToken":
			t := ch.(Token)
			if t.SymbolName() == "Identifier" {
				lst.Label = t
				continue
			}
			if t.SymbolName() == "Colon" {
				lst.Colon = t
				continue
			}
			//  Statement can be idenifier also?
			panic(ch)

		case "Block":
			lst.Statement = ch
		default:
			panic(ch.Type())

		}

	}
	return lst
}

// block
//     : '{' statementList? '}'
//     ;
type Block struct {
	*SourceInfo
	// *StatementList
	children   VNode
	prev, next VNode
}

var _ VNode = (*Block)(nil)

func (i *Block) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *Block) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func (i *Block) Type() string {
	return "Block"
}
func (i *Block) Code() string {
	return CodeDef(i)
}
func (i *Block) Children() []VNode {
	return children(i.children)
}
func (v *Visitor) VisitBlock(ctx *base.BlockContext) interface{} {
	// log.Println("VisitBlock", ctx.GetText())
	b := &Block{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if b.children == nil {
			b.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
	}
	return b
}
