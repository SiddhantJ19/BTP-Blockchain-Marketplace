const path = require('path');

let ccp = {};
let caClient = {}
let walletPath = ""
let wallet = {}
let gateway = {}
let network = {}
let contract = {}
let curUser = {}
let channelName = "mychannel"
let chaincodeName = "mychaincode"
let msp = "Org2MSP";

module.exports = {
    ccp: ccp,
    caClient : caClient,
    walletPath: walletPath,
    wallet : wallet,
    gateway: gateway,
    network: network,
    contract: contract,
    curUser: curUser,
    channelName: channelName,
    chaincodeName:chaincodeName,
    msp:msp
}
