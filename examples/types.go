package main

// interface Identifier <: Expression, Pattern {
//     type: "Identifier";
//     name: string;
// }

// An identifier. Note that an identifier may be an expression or a destructuring pattern.
type Identifier struct {
	Expression
}

// Any expression node.
// Since the left-hand side of an assignment may be any expression in general, an expression can also be a pattern.
type Expression struct {
	VNode
}
type Literal struct {
	typeOf string
	Value  interface{} //string | boolean | null | number | RegExp;
}

// A literal token. Note that a literal can be an expression.
