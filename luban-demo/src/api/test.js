import request from "../plugin/utils/request"

export const Test = (params) => request('get','/test/', params)