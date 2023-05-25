package main

import (
	"errors"
	"fmt"
	"github.com/KevinZonda/monkey/lexer"
	"github.com/KevinZonda/monkey/token"
	"github.com/chzyer/readline"
)

func main() {
	for {
		ln, err := readline.Line(">> ")
		if err != nil {
			if errors.Is(err, readline.ErrInterrupt) {
				break
			}
			panic(err)
		}
		lex := lexer.New(ln)
		for tok := lex.NextToken(); tok.Kind != token.EOF; tok = lex.NextToken() {
			fmt.Println(tok.Literal, "::", tok.Kind)
		}
	}
}
