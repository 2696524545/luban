package k8s

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/dnsjia/luban/api/types"
)

func GetPodLog(c *gin.Context) {
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

	podLogs, err := clientSet.CoreV1().Pods(requestOptions.Namespace).GetLogs(requestOptions.Name, &v1.PodLogOptions{}).Stream(context.TODO())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "操作成功", "data": buf.String()})

}

// DownloadContainerLog  /api/v1/kubernetes/test/kube-system/pods/logfile/nginx-ingress-controller-798b5df965-g7zsq
// Query params ?container=init-sysctl&previous=false
// /api/v1/kubernetes/test/kube-system/pods/logfile/nginx-ingress-controller-798b5df965-g7zsq?container=init-sysctl&previous=false
func DownloadContainerLog(c *gin.Context) {
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

	opts := new(v1.PodLogOptions)
	containerID := c.Query("container")
	// previous 为true表示获取容器之前退出的日志
	opts.Previous = c.Query("previous") == "true"
	// timestamps为true表示日志中添加详细的时间
	opts.Timestamps = c.Query("timestamps") == "true"
	logStream, err := GetLogFile(clientSet, requestOptions.Namespace, requestOptions.Name, containerID, opts)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

	// 添加context-type为text/plain
	c.Writer.Header().Add("Content-Type", "text/plain")
	defer logStream.Close()
	_, err = io.Copy(c.Writer, logStream)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error(), "data": ""})
		return
	}

}

func GetLogFile(client kubernetes.Interface, namespace, podID string, container string, opts *v1.PodLogOptions) (io.ReadCloser, error) {
	logOptions := &v1.PodLogOptions{
		Container:  container,
		Follow:     false,
		Previous:   opts.Previous,
		Timestamps: opts.Timestamps,
	}
	logStream, err := openStream(client, namespace, podID, logOptions)
	return logStream, err
}

func openStream(clientSet kubernetes.Interface, namespace, podId string, logOptions *v1.PodLogOptions) (io.ReadCloser, error) {
	return clientSet.CoreV1().RESTClient().Get().
		Namespace(namespace).
		Name(podId).
		Resource("pods").
		SubResource("log").
		VersionedParams(logOptions, scheme.ParameterCodec).
		Stream(context.TODO())
}
