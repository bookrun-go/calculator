package scanner

import (
	"math"

	"github.com/bookrun-go/calculator/token"
)

var _ IScanner = SeparatorScanner{}

var separatorMap = map[rune]*token.TokenValue{
	'(': {Tok: token.LeftParentheses, Value: floatNaN},
	')': {Tok: token.RightParentheses, Value: math.NaN()},
	'[': {Tok: token.LeftBracket, Value: floatNaN},
	']': {Tok: token.LeftBracket, Value: floatNaN},
}

type SeparatorScanner struct {
}

func (SeparatorScanner) Scan(formula []rune, startPos int) (*token.TokenValue, int) {
	tv, ok := separatorMap[formula[startPos]]
	if !ok {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: floatNaN,
		}, startPos
	}

	startPos++
	return tv, startPos
}
func (SeparatorScanner) Of(char rune) bool {
	_, ok := separatorMap[char]
	return ok
}
func (SeparatorScanner) Precedence() int {
	return 2
}
