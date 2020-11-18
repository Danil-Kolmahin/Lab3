import React from 'react'
import {Progress, List, Avatar} from 'antd'
import {
    MenuOutlined,
    PoweroffOutlined,
    ReloadOutlined,
} from '@ant-design/icons'
import 'antd/dist/antd.css'
import imVM1 from '../assets/imVM1.png'

export const HomeContent = ({machinesList1, machinesList2, onPowerClick}) => {
    let listData = machinesList1.map(val => ({
        id: val,
        title: `Virtual Machine №${val}`,
        description: 'The user can initialize several machines that perform' +
            'the same function and connect them to a' +
            'balancer that will evenly distribute requests to these machines.',
        workload: 100
    }))
    listData = listData.concat(machinesList2.map(val => ({
        id: val,
        title: `Virtual Machine №${val}`,
        description: 'The user can initialize several machines that perform' +
            'the same function and connect them to a' +
            'balancer that will evenly distribute requests to these machines.',
        workload: 0
    })))
    return <div>
        <List
            itemLayout="vertical"
            size="large"
            pagination={{
                onChange: page => {
                    console.log(page)
                },
                pageSize: 10,
            }}
            dataSource={listData}
            renderItem={item => (
                <List.Item
                    key={item.title}
                    actions={[
                        <MenuOutlined/>,
                        <PoweroffOutlined spin={!item.workload} onClick={() => onPowerClick(item.id)}/>,
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