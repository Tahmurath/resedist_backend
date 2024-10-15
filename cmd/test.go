package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  "test & output html",
	Run: func(cmd *cobra.Command, args []string) {
		htmltest()
	},
}

func htmltest() {

	cmd := exec.Command("go", "test", "./...", "-coverprofile", "profile.out")

	output, err := cmd.CombinedOutput() // برای گرفتن هم خروجی stdout و هم stderr
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))

	cmd2 := exec.Command("go", "tool", "cover", "-html=profile.out")
	//cmd2 := exec.Command("go", "tool", "cover", "-html=profile.out", "-o", "coverage.html")
	output2, err2 := cmd2.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err2)
		return
	}

	fmt.Println(string(output2))

	fmt.Println("Coverage report generated in HTML format.")
}
