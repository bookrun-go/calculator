package ast

type Node interface {
	UnmarshalValue(interface{}) error
}
