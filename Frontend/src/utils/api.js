import request from "./request";

const register = (data = {}) => {
    return request({
        url: '/auth/register',
        method: 'post',
        data,
        contentType: 'application/json;charset=UTF-8'
    })
}

const login = (data = {}) => {
    return request({
        url: '/auth/login',
        method: 'post',
        data
    })
}

export {
    register,
    login
}