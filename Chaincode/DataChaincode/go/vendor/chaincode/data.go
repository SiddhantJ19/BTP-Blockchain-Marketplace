package chaincode

import (
    "encoding/json"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) AddDeviceData(ctx contractapi.TransactionContextInterface) error {

    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil { }

    // 2.1 get Device from transientMap
    deviceAsBytes := transientMap["_Device"]
    if deviceAsBytes == nil {}

    // 2.2 unmarshal json to an object

    var deviceInput DeviceData
    err = json.Unmarshal(deviceAsBytes, &deviceInput)
    if err != nil {}

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}

    // ----------------- add Data -------------------

    // 4. getPrivateDetailsCollectionName(ctx)
    privateDetailsCollection, err := getPrivateDetailsCollectionName()
    if err != nil {}

    //4. check if device exist
    deviceAsBytes, err = ctx.GetStub().GetPrivateData(privateDetailsCollection, deviceInput.ID)
    if err != nil {}

    // todo = key for collection
    // push to collectionDeviceDataOrg1
    return nil
}

