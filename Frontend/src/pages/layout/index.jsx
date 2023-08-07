import React, {useState, useEffect} from 'react';
import { Outlet, Link, useNavigate} from 'react-router-dom';
import { mainRoutes } from '../../routes';
import { Layout, Menu, theme, Avatar, Dropdown, Button, Popover } from 'antd';
import { UserOutlined } from '@ant-design/icons';
import './index.less';
import { debounce, removeToken } from 'utils';
import { getUserInfo } from 'utils/api';

const { Header, Content, Footer, Sider } = Layout;
const { SubMenu } = Menu;

export default function Index() {
    const [userOptionsVisible, setUserOptionsVisible] = useState(true);
    const [firstLetter, setFirstLetter] = useState('');
    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const navigate = useNavigate();

    const getUser = async () => {
        const res = await getUserInfo();
        console.log(res);
        setFirstLetter(res?.data?.user?.name?.slice(0, 1));
    };

    useEffect(() => {
        getUser()
    }, [])

    const createMenu = (routes, path = '/') => {
        return routes.map((route) => {
            if (route.children) {
                return {
                    key: route.label,
                    icon: route.icon,
                    children: createMenu(route.children, route.path),
                    label: route.label,
                };
            } else if (route.label) {
                return {
                    key: route.label,
                    icon: route.icon,
                    label: (
                        <Link to={path === '/' ? route.path : `${path}/${route.path}`}>
                            {route.label}
                        </Link>
                    ),
                };
            }
        });
    };
    const logout = () => {
        console.log('logout');
    };
    const userInfoClick = () => {
        console.log('userInfoClick');
    };
    const logoutClick = () => {
        console.log('logoutClick');
        removeToken();
        navigate('/login');
    };
    const userOptionsItems = (
        <div className="userOption">
            <div className="userOption-item" onClick={userInfoClick}>用户中心</div>
            <div className="userOption-item" onClick={logoutClick}>退出登录</div>
        </div>
    )


    const userOptionsOver = debounce(() => {
        console.log('userOptionsOver');
        setUserOptionsVisible(true);
    }, 200);
    const userOptionsLeave = debounce(() => {
        console.log('userOptionsLeave');
        setUserOptionsVisible(false);
    }, 200);

    return (
        <Layout className="custom-layout">
            <Sider>
                <div className="logo-vertical">React-Go-Admin</div>
                <Menu
                    theme="dark"
                    mode="inline"
                    defaultSelectedKeys={[mainRoutes.children[0].label]}
                    items={createMenu(mainRoutes.children)}
                />
            </Sider>
            <Layout>
                <Header
                    style={{
                        background: colorBgContainer,
                    }}
                    className="custom-header"
                >
                    <Popover
                        content={userOptionsItems}
                    >
                        {
                            firstLetter ?
                                <Avatar>{firstLetter}</Avatar>
                                :
                                <Avatar
                                    icon={<UserOutlined />}
                                />
                        }
                    </Popover>
                </Header>
                <Content
                    style={{
                        margin: '24px 16px 16px',
                    }}
                >
                    <div
                        style={{
                            padding: 24,
                            minHeight: 360,
                            background: colorBgContainer,
                            height: '100%',
                        }}
                    >
                        <Outlet />
                    </div>
                </Content>
            </Layout>
        </Layout>
    );
}
