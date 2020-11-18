import {call, put, takeEvery} from 'redux-saga/effects'

const PUT_DATA = '/home/PUT_DATA'
const LOAD_DATA = '/home/LOAD_DATA'

const initialState = {
    data: []
}

export const homeReducer = (state = initialState, action) => {
    switch (action.type) {
        case PUT_DATA:
            return {
                ...state,
                data: action.data
            }
        default:
            return state
    }
}

export const putData = (data) => ({type: PUT_DATA, data})
export const loadData = () => ({type: LOAD_DATA})

export const getBalancers = state => state.homeReducer.data

function* workerLoadData() {
    const data = yield call(
        () => fetch('http://localhost:3001/getbalancers').then(res => res.json())
    )
    yield put(putData(data))
}

export function* watchLoadData() {
    yield takeEvery(LOAD_DATA, workerLoadData)
}