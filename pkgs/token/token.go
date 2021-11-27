package token

type Type string

const (
	EOF Type = "EOF" // EOF so we know when to stop running

	// operators
	MovePtrRight  = ">" // Increment the data pointer (to point to the next cell to the right).
	MovePtrLeft   = "<" // Decrement the data pointer (to point to the next cell to the left).
	IncrementCell = "+" // Increment (increase by one) the byte at the data pointer.
	DecrementCell = "-" // Decrement (decrease by one) the byte at the data pointer.
	OutputCell    = "." // Output the byte at the data pointer.
	InputCell     = "," // Accept one byte of input, storing its value in the byte at the data pointer.
	WhileStart    = "[" // If the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching WhileEnd command.
	WhileEnd      = "]" // If the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching WhileStart command.
)

var AllValidTokens = []Type{
	MovePtrRight,
	MovePtrLeft,
	IncrementCell,
	DecrementCell,
	OutputCell,
	InputCell,
	WhileStart,
	WhileEnd,
}

type Token struct {
	Type    Type
	Literal string
}

func New(t Type, l string) Token {
	return Token{
		Type:    t,
		Literal: l,
	}
}

func IsValidToken(ch byte) bool {
	// todo: unit test

	// EOF command will fail
	if ch == '0' {
		return true
	}

	for _, validValue := range AllValidTokens {
		if string(validValue) == string(ch) {
			return true
		}
	}

	return false
}
