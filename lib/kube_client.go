package lib

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var loadOnce sync.Once
var kubeconfig clientcmd.ClientConfig

// Loads kubeconfig object and returns rest client config
func getKubeRestConfig() (*rest.Config, error) {
	loadOnce.Do(func() {
		kubeconfig = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
	})
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return restconfig, nil
}

// Returns the kubernetes client set config object
func getKubeClientSet() (*kubernetes.Clientset, error) {
	restconfig, err := getKubeRestConfig()
	if err != nil {
		return nil, err
	}
	kubeclientset, err := kubernetes.NewForConfig(restconfig)
	return kubeclientset, err
}

// Read file and return data in form of bytes
func getFileData(fileName string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return fileData, nil
}

// List all the pods in particular namespace
func GetPodsInNamespace(namespace string) (*v1.PodList, error) {
	kubeclientset, err := getKubeClientSet()
	if err != nil {
		return nil, err
	}

	pods, err := kubeclientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

// Create secret from credentials file in a namespace
func CreateSecret(credsFileName string, namespace string, secretName string) error {
	kubeclientset, err := getKubeClientSet()
	if err != nil {
		return err
	}
	credsData, err := getFileData(credsFileName)

	// Define a struct for holding secret configs
	secretConfig := coreV1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: namespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: metav1.SchemeGroupVersion.String(),
		},
		Data: map[string][]byte{
			"cloud": credsData,
		},
		Type: coreV1.SecretTypeOpaque,
	}

	secret, err := kubeclientset.CoreV1().Secrets(namespace).Create(context.TODO(), &secretConfig, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return err
	}

	fmt.Printf("New secret created with name - %s ", secret.Name)
	return nil
}
