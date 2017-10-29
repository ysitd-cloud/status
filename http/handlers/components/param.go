package components

import (
	"path/filepath"

	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientSet *kubernetes.Clientset
var coreV1 v1.CoreV1Interface
var namespace string = os.Getenv("SERVICE_NAMESPACE")

func init() {
	var config *rest.Config
	var err error
	config, err = loadInOfClusterConfig()
	if err != nil {
		config, err = loadOutOfClusterConfig()
		if err != nil {
			panic(err)
		}
	}

	clientSet = kubernetes.NewForConfigOrDie(config)
	coreV1 = clientSet.CoreV1()
}

func loadInOfClusterConfig() (config *rest.Config, err error) {
	return rest.InClusterConfig()
}

func loadOutOfClusterConfig() (config *rest.Config, err error) {
	file := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	return clientcmd.BuildConfigFromFlags("", file)
}
