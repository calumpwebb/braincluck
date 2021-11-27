package runner

import (
	"braincluck/pkgs/environment"
	"braincluck/pkgs/evaluator"
	"braincluck/pkgs/lexer"
	"braincluck/pkgs/parser"
	"braincluck/pkgs/storage"
	"braincluck/pkgs/utils"
	"bufio"
	"fmt"
	"io"
)

func Run(in io.Reader, filePath string) error {
	scanner := bufio.NewScanner(in)

	s := storage.NewInMemoryStorage()
	env := environment.NewEnvironment(s, scanner)
	e := evaluator.NewEvaluator(env)

	fileStr, err := utils.ReadFileToString(filePath)
	if err != nil {
		return err
	}

	fmt.Printf("Running file at %s\n", filePath)
	err = run(e, fileStr)
	if err != nil {
		return err
	}

	return nil
}

func run(e *evaluator.Evaluator, t string) error {
	l := lexer.New(t)
	p := parser.New(l)

	program, err := p.ParseProgram()
	if err != nil {
		return fmt.Errorf("ERROR: %s", err)
	}

	if len(program.Commands) == 0 {
		fmt.Println("Nothing to run. Try again or type help")
		return nil
	}

	e.Eval(program)
	fmt.Print("\n")
	return nil
}
