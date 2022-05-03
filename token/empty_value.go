package token

var EmptyValue = emptyValue{}

type emptyValue struct {
}

func (emptyValue) UnmarshalValue(result interface{}) error {
	return nil
}
