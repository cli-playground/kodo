package cmd

import (
	"fmt"
	"log"
	"testing"

	routev1 "github.com/openshift/api/route/v1"
	fakeRouteClientset "github.com/openshift/client-go/route/clientset/versioned/fake"
	appsv1 "k8s.io/api/apps/v1" //  alias this as appsv1
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	fakeKubeClientset "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientSet() (*kubernetes.Clientset, error) {
	clientConfig, err := GetRESTConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get client config due to %w", err)
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get APIs client due to %w", err)
	}
	return clientSet, nil
}

func GetRESTConfig() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	return kubeconfig.ClientConfig()
}

func TypeMeta(kind, apiVersion string) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       kind,
		APIVersion: apiVersion,
	}
}

func Create(name string) *corev1.Namespace {
	ns := &corev1.Namespace{
		TypeMeta: TypeMeta("Namespace", "v1"),
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	return ns
}

func TestGenerateUniqueIdentifiers(t *testing.T) {
	got := GenerateUniqueIdentifiers()

	want := DeploymentIdentifiers{
		DeploymentIdentifierName: "sr452vf62vdyd78j",
	}

	if len(got.DeploymentIdentifierName) != len(want.DeploymentIdentifierName) {
		log.Fatal("deployment identifier name mismatch")
	}

}

func TestObjectMeta(t *testing.T) {
	got := ObjectMeta("sr452vf62vdyd78j")

	want := metav1.ObjectMeta{
		Name: "sr452vf62vdyd78j",
	}

	if got.Name != want.Name {
		log.Fatal("object meta mismatch ")
	}
}

func TestDeploymentSpec(t *testing.T) {
	var replicas int32 = 11

	got := DeploymentSpec(replicas, "sr452vf62vdyd78j", "openshift/hello-openshift:latest", 8080)

	want := appsv1.DeploymentSpec{
		Replicas: &replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "sr452vf62vdyd78j",
			},
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "sr452vf62vdyd78j",
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "sr452vf62vdyd78j",
						Image: "openshift/hello-openshift:latest", // should come from flag
						Ports: []corev1.ContainerPort{
							{
								Name:          "http",
								Protocol:      corev1.ProtocolTCP,
								ContainerPort: 8080, // should come from flag
							},
						},
					},
				},
			},
		},
	}

	if *got.Replicas != *want.Replicas {
		log.Fatal("replica count mismatch")
	} else if got.Selector.MatchLabels["app"] != want.Selector.MatchLabels["app"] {
		log.Fatal("selector name mismatch")
	} else if got.Template.ObjectMeta.Labels["app"] != got.Template.ObjectMeta.Labels["app"] {
		log.Fatal("Object meta label mismatch")
	} else if got.Template.Spec.Containers[0].Name != want.Template.Spec.Containers[0].Name {
		log.Fatal("container name mismatch")
	} else if got.Template.Spec.Containers[0].Image != want.Template.Spec.Containers[0].Image {
		log.Fatal("container image mismatch")
	} else if got.Template.Spec.Containers[0].Ports[0].ContainerPort != want.Template.Spec.Containers[0].Ports[0].ContainerPort {
		log.Fatal("container port mismatch")
	}
}

func TestDeploymentTemplate(t *testing.T) {

	got := DeploymentTemplate("sr452vf62vdyd78j", "openshift/hello-openshift:latest", 8080)

	want := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": "sr452vf62vdyd78j",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "sr452vf62vdyd78j",
					Image: "openshift/hello-openshift:latest", // should come from flag
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: 8080, // should come from flag
						},
					},
				},
			},
		},
	}

	if got.ObjectMeta.Labels["app"] != got.ObjectMeta.Labels["app"] {
		log.Fatal("Object meta label mismatch")
	} else if got.Spec.Containers[0].Name != want.Spec.Containers[0].Name {
		log.Fatal("container name mismatch")
	} else if got.Spec.Containers[0].Image != want.Spec.Containers[0].Image {
		log.Fatal("container image mismatch")
	} else if got.Spec.Containers[0].Ports[0].ContainerPort != want.Spec.Containers[0].Ports[0].ContainerPort {
		log.Fatal("container port mismatch")
	}
}

func TestDeploymentTemplateSpec(t *testing.T) {

	got := DeploymentTemplateSpec("sr452vf62vdyd78j", "openshift/hello-openshift:latest", 8080)

	want := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:  "sr452vf62vdyd78j",
				Image: "openshift/hello-openshift:latest", // should come from flag
				Ports: []corev1.ContainerPort{
					{
						Name:          "http",
						Protocol:      corev1.ProtocolTCP,
						ContainerPort: 8080, // should come from flag
					},
				},
			},
		},
	}

	if got.Containers[0].Name != want.Containers[0].Name {
		log.Fatal("container name mismatch")
	} else if got.Containers[0].Image != want.Containers[0].Image {
		log.Fatal("container image mismatch")
	} else if got.Containers[0].Ports[0].ContainerPort != want.Containers[0].Ports[0].ContainerPort {
		log.Fatal("container port mismatch")
	}

}

func TestServiceSpec(t *testing.T) {

	got := ServiceSpec(8080, "sr452vf62vdyd78j")

	want := corev1.ServiceSpec{
		Ports: []corev1.ServicePort{
			{
				Port:       80, // use correct datatype, hint: int32
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(8080), // port is to be obtained from the command flag.
			},
		},
		Selector: map[string]string{
			"app": "sr452vf62vdyd78j",
		},
	}

	if got.Ports[0].TargetPort != want.Ports[0].TargetPort {
		log.Fatal("error")
	} else if got.Selector["apps"] != want.Selector["apps"] {
		log.Fatal("selector app mismatch")
	}
}

func TestRouteSpec(t *testing.T) {

	got := RouteSpec(8080, "sr452vf62vdyd78j")

	want := routev1.RouteSpec{
		To: routev1.RouteTargetReference{
			Kind: "Service",
			Name: "sr452vf62vdyd78j",
		},
		Port: &routev1.RoutePort{
			TargetPort: intstr.IntOrString{IntVal: 8080}, // conventionalPort is 80
		},
	}

	if got.To.Name != want.To.Name {
		log.Fatal("route name mismatch")
	} else if got.Port.TargetPort != want.Port.TargetPort {
		log.Fatal("Target port mismatch")
	}

}

func TestDeploy(t *testing.T) {

	var deployVar = DeploymentVariables{
		Image:    "openshift/hello-openshift:latest",
		Replicas: 11,
		Port:     8080,
	}

	var envVar = EnvironmentVariables{
		Host:        "",
		Bearertoken: "",
		Namespace:   "jaideep-test-june-23",
	}

	deploymentID := GenerateUniqueIdentifiers()

	got, err := Deploy(fakeKubeClientset.NewSimpleClientset().AppsV1(), &deployVar, &envVar, deploymentID)

	want := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentID.DeploymentIdentifierName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deployVar.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentID.DeploymentIdentifierName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentID.DeploymentIdentifierName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentID.DeploymentIdentifierName,
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

	if err != nil {
		log.Fatal(err)
	} else if got.ObjectMeta.Name != want.ObjectMeta.Name {
		log.Fatal("object meta mismatch ")
	} else if *got.Spec.Replicas != *want.Spec.Replicas {
		log.Fatal("replica count mismatch")
	} else if got.Spec.Selector.MatchLabels["app"] != want.Spec.Selector.MatchLabels["app"] {
		log.Fatal("selector name mismatch")
	} else if got.Spec.Template.ObjectMeta.Labels["app"] != got.Spec.Template.ObjectMeta.Labels["app"] {
		log.Fatal("Object meta label mismatch")
	} else if got.Spec.Template.Spec.Containers[0].Name != want.Spec.Template.Spec.Containers[0].Name {
		log.Fatal("container name mismatch")
	} else if got.Spec.Template.Spec.Containers[0].Image != want.Spec.Template.Spec.Containers[0].Image {
		log.Fatal("container image mismatch")
	} else if got.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort != want.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort {
		log.Fatal("container port mismatch")
	}

}

func TestService(t *testing.T) {

	var deployVar = DeploymentVariables{
		Image:    "openshift/hello-openshift:latest",
		Replicas: 11,
		Port:     8080,
	}

	var envVar = EnvironmentVariables{
		Host:        "",
		Bearertoken: "",
		Namespace:   "jaideep-test-june-22",
	}

	deploymentID := GenerateUniqueIdentifiers()

	got, err := Service(fakeKubeClientset.NewSimpleClientset().CoreV1(), &deployVar, &envVar, deploymentID)

	want := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentID.DeploymentIdentifierName,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:       80, // use correct datatype, hint: int32
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.FromInt(int(deployVar.Port)), // port is to be obtained from the command flag.
				},
			},
			Selector: map[string]string{
				"app": deploymentID.DeploymentIdentifierName,
			},
		},
	}

	if err != nil {
		log.Fatal(err)
	} else if got.ObjectMeta.Name != want.ObjectMeta.Name {
		log.Fatal("object meta mismatch ")
	} else if got.Spec.Ports[0].TargetPort != want.Spec.Ports[0].TargetPort {
		log.Fatal("Target port mismatch")
	} else if got.Spec.Selector["apps"] != want.Spec.Selector["apps"] {
		log.Fatal("selector app mismatch")
	}

}

func TestRoute(t *testing.T) {

	var deployVar = DeploymentVariables{
		Image:    "openshift/hello-openshift:latest",
		Replicas: 11,
		Port:     8080,
	}

	var envVar = EnvironmentVariables{
		Host:        "",
		Bearertoken: "",
		Namespace:   "jaideep-test-june-23",
	}

	deploymentID := GenerateUniqueIdentifiers()
	serviceObj, _ := Service(fakeKubeClientset.NewSimpleClientset().CoreV1(), &deployVar, &envVar, deploymentID)

	routeFakeClient := fakeRouteClientset.NewSimpleClientset().RouteV1()

	got, err := Route(routeFakeClient, &deployVar, &envVar, serviceObj, deploymentID)

	want := &routev1.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentID.DeploymentIdentifierName,
		},
		Spec: routev1.RouteSpec{
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: deploymentID.DeploymentIdentifierName,
			},
			Port: &routev1.RoutePort{
				TargetPort: intstr.IntOrString{IntVal: deployVar.Port}, // conventionalPort is 80
			},
		},
	}

	if err != nil {
		log.Fatal(err)
	} else if got.ObjectMeta.Name != want.ObjectMeta.Name {
		log.Fatal("object meta mismatch ")
	} else if got.Spec.To.Name != want.Spec.To.Name {
		log.Fatal("route name mismatch")
	} else if got.Spec.Port.TargetPort != want.Spec.Port.TargetPort {
		log.Fatal("Target port mismatch")
	}

}
