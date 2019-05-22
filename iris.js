var jsonFilter = function(event) {
    console.log('jsonFilter')
    console.log(this)
}
var namespaceFilter = function() {
    console.log('namespaceFilter')
    console.log(this)
}

var buidIRISObject = function() {
    return JSON.stringify({
        filters: [
            {
                name: 'jsonFilter',
            },
            {
                name: 'namespaceFilter',
            },
        ],
        destinations: [{
            name: '',
            url: '',
        }],
        integrations: [{
            name: '',
            destinations: [],
            filters: [],
        }]
    })
}