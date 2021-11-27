package repl

import (
	"braincluck/pkgs/environment"
	"braincluck/pkgs/evaluator"
	"braincluck/pkgs/lexer"
	"braincluck/pkgs/parser"
	"braincluck/pkgs/storage"
	"bufio"
	"fmt"
	"io"
	"strings"
)

const PROMPT = ">> "
const REPL_VERSION = "1.0.0"
const BF_VERSION = "1.0.0"

func Run(in io.Reader) error {
	scanner := bufio.NewScanner(in)

	s := storage.NewInMemoryStorage()
	env := environment.NewEnvironment(s, scanner)
	e := evaluator.NewEvaluator(env)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return nil
		}

		line := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch line {
		case "":
			continue
		case "c", "clear":
			s = storage.NewInMemoryStorage()
			env = environment.NewEnvironment(s, scanner)
			e = evaluator.NewEvaluator(env)
			PrintClear()
			continue
		case "h", "help":
			PrintHelp(true)
			continue
		case "v", "version":
			PrintVersion()
		case "t", "tape":
			// todo: implement printtape
			PrintTape()
		case "f", "file":
			// todo: implement FILE!
			PrintFile()
		default:
			l := lexer.New(line)
			p := parser.New(l)

			program, err := p.ParseProgram()
			if err != nil {
				return fmt.Errorf("ERROR: %s", err)
			}

			if len(program.Commands) == 0 {
				fmt.Println("Nothing to run. Try again or type help")
			}

			e.Eval(program)
			fmt.Print("\n")
		}
	}
}

func PrintHelp(showTitle bool) {
	if showTitle {
		fmt.Println("Help Command: type either the long or short version of the commands into the REPL!")
	}

	fmt.Println("	help / h  	: shows help information again if you need it!")
	fmt.Println("	tape / t 	: prints the currently stored tape (first 25 max)")
	fmt.Println("	clear / c 	: clears the storage and creates a new one (ptr index will now be at 0)")
	fmt.Println("	version / v : prints the current version of the REPL and BrainFuck")
}

func PrintHeader() {
	fmt.Println("This is the Brainfuck programming language REPL!")
}

func PrintVersion() {
	fmt.Printf("REPL Version %s and Brainfuck Version %s\n", REPL_VERSION, BF_VERSION)
}

func PrintClear() {
	fmt.Println("Storage Tape cleared [0, 0, 0,...]")
}

func PrintTape() {
	fmt.Println("Storage Tape: [...]")
}

func PrintFile() {
	fmt.Print("File: NOT IMPLEMENTED YET!")
}
