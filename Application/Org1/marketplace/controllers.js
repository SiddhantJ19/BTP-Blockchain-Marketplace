const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');
const config = require('../config/base')


function prettyJSONString(inputString) {
    return JSON.stringify(JSON.parse(inputString), null, 2);
}

exports.getOnSaleDevices = async (req, res) => {


    console.log('\n--> Submit Transaction: QueryOnSaleDataMarketplace');

    const txResult = await config.contract.evaluateTransaction('QueryOnSaleDataMarketplace');
    console.log(`*** Result: ${prettyJSONString(txResult.toString())}`);

    res.status(200).send({"status":"Query Successful", "data": JSON.parse(prettyJSONString(txResult.toString()))})
}