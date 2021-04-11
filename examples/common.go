package main

import antlr "github.com/padraicbc/antlr4"

// just pulls original source code info
func getSourceInfo(ctx antlr.BaseParserRuleContext) *SourceInfo {

	return &SourceInfo{Line: ctx.GetStart().GetLine(), Start: ctx.GetStart().GetStart(), End: ctx.GetStop().GetStart(),
		Source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: ctx.GetStart().GetStart(), Stop: ctx.GetStop().GetStop() + 1})}

}
