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

exports.wishToBuy = async (req, res) => {
    // First Create agreetoBuy then create interest Token

    const tradeDetails = {
        'tradeId': req.body.tradeId,
        'deviceId': req.body.deviceId,
        'tradePrice': req.body.tradePrice,
        'seller_id' : req.body.seller_id,  
        'revoke_time': new Date(req.body.revoke_time * 1000)
    }
    console.log("tradeDetails\n", tradeDetails)
    if (!(tradeDetails.tradeId && tradeDetails.deviceId)) {
        return res.status(400).send({"status":"invalid input", "required_fields":"deviceId, description, dataDescription, onSale"})
    }


    // Agree to buy tx
    console.log('\n--> Submit Transaction: Agree to buy, ');
    let agreeToBuyTx = config.contract.createTransaction('AgreeToBuy')
    const ABTransientMapData = Buffer.from(JSON.stringify(tradeDetails));
    agreeToBuyTx.setTransient({
        _TradeAgreement: ABTransientMapData
    });
    const ABResult = await agreeToBuyTx.submit(tradeDetails.deviceId);
    console.log('*** Agree To Buy Result:');
    console.log(ABResult)


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