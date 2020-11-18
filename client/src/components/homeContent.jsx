import React from 'react'
import {Progress, List, Avatar} from 'antd'
import {
    MenuOutlined,
    PoweroffOutlined,
    ReloadOutlined,
} from '@ant-design/icons'
import 'antd/dist/antd.css'
import imVM1 from '../assets/imVM1.png'

export const HomeContent = ({machinesList}) => {
    const listData = machinesList.map(val => ({
        title: `Virtual Machine №${val}`,
        description: 'The user can initialize several machines that perform' +
            'the same function and connect them to a' +
            'balancer that will evenly distribute requests to these machines.',
        workload: 100
    }))
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
                        avatar={<Avatar src={imVM1}/>}
                        title={item.title}
                        description={item.description}
                    />
                </List.Item>
            )}
        />
    </div>
}