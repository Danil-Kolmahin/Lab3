import {applyMiddleware, createStore, combineReducers, compose} from 'redux'
import createSagaMiddleware from 'redux-saga'
import {homeReducer, watchLoadData} from './homeReducer'

let reducers = combineReducers({
    homeReducer
})
const sagaMiddleware = createSagaMiddleware()
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose
const store = createStore(reducers, composeEnhancers(applyMiddleware(sagaMiddleware)))
sagaMiddleware.run(watchLoadData)

export default store