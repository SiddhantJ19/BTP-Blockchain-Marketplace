package chaincode

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) invokeDataDistribution(ctx contractapi.TransactionContextInterface, tradeId string) error {
    // verify client's org == peer's org
    err := verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}

    // get Device Id from interest token
    bidderIntrestToken, err := s.QueryInterestTokenFromTradeId(ctx, tradeId)
    if err != nil {}

    deviceId := bidderIntrestToken[0].DeviceId
    bidderId := bidderIntrestToken[0].BidderID
    bidderTradeAgreementCollection := bidderIntrestToken[0].TradeAgreementCollection

    // check client org is the owner
    err = verifyClientOrgMatchesOwner(ctx, deviceId)
    if err != nil {}

    // getOwner's trade agreement collection
    ownerTradeAgreementCollection,err := getTradeAgreementCollection()
    if err != nil {}

    // verify trade conditions
    err = verifyTradeConditions(ctx, bidderTradeAgreementCollection, ownerTradeAgreementCollection, tradeId)
    if err != nil {
        return err
    }

    // updateACL
    err = AddToACL(ctx, bidderId, tradeId, deviceId)
    if err != nil {}

    return nil
}


// ----------------------------- Data Sharing Utils --------------------------------------------
func AddToACL(ctx contractapi.TransactionContextInterface, bidderId string, tradeId string, deviceId string) error {
    newACLObject := ACLObject{
        TradeID: tradeId,
        BuyerId: bidderId,
    }
    aclCollection, err := getACLCollection()

    aclAsBytes, err := ctx.GetStub().GetPrivateData(aclCollection, deviceId)
    if err != nil {}
    var acl DeviceACL
    err = json.Unmarshal(aclAsBytes, &acl)

    acl.List = append(acl.List, newACLObject)

    aclAsBytes, err = json.Marshal(acl)
    if err != nil {}

    err = ctx.GetStub().PutPrivateData(aclCollection, deviceId, aclAsBytes)
    if err != nil {}
    return nil
}

func verifyTradeConditions(ctx contractapi.TransactionContextInterface, bidderCollection string, sellerCollection string, key string) error {
    bidderAgreementHash, err := ctx.GetStub().GetPrivateDataHash(bidderCollection, key)
    if err != nil {}

    sellerAgreementHash, err := ctx.GetStub().GetPrivateDataHash(sellerCollection, key)
    if err != nil {}

    if !bytes.Equal(bidderAgreementHash, sellerAgreementHash){
        return fmt.Errorf("Agreements do not match")
    }
    return nil
}

func verifyClientOrgMatchesOwner(ctx contractapi.TransactionContextInterface, deviceId string) error {
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}
    deviceKey := generateKeyForDevice(deviceId)
    device := readDevicePublicDetails(ctx, marketplaceCollection, deviceKey) // returns marshalled data

    clientOrg, err := ctx.GetClientIdentity().GetMSPID()
    if device.Owner != clientOrg {
        return fmt.Errorf("clientOrg %v doesnot match Owner %v ", clientOrg, device.Owner)
    }
    return nil
}

func readDevicePublicDetails (ctx contractapi.TransactionContextInterface, collection string, key string) DevicePublicDetails {
    deviceAsBytes, err := ctx.GetStub().GetPrivateData(collection, key)
    if err != nil {}
    var device DevicePublicDetails
    err = json.Unmarshal(deviceAsBytes, &device)
    if err != nil {}
    return device
}


// todo
func (s *SmartContract) revokeDataDistribution(tradeId string) error {
    return nil
}
// TODO
// 1. chain of custody
// 2. what if one of the orgs changes contract details later
//      - can we prevent any updates on existing trade contract

