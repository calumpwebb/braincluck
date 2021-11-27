package environment

import (
	"braincluck/pkgs/storage"
	"bufio"
)

type Environment struct {
	Store   storage.Storage
	Scanner *bufio.Scanner
}

func NewEnvironment(store storage.Storage, scanner *bufio.Scanner) *Environment {
	return &Environment{
		Store:   store,
		Scanner: scanner,
	}
}
