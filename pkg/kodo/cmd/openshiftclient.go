package cmd

import (
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
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

func NewOpenShiftClient(envVar *EnvironmentVariables) (*kubernetes.Clientset, error) {
	config := rest.Config{
		Host:        envVar.Host,
		BearerToken: envVar.Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	return kubernetes.NewForConfig(&config)
}

func NewRouteClient(envVar *EnvironmentVariables) (*routev1client.RouteV1Client, error) {
	config := rest.Config{
		Host:        envVar.Host,
		BearerToken: envVar.Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	routeClient, routev1ClientError := routev1client.NewForConfig(&config)

	if routev1ClientError != nil {
		return nil, routev1ClientError
	}

	return routeClient, routev1ClientError
}
