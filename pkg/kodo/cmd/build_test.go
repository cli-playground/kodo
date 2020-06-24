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

var envVar = EnvironmentVariables{
	Namespace:   "buildtest",
	Host:        "",
	Bearertoken: "",
}

var deployVar = DeploymentVariables{
	Source: "https://github.com/openshift/ruby-hello-world.git",
}

func getBuildConfig() buildv1api.BuildConfig {
	return buildv1api.BuildConfig{
		TypeMeta:   createTypeMeta("BuildConfig", "build.openshift.io/v1"),
		ObjectMeta: createObjectType("my-app-docker-build", "buildtest"),
		Spec:       createBuildSpec("https://github.com/openshift/ruby-hello-world.git"),
	}
}

func getBuildConfigSpec() buildv1api.BuildConfigSpec {
	return buildv1api.BuildConfigSpec{
		CommonSpec: buildv1api.CommonSpec{
			Source: buildv1api.BuildSource{
				Type: buildv1api.BuildSourceType("Git"),
				Git: &buildv1api.GitBuildSource{
					URI: "gitUrl",
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

func TestCreateTypeMeta(t *testing.T) {
	want := metav1.TypeMeta{
		Kind:       "testKind",
		APIVersion: "testAPIVersion",
	}
	got := createTypeMeta("testKind", "testAPIVersion")
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The TypeMeta didnt match")
	}

}

func TestCreateObjectType(t *testing.T) {
	want := metav1.ObjectMeta{
		Name:      "test",
		Namespace: "test-namespace",
	}
	got := createObjectType("test", "test-namespace")
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The ObjectTypes didnt match")
	}
}

func TestCreateBuildSpec(t *testing.T) {
	want := getBuildConfigSpec()
	got := createBuildSpec("gitUrl")
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The BuildConfigSpecs didnt match")
	}
}

func TestBuildConfig(t *testing.T) {
	want := getBuildConfig()
	got := createBuildConfig(&envVar, &deployVar)
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The BuildConfigs didnt match")
	}
}

func TestImageStream(t *testing.T) {
	want := getImageStream()
	got := createImageStream(&envVar)
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Println(diff)
		t.Fatalf("The ImageStreams didnt match")
	}
}
