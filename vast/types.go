package vast

type SourceInfo struct {
	Start, End, Line, Column int
	Source                   string
}

func (s *SourceInfo) GetInfo() *SourceInfo {
	return s
}

// Parent will have type VNode
type Token interface {
	VNode
	SetValue(string)
	Value() string
	rname(string) string
	SymbolName() string
}

type LToken struct {
	value string
	*SourceInfo
	// From .. StringLiteral...
	sn string
	// RuleName .. reservedWord...
	rn         string
	prev, next VNode
}

var _ Token = (*LToken)(nil)
var _ VNode = (*LToken)(nil)

func (i *LToken) Value() string {
	return i.value
}
func (i *LToken) SetValue(s string) {
	i.value = s
}
func (i *LToken) SymbolName() string {
	return i.sn
}
func (i *LToken) Code() string {
	return i.value
}
func (i *LToken) rname(s string) string {
	if s != "" {
		i.rn = s
		return ""
	}
	return i.rn
}
func (i *LToken) FirstChild() VNode {
	return nil
}
func (i *LToken) SetChild(ch, prev VNode) {
	return
}
func (i *LToken) Next() VNode {

	return i.next
}
func (i *LToken) SetNext(v VNode) {
	i.next = v
}

func (i *LToken) Prev() VNode {

	return i.prev
}
func (i *LToken) SetPrev(v VNode) {
	i.prev = v
}

// keyword, reservedword, identifier
func (i *LToken) Type() string {
	return "LToken"
}
func (i *LToken) GetInfo() *SourceInfo {
	return i.SourceInfo
}
