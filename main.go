package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rcommand = &cobra.Command{
	Use: "kodo",
}

var versionCommand = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 2")
	},
}

func init() {
	rcommand.AddCommand(versionCommand)
}
func main() {
	fmt.Println("Hello World")
	rcommand.Execute()
}
