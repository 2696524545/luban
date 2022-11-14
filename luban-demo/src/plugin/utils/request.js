import axios from 'axios'
import {Message} from 'element-ui'
import store from '../store'
// 默认请求头信息
axios.defaults.headers.post['Content-Type'] = 'application/json;charset=UTF-8';
axios.defaults.headers.put['Content-Type'] = 'application/json;charset=UTF-8';

// 创建axios实例
const instance = axios.create({
    baseURL: process.env.BASE_URL,
    timeout: 3000,
})


// request 拦截器
instance.interceptors.request.use(config => {
    if (store.getters.token) {
        config.headers['token'] = 'jwt ' + store.getters.token
    }
    return config
    },
    error => {
        Promise.reject(error)
})

// 响应拦截器
instance.interceptors.response.use(
    response => {
        if (response.status === 200) {
            return Promise.resolve(response)
        } else {
            Message({
                message: response.msg,
                type: 'error',
                duration: 5 * 1000,
            })
            return Promise.reject(response)
        }
    },
    error => {
        const {response} = error
        if (response) {
            Message({
                message: response.msg,
                type: 'error',
                duration: 5 * 1000
            })
            return Promise.reject(response)
        }else {
            Message({
                message: "系统异常",
                type: 'error',
                duration: 5 * 1000
            })
            return Promise.reject(error)
        }
    }
)


export default function (method, url, data = null) {
    method = method.toLowerCase()
    if (method === 'post') {
        return instance.post(url, data)
    } else if (method === 'get') {
        return instance.get(url, { params: data })
    } else if (method === 'delete') {
        return instance.delete(url, { params: data })
    } else if (method === 'put') {
        return instance.put(url, data)
    } else {
        console.error("未知的方法:" + method)
        return false
    }
}

