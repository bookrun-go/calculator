package costumize

import (
	"github.com/bookrun-go/calculator/ast"
	"github.com/bookrun-go/calculator/token"
)

func NewParser(pa *ast.ParserAbstract, newNodeFunc func(char rune) ast.Node) ast.Parser {
	return &Parser{
		ParserAbstract: pa,
		newNodeFunc:    newNodeFunc,
	}
}

type Parser struct {
	*ast.ParserAbstract
	newNodeFunc func(char rune) ast.Node
}

func (np Parser) GenNode() error {
	curTv := np.CurTv()
	val := rune(0)
	err := curTv.Value.UnmarshalValue(&val)
	if err != nil {
		return err
	}

	node := np.newNodeFunc(rune(val))
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
	if curTv.Tok.IsLeft() || curTv.Tok.IsLiteral() {
		np.Back() // 2(5+2) 这种表达式，默认为*需要回退。
		tok = token.MUL
	} else if !curTv.Tok.IsOperator() {
		return ast.ErrorFomulaFormat
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
