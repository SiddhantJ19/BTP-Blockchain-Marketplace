const {Gateway, Wallets} = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');
const {buildCAClient, registerAndEnrollUser, enrollAdmin} = require('../base/CAUtil');
const {buildCCPOrg1, buildWallet} = require('../base/AppUtil');
const config = require('../config/base')
const fs = require('fs')

const fsPromises = fs.promises
const channelName = config.channelName;
const chaincodeName = config.chaincodeName;
const msp = config.msp

exports.saveUserIdentityInWallet = async (username, identityString) => {
    if (username === "admin") {
        throw "Can not overwrite admin identity"
    } else {
        await fsPromises.writeFile(path.join(config.walletPath, `${username}.id`), identityString)
    }
}

exports.readUserIdentity = async (username) => {
    if (username === "admin") {
        throw "Reading admin identity is not permitted"
    } else {
        const filePath = path.join(config.walletPath, `${username}.id`)
        if (fs.existsSync(filePath)) {
            return await fsPromises.readFile(filePath)
        }
        else {
            throw "This identity does not exist"
        }

    }
}
exports.connectGatewayForUser = async (username, certificate) => {
    const curGateway = new Gateway();
    const wallet = config.wallet
    await this.saveUserIdentityInWallet(username, certificate)

    await curGateway.connect(config.ccp, {
        wallet,
        identity: username,
        discovery: {enabled: true, asLocalhost: true} // using asLocalhost as this gateway is using a fabric network deployed locally
    });

    return curGateway
}

exports.getContractForUser = async (username, certificate) => {
    const gateway = await this.connectGatewayForUser(username, certificate)


    const network = await gateway.getNetwork(channelName);
    const contract = network.getContract(chaincodeName);
    contract.addDiscoveryInterest({name: 'mychaincode', collectionNames: ['Org2MSP_aclCollection']});

    return contract
}

exports.createNewuser = async (username) => {
    await registerAndEnrollUser(config.caClient, config.wallet, msp, username, 'org1.department1');
}
