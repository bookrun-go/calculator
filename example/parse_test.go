package example

import (
	"fmt"
	"testing"

	"github.com/bookrun-go/calculator/ast"
	"github.com/bookrun-go/calculator/scanner"
)

//"2-2*3+10"
func TestParse1(t *testing.T) {
	ok, err := excuteFomula("2-2*3+10", 6)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse2(t *testing.T) {
	ok, err := excuteFomula("2-2*(3+10)+24*2", 24)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse3(t *testing.T) {
	ok, err := excuteFomula("(2-1)*(3+10)+24*2", 61)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}

}

func TestParse4(t *testing.T) {
	ok, err := excuteFomula("9+(-1*10)-100", -101)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse5(t *testing.T) {
	ok, err := excuteFomula("9+(-1*10)-100", 0)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ok {
		panic("result not expect")
	}
}

func TestParse6(t *testing.T) {
	ok, err := excuteFomula("(-100)+67", -33)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse7(t *testing.T) {
	ok, err := excuteFomula("3-2*(-10+1)+9", 30)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse8(t *testing.T) {
	ok, err := excuteFomula("3-2*(-10*2)+9", 52)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse9(t *testing.T) {
	ok, err := excuteFomula(" 3-2*(-10*(2+3))+9", 112)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse10(t *testing.T) {
	ok, err := excuteFomula(" 3-2*((-10)*(2+3))+9", 112)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func TestParse11(t *testing.T) {
	ok, err := excuteFomula("3+(-1)", 2)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ok {
		panic("result not expect")
	}
}

func excuteFomula(str string, res float64) (bool, error) {
	node, err := getNode(str)
	if err != nil {
		return false, err
	}

	v, err := node.Result()
	if err != nil {
		return false, err
	}

	var f float64

	v.UnmarshalValue(&f)

	return f == res, nil
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
		panic(fmt.Sprintf("%+v", err))
	}

	node := op.Root()

	return node, nil
}
