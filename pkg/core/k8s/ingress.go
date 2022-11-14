package k8s

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/api/types"
)

func ListIngress(c *gin.Context) {
	var namespaceOptions types.NamespaceOptions
	if err := c.ShouldBindUri(&namespaceOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	clientSet, err := NewClientSet(namespaceOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	ingressList, err := clientSet.NetworkingV1().Ingresses(namespaceOptions.Namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": ingressList})
}

func GetIngress(c *gin.Context) {
	var requestOptions types.RequestOptions
	if err := c.ShouldBindUri(&requestOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	clientSet, err := NewClientSet(requestOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	ingress, err := clientSet.NetworkingV1().Ingresses(requestOptions.Namespace).Get(context.TODO(), requestOptions.Name, metaV1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": ingress})
}
