package ast

// import (
// 	"fmt"

// 	"github.com/bookrun-go/calculator/token"
// )

// type OperatorParse struct {
// }

// /*
// 保证第一个是数字

// 判断后面是否为这种表达式：'5+6+'，如果是则判断符号优先级。若不是，则逐个处理，遇到左括号处理一下，右括号直接返回。

// */
// func (op OperatorParse) Parse(tvs []*token.TokenValue, startIndex int) (endIndex int, _ Node, _ error) {
// 	maxIndex := len(tvs) - 1
// 	if startIndex > maxIndex {
// 		return startIndex, nil, nil
// 	}

// 	if maxIndex-startIndex >= 3 && tvs[startIndex].Tok.IsLiteral() && tvs[startIndex+1].Tok.IsOperator() && tvs[startIndex+2].Tok.IsLiteral() && tvs[startIndex+3].Tok.IsOperator() {
// 		if tvs[startIndex+1].Tok.Precedence() > tvs[startIndex+3].Tok.Precedence() {
// 			leftNode := OperatorNode{
// 				left:  NumberNode{val: tvs[startIndex].Value},
// 				right: NumberNode{val: tvs[startIndex+2].Value},
// 			}

// 			root := OperatorNode{}
// 			root.left = leftNode

// 			root.tok = tvs[startIndex+3].Tok

// 			startIndex = startIndex + 3 + 1
// 			if maxIndex < startIndex {
// 				return startIndex, nil, fmt.Errorf("express not support,index[%d]", startIndex)
// 			}

// 			p, err := ParseResgister.GetParser(tvs[startIndex].Tok)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}

// 			endIndex, rightNode, err := p.Parse(tvs, startIndex)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}
// 			root.right = rightNode

// 			return endIndex, root, nil
// 		} else {
// 			root := OperatorNode{}
// 			root.left = NumberNode{val: tvs[startIndex].Value}
// 			root.tok = tvs[startIndex+1].Tok

// 			startIndex += 2
// 			p, err := ParseResgister.GetParser(tvs[startIndex].Tok)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}

// 			endIndex, rightNode, err := p.Parse(tvs, startIndex)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}
// 			root.right = rightNode

// 			startIndex = endIndex
// 			return startIndex, root, nil
// 		}
// 	}

// 	root := OperatorNode{}
// 	for startIndex <= maxIndex {
// 		tv := tvs[startIndex]

// 		if tv.Tok.IsLeft() {
// 			p, err := ParseResgister.GetParser(tv.Tok)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}

// 			endIndex, node, err := p.Parse(tvs, startIndex)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}

// 			err = root.AddNode(node)
// 			if err != nil {
// 				return startIndex, nil, err
// 			}

// 			startIndex = endIndex + 1
// 			continue
// 		}

// 		if tv.Tok.IsRigh() {
// 			return startIndex, root, nil
// 		}

// 		if tv.Tok.IsLiteral() {
// 			node := NumberNode{val: tv.Value}
// 			root.AddNode(node)

// 			startIndex++
// 			continue
// 		}

// 		if tv.Tok.IsOperator() {
// 			if !root.tok.IsIllegal() || root.left == nil || root.right != nil {
// 				return startIndex, nil, fmt.Errorf("express not support,index[%d]", startIndex)
// 			}

// 			root.tok = tv.Tok
// 			startIndex++
// 			continue
// 		}

// 		return startIndex, nil, fmt.Errorf("express not support,index[%d]", startIndex)
// 	}

// 	return startIndex, root, nil
// }
