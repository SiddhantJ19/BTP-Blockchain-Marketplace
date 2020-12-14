#!/bin/bash

TEST_NETWORK_DIR=~/HLFabric/fabric-samples/test-network
# CHAINCODE_DIR=~/HLFabric/fabric-samples/Blockchain-marketplace/Chaincode/marbles02_private/go
CHAINCODE_DIR=~/HLFabric/fabric-samples/Blockchain-marketplace/Chaincode/DataChaincode/go
COLLECTION_PATH=~/HLFabric/fabric-samples/Blockchain-marketplace/Chaincode/DataChaincode/collections.json


# shuts down the network
function networkDown {
    cd $TEST_NETWORK_DIR
    ./network.sh down
}

function extraSetup {
    cd $CHAINCODE_DIR
    GO111MODULE=on go mod vendor
    cd $TEST_NETWORK_DIR
docker rm -f $(docker ps -a | awk '($2 ~ /dev-peer.*.mychaincode.*/) {print $1}')
docker rmi -f $(docker images | awk '($1 ~ /dev-peer.*.mychaincode.*/) {print $3}')
}


function networkUp {
    # sets up orgs, peers, admins and channel
    cd $TEST_NETWORK_DIR
    ./network.sh up createChannel -s couchdb -ca
}

function setOrg1Vars {
    cd $TEST_NETWORK_DIR
    export PATH=${PWD}/../bin:$PATH
    export FABRIC_CFG_PATH=$PWD/../config/
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
}

function setOrg2Vars {
    cd $TEST_NETWORK_DIR
    export PATH=${PWD}/../bin:$PATH
    export FABRIC_CFG_PATH=$PWD/../config/
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051
}

function packageChainCode {
    setOrg1Vars
    peer lifecycle chaincode package mychaincode.tar.gz --path $CHAINCODE_DIR --lang golang --label mychaincodev1
}

function installChaincodeOnBothOrgs {
    setOrg1Vars
    peer lifecycle chaincode install mychaincode.tar.gz
    setOrg2Vars
    peer lifecycle chaincode install mychaincode.tar.gz
}

function setInstalledPackageId {
    export CC_PACKAGE_ID=`peer lifecycle chaincode queryinstalled | awk 'NR > 1 {print $3}' | awk 'BEGIN{FS=","} {print $1}'`
}

function approveChainCode {
    peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name mychaincode --version 1.0 --collections-config $COLLECTION_PATH --signature-policy "OR('Org1MSP.member','Org2MSP.member')" --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile $ORDERER_CA
}

function approveChaincodeOnBothOrgs {
    cd $TEST_NETWORK_DIR
    export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
    
    setOrg1Vars
    setInstalledPackageId
    approveChainCode

    setOrg2Vars
    setInstalledPackageId
    approveChainCode
}

function commitChainCode {
    cd $TEST_NETWORK_DIR
    setOrg1Vars
    export ORG1_CA=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export ORG2_CA=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

    peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name mychaincode --version 1.0 --sequence 1 --collections-config $COLLECTION_PATH --signature-policy "OR('Org1MSP.member','Org2MSP.member')" --tls --cafile $ORDERER_CA --peerAddresses localhost:7051 --tlsRootCertFiles $ORG1_CA --peerAddresses localhost:9051 --tlsRootCertFiles $ORG2_CA
}


# Function Calls below

networkDown
extraSetup
networkUp
packageChainCode
installChaincodeOnBothOrgs
approveChaincodeOnBothOrgs
commitChainCode
