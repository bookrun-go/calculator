package ast

import "github.com/bookrun-go/calculator/token"

type NumberParser struct {
	*ParserAbstract
}

// 4+5+ or 4+(....)+
func (np *NumberParser) GenNode() error {
	curTv := np.CurTv()

	if curTv.Tok == token.ADD || curTv.Tok == token.SUB {
		np.Step()
		return np.AddSymbolNumberNode(curTv.Tok)
	}

	node := &NumberNode{Val: curTv.Value}
	return np.AddNumberNode(node)
}
