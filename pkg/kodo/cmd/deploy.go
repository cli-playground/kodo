package cmd

import (
	"context"
	"fmt"
	"log"

	routev1 "github.com/openshift/api/route/v1"
	routev1client "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	appsv1 "k8s.io/api/apps/v1" //  alias this as appsv1
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	intstr1 "k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/rest"
)

type DeploymentVariables struct { //New struct for deployment creation variables
	Image    string
	Replicas int32
	Port     int32
}

func Deploy(deployVar *DeploymentVariables, envVar *EnvironmentVariables) error {

	client, clientError := newOpenShiftClient(envVar)

	if clientError != nil {
		log.Fatal(clientError)
		return clientError
	}

	deploymentObj := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "app-image-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deployVar.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "app-image-deployment",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "app-image-deployment",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container-image",
							Image: deployVar.Image, // should come from flag
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: deployVar.Port, // should come from flag
								},
							},
						},
					},
				},
			},
		},
	}
	_, deploymentError := client.AppsV1().Deployments(envVar.Namespace).Create(context.TODO(), deploymentObj, metav1.CreateOptions{})

	if deploymentError == nil {
		fmt.Printf("\nDeployment created")
	} else {
		log.Fatal(deploymentError)
	}

	return deploymentError

}

func Service(deployVar *DeploymentVariables, envVar *EnvironmentVariables) (*corev1.Service, error) {

	client, clientError := newOpenShiftClient(envVar)

	if clientError != nil {
		log.Fatal(clientError)
		return nil, clientError
	}
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "app-image-service",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:       80, // use correct datatype, hint: int32
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr1.FromInt(int(deployVar.Port)), // port is to be obtained from the command flag.
				},
			},
			Selector: map[string]string{
				"app": "app-image-deployment",
			},
		},
	}
	_, serviceError := client.CoreV1().Services(envVar.Namespace).Create(context.TODO(), svc, metav1.CreateOptions{})
	if serviceError == nil {
		fmt.Printf("\nService created")
	} else {
		log.Fatal(serviceError)
	}
	return svc, serviceError

}

func Route(deployVar *DeploymentVariables, envVar *EnvironmentVariables, svc *corev1.Service) error {

	_, clientError := newOpenShiftClient(envVar)

	if clientError != nil {
		log.Fatal(clientError)
		return clientError
	}

	routeObj := &routev1.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name: "myroute",
		},
		Spec: routev1.RouteSpec{
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: svc.Name,
			},
			Port: &routev1.RoutePort{
				TargetPort: intstr.IntOrString{IntVal: deployVar.Port}, // conventionalPort is 80
			},
		},
	}

	config := rest.Config{
		Host:        envVar.Host,
		BearerToken: envVar.Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	routeClient, _ := routev1client.NewForConfig(&config)

	_, routeClientError := routeClient.Routes(envVar.Namespace).Create(context.TODO(), routeObj, metav1.CreateOptions{})

	return routeClientError
}
