package ast

import "github.com/KevinZonda/monkey/token"

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (r ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r ReturnStatement) statementNode() {}

var _ Statement = (*ReturnStatement)(nil)
