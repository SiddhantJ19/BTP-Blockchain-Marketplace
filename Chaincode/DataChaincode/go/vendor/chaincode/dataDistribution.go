package chaincode

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "time"
)

/*
1. QueryInterestTokenFromTradeId(tradeId, string)
2.
*/



func (s *SmartContract) GetAndVerifyTradeAgreements(ctx contractapi.TransactionContextInterface, tradeId string) (AgreementDetails, error) {
    err := verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}

    bidderIntrestToken, err := s.QueryInterestTokenFromTradeId(ctx, tradeId)
    if err != nil {
        return AgreementDetails{} ,fmt.Errorf("Cannot get BidderInterestToken, %v", err.Error())
    }
    bidderId := bidderIntrestToken[0].BidderID
    deviceId := bidderIntrestToken[0].DeviceId
    bidderTradeAgreementCollection := bidderIntrestToken[0].TradeAgreementCollection
    fmt.Println(bidderIntrestToken[0])

    err = verifyClientOrgMatchesOwner(ctx, deviceId)
    if err != nil {}

    ownerTradeAgreementCollection,err := getTradeAgreementCollection(ctx)
    if err != nil {}

    sellerAgreementHash, err := getAgreementHash(ctx, ownerTradeAgreementCollection, tradeId)
    buyerAgreementHash, err := getAgreementHash(ctx, bidderTradeAgreementCollection, tradeId)
    if !bytes.Equal(sellerAgreementHash, buyerAgreementHash) {
        return AgreementDetails{}, fmt.Errorf("Agreements do not match")
    }

    agreementDetails := AgreementDetails{
        TradeId: tradeId,
        BuyerID: bidderId,
        SellerAgreementHash: string(sellerAgreementHash),
        BuyerAgreementHash: string(buyerAgreementHash),
    }
    return agreementDetails, nil
}


func (s *SmartContract) AddToACL(ctx contractapi.TransactionContextInterface, bidderId string, tradeId string, deviceId string) error {
    revokeTime, err := s.GetRevokeTime(ctx, tradeId)
    newACLObject := ACLObject{
        TradeID: tradeId,
        BuyerId: bidderId,
        RevokeTime: revokeTime,
    }
    fmt.Println("newAClObject\n")
    fmt.Println(newACLObject)

    aclCollection, err := getACLCollection(ctx)
    fmt.Println(aclCollection)
    fmt.Printf("%s %s %s \n\n", bidderId, tradeId, deviceId)
    aclAsBytes, err := ctx.GetStub().GetPrivateData(aclCollection, deviceId)
    if err != nil {
        fmt.Println(err)
    }

    var acl DeviceACL
    err = json.Unmarshal(aclAsBytes, &acl)
    fmt.Println("acl\n")
    fmt.Println(acl)

    acl.ID = deviceId
    acl.List = append(acl.List, newACLObject)
    fmt.Println("acl\n")
    fmt.Println(acl)

    aclAsBytes, err = json.Marshal(acl)
    if err != nil {
        return fmt.Errorf("Marshalling Error %v", err.Error())
    }

    err = ctx.GetStub().PutPrivateData(aclCollection, deviceId, aclAsBytes)
    if err != nil {
        fmt.Println("Error while putting private data")
       return fmt.Errorf("Error Putting in ACL %v", err.Error())
    }
    return nil
}

func (s *SmartContract) GenerateReceipt(ctx contractapi.TransactionContextInterface, ad AgreementDetails) error {

    revokeTime, err := s.GetRevokeTime(ctx, ad.TradeId)
    tradeConfirmation := TradeConfirmation{
        Type: "TRADE_CONFIRMATION",
        SellerAgreementHash: ad.SellerAgreementHash,
        BuyerAgreementHash: ad.BuyerAgreementHash,
        RevokeTime: revokeTime,
    }

    tradeConfirmationAsBytes, err := ctx.GetStub().GetState(ad.TradeId)
    if err != nil {}
    if tradeConfirmationAsBytes != nil {
        return fmt.Errorf("TradeId Already Exists on Blockchain")
    }

    tradeConfirmationAsBytes, err = json.Marshal(tradeConfirmation)
    if err != nil {return err}
    err = ctx.GetStub().PutState(ad.TradeId, tradeConfirmationAsBytes)
    // check transactionid in database
    transactionId := ctx.GetStub().GetTxID()
    sellerId, err := shim.GetMSPID()
    tradeEventPayload := Receipt{
        Type: "Trade Receipt",
        Buyer: ad.BuyerID,
        Seller: sellerId,
        TransactionId: transactionId,
        TimeStamp: time.Now(),
        TradeId: ad.TradeId,
        RevokeTime: revokeTime,
    }
    tradeEventPayloadAsBytes, err := json.Marshal(tradeEventPayload)
    fmt.Println("INSIDE RECEIPT CONTRACT")
    return ctx.GetStub().SetEvent("RECEIPT-EVENT", tradeEventPayloadAsBytes)
}

func getAgreementHash(ctx contractapi.TransactionContextInterface, collection string, tradeId string) ([]byte, error) {
    agreementHashAsBytes, err := ctx.GetStub().GetPrivateDataHash(collection, tradeId)
    if err != nil {
        return nil, err
    }
    return agreementHashAsBytes, nil
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

