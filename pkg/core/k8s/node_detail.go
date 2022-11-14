package k8s

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"

	"github.com/dnsjia/luban/api/types"
)

func GetNodeDetail(c *gin.Context) {
	var nameOptions types.NameOptions
	if err := c.ShouldBindUri(&nameOptions); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  err.Error(),
		})
		return
	}

	// 3. 通过kube config生成clientSet
	clientSet, err := NewClientSet(nameOptions.Cluster)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	node, err := clientSet.CoreV1().Nodes().Get(context.TODO(), nameOptions.Name, metaV1.GetOptions{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "",
			"msg":  "获取节点详情出错",
		})
		return
	}

	//  获取node节点上的pod
	pods, err := getNodePods(clientSet, *node)
	if err != nil {

	}

	// 获取节点详情，状态
	nodeDetail := toNodeDetail(*node, pods)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nodeDetail,
	})

}

func getNodePods(client *kubernetes.Clientset, node v1.Node) (*v1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + node.Name +
		",status.phase!=" + string(v1.PodSucceeded) +
		",status.phase!=" + string(v1.PodFailed))

	if err != nil {
		return nil, err
	}

	return client.CoreV1().Pods(v1.NamespaceAll).List(context.TODO(), metaV1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
}

func toNodeDetail(node v1.Node, pods *v1.PodList) types.NodeDetail {
	return types.NodeDetail{
		Node: types.Node{
			Name:       node.Name,
			Version:    node.Status.NodeInfo.KubeletVersion,
			Ready:      getNodeStatus(node, v1.NodeReady),
			NodeIP:     getNodeIP(node),
			Role:       getNodeRole(node),
			NodeInfo:   node.Status.NodeInfo,
			CreateAt:   node.CreationTimestamp.Format("2006-01-02 15:04:05"),
			ObjectMeta: node.ObjectMeta,
		},

		ProviderID:      node.Spec.ProviderID,
		PodCIDR:         node.Spec.PodCIDR,
		Unschedulable:   node.Spec.Unschedulable,
		NodeInfo:        node.Status.NodeInfo,
		Conditions:      getNodeConditions(node),
		PodList:         *pods,
		Taints:          node.Spec.Taints,
		Addresses:       node.Status.Addresses,
		Ready:           getNodeStatus(node, v1.NodeReady),
		NodeIP:          getNodeIP(node),
		UID:             string(node.UID),
		ContainerImages: getContainerImages(node),
	}
}

func getNodeConditions(node v1.Node) []types.Condition {
	var conditions []types.Condition
	for _, condition := range node.Status.Conditions {
		conditions = append(conditions, types.Condition{
			Type:               string(condition.Type),
			Status:             v1.ConditionStatus(condition.Status),
			LastProbeTime:      condition.LastHeartbeatTime,
			LastTransitionTime: condition.LastTransitionTime,
			Reason:             condition.Reason,
			Message:            condition.Message,
		})
	}
	return conditions
}

// getContainerImages 获取节点上所有的镜像
func getContainerImages(node v1.Node) []string {
	var containerImages []string
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			containerImages = append(containerImages, name)
		}
	}
	return containerImages
}
