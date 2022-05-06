package ast

type NumberParser struct {
	*ParserAbstract
}

// 4+5+ or 4+(....)+
func (np *NumberParser) GenNode() error {
	curTv := np.CurTv()

	node := &NumberNode{Val: curTv.Value}
	return np.AddNumberNode(node)
}
