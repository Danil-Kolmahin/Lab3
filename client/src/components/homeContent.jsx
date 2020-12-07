import React from 'react'
import {Progress, List, Avatar} from 'antd'
import {
    MenuOutlined,
    PoweroffOutlined,
    ReloadOutlined,
} from '@ant-design/icons'
import 'antd/dist/antd.css'
import imVM1 from '../assets/imVM1.png'
import s from "./home.module.css";

export const HomeContent = ({machinesList, onPowerClick}) => {
    let listData = machinesList.map(val => ({
        id: val.id,
        title: `Virtual Machine №${val.id}`,
        description: 'The user can initialize several machines that perform' +
            'the same function and connect them to a' +
            'balancer that will evenly distribute requests to these machines.',
        workload: val.isUsed ? 100 : 0
    }))
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
                        <PoweroffOutlined
                            className={`${!item.workload ? s.activeButton : s.disactiveButton}
                             ${s.button}`}
                            onClick={
                            () => onPowerClick(item.id, !item.workload)
                        }/>,
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