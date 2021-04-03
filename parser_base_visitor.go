package gojsp // JavaScriptParser

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
)

// v.VisitChildren(ctx) for all below not implemented in own "visitor" calls antlr.BaseParseTreeVisitors VisitChildren
// i.e func (v *BaseParseTreeVisitor) VisitChildren(node RuleNode) interface{}     { return nil }
// so does nothing basically...
type BaseJavaScriptParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJavaScriptParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitSourceElement(ctx *SourceElementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitStatement(ctx *StatementContext) interface{} {

	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitImportFromBlock(ctx *ImportFromBlockContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitModuleItems(ctx *ModuleItemsContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitImportDefault(ctx *ImportDefaultContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitImportNamespace(ctx *ImportNamespaceContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitImportFrom(ctx *ImportFromContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAliasName(ctx *AliasNameContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitExportDeclaration(ctx *ExportDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitExportFromBlock(ctx *ExportFromBlockContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	log.Printf("%+v\n", v)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitEmptyStatement(ctx *EmptyStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitForInStatement(ctx *ForInStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitForOfStatement(ctx *ForOfStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitVarModifier(ctx *VarModifierContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitYieldStatement(ctx *YieldStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitWithStatement(ctx *WithStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitCaseClauses(ctx *CaseClausesContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLabelledStatement(ctx *LabelledStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitCatchProduction(ctx *CatchProductionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFinallyProduction(ctx *FinallyProductionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitClassTail(ctx *ClassTailContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitClassElement(ctx *ClassElementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitMethodDefinition(ctx *MethodDefinitionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitSourceElements(ctx *SourceElementsContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArrayElement(ctx *ArrayElementContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFunctionProperty(ctx *FunctionPropertyContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPropertyGetter(ctx *PropertyGetterContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPropertySetter(ctx *PropertySetterContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitMetaExpression(ctx *MetaExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAwaitExpression(ctx *AwaitExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitTypeofExpression(ctx *TypeofExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitDeleteExpression(ctx *DeleteExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitImportExpression(ctx *ImportExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitSuperExpression(ctx *SuperExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitYieldExpression(ctx *YieldExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBitNotExpression(ctx *BitNotExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitNewExpression(ctx *NewExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitClassExpression(ctx *ClassExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitVoidExpression(ctx *VoidExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitCoalesceExpression(ctx *CoalesceExpressionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAssignable(ctx *AssignableContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArrowFunction(ctx *ArrowFunctionContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitBigintLiteral(ctx *BigintLiteralContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitGetter(ctx *GetterContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitSetter(ctx *SetterContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitIdentifierName(ctx *IdentifierNameContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitLet_(ctx *Let_Context) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParserVisitor) VisitEos(ctx *EosContext) interface{} {
	log.Println(ctx)
	return ctx.GetText()
}
