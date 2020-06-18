package cmd

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1" //  alias this as appsv1
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
