package ast

type NumberParser struct {
	*ParserAbstract
}

// 4+5+ or 4+(....)+

func (np *NumberParser) Parse2() (Node, error) {
	if np.maxIndex < np.startIndex {
		return nil, nil
	}

	if np.startIndex == np.maxIndex {
		return NumberNode{val: np.tvs[np.startIndex].Value}, nil
	}

	if np.maxIndex-np.startIndex < 2 {
		return nil, FomulaFormatError
	}

	if !np.tvs[np.startIndex+1].Tok.IsOperator() {
		return nil, FomulaFormatError
	}

	root := OperatorNode{
		tok:  np.tvs[np.startIndex+1].Tok,
		left: NumberNode{val: np.tvs[np.startIndex].Value},
	}

	np.startIndex += 2
	if np.tvs[np.startIndex].Tok.IsSeparator() {
		p, err := ParseResgister.GetParser2(np.tvs[np.startIndex].Tok, np.ParserAbstract)
		if err != nil {
			return nil, err
		}

		node, err := p.Parse2()
		if err != nil {
			return nil, err
		}

		root.right = node
	} else if np.tvs[np.startIndex].Tok.IsLiteral() {
		root.right = NumberNode{val: np.tvs[np.startIndex].Value}
	} else {
		return nil, FomulaFormatError
	}

	if np.maxIndex < np.startIndex+1 {
		return root, nil // 结束
	}

	if np.tvs[np.startIndex+1].Tok.IsSeparator() {
		np.startIndex++
		return root, nil
	}

	if !np.tvs[np.startIndex+1].Tok.IsOperator() {
		return nil, FomulaFormatError
	}

	if root.tok.Precedence() > np.tvs[np.startIndex+1].Tok.Precedence() {
		newRoot := OperatorNode{
			tok:  np.tvs[np.startIndex+1].Tok,
			left: root,
		}
		root = newRoot

		if np.startIndex+2 > np.maxIndex { // 结束
			return root, nil
		}

		np.startIndex += 2
		p, err := ParseResgister.GetParser2(np.tvs[np.startIndex].Tok, np.ParserAbstract)
		if err != nil {
			return nil, err
		}

		node, err := p.Parse2()
		if err != nil {
			return nil, err
		}
		root.right = node

	} else {
		newRight := OperatorNode{
			tok:  np.tvs[np.startIndex+1].Tok,
			left: root.right,
		}

		if np.startIndex+2 > np.maxIndex { // 结束
			return root, nil
		}

		np.startIndex += 2
		p, err := ParseResgister.GetParser2(np.tvs[np.startIndex].Tok, np.ParserAbstract)
		if err != nil {
			return nil, err
		}

		node, err := p.Parse2()
		if err != nil {
			return nil, err
		}

		newRight.right = node
		root.right = newRight
	}

	return root, nil
}

// func (NumberParser) Parse(tvs []*token.TokenValue, startIndex int) (endIndex int, _ Node, _ error) {
// 	tv := tvs[startIndex]

// 	maxIndex := len(tvs) - 1

// 	if startIndex == maxIndex {
// 		return startIndex, NumberNode{val: tv.Value}, nil
// 	}

// 	if maxIndex-startIndex < 2 {
// 		return // todo err
// 	}

// 	if !tvs[startIndex+1].Tok.IsOperator() {
// 		return // todo err
// 	}

// 	if tvs[startIndex+2].Tok.IsOperator() {
// 		return // todo err
// 	}

// 	root := OperatorNode{}

// 	if tvs[startIndex+2].Tok.IsOperator() {
// 		p, err := ParseResgister.GetParser(tvs[startIndex+2].Tok)
// 		if err != nil {
// 			return startIndex, nil, err
// 		}

// 		endIndex, node, err := p.Parse(tvs, startIndex+2)
// 		if err != nil {
// 			return startIndex, nil, err
// 		}

// 	}

// }
