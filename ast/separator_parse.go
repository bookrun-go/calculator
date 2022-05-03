package ast

// import (
// 	"errors"

// 	"github.com/bookrun-go/calculator/token"
// )

// type SeparatorParse struct {
// 	leftToken  token.Token
// 	rightToken token.Token
// }

// func (sep SeparatorParse) Parse(tvs []*token.TokenValue, startIndex int) (endIndex int, _ Node, _ error) {
// 	startIndex++

// 	opp := new(OperatorParse)

// 	endIndex, node, err := opp.Parse(tvs, startIndex)
// 	if err != nil {
// 		return 0, nil, err
// 	}

// 	if tvs[endIndex].Tok != sep.rightToken {
// 		return endIndex, nil, errors.New("unusual end")
// 	}

// 	return endIndex, node, nil
// }
