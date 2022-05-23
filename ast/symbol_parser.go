package ast

import "github.com/bookrun-go/calculator/token"

type SymbolParser struct {
	*ParserAbstract
}

func (sp *SymbolParser) GenNode() error {
	curTv := sp.CurTv()
	sp.Step()

	if curTv.Tok == token.ADD {
		return nil
	}

	if curTv.Tok != token.SUB {
		return ErrorFomulaFormat
	}

	if sp.curOpNode.left != nil {
		return ErrorFomulaFormat
	}

	sp.curOpNode.left = &NumberNode{Val: &token.F64Value{Val: -1}}
	sp.curOpNode.tok = token.MUL

	return nil
}
