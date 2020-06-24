package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/dchest/uniuri"
	routev1 "github.com/openshift/api/route/v1"
	rv1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	appsv1 "k8s.io/api/apps/v1" //  alias this as appsv1
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	cv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

//DeploymentVariables - variables associated with deployment
type DeploymentVariables struct { //New struct for deployment creation variables
	Image    string
	Replicas int32
	Port     int32
	Source   string
}

type DeploymentIdentifiers struct { //New struct to hold unique identifiers for deployment/service/route
	DeploymentIdentifierName string
}

func GenerateUniqueIdentifiers() *DeploymentIdentifiers { // function to generate unique strings and put them in a struct
	deploymentIDs := DeploymentIdentifiers{
		DeploymentIdentifierName: strings.ToLower(uniuri.New()),
	}

	return &deploymentIDs
}

func Deploy(client v1.AppsV1Interface, deployVar *DeploymentVariables, envVar *EnvironmentVariables, deploymentID *DeploymentIdentifiers) (*appsv1.Deployment, error) {

	deploymentObj := &appsv1.Deployment{
		ObjectMeta: ObjectMeta(deploymentID.DeploymentIdentifierName),
		Spec:       DeploymentSpec(deployVar.Replicas, deploymentID.DeploymentIdentifierName, deployVar.Image, deployVar.Port),
	}
	deploymentPointer, deploymentError := client.Deployments(envVar.Namespace).Create(context.TODO(), deploymentObj, metav1.CreateOptions{})

	if deploymentError == nil {
		fmt.Printf("\nDeployment created")
	}

	return deploymentPointer, deploymentError

}

func Service(client cv1.CoreV1Interface, deployVar *DeploymentVariables, envVar *EnvironmentVariables, deploymentID *DeploymentIdentifiers) (*corev1.Service, error) {

	svc := &corev1.Service{
		ObjectMeta: ObjectMeta(deploymentID.DeploymentIdentifierName),
		Spec:       ServiceSpec(deployVar.Port, deploymentID.DeploymentIdentifierName),
	}
	servicePointer, serviceError := client.Services(envVar.Namespace).Create(context.TODO(), svc, metav1.CreateOptions{})
	if serviceError == nil {
		fmt.Printf("\nService created")
	}
	return servicePointer, serviceError

}

func Route(routeClient rv1.RouteV1Interface, deployVar *DeploymentVariables, envVar *EnvironmentVariables, svc *corev1.Service, deploymentID *DeploymentIdentifiers) (*routev1.Route, error) {

	routeObj := &routev1.Route{
		ObjectMeta: ObjectMeta(deploymentID.DeploymentIdentifierName),
		Spec:       RouteSpec(deployVar.Port, svc.Name),
	}

	routePointer, routeClientError := routeClient.Routes(envVar.Namespace).Create(context.TODO(), routeObj, metav1.CreateOptions{})
	if routeClientError == nil {
		fmt.Printf("\nRoute created")
	}
	return routePointer, routeClientError
}

func ObjectMeta(resourceName string) metav1.ObjectMeta {
	ObjectMeta := metav1.ObjectMeta{
		Name: resourceName,
	}
	return ObjectMeta
}

func DeploymentSpec(Replicas int32, DeploymentIdentifierName string, Image string, Port int32) appsv1.DeploymentSpec {
	Spec := appsv1.DeploymentSpec{
		Replicas: &Replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": DeploymentIdentifierName,
			},
		},
		Template: DeploymentTemplate(DeploymentIdentifierName, Image, Port),
	}

	return Spec
}

func DeploymentTemplate(DeploymentIdentifierName string, Image string, Port int32) corev1.PodTemplateSpec {
	Template := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": DeploymentIdentifierName,
			},
		},
		Spec: DeploymentTemplateSpec(DeploymentIdentifierName, Image, Port),
	}

	return Template
}

func DeploymentTemplateSpec(DeploymentIdentifierName string, Image string, Port int32) corev1.PodSpec {

	Spec := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  DeploymentIdentifierName,
				Image: Image, // should come from flag
				Ports: []corev1.ContainerPort{
					{
						Name:          "http",
						Protocol:      corev1.ProtocolTCP,
						ContainerPort: Port, // should come from flag
					},
				},
			},
		},
	}

	return Spec
}

func ServiceSpec(Port int32, DeploymentIdentifierName string) corev1.ServiceSpec {

	Spec := corev1.ServiceSpec{
		Ports: []corev1.ServicePort{
			{
				Port:       80, // use correct datatype, hint: int32
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(int(Port)), // port is to be obtained from the command flag.
			},
		},
		Selector: map[string]string{
			"app": DeploymentIdentifierName,
		},
	}

	return Spec
}

func RouteSpec(Port int32, DeploymentIdentifierName string) routev1.RouteSpec {

	Spec := routev1.RouteSpec{
		To: routev1.RouteTargetReference{
			Kind: "Service",
			Name: DeploymentIdentifierName,
		},
		Port: &routev1.RoutePort{
			TargetPort: intstr.IntOrString{IntVal: Port}, // conventionalPort is 80
		},
	}

	return Spec
}
