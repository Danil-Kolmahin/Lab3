import {call, put, takeEvery, take} from 'redux-saga/effects'

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
export const changeStatus = (id, changeOn) => ({type: CHANGE_STATUS, id, changeOn})

export const getBalancers = state => {
    const res = []
    state.homeReducer.data.forEach((currentValue) => {
        const timeRes = {id: currentValue.id, machines: []}
        let a = currentValue.usedMachines
        !a && (a = [])
        a.forEach((currentValue) => {
            timeRes.machines.push({
                isUsed: true,
                id: currentValue
            })
        })
        let b = currentValue.notUsedMachines
        !b && (b = [])
        b.forEach((currentValue) => {
            timeRes.machines.push({
                isUsed: false,
                id: currentValue
            })
        })
        res.push(timeRes)
    })
    return res
}

function* workerLoadData() {
    const data = yield call(
        () => fetch('/getbalancers').then(res => res.json())
    )
    yield put(putData(data))
}

function* workerChangeStatus() {
    const action = yield take(CHANGE_STATUS)
    console.log(action.id)
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
                isWork: action.changeOn,
                machineId: action.id
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