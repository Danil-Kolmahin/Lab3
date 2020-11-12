import React, {useEffect, useState} from 'react'
import {Avatar, List, Spin} from 'antd'
import InfiniteScroll from 'react-infinite-scroller'
import 'antd/dist/antd.css'
import imBL1 from '../assets/imBL1.png'

const messages = []
for (let i = 1; i < 5; i++) {
    messages.push({
        title: `Balancer №${i}`,
        avatar: imBL1,
        description: 'Virtual Machine Management Support System and load balancers.'
    })
}
const watchId = 1

export const HomeBalancerBar = () => {
    const [data, setData] = useState(messages)
    const [loading, setLoading] = useState(false)
    const [hasMore, setHasMore] = useState(true)

    const handleInfiniteOnLoad = () => {
        setLoading(true)
        if (data.length > 30) {
            setHasMore(false)
            setLoading(false)
            return
        }
        const newData = [] //getData()
        setData(newData.concat(data))
        setLoading(false)
    }

    useEffect(() => {
        setData(messages)
    }, [watchId, messages])

    return <div>
        <InfiniteScroll
            initialLoad={false}
            pageStart={0}
            loadMore={handleInfiniteOnLoad}
            hasMore={!loading && hasMore}
            useWindow={false} // isReverse={true}
        >
            <List
                dataSource={data}
                renderItem={item => (
                    <List.Item key={item.title}>
                        <List.Item.Meta
                            avatar={<Avatar src={item.avatar} alt={'!'}/>}
                            title={item.title}
                            description={item.description}
                        />
                    </List.Item>
                )}
            >
                {loading && hasMore && <Spin/>}
            </List>
        </InfiniteScroll>
    </div>
}