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



func (s *SmartContract) InvokeDataDistribution(ctx contractapi.TransactionContextInterface, tradeId string) ([]byte, []byte, string, error) {
    // verify client's org == peer's org
    err := verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}

    // get Device Id from interest token
    bidderIntrestToken, err := s.QueryInterestTokenFromTradeId(ctx, tradeId)
    if err != nil {
        return nil, nil, "ERROR" ,fmt.Errorf("Cannot get BidderInterestToken, %v", err.Error())
    }

    bidderId := bidderIntrestToken[0].BidderID
    deviceId := bidderIntrestToken[0].DeviceId
    bidderTradeAgreementCollection := bidderIntrestToken[0].TradeAgreementCollection
    fmt.Println(bidderIntrestToken[0])
    // check client org is the owner
    err = verifyClientOrgMatchesOwner(ctx, deviceId)
    if err != nil {}

    // getOwner's trade agreement collection
    ownerTradeAgreementCollection,err := getTradeAgreementCollection(ctx)
    if err != nil {}

    sellerAgreementHash, err := getAgreementHash(ctx, ownerTradeAgreementCollection, tradeId)
    buyerAgreementHash, err := getAgreementHash(ctx, bidderTradeAgreementCollection, tradeId)
    // verify trade conditions
    if !bytes.Equal(sellerAgreementHash, buyerAgreementHash) {
        return nil, nil, "ERROR" , fmt.Errorf("Agreements do not match")
    }

    return sellerAgreementHash, buyerAgreementHash, bidderId, nil
}

func (s *SmartContract) AddToACL(ctx contractapi.TransactionContextInterface, bidderId string, tradeId string, deviceId string) error {
    newACLObject := ACLObject{
        TradeID: tradeId,
        BuyerId: bidderId,
    }


    aclCollection, err := getACLCollection(ctx)
    fmt.Println(aclCollection)
    fmt.Printf("%s %s %s \n\n", bidderId, tradeId, deviceId)
    aclAsBytes, err := ctx.GetStub().GetPrivateData(aclCollection, deviceId)
    if err != nil {
        fmt.Println(err)
    }

    var acl DeviceACL
    err = json.Unmarshal(aclAsBytes, &acl)
    fmt.Printf("ACL %v \n", acl)

    acl.List = append(acl.List, newACLObject)
    fmt.Printf("ACL %v \n", acl)

    aclAsBytes, err = json.Marshal(acl)
    if err != nil {
        return fmt.Errorf("Marshalling Error %v", err.Error())
    }

    err = ctx.GetStub().PutPrivateData(aclCollection, deviceId, aclAsBytes)
    if err != nil {
        fmt.Println("Error while putting private data")
        fmt.Println(err.Error())
       return fmt.Errorf("Error Putting in ACL %v", err.Error())
    }
    fmt.Println("^^^^^^^^^^^^")
    return nil
}

func (s *SmartContract) CreateTradeConfirmationReceipt(ctx contractapi.TransactionContextInterface,
    sellerAgreementHash []byte, buyerAgreementHash []byte, tradeId string, buyerId string) error {

    tradeConfirmation := TradeConfirmation{
        Type: "TRADE_CONFIRMATION",
        SellerAgreementHash: string(sellerAgreementHash),
        BuyerAgreementHash: string(buyerAgreementHash),
    }
    tradeConfirmationAsBytes, err := json.Marshal(tradeConfirmation)
    if err != nil {return err}
    err = ctx.GetStub().PutState(tradeId, tradeConfirmationAsBytes)
    // check transactionid in database
    transactionId := ctx.GetStub().GetTxID()
    sellerId, err := shim.GetMSPID()
    tradeEventPayload := Receipt{
        Type: "Trade Receipt",
        Buyer: buyerId,
        Seller: sellerId,
        TransactionId: transactionId,
        TimeStamp: time.Now(),
        TradeId: tradeId,
    }
    tradeEventPayloadAsBytes, err := json.Marshal(tradeEventPayload)

    return ctx.GetStub().SetEvent("RECEIPT-EVENT", tradeEventPayloadAsBytes)
    // RETURN string, string, buyer's agreement hash, sellers agreement hash
}
// ----------------------------- Data Sharing Utils --------------------------------------------
//func verifyTradeConditions(ctx contractapi.TransactionContextInterface, bidderCollection string, sellerCollection string, key string) error {
//    bidderAgreementHash, err := ctx.GetStub().GetPrivateDataHash(bidderCollection, key)
//    if err != nil {}
//
//    sellerAgreementHash, err := ctx.GetStub().GetPrivateDataHash(sellerCollection, key)
//    if err != nil {}
//
//    if !bytes.Equal(bidderAgreementHash, sellerAgreementHash){
//        return fmt.Errorf("Agreements do not match")
//    }
//    return nil
//}

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

