package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

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
		log.Println("VisitLiteralExpression", ctx.GetText())
	}
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}
