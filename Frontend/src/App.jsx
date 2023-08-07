import { Suspense, useState } from 'react';
import './App.css';
import { Provider as MobxProvider } from 'mobx-react';
import { BrowserRouter as Router, useRoutes } from 'react-router-dom';
import { ConfigProvider, Spin } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import { routes } from './routes';
import store from './store';

function App() {
    const Element = () => useRoutes(routes);

    return (
        <>
            <ConfigProvider locale={zhCN}>
                <MobxProvider {...store}>
                    <Suspense fallback={<Spin />}>
                        <Router>
                            <Element />
                        </Router>
                    </Suspense>
                </MobxProvider>
            </ConfigProvider>
        </>
    );
}

export default App;
