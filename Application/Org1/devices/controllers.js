const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');
const config = require('../config/base')


function prettyJSONString(inputString) {
	return JSON.stringify(JSON.parse(inputString), null, 2);
}

exports.registerDevice = async (req, res) => {
    const assetDetails = {
        'deviceId': req.body.deviceId,
        'description': req.body.description,
        'dataDescription': req.body.dataDescription,
        'deviceSecret': req.body.deviceSecret
    }

    if (!(assetDetails.deviceId && assetDetails.description && assetDetails.dataDescription && assetDetails.deviceSecret)) {
        return res.status(400).send({"status":"invalid input", "required_fields":"deviceId, description, dataDescription, secret"})
    }
    console.log('\n--> Submit Transaction: RegisterDevice, Initialize Device Details');
    let createDeviceTxn = config.contract.createTransaction('CreateDevice')
    const transientMapData = Buffer.from(JSON.stringify(assetDetails));
    createDeviceTxn.setTransient({
        _Device: transientMapData
    });

    const result = await createDeviceTxn.submit();
    console.log('*** Result:');
    console.log(result.toString())

    console.log('\n--> Submit Transaction: GetDeviceDetails');
    const txResult = await config.contract.evaluateTransaction('GetDeviceDetails', assetDetails.deviceId);
    console.log(`*** Result: ${prettyJSONString(txResult.toString())}`);

    res.status(200).send({"status":"Device Registered", "data": JSON.parse(prettyJSONString(txResult.toString()))})
}

exports.updateDevice = async (req, res) => {
    const assetDetails = {
        'deviceId': req.body.deviceId,
        'description': req.body.description,
        // 'dataDescription': req.body.dataDescription,
        'on_sale': req.body.on_sale
    }
    //
    // if (!(assetDetails.deviceId && assetDetails.description && assetDetails.dataDescription && assetDetails.on_sale)) {
    //     return res.status(400).send({"status":"invalid input", "required_fields":"deviceId, description, dataDescription, onSale"})
    // }
    console.log('\n--> Submit Transaction: UpdateDeviceDetails, ');
    let createDeviceTxn = config.contract.createTransaction('UpdateDeviceDetails')
    const transientMapData = Buffer.from(JSON.stringify(assetDetails));
    createDeviceTxn.setTransient({
        _Device: transientMapData
    });

    const result = await createDeviceTxn.submit();
    console.log('*** Result:');
    console.log(result)

    console.log('\n--> Submit Transaction: GetDeviceDetails');
    const txResult = await config.contract.evaluateTransaction('GetDeviceDetails', assetDetails.deviceId);
    console.log(`*** Result: ${prettyJSONString(txResult.toString())}`);

    res.status(200).send({"status":"Device Updated", "data": JSON.parse(prettyJSONString(txResult.toString()))})
}

/*
exports.updateDeviceData = async (req, res) => {
    const data = req.body.data
    const deviceID = req.body.deviceID

    console.log('\n--> Submit Transaction: Update Data');
    await config.contract.submitTransaction('UpdateData', data, deviceID);
    console.log('*** Result: committed');
    
    result = await config.contract.evaluateTransaction('GetDeviceData', deviceID);
    console.log(`*** Result: ${prettyJSONString(result.toString())}`);

    res.status(200).send({"status":"Data Updated", "data": JSON.parse(prettyJSONString(result.toString()))})
}

exports.deleteDeviceData = async (req, res) => {
    const deviceID = req.body.deviceID
    console.log('\n--> Submit Transaction: DeleteDevice');
    await config.contract.submitTransaction('DeleteDevice', deviceID);
    console.log('*** Result: committed');
    res.status(200).send({"status":"Device Deleted"})
}

exports.getDeviceData = async (req, res) => {
    
    const deviceID = req.body.deviceID
    console.log('\n--> Submit Transaction: Get Data');
    result = await config.contract.evaluateTransaction('GetDeviceData', deviceID);
    console.log(`*** Result: ${prettyJSONString(result.toString())}`);
            
    res.status(200).send({"status":"Device Registered", "data": JSON.parse(prettyJSONString(result.toString()))})
}

exports.getAllDevices = async (req, res) => {
    console.log('\n--> Submit Transaction: GetAllDevices');
    result = await config.contract.evaluateTransaction('GetAllDevices');
    console.log(`*** Result: ${prettyJSONString(result.toString())}`);
    res.status(200).send({"status":"All Devices", "data": JSON.parse(prettyJSONString(result.toString()))})
}


exports.getHistoricalValues = async (req, res) => {
    const deviceID = req.body.deviceID
    console.log('\n--> Submit Transaction: GetHistoricalValues');
    result = await config.contract.evaluateTransaction('GetHistory', deviceID);
    console.log(result)
    res.status(200).send({"status":"History", "data": JSON.parse(prettyJSONString(result.toString()))})
}

*/
