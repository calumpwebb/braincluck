package parser

import (
	"braincluck/pkgs/ast"
	"braincluck/pkgs/lexer"
	"braincluck/pkgs/token"
	"errors"
)

type (
	Parser struct {
		l *lexer.Lexer

		currentToken token.Token

		astCurrentDepthLevel int // used to check if we're currently inside a loop mainly
	}
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,

		astCurrentDepthLevel: 0,
	}

	// read one token so currentToken is set
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := ast.NewProgram()
	program.Commands = []ast.Command{}

	for p.currentToken.Type != token.EOF {
		command, err := p.parseCommand()
		if err != nil {
			return nil, err
		}

		if command != nil {
			program.Commands = append(program.Commands, command)
		}
		p.nextToken()
	}

	return program, nil
}

func (p *Parser) parseCommand() (ast.Command, error) {
	switch t := p.currentToken.Type; t {
	case token.MovePtrRight, token.MovePtrLeft, token.IncrementCell,
		token.DecrementCell, token.OutputCell, token.InputCell, token.WhileEnd:
		return p.parseExecution()
	case token.WhileStart:
		return p.parseLoop()
	default:
		return nil, nil
	}
}

func (p *Parser) parseExecution() (ast.Command, error) {
	if p.currentToken.Type == token.WhileEnd {
		// if whileEnd then we need to check if we're inside a loop
		if p.astCurrentDepthLevel == 0 {
			return nil, errors.New("parsing error: Unable to find opening while loop")
		}
	}

	return ast.NewExecution(p.currentToken), nil
}

func (p *Parser) parseLoop() (ast.Command, error) {
	p.astCurrentDepthLevel += 1 // update so we know when how deep we are in the ast
	foundWhileEndToken := false

	totalPairings := 1 // total number of left WhileStarts we've seen. Once we've seen the same number of WhileEnds we've found the closing
	var commands []ast.Command

	startToken := p.currentToken

	// loop until we find the relevant WhileEnd Token or the we get to the EOF
	for !foundWhileEndToken && p.currentToken.Type != token.EOF {
		p.nextToken()

		switch p.currentToken.Type {
		case token.WhileStart:
			totalPairings += 1
		case token.WhileEnd:
			totalPairings -= 1
		}

		command, err := p.parseCommand()
		if err != nil {
			return nil, err
		}

		if command != nil {
			switch command.(type) {
			case *ast.Loop:
				totalPairings -= 1 // completed a loop so we can decrement a bracket that was part of that one
			}

			commands = append(commands, command)
		}

		if totalPairings == 0 {
			foundWhileEndToken = true
			break
		}
	}

	if !foundWhileEndToken {
		return nil, errors.New("parsing error: Unable to find closing while loop")
	}

	p.astCurrentDepthLevel -= 1
	return ast.NewLoop(startToken, commands), nil
}
