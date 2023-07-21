import { Suspense, useState } from 'react';
import './App.css';
import { register, login } from './utils/api.js';
import { stringify } from './utils/index.js';
import { Provider } from 'mobx-react';
import { BrowserRouter as Router, useRoutes } from 'react-router-dom';
import { ConfigProvider, Spin } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import { routes } from './routes';

function App() {
    const Element = () => useRoutes(routes);

    return (
        <>
            <ConfigProvider locale={zhCN}>
                <Suspense fallback={<Spin />}>
                    <Router>
                        <Element />
                    </Router>
                </Suspense>
            </ConfigProvider>
        </>
    );
}

export default App;
