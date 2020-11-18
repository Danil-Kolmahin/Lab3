import React, {useState} from 'react'
import {HomeContent} from './homeContent'
import {HomeBalancerBar} from './homeBalancerBar'
import s from './home.module.css'
import {useDispatch, useSelector} from 'react-redux'
import {Spin} from 'antd'
import {changeStatus, getBalancers} from '../store/homeReducer'

export const Home = () => {
    const data = useSelector(getBalancers)
    const dispatch = useDispatch()
    const [watchId, setWatchId] = useState(0)
    return <div>
        {data[0] ?
            <div className={s.home}>
                <HomeContent machinesList1={data[watchId].usedMachines}
                             machinesList2={data[watchId].notUsedMachines ?
                                 data[watchId].notUsedMachines : []}
                             onPowerClick={(id) => dispatch(changeStatus(id))}/>
                <HomeBalancerBar balancersList={data.map(val => val.id)}
                                 watchId={watchId} setWatchId={setWatchId}/>
            </div>
            : <Spin/>}
    </div>
}