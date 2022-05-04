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
	LeftBracket
	separator_mid
	RightParentheses
	RighBracket
	separator_end

	operator_beg
	ADD             // +
	SUB             // -
	MUL             // *
	QUO             // /
	OperatorReserve // 保留操作符
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

func (tok Token) IsLeft() bool {
	return tok < separator_mid && tok > separator_beg
}

func (tok Token) IsRigh() bool {
	return tok > separator_mid && tok < separator_end
}

func (tok Token) IsOperator() bool {
	return tok < operator_end && tok > operator_beg
}

func (tok Token) Precedence() int {
	switch tok {
	case ADD, SUB:
		{
			return 1
		}
	case MUL, QUO:
		{
			return 2
		}
	default:
		{
			return 3
		}
	}
}
