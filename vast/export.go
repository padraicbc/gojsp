package vast

import (
	"log"

	"github.com/padraicbc/gojsp/base"
)

// exportStatement
//     : Export (exportFromBlock | declaration) eos    # ExportDeclaration
//     | Export Default singleExpression eos           # ExportDefaultDeclaration
//     ;
type ExportDeclarationStatement struct {
	*SourceInfo
	Export          Token
	ExportFromBlock *ExportFromBlock
	Declaration     *Declaration
	Eos             Token
	firstChild      VNode
	prev, next      VNode
}

var _ VNode = (*ExportDeclarationStatement)(nil)

func (i *ExportDeclarationStatement) Type() string {
	return "ExportDeclarationStatement"
}
func (i *ExportDeclarationStatement) Code() string {
	return CodeDef(i)
}
func (i *ExportDeclarationStatement) Next() VNode {
	return i.next
}
func (i *ExportDeclarationStatement) SetNext(v VNode) {
	i.next = v
}
func (i *ExportDeclarationStatement) Prev() VNode {
	return i.prev
}
func (i *ExportDeclarationStatement) SetPrev(v VNode) {
	i.prev = v
}
func (i *ExportDeclarationStatement) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitExportDeclaration(ctx *base.ExportDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitExportDeclaration", ctx.GetText())
	}

	ed := &ExportDeclarationStatement{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if ed.firstChild == nil {
			ed.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {

		case "ExportFromBlock":
			ed.ExportFromBlock = ch.(*ExportFromBlock)
		case "Declaration":
			ed.Declaration = ch.(*Declaration)

		case "LToken":
			tk := ch.(Token)
			if tk.SymbolName() == "Export" {
				ed.Export = tk
				continue
			}
			if tk.SymbolName() == "SemiColon" {
				ed.Eos = tk
			}

			log.Panicf("%+v %s\n", ch, ch.Type())
		default:
			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}

	return v.VisitChildren(ctx)

}

// Export Default singleExpression eos
type ExportDefaultDeclaration struct {
	*SourceInfo
	Export     Token
	Default    Token
	SingleExp  VNode
	firstChild VNode
	prev, next VNode
}

var _ VNode = (*ExportDefaultDeclaration)(nil)

func (i *ExportDefaultDeclaration) Type() string {
	return "ExportDefaultDeclaration"
}
func (i *ExportDefaultDeclaration) Code() string {
	return CodeDef(i)
}
func (i *ExportDefaultDeclaration) Next() VNode {
	return i.next
}
func (i *ExportDefaultDeclaration) SetNext(v VNode) {
	i.next = v
}
func (i *ExportDefaultDeclaration) Prev() VNode {
	return i.prev
}
func (i *ExportDefaultDeclaration) SetPrev(v VNode) {
	i.prev = v
}
func (i *ExportDefaultDeclaration) FirstChild() VNode {
	return i.firstChild
}

func (v *Visitor) VisitExportDefaultDeclaration(ctx *base.ExportDefaultDeclarationContext) interface{} {
	if v.Debug {
		log.Println("VisitExportDefaultDeclaration", ctx.GetText())
	}
	ed := &ExportDefaultDeclaration{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}
	ed.Export = ident(v, ctx.Export().GetSymbol())
	ed.Default = ident(v, ctx.Default().GetSymbol())
	ed.SingleExp = v.Visit(ctx.SingleExpression()).(VNode)
	setAllSibs(ed.Export, ed.Default, ed.SingleExp)

	return ed
}

// exportFromBlock
//     : importNamespace importFrom eos
//     | moduleItems importFrom? eos
//     ;
type ExportFromBlock struct {
	*SourceInfo
	ImportNamespace *ImportNamespace
	ImportFrom      *ImportFrom
	ModuleItems     *ModuleItems
	Eos             Token
	firstChild      VNode
	prev, next      VNode
}

var _ VNode = (*ExportFromBlock)(nil)

func (i *ExportFromBlock) Type() string {
	return "ExportFromBlock"
}
func (i *ExportFromBlock) Code() string {
	return CodeDef(i)
}
func (i *ExportFromBlock) Next() VNode {
	return i.next
}
func (i *ExportFromBlock) SetNext(v VNode) {
	i.next = v
}
func (i *ExportFromBlock) Prev() VNode {
	return i.prev
}
func (i *ExportFromBlock) SetPrev(v VNode) {
	i.prev = v
}
func (i *ExportFromBlock) FirstChild() VNode {
	return i.firstChild
}
func (v *Visitor) VisitExportFromBlock(ctx *base.ExportFromBlockContext) interface{} {
	if v.Debug {
		log.Println("VisitExportFromBlock", ctx.GetText())
	}
	eb := &ExportFromBlock{SourceInfo: getSourceInfo(*ctx.BaseParserRuleContext)}

	var prev VNode
	// iterate here as some are there, some not.
	for _, ch := range v.VisitChildren(ctx).([]VNode) {
		if eb.firstChild == nil {
			eb.firstChild = ch
		}
		prev = setSib(prev, ch)

		switch ch.Type() {

		case "ImportNamespace":
			eb.ImportNamespace = ch.(*ImportNamespace)
		case "ModuleItems":
			eb.ModuleItems = ch.(*ModuleItems)
		case "ImportFrom":
			eb.ImportFrom = ch.(*ImportFrom)
		case "LToken":
			tk := ch.(Token)
			if tk.SymbolName() == "SemiColon" {
				eb.Eos = tk
				continue
			}

			log.Panicf("%+v %s\n", ch, ch.Type())

		}
	}

	return eb
}
