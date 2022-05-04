package ast

type NumberParser struct {
	*ParserAbstract
}

// 4+5+ or 4+(....)+

func (np *NumberParser) Parse() error {
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
	if !np.tvs[np.startIndex].Tok.IsOperator() {
		return ErrorFomulaFormat
	}
	err := np.AddNode(node, np.tvs[np.startIndex].Tok)
	if err != nil {
		return err
	}

	np.startIndex++
	return nil
}
