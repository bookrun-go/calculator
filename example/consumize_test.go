package example

import (
	"fmt"
	"testing"

	"github.com/bookrun-go/calculator/ast"
	"github.com/bookrun-go/calculator/example/consumize"
	"github.com/bookrun-go/calculator/scanner"
	"github.com/bookrun-go/calculator/token"
)

func TestConsumized(t *testing.T) {
	scanner := scanner.NewScanner("A2(5)", scanner.WithAddScanners(
		scanner.NumberScanner{}, scanner.SeparatorScanner{}, scanner.OperatorScanner{},
		consumize.MyScanner{}))

	tk, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	ast.ParseResgister.Registe(token.NumberReserve, func(pa *ast.ParserAbstract) ast.Parser {
		return consumize.NewParser(pa, consumize.NewNode)
	})

	op, err := ast.NewParser(tk, 0)
	if err != nil {
		panic(err)
	}

	err = op.Doing()
	if err != nil {
		panic(err)
	}

	node := op.Root()

	fmt.Println(node.Result())

	consumize.UpdateCastMap('A', 2)

	fmt.Println(node.Result())

}
