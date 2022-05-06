package token

import (
	"errors"
	"math"
)

var ErrorZeroAsDivisor = errors.New("0 cannot be used as a divisor")

type Value interface {
	UnmarshalValue(interface{}) error
	Operate(v2 Value, opTok Token) (Value, error)
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

func (nv F64Value) Operate(v1 Value, opTok Token) (Value, error) {
	v1Val := math.NaN()

	err := v1.UnmarshalValue(&v1Val)
	if err != nil {
		return nil, err
	}

	op, ok := DefaultTokenOperatorFunc[opTok]
	if !ok {
		return nil, errors.New("not suppert operator")
	}
	return op(v1Val, nv.Val)
}

var DefaultTokenOperatorFunc = map[Token]func(v1, v2 float64) (Value, error){
	ADD: func(v1, v2 float64) (Value, error) {
		return F64Value{v1 + v2}, nil
	},
	SUB: func(v1, v2 float64) (Value, error) {
		return F64Value{v1 - v2}, nil
	},
	MUL: func(v1, v2 float64) (Value, error) {
		return F64Value{v1 * v2}, nil
	},
	QUO: func(v1, v2 float64) (Value, error) {
		if v2 == 0 {
			return EmptyValue{}, ErrorZeroAsDivisor
		}
		return F64Value{v1 / v2}, nil
	},
}

type EmptyValue struct {
}

func (EmptyValue) UnmarshalValue(data interface{}) error { return nil }

func (EmptyValue) Operate(v2 Value, opTok Token) (Value, error) { return v2, nil }
