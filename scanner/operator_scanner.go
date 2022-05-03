package scanner

import "github.com/bookrun-go/calculator/token"

var operatorMap = map[rune]*token.TokenValue{
	'+': {Tok: token.ADD, Value: token.EmptyValue},
	'-': {Tok: token.SUB, Value: token.EmptyValue},
	'*': {Tok: token.MUL, Value: token.EmptyValue},
	'/': {Tok: token.QUO, Value: token.EmptyValue},
}

type OperatorScanner struct {
}

func (OperatorScanner) Scan(formula []rune, startPos int) (*token.TokenValue, int) {
	tv, ok := operatorMap[formula[startPos]]
	if !ok {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: token.IllegalValue,
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
