package k8s

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/api/types"
)

func ListPods(c *gin.Context) {
	var namespaceOptions types.NamespaceOptions
	if err := c.ShouldBindUri(&namespaceOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	var request Request
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	clientSet, err := NewClientSet(namespaceOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	podList, err := clientSet.CoreV1().Pods(namespaceOptions.Namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	// 序列化
	podData, _ := json.Marshal(podList)

	var listObj K8sListObj
	// 反序列化
	if err = json.Unmarshal(podData, &listObj); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	//分页
	PageResult, err := pagerAndSearch(request.Page, request.PageSize, listObj.Items, request.Search)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     PageResult.Items,
		"total":    PageResult.Total,
		"page":     PageResult.Page,
		"pageSize": PageResult.PageSize,
		"code":     0,
		"msg":      "操作成功",
	})

}

func GetPod(c *gin.Context) {
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

	pod, err := clientSet.CoreV1().Pods(requestOptions.Namespace).Get(context.TODO(), requestOptions.Name, metaV1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": pod})
}

func DeletePod(c *gin.Context) {
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

	err = clientSet.CoreV1().Pods(requestOptions.Namespace).Delete(context.TODO(), requestOptions.Name, metaV1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": nil})
}
