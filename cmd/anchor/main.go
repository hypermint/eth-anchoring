package main

import (
	"log"

	"eth-anchoring/pkg/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use: "anchoring",
	}
	rootCmd.AddCommand(
		cmd.MakeRunCMD(),
		cmd.MakeSubmitCMD(),
		cmd.MakeL1CMD(),
	)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
