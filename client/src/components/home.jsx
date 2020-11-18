import React, {useState} from 'react'
import {HomeContent} from './homeContent'
import {HomeBalancerBar} from './homeBalancerBar'
import s from './home.module.css'
import {useSelector} from 'react-redux'
import {Spin} from 'antd'
import {getBalancers} from '../store/homeReducer'

export const Home = () => {
    const data = useSelector(getBalancers)
    const [watchId, setWatchId] = useState(0)
    return <div>
        {data[0] ?
            <div className={s.home}>
                <HomeContent machinesList={data[watchId].usedMachines}/>
                <HomeBalancerBar balancersList={data.map(val => val.id)}
                                 watchId={watchId} setWatchId={setWatchId}/>
            </div>
            : <Spin/>}
    </div>
}