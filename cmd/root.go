package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Help",
	Short: "Help command",
	Long:  `Display help`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
