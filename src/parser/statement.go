package parser

import (
	"github.com/dtasada/doodle/src/ast"
	"github.com/dtasada/doodle/src/lexer"
)

func parseStatement(p *parser) ast.Statement {
	statementFunc, exists := statementLookup[p.currentToken().Kind]

	if exists {
		return statementFunc(p)
	}

	expression := parseExpression(p, defaultBp)
	p.expect(lexer.SEMICOLON)

	return ast.ExpressionStatement{
		Expression: expression,
	}
}

func parseVarDeclStatement(p *parser) ast.Statement {
	var explicitType ast.Type
	var assignedValue ast.Expression

	isMutable := p.advance().Kind == lexer.MUT
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable declaration: expected to find variable name").Value

	if p.currentToken().Kind == lexer.COLON {
		p.advance()
		explicitType = parseType(p, defaultBp)
	}

	if p.currentToken().Kind != lexer.SEMICOLON {
		p.expect(lexer.ASSIGNMENT)
		assignedValue = parseExpression(p, assignment)
	} else if explicitType == nil {
		lexer.Panic("Missing either right-hand side or type declaration in var declaration")
	}

	p.expect(lexer.SEMICOLON)

	if !isMutable && assignedValue == nil {
		lexer.Panic("Cannot define constant without providing value")
	}

	return ast.VarDeclStatement{
		IsMutable:          isMutable,
		VariableIdentifier: varName,
		AssignedValue:      assignedValue,
		ExplicitType:       explicitType,
	}
}
