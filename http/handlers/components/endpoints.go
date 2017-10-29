package components

import (
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

func createPodInfo(addresses []v1.EndpointAddress) *podInfo {
	names := make([]string, 0)
	for _, pod := range addresses {
		names = append(names, pod.TargetRef.Name)
	}

	return &podInfo{
		Total: len(addresses),
		Pods:  names,
	}
}

func createComponentStatusEndpoint(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint, err := coreV1.Endpoints(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		subset := endpoint.Subsets[0]

		notReady := subset.NotReadyAddresses
		notReadyPods := createPodInfo(notReady)

		ready := subset.Addresses
		readyPods := createPodInfo(ready)

		resp := componentEndpoints{
			Total:        len(ready) + len(notReady),
			Available:    readyPods,
			NotAvailable: notReadyPods,
		}

		c.JSON(http.StatusOK, &resp)
	}
}
