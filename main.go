package main

import (
	"fmt"
	"log"

	"github.com/cli-playground/kodo/pkg/kodo/cmd"
	"github.com/spf13/cobra"
)

var envVar = new(cmd.EnvironmentVariables)   // creating instance of struct EnvironmentVariables from cmd.openshiftclient.go
var deployVar = new(cmd.DeploymentVariables) // creating instance of struct EnvironmentVariables from cmd.deploy.go

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
	rcommand.AddCommand(countCommand)
	countCommand.AddCommand(podCommand)
	rcommand.AddCommand(deployCommand)

	rcommand.PersistentFlags().StringVarP(&envVar.Host, "server", "s", "myurl", "this is the cluster url")
	rcommand.PersistentFlags().StringVarP(&envVar.Bearertoken, "token", "t", "usertoken", "this is the user token")
	rcommand.PersistentFlags().StringVarP(&envVar.Namespace, "namespace", "n", "", "this is the namespace")

	/* Changed default namespace from 'namespace' to "" because
	'namespace' causes some internal error and returns 0 pods
	even if login credentials correct but namespace not explicitly set
	whereas "" works fine */

	rcommand.PersistentFlags().StringVarP(&deployVar.Image, "image", "i", "myimage", "this is the tagged image for deployment")
	rcommand.PersistentFlags().Int32VarP(&deployVar.Replicas, "replicas", "r", 1, "number of replicas")
	rcommand.PersistentFlags().Int32VarP(&deployVar.Port, "port", "p", 8000, "port at which app should run")

	rcommand.MarkFlagRequired("server")
}

func main() {

	rcommand.Execute()
}

var countCommand = &cobra.Command{
	Use:   "count",
	Short: "Command to count <resources>",
}

var podCommand = &cobra.Command{
	Use: "pods",
	Run: func(cm *cobra.Command, args []string) {
		fmt.Println("List All Kubernetes Applications")
		fmt.Printf("\nFetching all applications from %s in namespace %s", envVar.Host, envVar.Namespace)
		cmd.List(envVar)
	},
}

var deployCommand = &cobra.Command{
	Use: "deploy",
	Run: func(cm *cobra.Command, args []string) {
		fmt.Println("Creating image deployment")

		deploymentID := cmd.GenerateUniqueIdentifiers()

		client, clientError := cmd.NewOpenShiftClient(envVar)

		if clientError != nil {
			log.Fatal(clientError)
		} else {
			_, deployError := cmd.Deploy(client.AppsV1(), deployVar, envVar, deploymentID)
			if deployError != nil {
				log.Fatal(deployError)
			} else {
				serviceObj, serviceObjError := cmd.Service(client.CoreV1(), deployVar, envVar, deploymentID)
				if serviceObjError != nil {
					log.Fatal(serviceObjError)
				} else {
					routeClient, routev1ClientError := cmd.NewRouteClient(envVar)
					if routev1ClientError != nil {
						log.Fatal(routev1ClientError)
					} else {
						_, routeError := cmd.Route(routeClient, deployVar, envVar, serviceObj, deploymentID)
						if routeError != nil {
							log.Fatal(routeError)
						}
					}
				}
			}
		}

	},
}