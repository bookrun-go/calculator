package ast

import "github.com/bookrun-go/calculator/token"

type NumberParser struct {
	*ParserAbstract
}

// 4+5+ or 4+(....)+
func (np *NumberParser) GenNode() error {
	curTv := np.CurTv()
	val := float64(0)
	err := curTv.Value.UnmarshalValue(&val)
	if err != nil {
		return err
	}

	node := &NumberNode{val: val}
	if np.IsLastOne() {
		np.Step()
		np.AddLastNode(node)
		return nil
	}

	np.Step()
	curTv = np.CurTv()
	if np.IsEndTok(curTv.Tok) {
		np.AddLastNode(node)
		return nil
	}

	tok := token.Illegal
	if curTv.Tok.IsLeft() || curTv.Tok == token.NumberReserve {
		np.Back() // 2(5+2) 这种表达式，默认为*需要回退。
		tok = token.MUL
	} else if !curTv.Tok.IsOperator() {
		return ErrorFomulaFormat
	} else {
		tok = curTv.Tok
	}

	err = np.AddNode(node, tok)
	if err != nil {
		return err
	}

	np.Step()
	return nil
}
