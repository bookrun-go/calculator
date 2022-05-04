package costumize

import "github.com/bookrun-go/calculator/ast"

var castMap = &map[rune]float64{'A': 1}

func GetCharNode(char rune) ast.Node {
	return Node{
		char:    char,
		castMap: castMap,
	}
}

func UpdateCastMap(key rune, val float64) {
	temp := *castMap
	temp[key] = val

	*castMap = temp
}
