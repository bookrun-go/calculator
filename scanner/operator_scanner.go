package scanner

import "github.com/bookrun-go/calculator/token"

var operatorMap = map[rune]*token.TokenValue{
	'+': {Tok: token.ADD, Value: token.F64Value{}},
	'-': {Tok: token.SUB, Value: token.F64Value{}},
	'*': {Tok: token.MUL, Value: token.F64Value{}},
	'/': {Tok: token.QUO, Value: token.F64Value{}},
}

type OperatorScanner struct {
}

func (OperatorScanner) Scan(formula []rune, startPos int) (*token.TokenValue, int) {
	tv, ok := operatorMap[formula[startPos]]
	if !ok {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: token.F64Value{},
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
