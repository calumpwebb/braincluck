package ast

import (
	"braincluck/pkgs/token"
	"fmt"
)

type Execution struct {
	Token token.Token
}

func NewExecution(tok token.Token) *Execution {
	return &Execution{Token: tok}
}

func (e *Execution) String() string {
	return fmt.Sprintf("<Execution: Token={Type:\"%s\"}>", e.Token.Type)
}
