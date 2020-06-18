package cmd

import (
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//List is a function to list number of pods in the cluster
func List(envVar *EnvironmentVariables) error {
	client, clientError := newOpenShiftClient(envVar)

	if clientError != nil {
		log.Fatal(clientError)
		return clientError
	}
	pods, podlisterror := client.CoreV1().Pods(envVar.Namespace).List(v1.ListOptions{})

	if podlisterror == nil {
		fmt.Printf("\nThe number of pods are %d \n", len(pods.Items))
	} else {
		log.Fatal(podlisterror)
	}

	return podlisterror
}
