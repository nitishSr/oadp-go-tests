package lib

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"
	oadpv1alpha1 "github.com/openshift/oadp-operator/api/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// The Kubernetes Go client (nested within the OpenShift Go client)
// automatically registers its types in scheme.Scheme, however the
// additional OpenShift types must be registered manually.
func init() {
	oadpv1alpha1.AddToScheme(scheme.Scheme)
}

// Returns the client set config object
func getClientSet() (client.Client, error) {
	restconfig, err := getKubeRestConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := client.New(restconfig, client.Options{})
	return clientset, err
}

// Create Velero Custome Resource Definition from yaml
func CreateVeleroCRD(yamlFile string, namespace string) error {
	// Define struct for holding Velero type
	veleroSpec := oadpv1alpha1.Velero{}

	// Read yaml file containing Velero CR details
	veleroYamlData, err := getFileData(yamlFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(veleroYamlData, &veleroSpec)
	if err != nil {
		return err
	}

	// Get client config to read/write resources through API
	clientset, err := getClientSet()
	if err != nil {
		return err
	}

	// Create custom resource (velero in this case)
	err = clientset.Create(context.Background(), &veleroSpec)
	if err != nil {
		return err
	}
	fmt.Print("New Velero CR created \n")
	return nil
}
