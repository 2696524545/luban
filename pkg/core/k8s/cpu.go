package k8s

import (
	"context"
	"fmt"
	"github.com/dnsjia/luban/api/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCpu(c *gin.Context) {
	var cluster types.ClusterOptions
	if err := c.ShouldBindUri(&cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	clientSet, err := NewClientSet(cluster.Cluster)
	if err != nil {
		fmt.Println(err)
	}

	/*
		原生指标
			/apis/metrics.k8s.io/v1beta1/namespaces/test/pods/tomcat-787dc9647d-9tdxn
				/nodes – all node metrics; type []NodeMetrics
				/nodes/{node} – metrics for a specified node; type NodeMetrics
				/namespaces/{namespace}/pods – all pod metrics within namespace with support for all-namespaces; type []PodMetrics
				/namespaces/{namespace}/pods/{pod} – metrics for a specified pod; type PodMetrics

		阿里云heapster
		/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/metrics/cpu/usage_rate
		/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/metrics/memory/usage
		/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/namespaces/kube-system/pods/coredns-76b5876745-lfp8c/metrics/cpu/usage_rate
		/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/namespaces/kube-system/pods/coredns-76b5876745-lfp8c/metrics/memory/working_set
		/api/v1/namespaces/kube-system/services/heapster/proxy/apis/metrics/v1alpha1/nodes

	*/
	raw, err := clientSet.RESTClient().
		Get().
		AbsPath("/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/namespaces/kube-system/pods/coredns-76b5876745-lfp8c/metrics/memory/working_set").
		Do(context.TODO()).
		Raw()
	if err != nil {
		fmt.Println(err, "===")
	}

	c.JSON(http.StatusOK, gin.H{"data": string(raw)})
}
