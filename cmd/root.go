package cmd

import (
	"fmt"
	"os"
	"resedist/pkg/bootstrap"

	"github.com/spf13/cobra"
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

func Serve() {
	bootstrap.Serve()
}
