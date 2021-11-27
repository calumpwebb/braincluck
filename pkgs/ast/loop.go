package ast

import (
	"braincluck/pkgs/token"
	"fmt"
	"strings"
)

type Loop struct {
	Token    token.Token
	Commands []Command
}

func NewLoop(tok token.Token, commands []Command) *Loop {
	if commands == nil {
		commands = []Command{}
	}

	return &Loop{
		Token:    tok,
		Commands: commands,
	}
}

func (l *Loop) String() string {
	var commandStrings []string
	for _, s := range l.Commands {
		commandStrings = append(commandStrings, s.String())
	}

	if commandStrings == nil {
		commandStrings = []string{}
	}

	return fmt.Sprintf("<Loop: Token={Type:\"%s\"} Commands=[%s]>", l.Token.Type, strings.Join(commandStrings, ","))
}
