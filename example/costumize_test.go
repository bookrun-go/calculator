package example

import (
	"fmt"
	"math"
	"testing"

	"github.com/bookrun-go/calculator/ast"
	"github.com/bookrun-go/calculator/example/costumize"
	"github.com/bookrun-go/calculator/scanner"
	"github.com/bookrun-go/calculator/token"
)

func TestConsumized(t *testing.T) {
	scanner := scanner.NewScanner("A2(5)", scanner.WithAddScanners(
		scanner.NumberScanner{}, scanner.SeparatorScanner{}, scanner.OperatorScanner{},
		costumize.MyScanner{CastMap: costumize.CastMap}))

	tk, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	ast.ParseResgister.Registe(token.NumberReserve, func(pa *ast.ParserAbstract) ast.Parser {
		return costumize.NewParser(pa)
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

	costumize.UpdateCastMap('A', 2)

	fmt.Println(node.Result())

}

func TestNan(t *testing.T) {
	f := math.NaN()

	ss := f / 0

	fmt.Println(math.IsNaN(ss))
}
