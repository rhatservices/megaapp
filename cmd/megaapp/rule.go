package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRuleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "rule rules the world",
		Run: runRuleCommand,
	}

	return cmd
}

func runRuleCommand(cmd *cobra.Command, args []string) {
	fmt.Println("I'm ruling the world!")
}
