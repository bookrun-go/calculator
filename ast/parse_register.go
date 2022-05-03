package ast

import (
	"errors"

	"github.com/bookrun-go/calculator/token"
)

var ParseResgister = parseResgister{}

type parseResgister struct {
}

// func (parseResgister) GetParser(tok token.Token) (Parser, error) {
// 	if tok.IsOperator() {
// 		return nil, errors.New("operator not parser")
// 	}

// 	if tok == token.Number {
// 		NumberParser{}
// 		return new(OperatorParse), nil
// 	}

// 	if tok == token.LeftParentheses {
// 		return SeparatorParse{leftToken: tok, rightToken: token.RightParentheses}, nil
// 	}

// 	return nil, errors.New("operator not found parser")
// }

func (parseResgister) GetParser2(tok token.Token, pa *ParserAbstract) (Parser, error) {
	if tok.IsOperator() {
		return nil, errors.New("operator not parser")
	}

	if tok == token.Number {
		return &NumberParser{ParserAbstract: pa}, nil
	}

	if tok == token.LeftParentheses {
		return &ParenthesesParser{ParserAbstract: pa}, nil
	}

	return nil, errors.New("operator not found parser")
}
