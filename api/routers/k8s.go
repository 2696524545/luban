package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/dnsjia/luban/pkg/core/k8s"
)

func KubernetesRouter(r *gin.RouterGroup) {
	k8sRouter := r.Group("kubernetes")
	{
		k8sRouter.POST("cluster", k8s.CreateCluster)
		k8sRouter.GET("clusters", k8s.ListCluster)
		k8sRouter.DELETE("cluster", k8s.DeleteCluster)
		k8sRouter.PUT("cluster", k8s.UpdateCluster)

		k8sRouter.GET("/:cluster/nodes", k8s.ListNodes)
		k8sRouter.GET("/:cluster/nodes/:name", k8s.GetNodeDetail)

		// deployment
		k8sRouter.GET("/:cluster/:namespace/deployments", k8s.ListDeployment)
		k8sRouter.GET("/:cluster/:namespace/deployments/:name", k8s.GetDeploymentDetail)
		k8sRouter.DELETE("/:cluster/:namespace/deployments/:name", k8s.DeleteDeployment)

		// daemonsets
		k8sRouter.GET("/:cluster/:namespace/daemonsets", k8s.ListDaemonsets)
		k8sRouter.GET("/:cluster/:namespace/daemonsets/:name", k8s.GetDaemonset)
		k8sRouter.DELETE("/:cluster/:namespace/daemonsets/:name", k8s.DeleteDaemonset)

		// statefulsets
		k8sRouter.GET("/:cluster/:namespace/statefulsets", k8s.ListStatefulSets)
		k8sRouter.GET("/:cluster/:namespace/statefulsets/:name", k8s.GetStatefulSet)
		k8sRouter.DELETE("/:cluster/:namespace/statefulsets/:name", k8s.DeleteStatefulSet)

		// job
		k8sRouter.GET("/:cluster/:namespace/jobs", k8s.ListJobs)
		k8sRouter.GET("/:cluster/:namespace/jobs/:name", k8s.GetJob)
		k8sRouter.DELETE("/:cluster/:namespace/jobs/:name", k8s.DeleteJob)

		// pod
		k8sRouter.GET("/:cluster/:namespace/pods", k8s.ListPods)
		k8sRouter.GET("/:cluster/:namespace/pods/:name", k8s.GetPod)
		k8sRouter.DELETE("/:cluster/:namespace/pods/:name", k8s.DeletePod)

		// pv pvc
		k8sRouter.GET("/:cluster/pvs", k8s.ListPv)
		k8sRouter.GET("/:cluster/pvs/:name", k8s.GetPv)

		k8sRouter.GET("/:cluster/:namespace/pvcs", k8s.ListPvc)
		k8sRouter.GET("/:cluster/:namespace/pvc/:name", k8s.GetPvc)

		// svc ingress
		k8sRouter.GET("/:cluster/:namespace/services", k8s.ListServices)
		k8sRouter.GET("/:cluster/:namespace/services/:name", k8s.GetService)

		k8sRouter.GET("/:cluster/:namespace/ingress", k8s.ListIngress)
		k8sRouter.GET("/:cluster/:namespace/ingress/:name", k8s.GetIngress)

		// pod logs
		k8sRouter.GET("/:cluster/:namespace/pods/log/:name", k8s.GetPodLog)
		k8sRouter.GET("/:cluster/:namespace/pods/logfile/:name", k8s.DownloadContainerLog)
		k8sRouter.GET("/:cluster/metric/cpu", k8s.GetCpu)

	}
}
