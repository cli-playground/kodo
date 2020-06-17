package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	Host        string
	Namespace   string
	Bearertoken string
)

func newOpenShiftClient() *kubernetes.Clientset {
	config := rest.Config{
		Host:        Host,
		BearerToken: Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	myClientSet, _ := kubernetes.NewForConfig(&config)

	return myClientSet
}
