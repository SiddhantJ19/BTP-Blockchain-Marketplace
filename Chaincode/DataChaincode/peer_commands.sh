cd ~/HLFabric/fabric-samples/test-network

# ORG1
    export PATH=${PWD}/../bin:$PATH
    export FABRIC_CFG_PATH=$PWD/../config/
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051

    # create a device
    export DEVICE=$(echo -n "{\"deviceId\":\"dev001\",\"dataDescription\":\"random data\", \"description\":\"Device not on sale\", \"deviceSecret\":\"--secret--\"}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"CreateDevice","Args":[]}' --transient "{\"_Device\":\"$DEVICE\"}"

    # updateDevieDetails
    export DEVICE=$(echo -n "{\"deviceId\":\"dev001\",\"description\":\"Device is on sale\", \"on_sale\":\"true\"}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"UpdateDeviceDetails","Args":[]}' --transient "{\"_Device\":\"$DEVICE\"}"

    # owner agreetosell
    export TRADE=$(echo -n "{\"tradeId\":\"trade001\",\"tradePrice\":100}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AgreeToSell","Args":["dev001"]}' --transient "{\"_TradeAgreement\":\"$TRADE\"}"


# ORG2

    export PATH=${PWD}/../bin:$PATH
    export FABRIC_CFG_PATH=$PWD/../config/
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051

    # buyer agreetobuy
    export TRADE=$(echo -n "{\"tradeId\":\"trade001\",\"tradePrice\":100}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AgreeToBuy","Args":["dev001"]}' --transient "{\"_TradeAgreement\":\"$TRADE\"}"

    export INTEREST_TOKEN=$(echo -n "{\"tradeId\":\"trade001\",\"deviceId\":dev001}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AgreeToBuy","Args":["dev001"]}' --transient "{\"_InterestToken\":\"$INTEREST_TOKEN\"}"

TODO
contract - 
    data add
    trade <===> (transfer-asset)

application
    Marketplace
        - api + UI
    data
        - datalistner 
        - datatransmitter

