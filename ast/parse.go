package ast

import (
	"github.com/bookrun-go/calculator/token"
)

type Parser interface {
	Parse2() (Node, error)
	// Parse(tvs []*token.TokenValue, startIndex int) (endIndex int, _ Node, _ error)
}

type ParserAbstract struct {
	tvs []*token.TokenValue

	startIndex int
	maxIndex   int
}

func NewParser(tvs []*token.TokenValue) (Parser, error) {
	op, err := ParseResgister.GetParser2(tvs[0].Tok, &ParserAbstract{tvs: tvs, startIndex: 0, maxIndex: len(tvs) - 1})
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

// type CommanParser struct {
// }

// func (CommanParser) Parse(tvs []*token.TokenValue, startIndex int) (endIndex int, _ Node, _ error) {

// }

// func (CommanParser) LeftNode(tvs []*token.TokenValue, startIndex int) (_ int, _ Node, _ error) {
// 	maxIndex := len(tvs) - 1
// 	if startIndex > maxIndex {
// 		return
// 	}

// 	firstTv := tvs[startIndex]

// 	if firstTv.Tok.IsSeparator() {

// 	}

// }
