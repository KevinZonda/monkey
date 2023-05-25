package token

type TokenKind string

const (
	EOF = "EOF"
	BAD = "BAD"

	IDENTITY = "IDENTITY"
	INT      = "INT"

	STRING = "STRING"
	FLOAT  = "FLOAT"

	ASSIGN   = ":="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	BANG     = "!"

	LESS_THAN    = "<"
	GREATER_THAN = ">"

	COMMA     = ","
	SEMICOLON = ";"

	L_BRACE = "{"
	R_BRACE = "}"
	L_PAREN = "("
	R_PAREN = ")"

	FUNC = "FUNC"
	LET  = "LET"

	IF     = "IF"
	ELSE   = "ELSE"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	RETURN = "RETURN"
	EQUAL  = "=="
	NOT_EQ = "!="
)
