package lib

import (
	"context"
	"fmt"
	"io/ioutil"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

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

	// Read data from credentials file
	credsData, err := getFileData(credsFileName)
	if err != nil {
		return err
	}

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

	// Create Secret using coreV1 API
	secret, err := kubeclientset.CoreV1().Secrets(namespace).Create(context.TODO(), &secretConfig, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		fmt.Printf("Secret with name %s already exists ! Hence, Skipping this step ...\n", secretName)
		return nil
	} else if err != nil {
		return err
	}
	fmt.Printf("New secret created with name - %s \n", secret.Name)
	return nil
}
