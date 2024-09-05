package parser

import (
	"strconv"

	"github.com/dtasada/doodle/src/ast"
	"github.com/dtasada/doodle/src/lexer"
)

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	// First parse the NUD
	tokenKind := p.currentToken().Kind
	nudFunc, exists := nudLookup[tokenKind]

	if !exists {
		lexer.Panic("NUD handler expected for token", tokenKind.ToString())
	}

	left := nudFunc(p)

	for bpLookup[p.currentToken().Kind] > bp {
		tokenKind = p.currentToken().Kind
		ledFunc, exists := ledLookup[tokenKind]

		if !exists {
			lexer.Panic("LED handler expected for token", tokenKind.ToString())
		}

		left = ledFunc(p, left, bpLookup[p.currentToken().Kind]) // recursion
	}

	return left
}

func parsePrimaryExpression(p *parser) ast.Expression {
	switch p.currentToken().Kind {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpression{Value: number}
	case lexer.STRING:
		return ast.StringExpression{Value: p.advance().Value}
	case lexer.IDENTIFIER:
		return ast.SymbolExpression{Value: p.advance().Value}
	default:
		lexer.Panic("Cannot create primaryExpression from", p.currentToken().Kind.ToString())
		return nil
	}
}

func parseBinaryExpression(p *parser, left ast.Expression, bp bindingPower) ast.Expression {
	operatorToken := p.advance()
	right := parseExpression(p, bp)

	return ast.BinaryExpression{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

func parsePrefixExpression(p *parser) ast.Expression {
	operatorToken := p.advance()
	rightHand := parseExpression(p, defaultBp)

	return ast.PrefixExpression{
		Operator:        operatorToken,
		RightExpression: rightHand,
	}
}

func parseGroupingExpression(p *parser) ast.Expression {
	p.advance() // skip groupint start
	expression := parseExpression(p, defaultBp)
	p.expect(lexer.CLOSE_PAREN)
	return expression
}

func parseAssignmentExpression(p *parser, left ast.Expression, bp bindingPower) ast.Expression {
	operatorToken := p.advance()
	rightHand := parseExpression(p, assignment)

	return ast.AssignmentExpression{
		Assignee:        left,
		Operator:        operatorToken,
		RightExpression: rightHand,
	}
}
