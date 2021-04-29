package vast

import (
	"errors"
	"io"
	"log"

	antlr "github.com/padraicbc/antlr4"
	"github.com/padraicbc/gojsp/base"
)

// identifierName
//     : identifier
//     | reservedWord
//     ;
func (v *Visitor) VisitIdentifierName(ctx *base.IdentifierNameContext) interface{} {
	if v.Debug {
		log.Println("VisitIdentifierName", ctx.GetText())
	}

	if ctx.Identifier() != nil {
		return v.VisitIdentifier(ctx.Identifier().(*base.IdentifierContext)).(*LToken)

	}

	return v.VisitReservedWord(ctx.ReservedWord().(*base.ReservedWordContext)).(*LToken)

}

func (v *Visitor) VisitKeyword(ctx *base.KeywordContext) interface{} {
	if v.Debug {
		log.Println("VisitKeyword", ctx.GetText(), ctx.GetChildCount())
	}

	return v.Visit(ctx.GetChild(0).(antlr.ParseTree)).(VNode)

}

// reservedWord
//     : keyword
//     | NullLiteral
//     | BooleanLiteral
//     ;
func (v *Visitor) VisitReservedWord(ctx *base.ReservedWordContext) interface{} {
	if v.Debug {
		log.Println("VisitReservedWord", ctx.GetText(), ctx.GetChildCount())
	}
	if ctx.Keyword() != nil {
		return v.VisitKeyword(ctx.Keyword().(*base.KeywordContext)).(VNode)
	}
	if ctx.NullLiteral() != nil {
		return ident(v, ctx.NullLiteral().GetSymbol())
	}
	if ctx.BooleanLiteral() != nil {
		return ident(v, ctx.BooleanLiteral().GetSymbol())
	}

	panic("")
}

// eos
//     : SemiColon
//     | EOF
//     | {p.lineTerminatorAhead()}?
//     | {p.closeBrace()}?
//     ;
func (v *Visitor) VisitEos(ctx *base.EosContext) interface{} {
	if v.Debug {
		log.Println("VisitEos", ctx.GetText())
	}
	//  no ; after closing } so empty string ?
	if ctx.GetChildCount() == 0 {
		return errors.New("Optional semicolon not there")
	}

	if ctx.EOF() != nil {
		return io.EOF
	}
	if ctx.SemiColon() != nil {
		return ident(v, ctx.SemiColon().GetSymbol())
	}
	// can be }?
	return ident(v, ctx.GetChild(0).(antlr.TerminalNode).GetSymbol())

}

// identifier
//     : Identifier
//     | Async
//     ;
func (v *Visitor) VisitIdentifier(ctx *base.IdentifierContext) interface{} {
	// log.Println("VisitIdentifier", ctx.GetText(), ctx.GetChildCount())
	if v.Debug {
		log.Println("VisitIdentifier", ctx.GetText())
	}
	if ctx.Identifier() != nil {
		return ident(v, ctx.Identifier().GetSymbol())
	}

	return ident(v, ctx.Async().GetSymbol())
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return ident(v, node.GetSymbol())

}

// LITERALS

// literal
//     : NullLiteral
//     | BooleanLiteral
//     | StringLiteral
//     | TemplateStringLiteral
//     | RegularExpressionLiteral
//     | numericLiteral
//     | bigintLiteral
//     ;
func (v *Visitor) VisitLiteral(ctx *base.LiteralContext) interface{} {
	if v.Debug {
		log.Println("VisitLiteral", ctx.GetText())
	}
	// first child again as only ne
	return v.Visit(ctx.GetChild(0).(antlr.ParseTree))
}

// numericLiteral
//     : DecimalLiteral
//     | HexIntegerLiteral
//     | OctalIntegerLiteral
//     | OctalIntegerLiteral2
//     | BinaryIntegerLiteral
//     ;

func (v *Visitor) VisitNumericLiteral(ctx *base.NumericLiteralContext) interface{} {
	if v.Debug {
		log.Println("VisitNumericLiteral", ctx.GetText())
	}
	// Unless I am mistaken, the first child should be non nil so no need to if ctx.whatever != nil for all or visit all children
	return ident(v, ctx.GetChild(0).(antlr.TerminalNode).GetSymbol())
}

//bigintLiteral
// : BigDecimalIntegerLiteral
// | BigHexIntegerLiteral
// | BigOctalIntegerLiteral
// | BigBinaryIntegerLiteral
// ;
func (v *Visitor) VisitBigintLiteral(ctx *base.BigintLiteralContext) interface{} {
	if v.Debug {
		log.Println("VisitBigintLiteral", ctx.GetText())
	}
	return ident(v, ctx.GetChild(0).(antlr.TerminalNode).GetSymbol())
}

// ASSIGNMENT

func (v *Visitor) VisitAssignmentOperator(ctx *base.AssignmentOperatorContext) interface{} {
	if v.Debug {
		log.Println("VisitAssignmentOperator", ctx.GetText())
	}

	return ident(v, ctx.GetChild(0).(antlr.TerminalNode).GetSymbol())
}
