package costumize

import (
	"github.com/bookrun-go/calculator/token"
)

type Node struct {
	val token.Value
}

func (n Node) Result() (token.Value, error) {
	return n.val, nil
}
