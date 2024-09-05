package ast

type BlockStatement struct {
	Body []Statement
}

func (n BlockStatement) statement() {}

type ExpressionStatement struct {
	Expression Expression
}

func (n ExpressionStatement) statement() {}

type VarDeclStatement struct {
	VariableIdentifier string
	IsMutable          bool
	AssignedValue      Expression
	ExplicitType       Type
}

func (n VarDeclStatement) statement() {}
