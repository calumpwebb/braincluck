package cmd

import (
	"braincluck/pkgs/repl"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var replCommand = &cobra.Command{
	Use:   "repl",
	Short: "Run the interpreter (REPL)",
	Run: func(cmd *cobra.Command, args []string) {
		repl.PrintHeader()
		repl.PrintHelp(false)

		if err := repl.Run(os.Stdin); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(replCommand)
}
