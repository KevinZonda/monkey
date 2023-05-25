package token

type TokenKind string

const (
	EOF = "EOF"
	BAD = "BAD"

	IDENTITY = "IDENTITY"
	INT      = "INT"

	STRING = "STRING"
	FLOAT  = "FLOAT"

	ASSIGN = ":="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	L_BRACE = "{"
	R_BRACE = "}"
	L_PAREN = "("
	R_PAREN = ")"

	FUNC = "FUNC"
	LET  = "LET"
)
