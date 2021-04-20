package main

import (
	"fmt"
	"runtime"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/parser"
)

// just pulls original source code info
func getSourceInfo(ctx antlr.BaseParserRuleContext) *SourceInfo {
	start, end := ctx.GetStart().GetStart(), ctx.GetStop().GetStop()+1
	return &SourceInfo{Line: ctx.GetStart().GetLine(), Start: start, End: end,
		Column: ctx.GetStart().GetColumn(),
		Source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: start, Stop: end})}

}

// maybe SourceInfo should just be on "Tokens" as each token has it's own positioning?
type VNode interface {
	Code() string
	GetInfo() *SourceInfo
	Type() string
	Children() []VNode
	Prev(VNode) VNode
	Next(VNode) VNode
}

func debug(v VNode) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s:%d %s %+v\n", file, line, v.Type(), v)
	}
}

// just to make life easier for implementing.
// will be the default return from VisitChildren
type BaseDefaultVNode struct {
	prev, next VNode
}

var _ VNode = &BaseDefaultVNode{}

func (i *BaseDefaultVNode) Type() string {
	return "BaseDefaultVNode"
}
func (i *BaseDefaultVNode) Code() string {
	return CodeDef(i)
}
func (i *BaseDefaultVNode) GetInfo() *SourceInfo {
	return nil
}
func (i *BaseDefaultVNode) Children() []VNode {
	return nil
}

func (i *BaseDefaultVNode) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}
func (i *BaseDefaultVNode) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}

func _fill(v VNode, cc chan VNode) {
	for _, c := range v.Children() {
		cc <- c
		_fill(c, cc)

	}

}

// get a flattened list of tokens
func fill(v VNode) chan VNode {
	chi := v.Children()

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
	// keep track of what we have seen so we don't concta twice
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

type PTree struct {
	Root      VNode
	LastChild VNode
}

func (p *PTree) NextNodes() chan VNode {

	if p == nil {
		panic("p cannot be nil")

	}
	nodes := make(chan VNode)
	go func() {
		next := p.Root
		for next != nil {
			nodes <- next
			next = next.Next(nil)
		}
		close(nodes)
	}()
	return nodes
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
func (i *LToken) Children() []VNode {
	return nil
}
func (i *LToken) Next(v VNode) VNode {
	if v != nil {
		i.next = v
		return nil
	}

	return i.next
}
func (i *LToken) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}

	return i.prev
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
	children       VNode
	prev, next     VNode
}

var _ VNode = (*AliasName)(nil)

func (i *AliasName) Next(v VNode) VNode {

	if v != nil {
		i.next = v
		return nil
	}
	return i.next
}

func (i *AliasName) Prev(v VNode) VNode {
	if v != nil {
		i.prev = v
		return nil
	}
	return i.prev
}
func (a *AliasName) Code() string {
	return CodeDef(a)
}
func (i *AliasName) Type() string {
	return "AliasName"
}

func (i *AliasName) Children() []VNode {

	return children(i.children)
}

func (v *Visitor) VisitAliasName(ctx *parser.AliasNameContext) interface{} {
	al := &AliasName{

		SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for i, ch := range v.VisitChildren(ctx).([]VNode) {
		if al.children == nil {
			al.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
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

func children(start VNode) []VNode {
	out := []VNode{}
	n := start

	for n != nil {
		out = append(out, n)
		n = n.Next(nil)

	}
	return out
}
