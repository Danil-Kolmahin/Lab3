const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        balancersInf: () => client.get('/balancers'),
        changeBalancersStatus: (machineId, isWork) => client.post('/balancers', {machineId, isWork})
    }

};

module.exports = { Client }