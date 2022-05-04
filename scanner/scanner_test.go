package scanner

import (
	"fmt"
	"testing"
)

func TestScanner(t *testing.T) {
	scanner := NewScanner("(100*487)+5434 / 45 * 90 / 0", WithAddScanners(NumberScanner{}, SeparatorScanner{}, OperatorScanner{}))

	tk, err := scanner.Scan()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", tk)
}
