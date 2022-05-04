package token

import "errors"

type Value interface {
	UnmarshalValue(interface{}) error
}

type TokenValue struct {
	Tok   Token
	Value Value
}

type F64Value struct {
	Val float64
}

func (nv F64Value) UnmarshalValue(data interface{}) error {
	val, ok := data.(*float64)
	if !ok {
		return errors.New("unmarshal value error")
	}

	*val = nv.Val
	return nil
}

type EmptyValue struct {
}

func (EmptyValue) UnmarshalValue(data interface{}) error { return nil }
