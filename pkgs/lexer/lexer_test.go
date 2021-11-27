package lexer_test

import (
	"braincluck/pkgs/lexer"
	"braincluck/pkgs/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `><+-.ABC123,[]` // should ignore the illegalCharacters (ABC123)

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{
			expectedType:    token.MovePtrRight,
			expectedLiteral: ">",
		},
		{
			expectedType:    token.MovePtrLeft,
			expectedLiteral: "<",
		},
		{
			expectedType:    token.IncrementCell,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.DecrementCell,
			expectedLiteral: "-",
		},
		{
			expectedType:    token.OutputCell,
			expectedLiteral: ".",
		},
		{
			expectedType:    token.InputCell,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.WhileStart,
			expectedLiteral: "[",
		},
		{
			expectedType:    token.WhileEnd,
			expectedLiteral: "]",
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
	}

	l := lexer.New(input)

	for i, tc := range tests {
		tok := l.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tc.expectedType, tok.Type)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tc.expectedLiteral, tok.Literal)
		}
	}
}
