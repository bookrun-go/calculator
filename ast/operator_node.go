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

func (op OperatorNode) Result() (float64, error) {
	if !op.tok.IsOperator() {
		return 0, errors.New("token not operator")
	}

	if op.left == nil {
		return 0, errors.New("left can't empty")
	}

	if op.right == nil {
		return op.left.Result()
	}

	opFunc, ok := tokenOperatorFunc[op.tok]
	if !ok {
		return 0, errors.New("not suppert operator")
	}

	leftVal, err := op.left.Result()
	if err != nil {
		return 0, err
	}
	rightVal, err := op.right.Result()
	if err != nil {
		return 0, err
	}

	return opFunc(leftVal, rightVal)
}

var tokenOperatorFunc = map[token.Token]func(v1, v2 float64) (float64, error){
	token.ADD: func(v1, v2 float64) (float64, error) {
		return v1 + v2, nil
	},
	token.SUB: func(v1, v2 float64) (float64, error) {
		return v1 - v2, nil
	},
	token.MUL: func(v1, v2 float64) (float64, error) {
		return v1 * v2, nil
	},
	token.QUO: func(v1, v2 float64) (float64, error) {
		if v2 == 0 {
			return 0, errors.New("0 cannot be used as a divisor")
		}
		return v1 / v2, nil
	},
}
