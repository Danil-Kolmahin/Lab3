const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        balancersInf: () => client.get('/getbalancers'),
        changeBalancersStatus: (machineId, isWork) => client.post('/status', {machineId, isWork})
    }

};

module.exports = { Client }