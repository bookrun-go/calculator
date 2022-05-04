package ast

type Node interface {
	Result() (float64, error)
}
