package costumize

import (
	"errors"
	"math"

	"github.com/bookrun-go/calculator/token"
)

type MyValue struct {
	str     rune
	castMap *map[rune]float64
}

func (mv MyValue) UnmarshalValue(data interface{}) error {
	fVal, ok := data.(*float64)
	if !ok {
		return errors.New("data must float64 pointer")
	}

	*fVal = mv.getFloatVal()

	return nil
}

func (mv MyValue) getFloatVal() float64 {
	return (*mv.castMap)[mv.str]
}

func (mv MyValue) Operate(v1 token.Value, opTok token.Token) (token.Value, error) {
	op, ok := token.DefaultTokenOperatorFunc[opTok]
	if !ok {
		return nil, errors.New("not suppert operator")
	}

	v1Val := math.NaN()
	err := v1.UnmarshalValue(&v1Val)
	if err != nil {
		return nil, err
	}

	return op(v1Val, mv.getFloatVal())
}
