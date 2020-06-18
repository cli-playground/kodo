package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

/* modified variables to be inside a struct instead, and using references to instances of struct
   to access them in order to avoid using global variables */
type EnvironmentVariables struct {
	Host        string
	Namespace   string
	Bearertoken string
}

func newOpenShiftClient(envVar *EnvironmentVariables) (*kubernetes.Clientset, error) {
	config := rest.Config{
		Host:        envVar.Host,
		BearerToken: envVar.Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	return kubernetes.NewForConfig(&config)
}
