package ast

import (
	"fmt"
	"strings"
)

// Program: Program node is the root node of every AST
type Program struct {
	Commands []Command
}

func NewProgram() *Program {
	return &Program{Commands: []Command{}}
}

func (p *Program) String() string {
	var commandStrings []string
	for _, s := range p.Commands {
		commandStrings = append(commandStrings, s.String())
	}

	if commandStrings == nil {
		commandStrings = []string{}
	}

	return fmt.Sprintf("<Program: Commands=[%s]>", strings.Join(commandStrings, ","))
}
