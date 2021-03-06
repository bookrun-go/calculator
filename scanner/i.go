package scanner

import (
	"math"

	"github.com/bookrun-go/calculator/token"
)

type IScanner interface {
	Scan(formula []rune, startPos int) (tv *token.TokenValue, nextPos int)
	Of(char rune) bool
	Precedence() int
}

var floatNaN = math.NaN()
