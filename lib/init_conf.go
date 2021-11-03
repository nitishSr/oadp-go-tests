package lib

import (
	"sync"

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
