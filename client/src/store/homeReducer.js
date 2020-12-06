import {call, put, takeEvery} from 'redux-saga/effects'

const PUT_DATA = '/home/PUT_DATA'
const LOAD_DATA = '/home/LOAD_DATA'
const CHANGE_STATUS = '/home/CHANGE_STATUS'

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
        // case CHANGE_STATUS:
        //     console.log(action.id)
        //     return state
        default:
            return state
    }
}

export const putData = (data) => ({type: PUT_DATA, data})
export const loadData = () => ({type: LOAD_DATA})
export const changeStatus = (id) => ({type: CHANGE_STATUS, id})

export const getBalancers = state => state.homeReducer.data

function* workerLoadData() {
    const data = yield call(
        () => fetch('/getbalancers').then(res => res.json())
    )
    yield put(putData(data))
}

function* workerChangeStatus() {
    yield call(() => fetch('/status',
        {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            //mode: 'no-cors',
            //credentials : "include",
            body: JSON.stringify({
                isWork: true,
                machineId: 43
            })
        }
    ))
}

export function* watchLoadData() {
    yield takeEvery(LOAD_DATA, workerLoadData)
}

export function* watchChangeStatus() {
    yield takeEvery(CHANGE_STATUS, workerChangeStatus)
}