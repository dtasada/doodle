package ast

import "github.com/dtasada/doodle/src/lexer"

type NumberExpression struct {
	Value float64
}

type StringExpression struct {
	Value string
}

type SymbolExpression struct {
	Value string
}

type BinaryExpression struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

type PrefixExpression struct {
	Operator        lexer.Token
	RightExpression Expression
}

type AssignmentExpression struct {
	Assignee        Expression
	Operator        lexer.Token
	RightExpression Expression
}

// Comply with Expression interface
func (n NumberExpression) expression()     {}
func (n StringExpression) expression()     {}
func (n SymbolExpression) expression()     {}
func (n BinaryExpression) expression()     {}
func (n PrefixExpression) expression()     {}
func (n AssignmentExpression) expression() {}
