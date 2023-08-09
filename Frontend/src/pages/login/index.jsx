import React, {useState} from 'react';
import { Form, Input, Button, Checkbox, message, Tabs } from 'antd';
import { UserOutlined, LockOutlined, PhoneOutlined } from '@ant-design/icons';
import './index.less';
import { login, register } from 'utils/api'
import { useNavigate } from 'react-router-dom';
import { useStores } from 'utils/hooks';
import { observer } from 'mobx-react';
import { setToken } from 'utils';

const Login = observer(() => {
    const [tabKey, setTabKey] = useState('1');
    const [form] = Form.useForm();
    const navigate = useNavigate();


    const onFinish = async (values) => {
        console.log('Received values of form: ', values);
        const res = await login(values)
        console.log(res)
        setToken(res?.data?.token || '');
        message.success('登录成功');
        navigate('/home')

    };
    const onRegisterFinish = async (values) => {
        console.log('Received values of form: ', values);
        const res = await register(values);
        console.log(res)
        message.success('注册成功');
        setTimeout(() => {
            setTabKey('1');
        }, 1000);
    }
    const tabChange = (key) => {
        console.log(key);
        setTabKey(key);
    }
    const LoginForm = () => (
        <Form
            name="normal_login"
            className="login-form"
            initialValues={{
                remember: true,
            }}
            onFinish={onFinish}
        >
            <Form.Item
                name="telephone"
                rules={[
                    {
                        required: true,
                        message: '请输入手机号码!',
                    },
                    {
                        // 手机号验证
                        pattern: /^(?:(?:\+|00)86)?1[3-9]\d{9}$/,
                        message: '请输入正确的手机号码!',
                    },
                ]}
            >
                <Input
                    prefix={<PhoneOutlined className="site-form-item-icon" />}
                    placeholder="手机号码"
                />
            </Form.Item>
            <Form.Item
                name="password"
                rules={[
                    {
                        required: true,
                        message: '请输入密码!',
                    },
                ]}
            >
                <Input
                    prefix={<LockOutlined className="site-form-item-icon" />}
                    type="password"
                    placeholder="Password"
                />
            </Form.Item>

            <Form.Item className="login-button">
                <Button
                    type="primary"
                    htmlType="submit"
                    className="login-form-button"
                >
            登录
                </Button>
            </Form.Item>
        </Form>
    )
    const RegisterForm = () => (
        <Form
            name="normal_register"
            className="login-form"
            initialValues={{
                remember: true,
            }}
            onFinish={onRegisterFinish}
        >
            <Form.Item
                name="name"
            >
                <Input
                    prefix={<UserOutlined className="site-form-item-icon" />}
                    placeholder="用户名（可选）"
                />
            </Form.Item>
            <Form.Item
                name="telephone"
                rules={[
                    {
                        required: true,
                        message: '请输入手机号码!',
                    },
                    {
                        // 手机号验证
                        pattern: /^(?:(?:\+|00)86)?1[3-9]\d{9}$/,
                        message: '请输入正确的手机号码!',
                    },
                ]}
            >
                <Input
                    prefix={<PhoneOutlined className="site-form-item-icon" />}
                    placeholder="请输入手机号"
                />
            </Form.Item>
            <Form.Item
                name="password"
                rules={[
                    {
                        required: true,
                        message: '请输入密码!',
                    },
                ]}
            >
                <Input
                    prefix={<LockOutlined className="site-form-item-icon" />}
                    type="password"
                    placeholder="请输入密码"
                />
            </Form.Item>
            <Form.Item className="login-button">
                <Button
                    className="login-form-button"
                    htmlType="submit"
                >
            注册
                </Button>
            </Form.Item>
        </Form>
    )
    const items = [
        {
            key: '1',
            label: '登录',
            children: <LoginForm />
        },
        {
            key: '2',
            label: '注册',
            children: <RegisterForm />
        }
    ]
    return (
        <div className="login-container">
            <div className='form_wrap'>
                <Tabs defaultActiveKey={tabKey} centered items={items} onChange={tabChange} activeKey={tabKey}/>
            </div>
        </div>
    );
})

export default Login;
