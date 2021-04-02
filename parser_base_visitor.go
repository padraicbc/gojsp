package gojsp // JavaScriptParser

import (
	"log"

	antlr "github.com/padraicbc/antlr4"
)

type BaseJavaScriptgojspisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJavaScriptgojspisitor) VisitProgram(ctx *ProgramContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitSourceElement(ctx *SourceElementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitStatement(ctx *StatementContext) interface{} {
	// //log.Println(ctx)
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBlock(ctx *BlockContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitImportFromBlock(ctx *ImportFromBlockContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitModuleItems(ctx *ModuleItemsContext) interface{} {
	log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitImportDefault(ctx *ImportDefaultContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitImportNamespace(ctx *ImportNamespaceContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitImportFrom(ctx *ImportFromContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAliasName(ctx *AliasNameContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitExportDeclaration(ctx *ExportDeclarationContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitExportFromBlock(ctx *ExportFromBlockContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitEmptyStatement(ctx *EmptyStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	// //log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitForInStatement(ctx *ForInStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitForOfStatement(ctx *ForOfStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitVarModifier(ctx *VarModifierContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitYieldStatement(ctx *YieldStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitWithStatement(ctx *WithStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitCaseClauses(ctx *CaseClausesContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLabelledStatement(ctx *LabelledStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitCatchProduction(ctx *CatchProductionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFinallyProduction(ctx *FinallyProductionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitClassTail(ctx *ClassTailContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitClassElement(ctx *ClassElementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitMethodDefinition(ctx *MethodDefinitionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitSourceElements(ctx *SourceElementsContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitElementList(ctx *ElementListContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArrayElement(ctx *ArrayElementContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFunctionProperty(ctx *FunctionPropertyContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPropertyGetter(ctx *PropertyGetterContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPropertySetter(ctx *PropertySetterContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitMetaExpression(ctx *MetaExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAwaitExpression(ctx *AwaitExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitTypeofExpression(ctx *TypeofExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitDeleteExpression(ctx *DeleteExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitImportExpression(ctx *ImportExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitSuperExpression(ctx *SuperExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitYieldExpression(ctx *YieldExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBitNotExpression(ctx *BitNotExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitNewExpression(ctx *NewExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitClassExpression(ctx *ClassExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitVoidExpression(ctx *VoidExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitCoalesceExpression(ctx *CoalesceExpressionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAssignable(ctx *AssignableContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArrowFunction(ctx *ArrowFunctionContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitBigintLiteral(ctx *BigintLiteralContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitGetter(ctx *GetterContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitSetter(ctx *SetterContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitIdentifierName(ctx *IdentifierNameContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitLet_(ctx *Let_Context) interface{} {
	//log.Println(ctx)
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptgojspisitor) VisitEos(ctx *EosContext) interface{} {
	// //log.Println(ctx)
	return v.VisitChildren(ctx)
}
