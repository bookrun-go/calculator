package ast

import (
	"errors"

	"github.com/bookrun-go/calculator/token"
)

type Parser interface {
	Doing() error
	Root() Node
	GenNode() error
}

type ParserAbstract struct {
	tvs  []*token.TokenValue
	root Node // recovery

	startIndex int
	maxIndex   int
	curOpNode  *OperatorNode // recovery
	endTok     token.Token   // recovery
}

func (np *ParserAbstract) AddNumberNode(node Node) error {
	if np.IsLastOne() {
		np.Step()
		np.AddLastNode(node)
		return nil
	}

	np.Step()
	curTv := np.CurTv()
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

	err := np.AddNode(node, tok)
	if err != nil {
		return err
	}

	np.Step()
	return nil
}
func (pa *ParserAbstract) AddLastNode(node Node) error {
	if pa.curOpNode.left == nil {
		pa.curOpNode.left = node
	} else {
		pa.curOpNode.right = node
	}

	return nil
}

func (pa *ParserAbstract) AddNode(node Node, nextOpToken token.Token) error {
	if node == nil {
		return nil
	}

	if pa.curOpNode.left == nil {
		pa.curOpNode.left = node
		pa.curOpNode.tok = nextOpToken
		return nil
	}

	if pa.curOpNode.tok.Precedence() >= nextOpToken.Precedence() {
		newLeft := &OperatorNode{
			left:  pa.curOpNode.left,
			right: node,
			tok:   pa.curOpNode.tok,
		}

		pa.curOpNode.left = newLeft
		pa.curOpNode.tok = nextOpToken
		return nil
	} else {
		newRigth := &OperatorNode{
			left: node,
			tok:  nextOpToken,
		}

		pa.curOpNode.right = newRigth

		pa.curOpNode = newRigth

		return nil
	}
}

type RecoverFunc func()

func (pa *ParserAbstract) NewChildParser(childEndTok token.Token) (Parser, func(), error) {
	curRoot := pa.root
	curOpNode := pa.curOpNode
	endTok := pa.endTok

	pa.endTok = childEndTok
	pa.curOpNode = &OperatorNode{}
	pa.root = pa.curOpNode
	op, err := ParseResgister.GetParser(pa.tvs[pa.startIndex].Tok, pa)
	if err != nil {
		return nil, nil, err
	}

	return op, func() {
		pa.curOpNode = curOpNode
		pa.root = curRoot
		pa.endTok = endTok
	}, nil
}

func (pa *ParserAbstract) Doing() error {
	preIndex := pa.startIndex - 1
	for pa.startIndex <= pa.maxIndex {
		if preIndex >= pa.startIndex {
			return errors.New("loop die")
		}
		preIndex = pa.startIndex

		if pa.tvs[pa.startIndex].Tok == pa.endTok {
			break
		}

		p, err := ParseResgister.GetParser(pa.tvs[pa.startIndex].Tok, pa)
		if err != nil {
			return err
		}

		err = p.GenNode()
		if err != nil {
			return err
		}
	}

	return nil
}

func (pa *ParserAbstract) Root() Node {
	return pa.root
}

func (pa *ParserAbstract) IsLastOne() bool {
	return pa.startIndex == pa.maxIndex
}

func (pa *ParserAbstract) Step() {
	pa.startIndex++
}

func (pa *ParserAbstract) Back() {
	pa.startIndex--
}

func (pa *ParserAbstract) CurTv() *token.TokenValue {
	return pa.tvs[pa.startIndex]
}

func (pa *ParserAbstract) IsEndTok(tok token.Token) bool {
	return pa.endTok == tok
}

func NewParser(tvs []*token.TokenValue, startIndex int) (Parser, error) {
	root := &OperatorNode{}
	pa := &ParserAbstract{tvs: tvs, startIndex: startIndex, maxIndex: len(tvs) - 1}
	pa.root = root
	pa.curOpNode = root

	op, err := ParseResgister.GetParser(tvs[0].Tok, pa)
	if err != nil {
		return nil, err
	}

	return op, err
}

/*
先找出左节点和右节点

第一个遇见左括号，则括号内到为左节点

第一个遇见的是number，则此number为左节点

第三个如果遇见左括号，则先把括号内当一个节点，算出后，需要对比两个节点的操作符合优先级，如果第一个操作符优先，则把两个节点合并下沉；如果第二个操作符优先，则保持原来第一个节点为左节点，生成新的操作节点，把右节点放到新操作节点左节点。

*/
