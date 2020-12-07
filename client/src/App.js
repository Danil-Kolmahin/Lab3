import React, {useEffect} from 'react'
import {Router, Route, Switch} from 'react-router-dom'
import {createHashHistory} from 'history'
import {Home} from './components/home'
import {useDispatch} from 'react-redux'
import {loadData} from './store/homeReducer'

const history = createHashHistory()

const App = () => {
    const dispatch = useDispatch()
    useEffect(() => {
        dispatch(loadData())
    }, [])
    const {location} = history
    return <Router history={history}>
        <Switch location={location}>
            <Route path="*" exact component={Home}/>
        </Switch>
    </Router>
}

export default App
