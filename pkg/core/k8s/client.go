package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
)

func NewClientSet(clusterName string) (*kubernetes.Clientset, error) {
	var k model.K8SCluster
	if err := options.DB.Table("").Where("cluster_name = ?", clusterName).First(&k).Error; err != nil {
		return nil, err
	}

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(k.KubeConfig))
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func GetClientSet(kubeConfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfig))
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}
