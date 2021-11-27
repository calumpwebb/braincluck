package cmd

import (
	"braincluck/pkgs/runner"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the interpreter (file mode)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Must supply relative file path as first argument e.g. brainfuck run ./programs/hello-world.bf")
			os.Exit(1)
		}

		if err := runner.Run(os.Stdin, args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
