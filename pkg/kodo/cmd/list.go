package cmd

import (
<<<<<<< HEAD
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
=======
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
>>>>>>> 54c068ce13a54d7f0750bc32c49d1268413fee1f
)

//List is a function to list number of pods in the cluster
//List
func List() error {
	client := newOpenShiftClient()
<<<<<<< HEAD
	pods, _ := client.CoreV1().Pods("").List(v1.ListOptions{})
=======
	pods, _ := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
>>>>>>> 54c068ce13a54d7f0750bc32c49d1268413fee1f
	fmt.Printf("\n The number of pods are %d \n", len(pods.Items))
	return nil
}
