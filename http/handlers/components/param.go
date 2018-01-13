package components

import (
	"os"

	"code.ysitd.cloud/k8s/utils/go"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

var clientSet kubernetes.Interface
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
