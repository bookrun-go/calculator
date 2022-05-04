package costumize

import "errors"

type MyValue struct {
	str rune
}

func (mv MyValue) UnmarshalValue(data interface{}) error {
	val, ok := data.(*rune)
	if !ok {
		return errors.New("type err")
	}

	*val = mv.str
	return nil
}
