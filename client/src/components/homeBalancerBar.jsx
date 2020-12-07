import React, {useState} from 'react'
import {Avatar, List, Spin} from 'antd'
import InfiniteScroll from 'react-infinite-scroller'
import 'antd/dist/antd.css'
import imBL1 from '../assets/imBL1.png'
import s from './home.module.css'

export const HomeBalancerBar = ({balancersList, watchId, setWatchId}) => {
    let messages = balancersList.reverse().map((val, i) => ({
        id: i,
        title: `Balancer №${val}`,
        description: 'Virtual Machine Management Support System and load balancers.'
    }))
    messages = messages.reverse()

    const [data, setData] = useState(messages)
    const [loading, setLoading] = useState(false)
    const [hasMore, setHasMore] = useState(true)

    const handleInfiniteOnLoad = () => {
        setLoading(true)
        if (data.length > 10) {
            setHasMore(false)
            setLoading(false)
            return
        }
        const newData = [] //getData()
        setData(newData.concat(data))
        setLoading(false)
    }
    return <div className={s.infiniteScroll}>
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
                    <List.Item key={item.id} className={item.id === watchId && s.active}
                    onClick={() => setWatchId(item.id)}>
                        <List.Item.Meta
                            avatar={<Avatar src={imBL1} alt={'!'}/>}
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