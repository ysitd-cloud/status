package components

import (
	"os"

	"github.com/ysitd-cloud/k8s-utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

var clientSet *kubernetes.Clientset
var coreV1 v1.CoreV1Interface
var namespace = os.Getenv("SERVICE_NAMESPACE")

func init() {
	var err error
	clientSet, err = utils.AutoConnect()
	if err != nil {
		panic(err)
	}
	coreV1 = clientSet.CoreV1()
}
