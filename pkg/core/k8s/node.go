package k8s

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/dnsjia/luban/api/types"
)

func ListNodes(c *gin.Context) {
	var clusterOptions types.ClusterOptions
	// 1. 绑定参数
	if err := c.ShouldBindUri(&clusterOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 2. 通过kube config生成clientSet
	clientSet, err := NewClientSet(clusterOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	// 3. 获取node
	nodeList, err := clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {

	}

	// 4. 返回数据
	var nodes []types.Node
	for _, item := range nodeList.Items {
		nodes = append(nodes, toNodeList(item))
	}

	// 定义匿名结构体
	data := struct {
		Total int         `json:"total"`
		Nodes interface{} `json:"nodes"`
	}{}

	data.Total = len(nodes)
	data.Nodes = nodes
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": data,
	})

}

func toNodeList(node v1.Node) types.Node {
	return types.Node{
		Name:     node.Name,
		Version:  node.Status.NodeInfo.KubeletVersion,
		Ready:    getNodeStatus(node, v1.NodeReady),
		NodeIP:   getNodeIP(node),
		Role:     getNodeRole(node),
		NodeInfo: node.Status.NodeInfo,
		CreateAt: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
	}
}

// getNodeStatus 获取节点状态
func getNodeStatus(node v1.Node, conditionType v1.NodeConditionType) v1.ConditionStatus {
	for _, condition := range node.Status.Conditions {
		if condition.Type == conditionType {
			return condition.Status
		}
	}
	return v1.ConditionUnknown
}

// getNodeIP 获取节点IP
func getNodeIP(node v1.Node) string {
	for _, addr := range node.Status.Addresses {
		if addr.Type == v1.NodeInternalIP {
			return addr.Address
		}
	}
	return ""
}

// getNodeRole 获取节点角色
func getNodeRole(node v1.Node) string {
	var role string
	// 新版本label可能是 control-plane
	if _, ok := node.ObjectMeta.Labels["node-role.kubernetes.io/master"]; ok {
		role = "Master"
	} else {
		role = "Worker"
	}
	return role
}
