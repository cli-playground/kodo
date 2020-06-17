package main

import (
	"fmt"
	"os"

	"github.com/cli-playground/kodo/pkg/kodo/cmd"
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
	rcommand.AddCommand(listCommand)
	rcommand.PersistentFlags().StringVarP(&cmd.Host, "server", "s", "myurl", "this is the cluster url")
	rcommand.PersistentFlags().StringVarP(&cmd.Bearertoken, "token", "t", "usertoken", "this is the user token")
	rcommand.PersistentFlags().StringVarP(&cmd.Namespace, "namespace", "n", "", "this is the namespace")
	rcommand.MarkFlagRequired("server")
}

func main() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg)
	rcommand.Execute()
}

var listCommand = &cobra.Command{
	Use: "list",
	Run: func(cm *cobra.Command, args []string) {
		fmt.Println("\nList All Kubernetes Applications")
		fmt.Printf("\nFetching all applications from %s in namespace %s", cmd.Host, cmd.Namespace)
		cmd.List()

	},
}
