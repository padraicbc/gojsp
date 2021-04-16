package main

import (
	"strings"

	antlr "github.com/padraicbc/antlr4"
)

// just pulls original source code info
func getSourceInfo(ctx antlr.BaseParserRuleContext) *SourceInfo {
	start, end := ctx.GetStart().GetStart(), ctx.GetStop().GetStop()+1
	return &SourceInfo{Line: ctx.GetStart().GetLine(), Start: start, End: end,
		Column: ctx.GetStart().GetColumn(),
		Source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: start, Stop: end})}

}

// maybe better just having maybe a "GetNode" and return that, forgetting all other methods...
type VNode interface {
	Code() string
	GetInfo() *SourceInfo
	Type() string
	GetChildren() []VNode
}

// just to make life easier for implementing.
// will be the default return from VisitChildren
type BaseDefaultVNode struct{}

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
func (i *BaseDefaultVNode) GetChildren() []VNode {
	return nil
}

func CodeDef(t VNode) string {
	if t == nil {
		return ""
	}
	var c []string
	for _, n := range t.GetChildren() {

		c = append(c, n.Code())

	}

	return strings.Join(c, " ")

}

type SourceInfo struct {
	Start, End, Line, Column int
	Source                   string
}

func (s *SourceInfo) GetInfo() *SourceInfo {
	return s
}

type PTree struct {
	Root      *SourceElement
	LastChild *SourceElement
}

func (p *PTree) NextNodes() chan *SourceElement {

	if p == nil {
		panic("p cannot be nil")

	}
	nodes := make(chan *SourceElement)
	go func() {
		next := p.Root
		for next != nil {
			nodes <- next
			next = next.Next
		}
		close(nodes)
	}()
	return nodes
}

type SourceElement struct {
	*SourceInfo
	// VNodes have their own next/prev. Can visit all children from here
	Children   []VNode
	Prev, Next *SourceElement
	FirstChild VNode
}

// An identifier. Note that an identifier may be an expression or a destructuring pattern.

type LToken struct {
	value string
	*SourceInfo
	// From, StringLiteral...
	sn string
	// reservedWord...
	rn string

	// SymbolName string
}

var _ Token = (*LToken)(nil)

func (i *LToken) Value() string {
	return i.value
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

// keyword, reservedword, identifier
func (i *LToken) Type() string {
	return "Token"
}
func (i *LToken) GetInfo() *SourceInfo {
	return i.SourceInfo
}
func (i *LToken) GetChildren() []VNode {
	return nil
}
