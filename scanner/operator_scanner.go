package scanner

import "github.com/bookrun-go/calculator/token"

var operatorMap = map[rune]*token.TokenValue{
	'+': {Tok: token.ADD, Value: floatNaN},
	'-': {Tok: token.SUB, Value: floatNaN},
	'*': {Tok: token.MUL, Value: floatNaN},
	'/': {Tok: token.QUO, Value: floatNaN},
}

type OperatorScanner struct {
}

func (OperatorScanner) Scan(formula []rune, startPos int) (*token.TokenValue, int) {
	tv, ok := operatorMap[formula[startPos]]
	if !ok {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: floatNaN,
		}, startPos
	}

	startPos++
	return tv, startPos
}
func (OperatorScanner) Of(char rune) bool {
	_, ok := operatorMap[char]
	return ok
}
func (OperatorScanner) Precedence() int {
	return 3
}
