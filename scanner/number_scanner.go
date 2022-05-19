package scanner

import (
	"strconv"

	"github.com/bookrun-go/calculator/token"
)

var _ IScanner = NumberScanner{}

const (
	NumberZero  = '0'
	NumberNine  = '9'
	NumberPoint = '.'
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
	var pointExist bool
	for ; startPos < len(formula); startPos++ {
		if ns.isPoint(formula[startPos]) { // 检查小数点是否重复
			if pointExist {
				break
			} else {
				pointExist = true
			}
		}
		legal := ns.of(formula[startPos])
		if !legal {
			break
		}
		runeVal = append(runeVal, formula[startPos])
	}
	if ns.isPoint(runeVal[len(runeVal)-1]) {
		startPos--
		runeVal = runeVal[:len(runeVal)-1]
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

func (NumberScanner) isPoint(char rune) bool {
	return char == NumberPoint
}

func (NumberScanner) of(char rune) bool {
	if char == NumberPoint {
		return true
	}

	return char >= NumberZero && char <= NumberNine
}

func (NumberScanner) Of(char rune) bool {
	return char >= NumberZero && char <= NumberNine
}
