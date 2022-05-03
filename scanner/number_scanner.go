package scanner

import (
	"strconv"

	"github.com/bookrun-go/calculator/token"
)

var _ IScanner = NumberScanner{}

const (
	NumberZero = '0'
	NumberNine = '9'
)

type NumberScanner struct {
}

func (ns NumberScanner) Scan(formula []rune, startPos int) (_ *token.TokenValue, _ int) {
	startLegal := ns.Of(formula[startPos]) // 检查第一个是否合法
	if !startLegal {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: floatNaN,
		}, startPos
	}

	float64Value := &token.Float64Value{}

	var runeVal []rune
	for ; startPos < len(formula); startPos++ {
		legal := ns.Of(formula[startPos])
		if !legal {
			break
		}
		runeVal = append(runeVal, formula[startPos])
	}

	float64Value.Val, _ = strconv.ParseFloat(string(runeVal), 64)

	return &token.TokenValue{
		Tok:   token.Number,
		Value: float64Value.Val,
	}, startPos
}

func (NumberScanner) Precedence() int {
	return 1
}

func (NumberScanner) Of(char rune) bool {
	return char >= NumberZero && char <= NumberNine
}
