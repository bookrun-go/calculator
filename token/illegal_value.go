package token

import "errors"

var IllegalValue = &illegalValue{}

type illegalValue struct {
	val string
}

func (illegal illegalValue) UnmarshalValue(result interface{}) error {
	strPointer, ok := result.(*string)
	if !ok {
		return errors.New("illegal")
	}
	*strPointer = illegal.val

	return nil
}
