package example

import (
	"fmt"
	"testing"

	"github.com/bookrun-go/calculator/ast"
	"github.com/bookrun-go/calculator/scanner"
)

func TestParse1(t *testing.T) {
	scanner := scanner.NewScanner("2-2*3+10", scanner.WithAddScanners(scanner.NumberScanner{}, scanner.SeparatorScanner{}, scanner.OperatorScanner{}))

	tk, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

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
}
