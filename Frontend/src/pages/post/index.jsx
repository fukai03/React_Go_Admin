import React, { useEffect, useState } from 'react';
import { getPostList } from 'utils/api';
import { Table } from 'antd';
import { set } from 'mobx';

export default function Post() {
    const [dataSource, setDataSource] = useState([]);
    const [pageSize, setPageSize] = useState(10);
    const [pageNum, setPageNum] = useState(1);
    const [total, setTotal] = useState(0);

    const getListData = async (params = {}) => {
        const res = await getPostList({
            pageSize: params.pageSize,
            pageNum: params.pageNum,
        });
        console.log(res);
        const data = res?.data?.data.map((i, index) => {
            return {
                ...i,
                key: index,
            };
        });
        setDataSource(data);
        setTotal(res?.data?.total);
    };

    const columns = [
        {
            title: '标题',
            dataIndex: 'title',
        },
        // {
        //     title: '图标',
        //     dateIndex: 'head_img',
        //     key: 'head_img',
        //     render: (text, record) => (
        //         <img src={record.head_img} alt="" style={{width: '100px'}}/>
        //     )
        // },
        {
            title: '内容',
            dataIndex: 'content',
        },
        {
            title: '创建时间',
            dataIndex: 'create_at',
        },
        {
            title: '更新时间',
            dataIndex: 'update_at',
        },
    ];

    useEffect(() => {
        console.log(pageSize, pageNum);
        getListData({
            pageSize,
            pageNum,
        });
    }, [pageSize, pageNum]);
    return (
        <div>
            <Table
                dataSource={dataSource}
                columns={columns}
                pagination={{
                    pageSize,
                    current: pageNum,
                    total,
                    onChange: (page, pageSize) => {
                        setPageNum(page);
                        setPageSize(pageSize);
                    }
                }}
            />
        </div>
    );
}
