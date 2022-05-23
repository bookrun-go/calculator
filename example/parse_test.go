package example

import (
	"fmt"
	"testing"

	"github.com/bookrun-go/calculator/ast"
	"github.com/bookrun-go/calculator/scanner"
)

//"2-2*3+10"
func TestParse1(t *testing.T) {
	node, err := getNode("2-2*3+10")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	fmt.Println(node.Result())
}

func TestParse2(t *testing.T) {
	node, err := getNode("2-2*(3+10)+24*2")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(node.Result())
}

func TestParse3(t *testing.T) {
	node, err := getNode("(2-1)*(3+10)+24*2")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(node.Result())
}

func getNode(str string) (ast.Node, error) {
	scanner := scanner.NewScanner(str, scanner.WithAddScanners(scanner.NumberScanner{}, scanner.SeparatorScanner{}, scanner.OperatorScanner{}))

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

	return node, nil
}
