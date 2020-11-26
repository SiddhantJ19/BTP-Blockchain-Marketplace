const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');
const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require('../base/CAUtil');
const { buildCCPOrg1, buildWallet } = require('../base/AppUtil');
const config = require('../config/base')

const channelName = config.channelName;
const chaincodeName = config.chaincodeName;
const msp = config.msp

exports.enrollAdmin = async (req, res) => {
    await enrollAdmin(config.caClient, config.wallet, msp);
    res.status(200).send({"status":"Admin enrolled successfully"})
}

exports.registerUser = async (req, res) => {
    const userName = req.body.userName
    await registerAndEnrollUser(config.caClient, config.wallet, msp, userName, 'org1.department1');
    config.curUser = userName
    res.status(200).send({"status":"User enrolled successfully", "userName": userName})
}

exports.connectGateway = async (req, res) => {
    const userName = req.body.userName;
    config.gateway = new Gateway();
    const wallet = config.wallet
    await config.gateway.connect(config.ccp, {
        wallet,
        identity: userName,
        discovery: { enabled: true, asLocalhost: true } // using asLocalhost as this gateway is using a fabric network deployed locally
    });
    config.network = await config.gateway.getNetwork(channelName);
    config.contract = config.network.getContract(chaincodeName);
    config.contract.addDiscoveryInterest({name: 'mychaincode', collectionNames: ['Org2MSP_privateDetailsCollection']});
    res.status(200).send({"status":"Gateway Connected", "userName": userName})

}

