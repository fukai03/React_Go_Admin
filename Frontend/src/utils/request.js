import { message } from 'antd'
import axios from 'axios'

const request = axios.create({
    baseURL: '/api',
    timeout: 5000
})

// 请求拦截器
request.interceptors.request.use(config => {
    // 在发送请求之前做些什么
    return config
}, error => {
    // 对请求错误做些什么
    return Promise.reject(error)
})

// 响应拦截器
request.interceptors.response.use(response => {
    // 对响应数据做点什么
    console.log('response', response);
    return response.data || response
}, error => {
    // 对响应错误统一处理
    message.error(error.response.data.msg || '请求错误，请重试！')
    return Promise.reject(error)
})

export default request