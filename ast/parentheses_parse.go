package ast

import "github.com/bookrun-go/calculator/token"

type ParenthesesParser struct {
	*ParserAbstract
}

func (pp *ParenthesesParser) GenNode() error {
	pp.startIndex++

	if pp.maxIndex < pp.startIndex {
		return nil
	}

	childParse, recoveryFunc, err := pp.NewChildParser(token.RightParentheses)
	if err != nil {
		return err
	}

	err = childParse.Doing()
	if err != nil {
		return err
	}
	childRoot := childParse.Root()
	recoveryFunc()

	if pp.maxIndex == pp.startIndex {
		pp.startIndex++
		return pp.AddLastNode(childRoot)
	}

	//
	pp.startIndex++
	if pp.tvs[pp.startIndex].Tok == pp.endTok {
		pp.AddLastNode(childRoot)
		return nil
	}

	if !pp.tvs[pp.startIndex].Tok.IsOperator() {
		return ErrorFomulaFormat
	}
	err = pp.AddNode(childRoot, pp.tvs[pp.startIndex].Tok)
	if err != nil {
		return err
	}

	//
	pp.startIndex++
	return nil
}
