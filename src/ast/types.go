package ast

type SymbolType struct {
	Name string // T
}

type ArrayType struct {
	Underlying Type // []T
}

func (t SymbolType) _type() {}
func (t ArrayType) _type()  {}
