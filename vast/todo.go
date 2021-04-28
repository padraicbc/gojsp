package vast

import (
	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

// todo: implement
// caseBlock
//     : '{' caseClauses? (defaultClause caseClauses?)? '}'
//     ;
type CaseBlock struct {
	*SourceInfo
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
	return v.VisitChildren(ctx)
}

// caseClauses
//     : caseClause+
//     ;
type CaseClauses struct {
	*SourceInfo
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
	return v.VisitChildren(ctx)
}

// caseClause
//     : Case expressionSequence ':' statementList?
//     ;
type CaseClause struct {
	*SourceInfo

	Case        LToken
	ExpSequence *ExpressionSequence
	Colon       Token
	// todo: create statement type?
	StatementList *[]VNode
	firstChild    VNode
	prev, next    VNode
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
	return v.VisitChildren(ctx)
}

// defaultClause
//     : Default ':' statementList?
type DefaultClause struct {
	*SourceInfo
	Default, Colon Token
	// todo: create statement type?
	StatementList *[]VNode
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
	return v.VisitChildren(ctx)
}

// throwStatement
//     : Throw {p.notLineTerminator()}? expressionSequence eos
//     ;
type ThrowStatement struct {
	*SourceInfo
	Throw        Token
	ExprSequence *ExpressionSequence
	Eos          Token
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*ThrowStatement)(nil)

func (i *ThrowStatement) Type() string {
	return "ThrowStatement"
}
func (i *ThrowStatement) Code() string {
	return CodeDef(i)
}
func (i *ThrowStatement) Next() VNode {
	return i.next
}
func (i *ThrowStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ThrowStatement) Prev() VNode {
	return i.prev
}
func (i *ThrowStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ThrowStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitThrowStatement(ctx *base.ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

// tryStatement
//     : Try block (catchProduction finallyProduction? | finallyProduction)
//     ;
type TryStatement struct {
	*SourceInfo
	Try        Token
	Block      *Block
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*TryStatement)(nil)

func (i *TryStatement) Type() string {
	return "TryStatement"
}
func (i *TryStatement) Code() string {
	return CodeDef(i)
}
func (i *TryStatement) Next() VNode {
	return i.next
}
func (i *TryStatement) SetNext(v VNode) {
	i.next = v
}
func (i *TryStatement) Prev() VNode {
	return i.prev
}
func (i *TryStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *TryStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitTryStatement(ctx *base.TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

// catchProduction
//     : Catch ('(' assignable? ')')? block
//     ;
type CatchProduction struct {
	*SourceInfo
	Catch      Token
	OpenParen  Token
	Assignable *VNode
	CloseParen Token
	Block      *Block
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*CatchProduction)(nil)

func (i *CatchProduction) Type() string {
	return "CatchProduction"
}
func (i *CatchProduction) Code() string {
	return CodeDef(i)
}
func (i *CatchProduction) Next() VNode {
	return i.next
}
func (i *CatchProduction) SetNext(v VNode) {
	i.next = v
}
func (i *CatchProduction) Prev() VNode {
	return i.prev
}
func (i *CatchProduction) SetPrev(v VNode) {
	i.prev = v
}
func (i *CatchProduction) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitCatchProduction(ctx *base.CatchProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

// finallyProduction
//     : Finally block
//     ;
type FinallyProduction struct {
	*SourceInfo
	Finally    Token
	Block      *Block
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*FinallyProduction)(nil)

func (i *FinallyProduction) Type() string {
	return "FinallyProduction"
}
func (i *FinallyProduction) Code() string {
	return CodeDef(i)
}
func (i *FinallyProduction) Next() VNode {
	return i.next
}
func (i *FinallyProduction) SetNext(v VNode) {
	i.next = v
}
func (i *FinallyProduction) Prev() VNode {
	return i.prev
}
func (i *FinallyProduction) SetPrev(v VNode) {
	i.prev = v
}
func (i *FinallyProduction) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitFinallyProduction(ctx *base.FinallyProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

// debuggerStatement
//     : Debugger eos
//     ;
type DebuggerStatement struct {
	*SourceInfo
	Debugger   Token
	Eos        Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*DebuggerStatement)(nil)

func (i *DebuggerStatement) Type() string {
	return "DebuggerStatement"
}
func (i *DebuggerStatement) Code() string {
	return CodeDef(i)
}
func (i *DebuggerStatement) Next() VNode {
	return i.next
}
func (i *DebuggerStatement) SetNext(v VNode) {
	i.next = v
}
func (i *DebuggerStatement) Prev() VNode {
	return i.prev
}
func (i *DebuggerStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *DebuggerStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitDebuggerStatement(ctx *base.DebuggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

//  singleExpression TemplateStringLiteral                       // ECMAScript 6
type TemplateStringExpression struct {
	*SourceInfo
	firstChild            VNode
	SingleExp             VNode
	TemplateStringLiteral Token
	prev, next            VNode
}

var _ VNode = (*TemplateStringExpression)(nil)

func (i *TemplateStringExpression) Type() string {
	return "TemplateStringExpression"
}
func (i *TemplateStringExpression) Code() string {
	return CodeDef(i)
}
func (i *TemplateStringExpression) Next() VNode {
	return i.next
}
func (i *TemplateStringExpression) SetNext(v VNode) {
	i.next = v
}
func (i *TemplateStringExpression) Prev() VNode {
	return i.prev
}
func (i *TemplateStringExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *TemplateStringExpression) FirstChild() VNode {
	return i.firstChild
}

func (v *Visitor) VisitTemplateStringExpression(ctx *base.TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// '++' singleExpression
type PreIncrementExpression struct {
	*SourceInfo
	PlusPlus   Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*PreIncrementExpression)(nil)

func (i *PreIncrementExpression) Type() string {
	return "PreIncrementExpression"
}
func (i *PreIncrementExpression) Code() string {
	return CodeDef(i)
}
func (i *PreIncrementExpression) Next() VNode {
	return i.next
}
func (i *PreIncrementExpression) SetNext(v VNode) {
	i.next = v
}
func (i *PreIncrementExpression) Prev() VNode {
	return i.prev
}
func (i *PreIncrementExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *PreIncrementExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPreIncrementExpression(ctx *base.PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

//  New '.' identifier
type MetaExpression struct {
	*SourceInfo
	New        Token
	Dot        Token
	Identifier VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*MetaExpression)(nil)

func (i *MetaExpression) Type() string {
	return "MetaExpression"
}
func (i *MetaExpression) Code() string {
	return CodeDef(i)
}
func (i *MetaExpression) Next() VNode {
	return i.next
}
func (i *MetaExpression) SetNext(v VNode) {
	i.next = v
}
func (i *MetaExpression) Prev() VNode {
	return i.prev
}
func (i *MetaExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *MetaExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitMetaExpression(ctx *base.MetaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// '!' singleExpression
type NotExpression struct {
	*SourceInfo
	Not        Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*NotExpression)(nil)

func (i *NotExpression) Type() string {
	return "NotExpression"
}
func (i *NotExpression) Code() string {
	return CodeDef(i)
}
func (i *NotExpression) Next() VNode {
	return i.next
}
func (i *NotExpression) SetNext(v VNode) {
	i.next = v
}
func (i *NotExpression) Prev() VNode {
	return i.prev
}
func (i *NotExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *NotExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitNotExpression(ctx *base.NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// '--' singleExpression
type PreDecreaseExpression struct {
	*SourceInfo
	MinusMinus Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*PreDecreaseExpression)(nil)

func (i *PreDecreaseExpression) Type() string {
	return "PreDecreaseExpression"
}
func (i *PreDecreaseExpression) Code() string {
	return CodeDef(i)
}
func (i *PreDecreaseExpression) Next() VNode {
	return i.next
}
func (i *PreDecreaseExpression) SetNext(v VNode) {
	i.next = v
}
func (i *PreDecreaseExpression) Prev() VNode {
	return i.prev
}
func (i *PreDecreaseExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *PreDecreaseExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPreDecreaseExpression(ctx *base.PreDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

type ThisExpression struct {
	*SourceInfo
	This       Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ThisExpression)(nil)

func (i *ThisExpression) Type() string {
	return "ThisExpression"
}
func (i *ThisExpression) Code() string {
	return CodeDef(i)
}
func (i *ThisExpression) Next() VNode {
	return i.next
}
func (i *ThisExpression) SetNext(v VNode) {
	i.next = v
}
func (i *ThisExpression) Prev() VNode {
	return i.prev
}
func (i *ThisExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *ThisExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitThisExpression(ctx *base.ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// '-' singleExpression
type UnaryMinusExpression struct {
	*SourceInfo
	Minus      Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*UnaryMinusExpression)(nil)

func (i *UnaryMinusExpression) Type() string {
	return "UnaryMinusExpression"
}
func (i *UnaryMinusExpression) Code() string {
	return CodeDef(i)
}
func (i *UnaryMinusExpression) Next() VNode {
	return i.next
}
func (i *UnaryMinusExpression) SetNext(v VNode) {
	i.next = v
}
func (i *UnaryMinusExpression) Prev() VNode {
	return i.prev
}
func (i *UnaryMinusExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *UnaryMinusExpression) FirstChild() VNode {
	return i.firstChild
}

func (v *Visitor) VisitUnaryMinusExpression(ctx *base.UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

type UnaryPlusExpression struct {
	*SourceInfo
	Plus       Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*UnaryPlusExpression)(nil)

func (i *UnaryPlusExpression) Type() string {
	return "UnaryPlusExpression"
}
func (i *UnaryPlusExpression) Code() string {
	return CodeDef(i)
}
func (i *UnaryPlusExpression) Next() VNode {
	return i.next
}
func (i *UnaryPlusExpression) SetNext(v VNode) {
	i.next = v
}
func (i *UnaryPlusExpression) Prev() VNode {
	return i.prev
}
func (i *UnaryPlusExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *UnaryPlusExpression) FirstChild() VNode {
	return i.firstChild
}

func (v *Visitor) VisitUnaryPlusExpression(ctx *base.UnaryPlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// singleExpression {p.notLineTerminator()}? '--'
type PostDecreaseExpression struct {
	*SourceInfo
	SingleExp  VNode
	MinusMinus Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*PostDecreaseExpression)(nil)

func (i *PostDecreaseExpression) Type() string {
	return "PostDecreaseExpression"
}
func (i *PostDecreaseExpression) Code() string {
	return CodeDef(i)
}
func (i *PostDecreaseExpression) Next() VNode {
	return i.next
}
func (i *PostDecreaseExpression) SetNext(v VNode) {
	i.next = v
}
func (i *PostDecreaseExpression) Prev() VNode {
	return i.prev
}
func (i *PostDecreaseExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *PostDecreaseExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPostDecreaseExpression(ctx *base.PostDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// Typeof singleExpression
type TypeofExpression struct {
	*SourceInfo
	Typeof     Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*TypeofExpression)(nil)

func (i *TypeofExpression) Type() string {
	return "TypeofExpression"
}
func (i *TypeofExpression) Code() string {
	return CodeDef(i)
}
func (i *TypeofExpression) Next() VNode {
	return i.next
}
func (i *TypeofExpression) SetNext(v VNode) {
	i.next = v
}
func (i *TypeofExpression) Prev() VNode {
	return i.prev
}
func (i *TypeofExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *TypeofExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitTypeofExpression(ctx *base.TypeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// Delete singleExpression
type DeleteExpression struct {
	*SourceInfo
	Delete     Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*DeleteExpression)(nil)

func (i *DeleteExpression) Type() string {
	return "DeleteExpression"
}
func (i *DeleteExpression) Code() string {
	return CodeDef(i)
}
func (i *DeleteExpression) Next() VNode {
	return i.next
}
func (i *DeleteExpression) SetNext(v VNode) {
	i.next = v
}
func (i *DeleteExpression) Prev() VNode {
	return i.prev
}
func (i *DeleteExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *DeleteExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitDeleteExpression(ctx *base.DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

type SuperExpression struct {
	*SourceInfo
	Super      Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*SuperExpression)(nil)

func (i *SuperExpression) Type() string {
	return "SuperExpression"
}
func (i *SuperExpression) Code() string {
	return CodeDef(i)
}
func (i *SuperExpression) Next() VNode {
	return i.next
}
func (i *SuperExpression) SetNext(v VNode) {
	i.next = v
}
func (i *SuperExpression) Prev() VNode {
	return i.prev
}
func (i *SuperExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *SuperExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitSuperExpression(ctx *base.SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// singleExpression {p.notLineTerminator()}? '++'
type PostIncrementExpression struct {
	*SourceInfo
	SingleExp  VNode
	PlusPlus   Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*PostIncrementExpression)(nil)

func (i *PostIncrementExpression) Type() string {
	return "PostIncrementExpression"
}
func (i *PostIncrementExpression) Code() string {
	return CodeDef(i)
}
func (i *PostIncrementExpression) Next() VNode {
	return i.next
}
func (i *PostIncrementExpression) SetNext(v VNode) {
	i.next = v
}
func (i *PostIncrementExpression) Prev() VNode {
	return i.prev
}
func (i *PostIncrementExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *PostIncrementExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitPostIncrementExpression(ctx *base.PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// singleExpression TemplateStringLiteral
type YieldExpression struct {
	*SourceInfo
	SingleExp             VNode
	TemplateStringLiteral Token
	firstChild            VNode
	prev, next            VNode
}

var _ VNode = (*YieldExpression)(nil)

func (i *YieldExpression) Type() string {
	return "YieldExpression"
}
func (i *YieldExpression) Code() string {
	return CodeDef(i)
}
func (i *YieldExpression) Next() VNode {
	return i.next
}
func (i *YieldExpression) SetNext(v VNode) {
	i.next = v
}
func (i *YieldExpression) Prev() VNode {
	return i.prev
}
func (i *YieldExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *YieldExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitYieldExpression(ctx *base.YieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// '~' singleExpression
type BitNotExpression struct {
	*SourceInfo
	BitNot     Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*BitNotExpression)(nil)

func (i *BitNotExpression) Type() string {
	return "BitNotExpression"
}
func (i *BitNotExpression) Code() string {
	return CodeDef(i)
}
func (i *BitNotExpression) Next() VNode {
	return i.next
}
func (i *BitNotExpression) SetNext(v VNode) {
	i.next = v
}
func (i *BitNotExpression) Prev() VNode {
	return i.prev
}
func (i *BitNotExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *BitNotExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitBitNotExpression(ctx *base.BitNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

//  | New singleExpression arguments?
type NewExpression struct {
	*SourceInfo
	New        Token
	SingleExp  VNode
	Arguments  *Arguments
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*NewExpression)(nil)

func (i *NewExpression) Type() string {
	return "NewExpression"
}
func (i *NewExpression) Code() string {
	return CodeDef(i)
}
func (i *NewExpression) Next() VNode {
	return i.next
}
func (i *NewExpression) SetNext(v VNode) {
	i.next = v
}
func (i *NewExpression) Prev() VNode {
	return i.prev
}
func (i *NewExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *NewExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitNewExpression(ctx *base.NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// Class identifier? classTail
type ClassExpression struct {
	*SourceInfo
	Class      Token
	Identifier VNode
	Tail       *ClassTail
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ClassExpression)(nil)

func (i *ClassExpression) Type() string {
	return "ClassExpression"
}
func (i *ClassExpression) Code() string {
	return CodeDef(i)
}
func (i *ClassExpression) Next() VNode {
	return i.next
}
func (i *ClassExpression) SetNext(v VNode) {
	i.next = v
}
func (i *ClassExpression) Prev() VNode {
	return i.prev
}
func (i *ClassExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *ClassExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitClassExpression(ctx *base.ClassExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMemberIndexExpression(ctx *base.MemberIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVoidExpression(ctx *base.VoidExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

// argument
//     : Ellipsis? (singleExpression | identifier)
//     ;
type Argument struct {
	*SourceInfo
	Ellipsis   Token
	SingleExp  VNode
	Identifier VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*Argument)(nil)

func (i *Argument) Type() string {
	return "Argument"
}
func (i *Argument) Code() string {
	return CodeDef(i)
}
func (i *Argument) Next() VNode {
	return i.next
}
func (i *Argument) SetNext(v VNode) {
	i.next = v
}
func (i *Argument) Prev() VNode {
	return i.prev
}
func (i *Argument) SetPrev(v VNode) {
	i.prev = v
}
func (i *Argument) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitArgument(ctx *base.ArgumentContext) interface{} {

	return v.VisitChildren(ctx)
}

// public T visit(ParseTree tree)
// Visit a parse tree, and return a user-defined result of the operation.
// The default implementation calls ParseTree.accept(org.antlr.v4.runtime.tree.ParseTreeVisitor<? extends T>) on the specified tree.

// Specified by:
// visit in interface ParseTreeVisitor<T>
// Parameters:
// tree - The ParseTree to visit.
// Returns:
//     the result of visiting the parse tree.
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitSourceElements is called when production sourceElements is entered.
// Program -> this -> SourceElement
func (v *Visitor) VisitSourceElements(ctx *base.SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

// maybe make an a files on Visitor and non nil implementation gets called...
func (v *Visitor) ShouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	return true
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return node
}

// variableStatement
//     : variableDeclarationList eos
//     ;
type VariableStatement struct {
	*SourceInfo
}

func (v *Visitor) VisitVariableStatement(ctx *base.VariableStatementContext) interface{} {
	// log.Println("VisitVariableStatement", ctx.GetText())
	return v.VisitChildren(ctx)
}

// classDeclaration
//     : Class identifier classTail
//     ;

func (v *Visitor) VisitClassDeclaration(ctx *base.ClassDeclarationContext) interface{} {

	return v.VisitChildren(ctx)
}

// classTail
//     : (Extends singleExpression)? '{' classElement* '}'
//     ;
type ClassTail struct {
	*SourceInfo
	Extends    Token
	SingleExp  VNode
	OpenBrace  Token
	Elements   *ClassElement
	CloseBrace Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ClassTail)(nil)

func (i *ClassTail) Type() string {
	return "ClassTail"
}
func (i *ClassTail) Code() string {
	return CodeDef(i)
}
func (i *ClassTail) Next() VNode {
	return i.next
}
func (i *ClassTail) SetNext(v VNode) {
	i.next = v
}
func (i *ClassTail) Prev() VNode {
	return i.prev
}
func (i *ClassTail) SetPrev(v VNode) {
	i.prev = v
}
func (i *ClassTail) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitClassTail(ctx *base.ClassTailContext) interface{} {

	return v.VisitChildren(ctx)
}

// classElement
//     : (Static | {p.n("static")}? identifier | Async)* (methodDefinition | assignable '=' objectLiteral ';')
//     | emptyStatement_
//     | '#'? propertyName '=' singleExpression
//     ;
// todo: split into 3 types
type ClassElement struct {
	*SourceInfo
	Static           Token
	Identifier       VNode
	Async            VNode
	MethodDefinition *MethodDefinition
	Assignable       VNode
	Equals           Token
	ObjectLiteral    *ObjectLiteral

	// Method
	EmptyStatement Token
	PropertyName   *PropertyName
	// equals ^^
	SingleExp VNode

	//
	Hashtag Token

	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ClassElement)(nil)

func (i *ClassElement) Type() string {
	return "ClassElement"
}
func (i *ClassElement) Code() string {
	return CodeDef(i)
}
func (i *ClassElement) Next() VNode {
	return i.next
}
func (i *ClassElement) SetNext(v VNode) {
	i.next = v
}
func (i *ClassElement) Prev() VNode {
	return i.prev
}
func (i *ClassElement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ClassElement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitClassElement(ctx *base.ClassElementContext) interface{} {

	return v.VisitChildren(ctx)
}

// methodDefinition
//     : '*'? '#'? propertyName '(' formalParameterList? ')' functionBody
//     | '*'? '#'? getter '(' ')' functionBody
//     | '*'? '#'? setter '(' formalParameterList? ')' functionBody
//     ;
type MethodDefinition struct {
	*SourceInfo
	Multiply     Token //Multiply
	Hashtag      Token
	PropertyName *PropertyName
	OpenParen    Token
	ParamsList   *FormalParameterList
	CloseParen   Token
	FunctionBody *FunctionBody

	//
	Getter *Getter
	Setter *Setter

	firstChild VNode
	prev, next VNode
}

var _ VNode = (*MethodDefinition)(nil)

func (i *MethodDefinition) Type() string {
	return "MethodDefinition"
}
func (i *MethodDefinition) Code() string {
	return CodeDef(i)
}
func (i *MethodDefinition) Next() VNode {
	return i.next
}
func (i *MethodDefinition) SetNext(v VNode) {
	i.next = v
}
func (i *MethodDefinition) Prev() VNode {
	return i.prev
}
func (i *MethodDefinition) SetPrev(v VNode) {
	i.prev = v
}
func (i *MethodDefinition) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitMethodDefinition(ctx *base.MethodDefinitionContext) interface{} {

	return v.VisitChildren(ctx)
}

//  '(' expressionSequence ')'
type ParenthesizedExpression struct {
	*SourceInfo
	OpenParen          Token
	ExpressionSequence *ExpressionSequence

	CloseParen Token
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ParenthesizedExpression)(nil)

func (i *ParenthesizedExpression) Type() string {
	return "ParenthesizedExpression"
}
func (i *ParenthesizedExpression) Code() string {
	return CodeDef(i)
}
func (i *ParenthesizedExpression) Next() VNode {
	return i.next
}
func (i *ParenthesizedExpression) SetNext(v VNode) {
	i.next = v
}
func (i *ParenthesizedExpression) Prev() VNode {
	return i.prev
}
func (i *ParenthesizedExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *ParenthesizedExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitParenthesizedExpression(ctx *base.ParenthesizedExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

// Await singleExpression
type AwaitExpression struct {
	*SourceInfo
	Await      Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*AwaitExpression)(nil)

func (i *AwaitExpression) Type() string {
	return "AwaitExpression"
}
func (i *AwaitExpression) Code() string {
	return CodeDef(i)
}
func (i *AwaitExpression) Next() VNode {
	return i.next
}
func (i *AwaitExpression) SetNext(v VNode) {
	i.next = v
}
func (i *AwaitExpression) Prev() VNode {
	return i.prev
}
func (i *AwaitExpression) SetPrev(v VNode) {
	i.prev = v
}
func (i *AwaitExpression) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitAwaitExpression(ctx *base.AwaitExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

// getter
//     : {p.n("get")}? identifier propertyName
//     ;
type Getter struct {
	*SourceInfo
	Identifier   VNode
	PropertyName *PropertyName
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*Getter)(nil)

func (i *Getter) Type() string {
	return "Getter"
}
func (i *Getter) Code() string {
	return CodeDef(i)
}
func (i *Getter) Next() VNode {
	return i.next
}
func (i *Getter) SetNext(v VNode) {
	i.next = v
}
func (i *Getter) Prev() VNode {
	return i.prev
}
func (i *Getter) SetPrev(v VNode) {
	i.prev = v
}
func (i *Getter) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitGetter(ctx *base.GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

// setter
//     : {p.n("set")}? identifier propertyName
//     ;
type Setter struct {
	*SourceInfo
	Identifier   VNode
	PropertyName *PropertyName
	firstChild   VNode
	prev, next   VNode
}

var _ VNode = (*Setter)(nil)

func (i *Setter) Type() string {
	return "Setter"
}
func (i *Setter) Code() string {
	return CodeDef(i)
}
func (i *Setter) Next() VNode {
	return i.next
}
func (i *Setter) SetNext(v VNode) {
	i.next = v
}
func (i *Setter) Prev() VNode {
	return i.prev
}
func (i *Setter) SetPrev(v VNode) {
	i.prev = v
}
func (i *Setter) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitSetter(ctx *base.SetterContext) interface{} {
	return v.VisitChildren(ctx)
}
