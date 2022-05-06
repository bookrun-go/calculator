package ast

import "github.com/bookrun-go/calculator/token"

type NumberNode struct {
	Val token.Value
}

func (num NumberNode) Result() (token.Value, error) {
	return num.Val, nil
}
