package token

type Token int

const (
	Illegal Token = iota
	illegal_end

	literal_beg
	Number
	NumberReserve
	literal_end

	separator_beg
	LeftParentheses
	RightParentheses
	LeftBracket
	RighBracket
	separator_end

	operator_beg
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	operator_end
)

const (
	Reserve Token = 10 << 10
)

func (tok Token) IsIllegal() bool {
	return tok < illegal_end
}

func (tok Token) IsLiteral() bool {
	return tok < literal_end && tok > literal_beg
}

func (tok Token) IsSeparator() bool {
	return tok < separator_end && tok > separator_beg
}

func (tok Token) IsOperator() bool {
	return tok < operator_end && tok > operator_beg
}
