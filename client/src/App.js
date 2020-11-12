import React from 'react'
import {Router, Route, Switch} from 'react-router-dom'
import {createHashHistory} from 'history'
import {Home} from './components/home'

const history = createHashHistory()

const App = () => {
  const {location} = history
  return <Router history={history}>
    <Switch location={location}>
      <Route path="*" exact component={Home} />
    </Switch>
  </Router>
}

export default App
