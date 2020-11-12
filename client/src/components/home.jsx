import React from 'react'
import {HomeContent} from './homeContent'
import {HomeBalancerBar} from './homeBalancerBar'
import s from './home.module.css'

export const Home = () => {
    return <div className={s.home}>
        <HomeContent/>
        <HomeBalancerBar/>
    </div>
}