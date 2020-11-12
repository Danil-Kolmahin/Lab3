import React from 'react'
import {Progress, List, Avatar} from 'antd'
import {
    MenuOutlined,
    PoweroffOutlined,
    ReloadOutlined,
} from "@ant-design/icons"
import 'antd/dist/antd.css'
import imVM1 from '../assets/imVM1.png'

const listData = []
for (let i = 1; i < 23; i++) {
    listData.push({
        href: 'https://ant.design',
        title: `Virtual Machine №${i}`,
        avatar: imVM1,
        description: 'The user can initialize several machines that perform the same function and connect them to a' +
            'balancer that will evenly distribute requests to these machines.',
        workload: i * 10 % 100
    })
}

export const HomeContent = () => {
    return <div>
        <List
            itemLayout="vertical"
            size="large"
            pagination={{
                onChange: page => {
                    console.log(page)
                },
                pageSize: 3,
            }}
            dataSource={listData}
            renderItem={item => (
                <List.Item
                    key={item.title}
                    actions={[
                        <MenuOutlined/>,
                        <PoweroffOutlined/>,
                        <ReloadOutlined/>
                    ]}
                    extra={<Progress type="circle" percent={item.workload}/>}
                >
                    <List.Item.Meta
                        avatar={<Avatar src={item.avatar}/>}
                        title={<a href={item.href}>{item.title}</a>}
                        description={item.description}
                    />
                </List.Item>
            )}
        />
    </div>
}