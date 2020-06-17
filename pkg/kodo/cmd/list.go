package cmd

import (
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//List is a function to list number of pods in the cluster
//List
func List() error {
	client, clientError := newOpenShiftClient()

	if clientError != nil {
		log.Fatal(clientError)
	}
	pods, podlisterror := client.CoreV1().Pods(Namespace).List(v1.ListOptions{})
	fmt.Printf("\nThe number of pods are %d \n", len(pods.Items))
	return podlisterror
}
