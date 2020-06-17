package cmd

import (
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//List is a function to list number of pods in the cluster
//List
func List() error {
	client := newOpenShiftClient()
	pods, _ := client.CoreV1().Pods("").List(v1.ListOptions{})
	fmt.Printf("\n The number of pods are %d \n", len(pods.Items))
	return nil
}
