package consumize

import "github.com/bookrun-go/calculator/ast"

type Node struct {
	char    rune
	castMap *map[rune]float64
}

func (n Node) Result() (float64, error) {
	res := (*n.castMap)[n.char]
	return res, nil
}

func NewNode(char rune) ast.Node {
	return &Node{char: char, castMap: castMap}
}
