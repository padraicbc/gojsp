package vast

import "github.com/padraicbc/gojsp/base"

// switchStatement
//     : Switch '(' expressionSequence ')' caseBlock
//     ;
type SwitchStatement struct {
	*SourceInfo
	Switch             Token
	OpenParen          Token
	ExpressionSequence *ExpressionSequence
	CloseParen         Token
	CaseBlock          *CaseBlock
	firstChild         VNode
	prev, next         VNode
}

var _ VNode = (*SwitchStatement)(nil)

func (i *SwitchStatement) Type() string {
	return "SwitchStatement"
}
func (i *SwitchStatement) Code() string {
	return CodeDef(i)
}
func (i *SwitchStatement) Next() VNode {
	return i.next
}
func (i *SwitchStatement) SetNext(v VNode) {
	i.next = v
}
func (i *SwitchStatement) Prev() VNode {
	return i.prev
}
func (i *SwitchStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *SwitchStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitSwitchStatement(ctx *base.SwitchStatementContext) interface{} {
	s := &SwitchStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	s.Switch = ident(v, ctx.Switch().GetSymbol())
	s.OpenParen = ident(v, ctx.OpenParen().GetSymbol())
	s.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	s.CloseParen = ident(v, ctx.CloseParen().GetSymbol())
	s.CaseBlock = v.VisitCaseBlock(ctx.CaseBlock().(*base.CaseBlockContext)).(*CaseBlock)
	setAllSibs(s.Switch, s.OpenParen, s.ExpressionSequence, s.CloseParen, s.CaseBlock)
	return s
}

// caseBlock
//     : '{' caseClauses? (defaultClause caseClauses?)? '}'
//     ;
type CaseBlock struct {
	*SourceInfo
	OpenBrace     Token
	CaseClauses   *CaseClauses
	DefaultClause *DefaultClause
	DCaseClauses  *CaseClauses
	CloseBrace    Token

	firstChild VNode
	prev, next VNode
}

var _ VNode = (*CaseBlock)(nil)

func (i *CaseBlock) Type() string {
	return "CaseBlock"
}
func (i *CaseBlock) Code() string {
	return CodeDef(i)
}
func (i *CaseBlock) Next() VNode {
	return i.next
}
func (i *CaseBlock) SetNext(v VNode) {
	i.next = v
}
func (i *CaseBlock) Prev() VNode {
	return i.prev
}
func (i *CaseBlock) SetPrev(v VNode) {
	i.prev = v
}
func (i *CaseBlock) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitCaseBlock(ctx *base.CaseBlockContext) interface{} {
	c := &CaseBlock{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	c.OpenBrace = ident(v, ctx.OpenBrace().GetSymbol())
	c.firstChild = c.OpenBrace
	if ctx.Main != nil {
		c.CaseClauses = v.VisitCaseClauses(ctx.Main.(*base.CaseClausesContext)).(*CaseClauses)

	}
	if ctx.DefaultClause() != nil {
		c.DefaultClause = v.VisitDefaultClause(
			ctx.DefaultClause().(*base.DefaultClauseContext)).(*DefaultClause)

		if ctx.Def != nil {
			c.DCaseClauses = v.VisitCaseClauses(ctx.Def.(*base.CaseClausesContext)).(*CaseClauses)

		}
	}

	c.CloseBrace = ident(v, ctx.CloseBrace().GetSymbol())
	setAllSibs(c.OpenBrace, c.CaseClauses, c.DefaultClause, c.DCaseClauses, c.CloseBrace)

	return c
}

// caseClauses
//     : caseClause+
//     ;
type CaseClauses struct {
	*SourceInfo
	// should probably just use firtsChild when there is a list of one type as only children
	Clauses    []*CaseClause
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*CaseClauses)(nil)

func (i *CaseClauses) Type() string {
	return "CaseClauses"
}
func (i *CaseClauses) Code() string {
	return CodeDef(i)
}
func (i *CaseClauses) Next() VNode {
	return i.next
}
func (i *CaseClauses) SetNext(v VNode) {
	i.next = v
}
func (i *CaseClauses) Prev() VNode {
	return i.prev
}
func (i *CaseClauses) SetPrev(v VNode) {
	i.prev = v
}
func (i *CaseClauses) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitCaseClauses(ctx *base.CaseClausesContext) interface{} {
	c := &CaseClauses{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	var prev VNode
	for _, cl := range ctx.AllCaseClause() {
		cc := v.VisitCaseClause(cl.(*base.CaseClauseContext)).(*CaseClause)
		if c.firstChild == nil {
			c.firstChild = cc
		}
		prev = setSib(prev, cc)
		c.Clauses = append(c.Clauses, cc)
	}
	return c
}

// caseClause
//     : Case expressionSequence ':' statementList?
//     ;
type CaseClause struct {
	*SourceInfo

	Case               *LToken
	ExpressionSequence *ExpressionSequence
	Colon              Token
	StatementList      *StatementList
	firstChild         VNode
	prev, next         VNode
}

var _ VNode = (*CaseClause)(nil)

func (i *CaseClause) Type() string {
	return "CaseClause"
}
func (i *CaseClause) Code() string {
	return CodeDef(i)
}
func (i *CaseClause) Next() VNode {
	return i.next
}
func (i *CaseClause) SetNext(v VNode) {
	i.next = v
}
func (i *CaseClause) Prev() VNode {
	return i.prev
}
func (i *CaseClause) SetPrev(v VNode) {
	i.prev = v
}
func (i *CaseClause) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitCaseClause(ctx *base.CaseClauseContext) interface{} {
	c := &CaseClause{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	c.Case = ident(v, ctx.Case().GetSymbol())
	c.firstChild = c.Case
	c.ExpressionSequence = v.VisitExpressionSequence(
		ctx.ExpressionSequence().(*base.ExpressionSequenceContext)).(*ExpressionSequence)
	c.Colon = ident(v, ctx.Colon().GetSymbol())
	if ctx.StatementList() != nil {
		c.StatementList = v.VisitStatementList(
			ctx.StatementList().(*base.StatementListContext)).(*StatementList)
	}
	setAllSibs(c.Case, c.ExpressionSequence, c.Colon, c.StatementList)

	return c

}

// defaultClause
//     : Default ':' statementList?
type DefaultClause struct {
	*SourceInfo
	Default, Colon Token

	StatementList *StatementList
	firstChild    VNode
	prev, next    VNode
}

var _ VNode = (*DefaultClause)(nil)

func (i *DefaultClause) Type() string {
	return "DefaultClause"
}
func (i *DefaultClause) Code() string {
	return CodeDef(i)
}
func (i *DefaultClause) Next() VNode {
	return i.next
}
func (i *DefaultClause) SetNext(v VNode) {
	i.next = v
}
func (i *DefaultClause) Prev() VNode {
	return i.prev
}
func (i *DefaultClause) SetPrev(v VNode) {
	i.prev = v
}
func (i *DefaultClause) FirstChild() VNode {
	return i.firstChild
}

func (v *Visitor) VisitDefaultClause(ctx *base.DefaultClauseContext) interface{} {
	d := &DefaultClause{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	d.Default = ident(v, ctx.Default().GetSymbol())
	d.firstChild = d.Default
	d.Colon = ident(v, ctx.Colon().GetSymbol())

	if ctx.StatementList() != nil {
		d.StatementList = v.VisitStatementList(
			ctx.StatementList().(*base.StatementListContext)).(*StatementList)
	}
	setAllSibs(d.Default, d.Colon, d.StatementList)
	return d
}
