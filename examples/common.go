package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

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

func CodeDef(t VNode) string {

	if t == nil {
		return ""
	}

	var c []string
	for _, n := range t.Children() {
		// ugly but zero value is not nil, this will change or be remove completely
		if reflect.ValueOf(n).Kind() == reflect.Ptr && !reflect.ValueOf(n).IsNil() {
			c = append(c, n.Code())
		}

	}

	return strings.Join(c, "")

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
		log.Println(next.Type(), next.Code())
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
	RName() string
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
func (i *LToken) RName() string {
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
		if al.children == nil {
			al.children = ch
		} else {
			prev.Next(ch)
		}
		ch.Prev(prev)
		prev = ch
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
