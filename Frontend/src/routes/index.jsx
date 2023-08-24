import React, {lazy, Suspense} from 'react';
import { Spin } from 'antd';
import { Navigate } from 'react-router-dom';
import {
    HomeOutlined,
    UserOutlined,
    SettingOutlined,
    BarsOutlined,
    FileTextOutlined
} from '@ant-design/icons'

const Loading = (element) => {
    return (
        <Suspense fallback={<Spin />}>
            {element}
        </Suspense>
    )
}

const Layout = lazy(() => import('../pages/layout'));
const Home = lazy(() => import('../pages/home'));
const Login = lazy(() => import('../pages/login'));
const User = lazy(() => import('../pages/user'));
const ErrorPage = lazy(() => import('../pages/404'));
const IndustryProject = lazy(() => import('../pages/project/industry'));
const GovernmentProject = lazy(() => import('../pages/project/government'));
const AITranslate = lazy(() => import('../pages/ai_translate'));
const Post = lazy(() => import('../pages/post'));



export const mainRoutes = {
    path: '/',
    element: <Layout />,
    children: [
        {
            path: 'home',
            label: '首页',
            element: Loading(<Home />),
            icon: <HomeOutlined />
        },
        {
            path: 'user',
            label: '用户管理',
            element: Loading(<User />),
            icon: <UserOutlined />
        },
        {
            path: 'project',
            label: '项目管理',
            icon: <BarsOutlined />,
            children: [
                {
                    path: 'industry',
                    label: '产业项目',
                    element: Loading(<IndustryProject />),
                    icon: <SettingOutlined />
                },
                {
                    path: 'government',
                    label: '政府项目',
                    element: Loading(<GovernmentProject />),
                    icon: <SettingOutlined />
                }
            ]

        },
        {
            path: 'ai_translate',
            label: 'AI翻译',
            element: Loading(<AITranslate />),
            icon: <SettingOutlined />
        },
        {
            path: 'post',
            label: '文章管理',
            element: Loading(<Post />),
            icon: <FileTextOutlined />
        }
    ]
}


export const routes = [
    {
        path: '/login',
        element: <Login />
    },
    mainRoutes,
    // 重定向
    {
        path: '/',
        element: <Navigate to="/login" />
    },
    // 404
    {
        path: '*',
        element: <ErrorPage />
    }
]