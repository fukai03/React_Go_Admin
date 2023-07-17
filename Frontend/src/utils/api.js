import request from "./request";
import {stringify} from "./index";

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

export {
    register,
    login
}