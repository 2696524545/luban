import request from "../plugin/utils/request"

const listPodsUrl = (clusterName, namespace) => {
    return `http://localhost:19999/api/v1/kubernetes/${clusterName}/${namespace}/pods`
}

export function ListPods(clusterName, namespace, params) {
    let url = listPodsUrl(clusterName, namespace)

    return request("get", url, params)
}

/*
    调用接口示例
    listPods(clusterName, namespace) {
       clusterName = "test"
      namespace = "default"
      ListPods(clusterName, namespace).then(res => {
        console.log(res, "pods all")
      })
    },

 */