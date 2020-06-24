package cmd

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//List is a function to list number of pods in the cluster
func List(envVar *EnvironmentVariables) error {
	client, clientError := NewOpenShiftClient(envVar)

	if clientError != nil {
		log.Fatal(clientError)
		return clientError
	}
	pods, podlisterror := client.CoreV1().Pods(envVar.Namespace).List(context.TODO(), metav1.ListOptions{})
	if podlisterror == nil {
		fmt.Printf("\nThe number of pods are %d \n", len(pods.Items))
	} else {
		log.Fatal(podlisterror)
	}

	return podlisterror
}
