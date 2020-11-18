package chaincode

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// to be called by seller (only owner can sell their asset)
// creates a trade agreement if seller is owner
func (s *SmartContract) AgreeToSell(ctx contractapi.TransactionContextInterface, deviceId string) error {
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

    deviceKey := generateKeyForDevice(deviceId)

    deviceAsBytes, err := ctx.GetStub().GetPrivateData(marketplaceCollection, deviceKey)
    if err != nil {
        return fmt.Errorf(err.Error())
    }

    var device DevicePublicDetails
    err = json.Unmarshal(deviceAsBytes, &device)
    if err != nil {return fmt.Errorf(err.Error())}

    ownerOrgId := device.Owner
    peerOrgId, err := shim.GetMSPID()
    if err != nil {return fmt.Errorf(err.Error())}

    if ownerOrgId != peerOrgId {
        return fmt.Errorf("Operation not permitted. Cannot sell someone else's asset")
    }
    return s.CreateTradeAgreement(ctx, deviceId)
}

// to be called by buyer
// creates a trade agreement
func (s *SmartContract) AgreeToBuy(ctx contractapi.TransactionContextInterface, deviceId string) error {
    return s.CreateTradeAgreement(ctx, deviceId)
}

// not to be called directly
func (s *SmartContract) CreateTradeAgreement(ctx contractapi.TransactionContextInterface, deviceId string) error {
    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil { }

    // 2.1 get Trade agreement from transientMap
    tradeAgreementAsBytes := transientMap["_TradeAgreement"]
    if tradeAgreementAsBytes == nil {}

    // 2.2 unmarshal json to an object
    type TradeAgreementInputTransient struct {
        ID          string `json:"tradeId"`
        Price       int    `json:"tradePrice"`
    }

    var tradeAgreementInput TradeAgreementInputTransient
    err = json.Unmarshal(tradeAgreementAsBytes, &tradeAgreementInput)
    if err != nil {}

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}


    // ----------------- TradeAgreement ---------------
    tradeAgreementCollection, err := getTradeAgreementCollection()
    if err != nil {}

    // check if tradeAgreement is present in ORG's TradeAgreements collection
    tradeAgreementAsBytes, err = ctx.GetStub().GetPrivateData(tradeAgreementCollection, tradeAgreementInput.ID)
    if err != nil {}
    if tradeAgreementAsBytes != nil {}

    // create trade agreement
    tradeAgreement := TradeAgreement{ID: tradeAgreementInput.ID, DeviceId: deviceId, Price: tradeAgreementInput.Price}

    // marshal the trade input
    tradeAgreementAsBytes, err = json.Marshal(tradeAgreement)

    // save trade agreement
    err = ctx.GetStub().PutPrivateData(tradeAgreementCollection, tradeAgreementInput.ID, tradeAgreementAsBytes)
    return nil
}

// to be called by buyer
// creates a bidder interest token on marketplace
func (s *SmartContract) CreateInterestToken (ctx contractapi.TransactionContextInterface) error {
    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil { }

    // 2.1 get Device from transientMap
    interestTokenAsBytes := transientMap["_InterestToken"]
    if interestTokenAsBytes == nil {}

    // 2.2 unmarshal json to an object
    type interestTokenInputTransient struct {
        ID              string `json:"tradeId"`
        deviceId        string `json:"deviceId"`
    }

    var interestTokenInput interestTokenInputTransient
    err = json.Unmarshal(interestTokenAsBytes, &interestTokenInput)
    if err != nil {}

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}

    // --------------------------- create interest token ---------------------------------------------

    // bidderId = clientId
    bidderOrgId, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {}

    // DealsCollection -> where all the deals are stored
    tradeAgreementCollection, err := getTradeAgreementCollection() // required to generate private-data hash for the bidder's agreement collection:tradeID

    // create Interest token
    interestToken := InterestToken{
        ID: interestTokenInput.ID,
        deviceId: interestTokenInput.deviceId,
        BidderID: bidderOrgId,
        TradeAgreementCollection: tradeAgreementCollection,
    }

    // marshal interest token obj to bytes[] and store in Marketplace with Key
    interestTokenAsBytes, err = json.Marshal(interestToken)
    if err != nil {}

    key:= generateKeyForInterestToken(interestToken.ID)

    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

    err = ctx.GetStub().PutPrivateData(marketplaceCollection,  key, interestTokenAsBytes)
    if err != nil {}

    return nil
}


