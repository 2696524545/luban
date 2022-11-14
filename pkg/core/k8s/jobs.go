package k8s

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/api/types"
)

func ListJobs(c *gin.Context) {
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

	jobList, err := clientSet.BatchV1().Jobs(namespaceOptions.Namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": jobList})

}

func GetJob(c *gin.Context) {
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

	job, err := clientSet.BatchV1().Jobs(requestOptions.Namespace).Get(context.TODO(), requestOptions.Name, metaV1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": job})
}

func DeleteJob(c *gin.Context) {
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

	err = clientSet.BatchV1().Jobs(requestOptions.Namespace).Delete(context.TODO(), requestOptions.Name, metaV1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功"})
}
