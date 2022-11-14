package k8s

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
)

func CreateCluster(c *gin.Context) {
	var cluster model.K8SCluster
	if err := c.ShouldBindJSON(cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 创建clientSet
	clientSet, err := GetClientSet(cluster.ClusterName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 获取node数量
	nodeList, err := clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 插入数据
	cluster.NodeNumber = len(nodeList.Items)
	if err := options.DB.Model(&model.K8SCluster{}).Create(&cluster).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "创建集群成功",
	})
}

func ListCluster(c *gin.Context) {
	var cluster []model.K8SCluster
	if err := options.DB.Model(&model.K8SCluster{}).Find(&cluster).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": cluster,
		"msg":  "创建集群成功",
	})
}

func DeleteCluster(c *gin.Context) {
	var cluster model.K8SCluster
	if err := c.ShouldBindJSON(cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	if err := options.DB.Model(&model.K8SCluster{}).Where("id = ?", cluster.Id).Delete(&cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "操作成功",
	})

}

func UpdateCluster(c *gin.Context) {
	var cluster model.K8SCluster
	if err := c.ShouldBindJSON(cluster); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 只更新集群kubeconfig字段
	if err := options.DB.Model(&model.K8SCluster{}).Where("id = ?", cluster.Id).Updates(
		map[string]interface{}{
			"kube_config": cluster.KubeConfig,
		}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": "",
		"msg":  "操作成功",
	})
}
