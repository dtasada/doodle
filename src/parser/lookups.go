package parser

import (
	"github.com/dtasada/doodle/src/ast"
	"github.com/dtasada/doodle/src/lexer"
)

type bindingPower int

const (
	// Do not change order of iota
	defaultBp bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	cal
	member
	primary
)

type (
	ledHandler       func(p *parser, left ast.Expression, bp bindingPower) ast.Expression
	nudHandler       func(p *parser) ast.Expression
	statementHandler func(p *parser) ast.Statement
)

var (
	bpLookup        = map[lexer.TokenKind]bindingPower{}
	ledLookup       = map[lexer.TokenKind]ledHandler{}
	nudLookup       = map[lexer.TokenKind]nudHandler{}
	statementLookup = map[lexer.TokenKind]statementHandler{}
)

func led(kind lexer.TokenKind, bp bindingPower, ledFunc ledHandler) {
	bpLookup[kind] = bp
	ledLookup[kind] = ledFunc
}

func nud(kind lexer.TokenKind, nudFunc nudHandler) {
	nudLookup[kind] = nudFunc
}

func statement(kind lexer.TokenKind, statementFunc statementHandler) {
	bpLookup[kind] = defaultBp
	statementLookup[kind] = statementFunc
}

func createTokenLookups() {
	led(lexer.ASSIGNMENT, assignment, parseAssignmentExpression)
	led(lexer.PLUS_EQUALS, assignment, parseAssignmentExpression)
	led(lexer.MINUS_EQUALS, assignment, parseAssignmentExpression)
	// add *=, /= and %=

	// Logical
	led(lexer.AND, logical, parseBinaryExpression)
	led(lexer.OR, logical, parseBinaryExpression)
	led(lexer.ELLIPSIS, logical, parseBinaryExpression)

	// Relational
	led(lexer.LESS, relational, parseBinaryExpression)
	led(lexer.LESS_EQUALS, relational, parseBinaryExpression)
	led(lexer.GREATER, relational, parseBinaryExpression)
	led(lexer.GREATER_EQUALS, relational, parseBinaryExpression)
	led(lexer.EQUALS, relational, parseBinaryExpression)
	led(lexer.NOT_EQUALS, relational, parseBinaryExpression)

	// Additive & Multiplicative
	led(lexer.PLUS, additive, parseBinaryExpression)
	led(lexer.DASH, additive, parseBinaryExpression)

	led(lexer.ASTERISK, multiplicative, parseBinaryExpression)
	led(lexer.SLASH, multiplicative, parseBinaryExpression)
	led(lexer.PERCENT, multiplicative, parseBinaryExpression)

	// Literals & Symbols
	nud(lexer.NUMBER, parsePrimaryExpression)
	nud(lexer.STRING, parsePrimaryExpression)
	nud(lexer.IDENTIFIER, parsePrimaryExpression)
	nud(lexer.OPEN_PAREN, parseGroupingExpression)
	nud(lexer.DASH, parsePrefixExpression)

	// Statements
	statement(lexer.LET, parseVarDeclStatement)
	statement(lexer.MUT, parseVarDeclStatement)
}
