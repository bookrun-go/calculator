package ast

import "github.com/bookrun-go/calculator/token"

type ParenthesesParser struct {
	*ParserAbstract
}

func (pp *ParenthesesParser) Parse2() (Node, error) {
	pp.startIndex++

	if pp.maxIndex < pp.startIndex {
		return nil, nil
	}

	root := OperatorNode{}

	p, err := ParseResgister.GetParser2(pp.tvs[pp.startIndex].Tok, pp.ParserAbstract)
	if err != nil {
		return nil, err
	}

	node, err := p.Parse2()
	if err != nil {
		return nil, err
	}

	root.left = node

	if pp.startIndex > pp.maxIndex { // 结束
		return root, nil
	}

	if pp.startIndex+2 > pp.maxIndex {
		return nil, FomulaFormatError
	}

	if !pp.tvs[pp.startIndex+1].Tok.IsOperator() {
		return nil, FomulaFormatError
	}
	root.tok = pp.tvs[pp.startIndex+1].Tok

	pp.startIndex += 2
	if pp.tvs[pp.startIndex].Tok.IsSeparator() {
		p, err := ParseResgister.GetParser2(pp.tvs[pp.startIndex].Tok, pp.ParserAbstract)
		if err != nil {
			return nil, err
		}

		node, err := p.Parse2()
		if err != nil {
			return nil, err
		}

		root.right = node
	} else if pp.tvs[pp.startIndex].Tok.IsLiteral() {
		root.right = NumberNode{val: pp.tvs[pp.startIndex].Value}
	} else {
		return nil, FomulaFormatError
	}

	if pp.maxIndex < pp.startIndex+1 {
		return root, nil // 结束
	}

	if pp.tvs[pp.startIndex+1].Tok.IsSeparator() {
		pp.startIndex++
		return root, nil
	}

	if !pp.tvs[pp.startIndex+1].Tok.IsOperator() {
		return nil, FomulaFormatError
	}

	if root.tok.Precedence() > pp.tvs[pp.startIndex+1].Tok.Precedence() {
		newRoot := OperatorNode{
			tok:  pp.tvs[pp.startIndex+1].Tok,
			left: root,
		}
		root = newRoot

		if pp.startIndex+2 > pp.maxIndex { // 结束
			return root, nil
		}

		pp.startIndex += 2
		p, err := ParseResgister.GetParser2(pp.tvs[pp.startIndex].Tok, pp.ParserAbstract)
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
			tok:  pp.tvs[pp.startIndex+1].Tok,
			left: root.right,
		}

		if pp.startIndex+2 > pp.maxIndex { // 结束
			return root, nil
		}

		pp.startIndex += 2
		p, err := ParseResgister.GetParser2(pp.tvs[pp.startIndex].Tok, pp.ParserAbstract)
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

func (ParenthesesParser) Parse(tvs []*token.TokenValue, startIndex int) (endIndex int, _ Node, _ error) {
	startIndex++

	maxIndex := len(tvs) - 1
	if maxIndex < startIndex {
		return startIndex, nil, nil
	}

	root := OperatorNode{}
	for tvs[startIndex].Tok != token.RightParentheses || startIndex > maxIndex {

	}

	return startIndex, root, nil
}
