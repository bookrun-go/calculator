package ast

import "github.com/bookrun-go/calculator/token"

type NumberParser struct {
	*ParserAbstract
}

// 4+5+ or 4+(....)+
func (np *NumberParser) GenNode() error {
	node := &NumberNode{val: np.tvs[np.startIndex].Value}
	if np.startIndex == np.maxIndex {
		np.startIndex++
		np.AddLastNode(node)
		return nil
	}

	np.startIndex++
	if np.tvs[np.startIndex].Tok == np.endTok {
		np.AddLastNode(node)
		return nil
	}

	tok := token.Illegal
	if np.tvs[np.startIndex].Tok.IsLeft() {
		np.startIndex-- // 2(5+2) 这种表达式，默认为*需要回退。
		tok = token.MUL
	} else if !np.tvs[np.startIndex].Tok.IsOperator() {
		return ErrorFomulaFormat
	} else {
		tok = np.tvs[np.startIndex].Tok
	}

	err := np.AddNode(node, tok)
	if err != nil {
		return err
	}

	np.startIndex++
	return nil
}
