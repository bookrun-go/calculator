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
			Value: token.F64Value{},
		}, startPos
	}

	var runeVal []rune
	for ; startPos < len(formula); startPos++ {
		legal := ns.Of(formula[startPos])
		if !legal {
			break
		}
		runeVal = append(runeVal, formula[startPos])
	}

	val, _ := strconv.ParseFloat(string(runeVal), 64)

	return &token.TokenValue{
		Tok:   token.Number,
		Value: token.F64Value{Val: val},
	}, startPos
}

func (NumberScanner) Precedence() int {
	return 1
}

func (NumberScanner) Of(char rune) bool {
	return char >= NumberZero && char <= NumberNine
}
