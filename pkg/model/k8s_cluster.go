package model

type K8SCluster struct {
	Id          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	ClusterName string `gorm:"comment:集群名称;index:idx_cluster_name,unique" binding:"required" json:"clusterName"`
	KubeConfig  string `gorm:"comment:集群凭证;type:text" binding:"required" json:"kubeConfig"`
	NodeNumber  int    `gorm:"comment:节点数" json:"nodeNumber" `
}

func (k K8SCluster) TableName() string {
	return "k8s_cluster"
}
