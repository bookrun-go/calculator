package ast

import (
	"errors"

	"github.com/bookrun-go/calculator/token"
)

var ParseResgister = parseResgister{}

type NewParserFunc func(*ParserAbstract) Parser

type parseResgister struct {
	parseMap map[token.Token]NewParserFunc
}

func (pr *parseResgister) GetParser(tok token.Token, pa *ParserAbstract) (Parser, error) {
	if tok == token.Number {
		return &NumberParser{ParserAbstract: pa}, nil
	}

	if tok == token.ADD || tok == token.SUB {
		return &SymbolParser{ParserAbstract: pa}, nil
	}

	if tok == token.LeftParentheses {
		return &ParenthesesParser{ParserAbstract: pa}, nil
	}

	if pr.parseMap == nil {
		return nil, errors.New("operator not found parser")
	}

	f, ok := pr.parseMap[tok]
	if !ok {
		return nil, errors.New("operator not found parser")
	}

	return f(pa), nil
}

func (pr *parseResgister) Registe(tok token.Token, f NewParserFunc) {
	if pr.parseMap == nil {
		pr.parseMap = make(map[token.Token]NewParserFunc)
	}

	pr.parseMap[tok] = f
}
