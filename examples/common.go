package main

import (
	"strings"

	antlr "github.com/padraicbc/antlr4"
)

// just pulls original source code info
func getSourceInfo(ctx antlr.BaseParserRuleContext) *SourceInfo {

	return &SourceInfo{Line: ctx.GetStart().GetLine(), Start: ctx.GetStart().GetStart(), End: ctx.GetStop().GetStart(),
		Source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: ctx.GetStart().GetStart(), Stop: ctx.GetStop().GetStop() + 1})}

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
	return ""
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
	Start, End, Line int
	Source           string
}

func (s *SourceInfo) GetInfo() *SourceInfo {
	return s
}

type PTree struct {
	Root      *SourceElement
	LastChild *SourceElement
}

func (p *PTree) NextNodes() chan *SourceElement {

	nodes := make(chan *SourceElement)
	next := p.Root
	go func() {
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
}
