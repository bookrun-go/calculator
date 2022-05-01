package token

type EmptyValue struct {
}

func (EmptyValue) UnmarshalValue(result interface{}) error {
	return nil
}
