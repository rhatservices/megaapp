package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "A mega application that is going to rule the world",
		Long:  "",
	}

	rootCmd.AddCommand(newRuleCommand())

	rootCmd.Execute()
}
