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

exports.getInterestTokensForDevice = async (req, res) => {

    const deviceId = req.body.deviceId
    console.log('\n--> Submit Transaction: QueryInterestTokensForDevice');

    const txResult = await config.contract.evaluateTransaction('QueryInterestTokensForDevice', deviceId );
    console.log(`*** Result: ${prettyJSONString(txResult.toString())}`);

    res.status(200).send({"status":"Query Successful", "data": JSON.parse(prettyJSONString(txResult.toString()))})
}

exports.submitInterestToken = async (req, res) => {
    const tradeDetails = {
        'tradeId': req.body.tradeId,
        'deviceId': req.body.deviceId,
    }

    if (!(tradeDetails.tradeId && tradeDetails.deviceId)) {
        return res.status(400).send({"status":"invalid input", "required_fields":"deviceId, description, dataDescription, onSale"})
    }

    console.log('\n--> Submit Transaction: CreateInterestToken, ');
    let tradeTx = config.contract.createTransaction('CreateInterestToken')
    const transientMapData = Buffer.from(JSON.stringify(tradeDetails));
    tradeTx.setTransient({
        _InterestToken: transientMapData
    });

    const result = await tradeTx.submit();
    console.log('*** Result:');
    console.log(result)


    console.log('\n--> Submit Transaction: GetInterestToken');
    const txResult = await config.contract.evaluateTransaction('GetInterestToken', tradeDetails.tradeId);
    console.log(`*** Result: ${prettyJSONString(txResult.toString())}`);

    res.status(200).send({"status":"Interest Token Created", "data": JSON.parse(prettyJSONString(txResult.toString()))})

}