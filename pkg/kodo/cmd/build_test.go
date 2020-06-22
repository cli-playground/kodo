package cmd

import (
	"fmt"
	"testing"

	cmp "github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	buildv1api "github.com/openshift/api/build/v1"
	imagev1api "github.com/openshift/api/image/v1"
	corev1 "k8s.io/api/core/v1"
)

func getBuildConfig() buildv1api.BuildConfig {
	return buildv1api.BuildConfig{
		TypeMeta: metav1.TypeMeta{
			Kind:       "BuildConfig",
			APIVersion: "build.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-app-docker-build",
		},
		Spec: buildv1api.BuildConfigSpec{
			CommonSpec: buildv1api.CommonSpec{
				Source: buildv1api.BuildSource{
					Type: buildv1api.BuildSourceType("Git"),
					Git: &buildv1api.GitBuildSource{
						URI: "https://github.com/openshift/ruby-hello-world.git",
					},
				},
				Strategy: buildv1api.BuildStrategy{
					Type: buildv1api.BuildStrategyType("Docker"),
				},
				Output: buildv1api.BuildOutput{
					To: &corev1.ObjectReference{
						Kind: "ImageStreamTag",
						Name: "my-ruby-image:latest",
					},
				},
			},
		},
	}
}

func getImageStream() imagev1api.ImageStream {
	return imagev1api.ImageStream{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ImageStream",
			APIVersion: "image.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-ruby-image",
			Namespace: "buildtest",
		},
	}
}

func TestBuildConfig(t *testing.T) {
	want := getBuildConfig()
	got := createBuildConfig("https://github.com/openshift/ruby-hello-world.git")
	// fmt.Println("Want-----> ", want)
	// fmt.Println("Got---->", got)
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The Build Configs didnt match")
	}
}

func TestImageStream(t *testing.T) {
	want := getImageStream()
	got := createImageStream()
	// fmt.Println("Want-----> ", want)
	// fmt.Println("Got---->", got)
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The ImageStreams didnt match")
	}
}
