package token

type TokenType int64

const (
	Plus TokenType = iota
	Subtract
	EndCodeBlock
	Start
	EOF
	Int
)