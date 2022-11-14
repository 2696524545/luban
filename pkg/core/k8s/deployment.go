package k8s

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/api/types"
)

func ListDeployment(c *gin.Context) {
	var namespaceOptions types.NamespaceOptions
	if err := c.ShouldBindUri(&namespaceOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	clientSet, err := NewClientSet(namespaceOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "获取deployment出错", "data": ""})
		return
	}

	deploymentList, err := clientSet.AppsV1().Deployments(namespaceOptions.Namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "获取deployment出错", "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": deploymentList.Items})
}

func GetDeploymentDetail(c *gin.Context) {
	var requestOptions types.RequestOptions
	if err := c.ShouldBindUri(&requestOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	clientSet, err := NewClientSet(requestOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "获取deployment出错", "data": ""})
		return
	}

	// 获取deployment详情
	deployment, err := clientSet.AppsV1().Deployments(requestOptions.Namespace).Get(context.TODO(), requestOptions.Name, metaV1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "data": "", "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 0, "data": deployment})
}

func DeleteDeployment(c *gin.Context) {
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

	// 删除deployment
	err = clientSet.AppsV1().Deployments(requestOptions.Namespace).Delete(context.TODO(), requestOptions.Name, metaV1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "操作成功"})
}
