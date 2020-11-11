'use strict';

const {Gateway, Wallets} = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');

const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require('../../../test-application/javascript/CAUtil');
const { buildCCPOrg1, buildWallet } = require('../../../test-application/javascript/AppUtil');

const constants = {
    CHANNEL_NAME: 'mychannel',
    CHAINCODE_NAME: '--------', // package name ?
    MSP_ORG1: 'Org1MSP',
    WALLET_PATH: path.join(__dirname, 'wallet'),
    USER_ID: 'appUserOrg1'
};

async function main() {
    try {
        // network config aka connection profile
        const ccp = buildCCPOrg1();
        
    }
    catch {

    }
}