package ast

type NumberNode struct {
	val float64
}

func (num NumberNode) Result() (float64, error) {
	return num.val, nil
}
