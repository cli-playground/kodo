package cmd

import (
	"context"

	"k8s.io/client-go/rest"

	buildv1api "github.com/openshift/api/build/v1"
	imagev1api "github.com/openshift/api/image/v1"

	buildv1clientapi "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1clientapi "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newBuildConfigClient() *buildv1clientapi.BuildV1Client {
	config := rest.Config{
		Host:        Host,
		BearerToken: Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	myClientSet, _ := buildv1clientapi.NewForConfig(&config)
	return myClientSet
}

func createBuildConfig() error {
	buildclient := newBuildConfigClient()
	buildConfig := buildv1api.BuildConfig{
		// populate with relevant values
	}

	_, err := buildclient.BuildConfigs(Namespace).Create(context.TODO(), &buildConfig, metav1.CreateOptions{})
	return err
}

func newImageStreamClient() *imagev1clientapi.ImageV1Client {
	config := rest.Config{
		Host:        Host,
		BearerToken: Bearertoken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	myClientSet, _ := imagev1clientapi.NewForConfig(&config)
	return myClientSet
}

func createImageStream() error {
	imagestreamClient := newImageStreamClient()
	imageStream := imagev1api.ImageStream{
		// populate with relevant values
	}

	_, err := imagestreamClient.ImageStreams(Namespace).Create(context.TODO(), &imageStream, metav1.CreateOptions{})
	return err

}
