const channels = require('./balancers/client')

const client = channels.Client('http://localhost:3001')

// Scenario 2: balancers list
client.balancersInf()
    .then(list => {
        console.log('=== Scenario 1 ===')
        console.log('Balancers:')
        list.forEach(b => console.log(b))
    })
    .catch(e => console.log(`Problem with list of balancers inf: ${e.message}`));

// Scenario 2: change balancer status
client.changeBalancersStatus(21, false)
    .then(resp => {
        console.log('=== Scenario 2 ===');
        console.log('Change balancer status:', resp);
        return client.balancersInf()
            .then(list => {
                list.forEach(b => console.log(b))
            })
    })
    .catch(e => console.log(`Problem with changing balancer status: ${e.message}`));