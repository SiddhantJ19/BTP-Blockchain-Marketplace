package chaincode

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// constants
const ORG1_DevicePrivateDetails = "collectionDevicePrivateDetailsOrg1"
const ORG_Marketplace = "collectionMarketplace"

// ----------------------- Device ------------------------------------------------
func (s *SmartContract) CreateDevice(ctx contractapi.TransactionContextInterface) error {

    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil { }

    // 2.1 get Device from transientMap
    deviceJSON := transientMap["_Device"]
    if deviceJSON == nil {}

    // 2.2 unmarshal json to an object
    type deviceTransientInput struct {
       ID          string `json:"deviceId"`
       Data        string `json:"dataDescription"`
       Description string `json:"description"`
       Secret      string `json:"deviceSecret"`
    }

    var device deviceTransientInput
    err = json.Unmarshal(deviceJSON, &device)
    if err != nil {}

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}



    // 4. check if device already exists
    // collection = ORG1_DevicePrivateDetails, key = deviceID
    deviceAsBytes, err := ctx.GetStub().GetPrivateData(ORG1_DevicePrivateDetails, device.ID)
    if err != nil {}
    if deviceAsBytes != nil {}


    //5. public details
    clientID, err := ctx.GetClientIdentity().GetID()
    if err != nil {}

    devicePublicDetails := DevicePublicDetails{
        ID:	device.ID,
        Owner: clientID,
        Data:	device.Data,
        Description: device.Description,
        OnSale: true,
    }

    deviceAsBytes, err = json.Marshal(devicePublicDetails)
    if err != nil {}

    // 5.1 save to collection
    // Marketplace => key : deviceID
    err = ctx.GetStub().PutPrivateData(ORG_Marketplace, device.ID, deviceAsBytes)
    if err != nil {}

    // 6. private details
    devicePrivateDetails := DevicePrivateDetails{
       ID: device.ID,
       Secret: device.Secret,
    }
    deviceAsBytes, err = json.Marshal(devicePrivateDetails)
    if err != nil {}

    // 6.1 save to db
    err = ctx.GetStub().PutPrivateData(ORG1_DevicePrivateDetails, device.ID, deviceAsBytes)
    if err != nil {}

    return nil
}

func (s *SmartContract) UpdateDeviceDetails(ctx contractapi.TransactionContextInterface) error {
   // 1. get transient map
   transientMap, err := ctx.GetStub().GetTransient()
   if err != nil { }

   // 2.1 get Device from transientMap
   deviceAsBytes := transientMap["_Device"]
   if deviceAsBytes == nil {}

   // 2.2 unmarshal json to an object
   type DeviceTransientInput struct {
       ID          string `json:"deviceId"`
       Description string `json:"description"`
       OnSale      bool `json:"on_sale"`
   }

   var deviceInput DeviceTransientInput
   err = json.Unmarshal(deviceAsBytes, &deviceInput)
   if err != nil {}

   // 2.3 validate non empty fields

   //3. verify if clientMSP = peerMSP
   err = verifyClientOrgMatchesPeerOrg(ctx)
   if err != nil {}

   // ---------------- update description ----------------
   // get devicePublicDetails
   deviceAsBytes, err = ctx.GetStub().GetPrivateData(ORG_Marketplace, deviceInput.ID)
   if err != nil {}

   // unmarshall to DevicePublicDetails
   var deviceMarketplace DevicePublicDetails
   err = json.Unmarshal(deviceAsBytes,&deviceMarketplace)
   if err != nil {}

   // change the description
    deviceMarketplace.Description = deviceInput.Description
    deviceMarketplace.OnSale = deviceInput.OnSale

    // marshall the device
   deviceAsBytes, err = json.Marshal(deviceMarketplace)
   if err != nil {}

   // put in the db
   err = ctx.GetStub().PutPrivateData(ORG_Marketplace, deviceInput.ID, deviceAsBytes)

   return nil
}

// ----------------------- Data ------------------------------------------------
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
    privateDetailsCollection, err := getPrivateDetailsCollectionName(ctx)
    if err != nil {}

    //4. check if device exist
    deviceAsBytes, err = ctx.GetStub().GetPrivateData(privateDetailsCollection, deviceInput.ID)
    if err != nil {}

    // todo = key for collection
    // push to collectionDeviceDataOrg1
    return nil
}


// ----------------------- Trade ------------------------------------------------
func (s *SmartContract) CreateTradeAgreement(ctx contractapi.TransactionContextInterface) error {
    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil { }

    // 2.1 get Device from transientMap
    tradeAgreementAsBytes := transientMap["_TradeAgreement"]
    if tradeAgreementAsBytes == nil {}

    // 2.2 unmarshal json to an object
    type TradeAgreementInputTransient struct {
        ID          string `json:"tradeId"`
        DeviceId    string `json:"deviceId"`
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

    // marshal the trade input
    tradeAgreementAsBytes, err = json.Marshal(tradeAgreementInput)

    // save tradeagreement
    err = ctx.GetStub().PutPrivateData(tradeAgreementCollection, tradeAgreementInput.ID, tradeAgreementAsBytes)
    return nil
}

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
        BidderID        string `json:"bidderId"`
        DealsCollection string `json:"dealsCollection"` // required to generate private-data hash for the bidder's agreement collection:tradeID
    }

    var interestTokenInput interestTokenInputTransient
    err = json.Unmarshal(interestTokenAsBytes, &interestTokenInput)
    if err != nil {}

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}



}

// ============================ UTILS =========================================

func getTradeAgreementCollection() (string, error) {
    msp, err := shim.GetMSPID()
    if err != nil {return "", err}

    return "collection_" + msp + "tradeAgreement", nil
}

func getPrivateDetailsCollectionName() (string, error) {
    msp, err := shim.GetMSPID()
    if err != nil {return "", err}

    return "collection_" + msp + "privateDetails", nil
}


func verifyClientOrgMatchesPeerOrg(ctx contractapi.TransactionContextInterface) error {
	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {}

	peerMSP, err := shim.GetMSPID();
	if err != nil {}

	if clientMSP != peerMSP {
		return fmt.Errorf("client MSP %v does not match PeerMSP %v", clientMSP, peerMSP)
	}
	return nil
}



// updateDeviceDescription - error
// addDeviceData

// createTradeAgreement
// createInterestToken
// InvokeDataSharing(BiddersInterestToken) <--> asset transfer
