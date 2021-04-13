package main

// ** Not sure if pointer receivers are better or not. Means checkign for nil a lot more
// but does allow easily manipulating struct values...**

// importStatement
//     : Import importFromBlock
//     ;
type ImportStatement struct {
	*SourceInfo
	// one child *ImportFromBlock
	ImportFromBlock *ImportFromBlock
	// one child *ImportFromBlock
	Children []VNode //
}

var _ VNode = (*ImportStatement)(nil)

func (i *ImportStatement) Type() string {
	return "ImportStatement"
}
func (i *ImportStatement) Code() string {
	return CodeDef(i)
}
func (i *ImportStatement) GetChildren() []VNode {
	return i.Children
}

// importFromBlock
//     : importDefault? (importNamespace | moduleItems) importFrom eos
//     | StringLiteral eos
// importDefault
// : aliasName ','
type ImportFromBlock struct {
	*SourceInfo

	Default         *ImportDefault
	ModulesItems    *ModulesItems
	ImportNamespace *ImportNamespace
	StringLiteral   string
	ImportFrom      *ImportFrom
}

var _ VNode = (*ImportFromBlock)(nil)

func (i *ImportFromBlock) Type() string {
	return "ImportFromBlock"
}
func (i *ImportFromBlock) Code() string {
	return CodeDef(i)
}
func (i *ImportFromBlock) GetChildren() []VNode {
	return nil
}

type ImportFrom struct {
	*SourceInfo

	From string
	Path string
}

var _ VNode = (*ImportFrom)(nil)

func f(v VNode) {

}
func (i *ImportFrom) Type() string {
	return "ImportFrom"
}

func (i *ImportFrom) Code() string {
	return CodeDef(i)
}
func (i *ImportFrom) GetChildren() []VNode {
	return nil
}

// import '(' singleExpression ')'                                       # ImportExpression
type ImportExpression struct {
	*SourceInfo
	SingleExpression string
	Children         []VNode //
}

var _ VNode = (*ImportExpression)(nil)

func (i *ImportExpression) Code() string {
	return CodeDef(i)
}

func (i *ImportExpression) Type() string {
	return "ImportExpression"
}
func (i *ImportExpression) GetChildren() []VNode {
	return i.Children
}

// aliasName
//     : identifierName (As identifierName)?
//     ;
type AliasName struct {
	*SourceInfo
	IdentifierName string
	Alias          string
	Children       []VNode //
}

var _ VNode = (*AliasName)(nil)

func (a *AliasName) Code() string {
	return CodeDef(a)
}
func (i *AliasName) Type() string {

	return "AliasName"
}
func (i *AliasName) GetChildren() []VNode {
	return nil
}

// moduleItems
//     : '{' (aliasName ',')* (aliasName ','?)? '}'
//     ;
type ModulesItems struct {
	*SourceInfo
	AliasNames []*AliasName
	// always AliasName(s)
	Children []VNode //
}

var _ VNode = (*ModulesItems)(nil)

func (m *ModulesItems) Code() string {

	return CodeDef(m)
}
func (m *ModulesItems) Type() string {
	return "ModulesItems"
}
func (m *ModulesItems) GetChildren() []VNode {
	return m.Children
}

// importNamespace
//     : ('*' | identifierName) (As identifierName)?
//     ;
type ImportNamespace struct {
	*SourceInfo
	// Star           string
	// IdentifierName string
	// AliasName      string
	Children []VNode //
}

var _ VNode = (*ImportNamespace)(nil)

func (in *ImportNamespace) Code() string {
	return CodeDef(in)
}
func (i *ImportNamespace) Type() string {
	return "ImportNamespace"
}
func (i *ImportNamespace) GetChildren() []VNode {
	return i.Children
}

// importDefault
// : aliasName ','
type ImportDefault struct {
	*SourceInfo
	From     string
	Children []VNode //
}

var _ VNode = (*ImportDefault)(nil)

func (i *ImportDefault) Type() string {
	return "ImportDefault"
}

func (i *ImportDefault) Code() string {
	return CodeDef(i)
}
func (i *ImportDefault) GetChildren() []VNode {
	return i.Children
}
