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

func (np *NumberParser) AddSymbolNumberNode(tok token.Token) error {
	if np.curOpNode.left != nil {
		return ErrorFomulaFormat
	}

	curTv := np.CurTv()

	if tok == token.ADD {
		node := &NumberNode{Val: curTv.Value}
		return np.AddNumberNode(node)
	}

	opNode := &OperatorNode{
		left:  &NumberNode{Val: &token.F64Value{Val: 0}},
		tok:   tok,
		right: &NumberNode{Val: curTv.Value},
	}

	return np.AddNumberNode(opNode)
}
