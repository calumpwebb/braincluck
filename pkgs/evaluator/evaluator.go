package evaluator

import (
	"braincluck/pkgs/ast"
	"braincluck/pkgs/environment"
	"braincluck/pkgs/token"
	"fmt"
	"strconv"
)

type Evaluator struct {
	Environment *environment.Environment
}

func NewEvaluator(environment *environment.Environment) *Evaluator {
	return &Evaluator{Environment: environment}
}

func (e *Evaluator) Eval(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		e.evalCommands(node)
	case *ast.Execution:
		e.evalExecution(node)
	case *ast.Loop:
		e.evalLoop(node)
	}
}

func (e *Evaluator) evalCommands(node *ast.Program) {
	for _, command := range node.Commands {
		e.Eval(command)
	}
}

func (e *Evaluator) evalExecution(node *ast.Execution) {
	switch node.Token.Type {
	case token.WhileEnd:
		// no-op
	case token.MovePtrRight:
		e.Environment.Store.MovePtrRight()
	case token.MovePtrLeft:
		e.Environment.Store.MovePtrLeft()
	case token.IncrementCell:
		e.Environment.Store.IncrementCell()
	case token.DecrementCell:
		e.Environment.Store.DecrementCell()
	case token.OutputCell:
		fmt.Print(string([]byte{e.Environment.Store.GetValue()}))
	case token.InputCell:
		// TODO: implement this!
		fmt.Print("\n(input): ")
		scanned := e.Environment.Scanner.Scan()
		if !scanned {
			return
		}
		line := e.Environment.Scanner.Text()

		val, _ := strconv.ParseInt(line, 10, 8)
		// todo: handle error
		e.Environment.Store.SetCellValue(uint8(val))
	default:
		fmt.Printf("Unhandled TokenType=%s", node.Token.Type)
	}
}

func (e *Evaluator) evalLoop(node *ast.Loop) {
	if e.Environment.Store.GetValue() == 0 {
		return
	}

	for {
		for _, command := range node.Commands {
			e.Eval(command)
		}

		if e.Environment.Store.GetValue() == 0 {
			break
		}
	}
}
