package ast

type RootNode struct {
	next Node
}

func (root RootNode) Result() (float64, error) {
	if root.next == nil {
		return 0, nil
	}

	return root.next.Result()
}
