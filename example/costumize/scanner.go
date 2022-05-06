package costumize

import "github.com/bookrun-go/calculator/token"

type MyScanner struct {
	CastMap *map[rune]float64
}

func (ms MyScanner) Scan(formula []rune, startPos int) (tv *token.TokenValue, nextPos int) {
	startLegal := ms.Of(formula[startPos]) // 检查第一个是否合法
	if !startLegal {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: MyValue{},
		}, startPos
	}

	tv = &token.TokenValue{
		Tok:   token.NumberReserve,
		Value: MyValue{str: formula[startPos], castMap: ms.CastMap},
	}

	startPos++
	return tv, startPos
}
func (MyScanner) Of(char rune) bool {
	return char >= 'A' && char <= 'Z'
}
func (MyScanner) Precedence() int {
	return 1
}
