package ast

import (
	"errors"

	"github.com/bookrun-go/calculator/token"
)

type OperatorNode struct {
	tok   token.Token
	left  Node
	right Node
}

func (op *OperatorNode) AddNode(node Node) error {
	if op.left == nil {
		op.left = node
	} else if op.right == nil {
		op.right = node
	} else {
		return errors.New("add node error")
	}

	return nil
}

func (op OperatorNode) Result() (token.Value, error) {
	if op.left == nil {
		return token.EmptyValue{}, errors.New("left can't empty")
	}

	if op.right == nil {
		if !op.tok.IsIllegal() {
			return token.EmptyValue{}, ErrorFomulaFormat
		}
		return op.left.Result()
	}

	leftVal, err := op.left.Result()
	if err != nil {
		return token.EmptyValue{}, err
	}
	rightVal, err := op.right.Result()
	if err != nil {
		return token.EmptyValue{}, err
	}

	return rightVal.Operate(leftVal, op.tok)
}
