package ast

import "github.com/bookrun-go/calculator/token"

type Node interface {
	Result() (token.Value, error)
}
