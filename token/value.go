package token

type Value interface {
	UnmarshalValue(interface{}) error
}

type TokenValue struct {
	Tok   Token
	Value float64
}
