package scanner

import "github.com/bookrun-go/calculator/token"

var _ IScanner = SeparatorScanner{}

var separatorMap = map[rune]*token.TokenValue{
	'(': {Tok: token.LeftParentheses, Value: token.EmptyValue},
	')': {Tok: token.RightParentheses, Value: token.EmptyValue},
	'[': {Tok: token.LeftBracket, Value: token.EmptyValue},
	']': {Tok: token.LeftBracket, Value: token.EmptyValue},
}

type SeparatorScanner struct {
}

func (SeparatorScanner) Scan(formula []rune, startPos int) (*token.TokenValue, int) {
	tv, ok := separatorMap[formula[startPos]]
	if !ok {
		return &token.TokenValue{
			Tok:   token.Illegal,
			Value: token.IllegalValue,
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
