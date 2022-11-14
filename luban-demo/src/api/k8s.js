import  request from '../plugin/utils/request'

export const ListNodes = (params) => request("get", "http://baidu.com/api/v1/k8s/clusterName/nodes", params)