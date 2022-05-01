package token

import "errors"

type Float64Value struct {
	Val float64
}

func (f Float64Value) UnmarshalValue(result interface{}) error {
	floatPointer, ok := result.(*float64)
	if !ok {
		return errors.New("illegal")
	}
	*floatPointer = f.Val

	return nil
}
