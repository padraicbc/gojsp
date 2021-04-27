package vast

import (
	"fmt"
	"reflect"
	"runtime"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

// just pulls original source code info
func getSourceInfo(ctx antlr.BaseParserRuleContext) *SourceInfo {
	start, end := ctx.GetStart().GetStart(), ctx.GetStop().GetStop()
	return &SourceInfo{Line: ctx.GetStart().GetLine(), Start: start, End: end,
		Column: ctx.GetStart().GetColumn(),
		Source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: start, Stop: end})}

}

type Node struct {
	firstChild VNode
	VNode
}

// maybe SourceInfo should just be on "Tokens" as each token has it's own positioning?
type VNode interface {
	Code() string
	GetInfo() *SourceInfo
	Type() string
	FirstChild() VNode
	SetPrev(VNode)
	Prev() VNode
	Next() VNode
	SetNext(VNode)
}

func debug(v VNode) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s:%d %s %+v\n", file, line, v.Type(), v)
	}
}

func _fill(v VNode, cc chan VNode) {
	for _, c := range Children(v.FirstChild()) {
		cc <- c
		_fill(c, cc)

	}

}

// get a flattened list of tokens
func fill(v VNode) chan VNode {
	chi := Children(v.FirstChild())

	cc := make(chan VNode, 1)
	go func() {
		for _, c := range chi {
			cc <- c
			if c.Type() != "LToken" {
				_fill(c, cc)
			}

		}
		close(cc)
	}()

	return cc

}

// prints out the source respecting any changes made to Tokens and original spacing etc..
func CodeDef(t VNode) string {

	if t == nil {
		return ""
	}

	orig := t.GetInfo().Source
	start := t.GetInfo().Start
	// keep track of what we have seen so we don't concat twice
	offset := 0
	var source string
	for n := range fill(t) {

		if tk, ok := n.(Token); ok {
			// need to subtract as these offsets are based on original source
			tkstart, tkend := tk.GetInfo().Start-start, tk.GetInfo().End-start
			fh := orig[offset:tkstart]

			source += fmt.Sprintf("%s%s", fh, tk.Value())
			// use tkend - tkstar as that is original token offsets
			// if we have changed it may have srhunk/grown...
			offset += len(fh) + tkend - tkstart

		}

	}

	return source

}

type SourceInfo struct {
	Start, End, Line, Column int
	Source                   string
}

func (s *SourceInfo) GetInfo() *SourceInfo {
	return s
}

// Parent will have type VNode
type Token interface {
	VNode
	SetValue(string)
	Value() string
	RName(string) string
	SymbolName() string
}

type LToken struct {
	value string
	*SourceInfo
	// From .. StringLiteral...
	sn string
	// rulename .. reservedWord...
	rn         string
	prev, next VNode
}

var _ Token = (*LToken)(nil)
var _ VNode = (*LToken)(nil)

func (i *LToken) Value() string {
	return i.value
}
func (i *LToken) SetValue(s string) {
	i.value = s
}
func (i *LToken) SymbolName() string {
	return i.sn
}
func (i *LToken) Code() string {
	return i.value
}
func (i *LToken) RName(s string) string {
	if s != "" {
		i.rn = s
		return ""
	}
	return i.rn
}
func (i *LToken) FirstChild() VNode {
	return nil
}
func (i *LToken) SetChild(ch, prev VNode) {
	return
}
func (i *LToken) Next() VNode {

	return i.next
}
func (i *LToken) SetNext(v VNode) {
	i.next = v
}

func (i *LToken) Prev() VNode {

	return i.prev
}
func (i *LToken) SetPrev(v VNode) {
	i.prev = v
}

// keyword, reservedword, identifier
func (i *LToken) Type() string {
	return "LToken"
}
func (i *LToken) GetInfo() *SourceInfo {
	return i.SourceInfo
}

// aliasName
//     : identifierName (As identifierName)?
//     ;
type AliasName struct {
	*SourceInfo
	IdentifierName Token
	Alias          Token
	As             Token
	Comma          Token
	firstChild     VNode

	prev, next VNode
}

var _ VNode = (*AliasName)(nil)

func (i *AliasName) Next() VNode {

	return i.next
}
func (i *AliasName) SetNext(v VNode) {
	i.next = v
}

func (i *AliasName) Prev() VNode {

	return i.prev
}
func (i *AliasName) SetPrev(v VNode) {
	i.prev = v
}
func (a *AliasName) Code() string {
	return CodeDef(a)
}
func (i *AliasName) Type() string {
	return "AliasName"
}

func (i *AliasName) FirstChild() VNode {

	return i.firstChild

}

func (v *Visitor) VisitAliasName(ctx *base.AliasNameContext) interface{} {
	al := &AliasName{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for i, ch := range v.VisitChildren(ctx).([]VNode) {
		if al.firstChild == nil {
			al.firstChild = ch
		} else {
			prev.SetNext(ch)

		}
		ch.SetPrev(prev)

		prev = ch
		t := ch.(Token)

		switch t.SymbolName() {

		case "Identifier":
			// always there
			if i == 0 {
				al.IdentifierName = t
				// > 0 means alias
			} else {
				al.Alias = t
			}
		case "As":
			al.As = t
		default:
			panic(t.SymbolName())

		}

	}
	return al
}

func Children(start VNode) []VNode {
	out := []VNode{}
	n := start

	for n != nil {
		out = append(out, n)
		n = n.Next()

	}
	return out
}

// todo: build tree from result nodes not each iteration and remove dupes
func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {

	var result []VNode
	for _, ch := range node.GetChildren() {

		res := ch.(antlr.ParseTree).Accept(v)

		switch rr := res.(type) {

		// satisifes both interfaces Token and VNode so check Token first as we set extra info specific to Token.
		case Token:
			rr.RName(v.ruleNames[node.GetRuleContext().GetRuleIndex()])
			result = append(result, rr)
		case VNode:
			result = append(result, rr)
		case []VNode:
			result = append(result, rr...)
		case error:
			// eof
			// log.Println(rr)
		case antlr.ErrorNode:
			// log.Fatal(rr)

		default:

			panic(reflect.TypeOf(rr))

		}

	}

	return result

}

// getter
//     : {p.n("get")}? identifier propertyName
//     ;
func (v *Visitor) VisitGetter(ctx *base.GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

// setter
//     : {p.n("set")}? identifier propertyName
//     ;
func (v *Visitor) VisitSetter(ctx *base.SetterContext) interface{} {
	return v.VisitChildren(ctx)
}
