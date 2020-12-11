cd ~/HLFabric/fabric-samples/test-network

# ORG1
    export PATH=${PWD}/../bin:$PATH
    export FABRIC_CFG_PATH=$PWD/../config/
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051


    export tradeId=\"trade001\"
    export deviceId=\"device001\"
    export bidderId=\"Org2MSP\"

    # create a device
    export DEVICE=$(echo -n "{\"deviceId\":${deviceId},\"dataDescription\":\"random data\", \"description\":\"Device not on sale\", \"deviceSecret\":\"--secret--\"}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mychaincode -c '{"function":"CreateDevice","Args":[]}' --transient "{\"_Device\":\"$DEVICE\"}"

    # updateDevieDetails
    export DEVICE=$(echo -n "{\"deviceId\":${deviceId},\"description\":\"Device is on sale\", \"on_sale\":true}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"UpdateDeviceDetails","Args":[]}' --transient "{\"_Device\":\"$DEVICE\"}"

    # owner agreetosell
    export TRADE=$(echo -n "{\"tradeId\":${tradeId},\"tradePrice\":100}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AgreeToSell","Args":['${deviceId}']}' --transient "{\"_TradeAgreement\":\"$TRADE\"}"

    # add data
    export DEVICE_DATA=$(echo -n "{\"deviceId\":${deviceId},\"dataJSON\":\"random data\"}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AddDeviceData","Args":[]}' --transient "{\"_Data\":\"$DEVICE_DATA\"}"


    peer chaincode query -C mychannel -n marblesp -c '{"Args":["GetDeviceAllData", '${deviceId}']}'
    peer chaincode query -C mychannel -n marblesp -c '{"Args":["GetDeviceLatestData", '${deviceId}']}'


    peer chaincode query -C mychannel -n marblesp -c '{"Args":["QueryInterestTokenFromTradeId", '${tradeId}']}'

    # invokeDatasharing
    # 1. verify agreements
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"InvokeDataDistribution","Args":['${tradeId}']}'
    # 2. add to ACL
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AddToACL","Args":['${bidderId}','${tradeId}','${deviceId}']}'

    # try adding data again

# ORG2

    export PATH=${PWD}/../bin:$PATH
    export FABRIC_CFG_PATH=$PWD/../config/
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051

    export tradeId=\"trade001\"
    export deviceId=\"device001\"

    # buyer agreetobuy
    export TRADE=$(echo -n "{\"tradeId\":${tradeId},\"tradePrice\":100}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"AgreeToBuy","Args":['${deviceId}']}' --transient "{\"_TradeAgreement\":\"$TRADE\"}"

    export INTEREST_TOKEN=$(echo -n "{\"tradeId\":${tradeId},\"deviceId\":${deviceId}, \"dealsCollection\":\"Org2MSP_tradeAgreementCollection\"}" | base64 | tr -d \\n)
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n marblesp -c '{"function":"CreateInterestToken","Args":[]}' --transient "{\"_InterestToken\":\"$INTEREST_TOKEN\"}"





















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

