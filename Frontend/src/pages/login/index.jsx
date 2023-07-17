import React from "react";
import { Form, Input, Button, Checkbox, message } from "antd";
import { UserOutlined, LockOutlined } from "@ant-design/icons";
import "./index.less";
import { login } from 'utils/api'
import { useNavigate } from "react-router-dom";

export default () => {
    const [form] = Form.useForm();
    const navigate = useNavigate();

  const onFinish = async (values) => {
    console.log("Received values of form: ", values);
    const res = await login(values)
    console.log(res)
    message.success('登录成功');
    navigate('/home')

  };
  return (
    <div className="login-container">
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
              message: "请输入手机号码!",
            },
          ]}
        >
          <Input
            prefix={<UserOutlined className="site-form-item-icon" />}
            placeholder="Username"
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[
            {
              required: true,
              message: "请输入密码!",
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
        <Form.Item className="login-button">
          <Button
            className="login-form-button"
          >
            注册
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
