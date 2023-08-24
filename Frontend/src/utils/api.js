import request from './request';
import {stringify} from './index';

const register = (data = {}) => {
    return request({
        url: '/auth/register',
        method: 'post',
        data: stringify(data),
        contentType: 'application/json;charset=UTF-8'
    })
}

const login = (data = {}) => {
    return request({
        url: '/auth/login',
        method: 'post',
        data: stringify(data)
    })
}

// 获取用户信息
const getUserInfo = () => {
    return request({
        url: '/auth/info',
        method: 'get',
    })
}

// 文章创建
const createPost = (data = {}) => {
    return request({
        url: '/post',
        method: 'post',
        data
    })
}

// 文章修改
const updatePost = (data = {}) => {
    const {id, ...rest} = data;
    return request({
        url: `/post/${id}`,
        method: 'put',
        data: rest
    })
}
// 文章查找
const getPost = (id) => {
    return request({
        url: `/post/${id}`,
        method: 'get',
    })
}

// 文章删除
const deletePost = (id) => {
    return request({
        url: `/post/${id}`,
        method: 'delete',
    })
}

// 文章列表
const getPostList = (params = {}) => {

    return request({
        url: `/post/list?pageNum=${params.pageNum}&pageSize=${params.pageSize}`,
        method: 'post',
    })
}

export {
    register,
    login,
    getUserInfo,
    createPost,
    updatePost,
    getPost,
    deletePost,
    getPostList
}