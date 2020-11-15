package chaincode

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/pkg/statebased"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)


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
    privateDetailsCollection, err := getPrivateDetailsCollectionName()
    if err != nil {}

    deviceAsBytes, err := ctx.GetStub().GetPrivateData(privateDetailsCollection, device.ID)
    if err != nil {}
    if deviceAsBytes != nil {}


    //5. public details
    clientID, err := ctx.GetClientIdentity().GetMSPID()
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
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

    err = ctx.GetStub().PutPrivateData(marketplaceCollection, device.ID, deviceAsBytes)
    if err != nil {}

    // 5.5 set the endorsement policy such that an owner's endorsement is required to update marketplace details of an asset
    // this is to prevent asset loss because of number of nodes wanting to change the asset details on a public marketplace
    err = setDeviceStateBasedEndorsement(ctx, device.ID, clientID, marketplaceCollection)
    if err != nil {}

    // 6. private details
    devicePrivateDetails := DevicePrivateDetails{
       ID: device.ID,
       Secret: device.Secret,
    }
    deviceAsBytes, err = json.Marshal(devicePrivateDetails)
    if err != nil {}

    // 6.1 save to db
    err = ctx.GetStub().PutPrivateData(privateDetailsCollection, device.ID, deviceAsBytes)
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
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

   deviceAsBytes, err = ctx.GetStub().GetPrivateData(marketplaceCollection, deviceInput.ID)
   if err != nil {}

   // unmarshall to DevicePublicDetails
   var deviceMarketplace DevicePublicDetails
   err = json.Unmarshal(deviceAsBytes,&deviceMarketplace)
   if err != nil {}

   // change the description if device's owner == clientOrgId -> done by the state based ep
    deviceMarketplace.Description = deviceInput.Description
    deviceMarketplace.OnSale = deviceInput.OnSale

    // marshall the device
   deviceAsBytes, err = json.Marshal(deviceMarketplace)
   if err != nil {}

   // put in the db
   err = ctx.GetStub().PutPrivateData(marketplaceCollection, deviceInput.ID, deviceAsBytes)

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
    privateDetailsCollection, err := getPrivateDetailsCollectionName()
    if err != nil {}

    //4. check if device exist
    deviceAsBytes, err = ctx.GetStub().GetPrivateData(privateDetailsCollection, deviceInput.ID)
    if err != nil {}

    // todo = key for collection
    // push to collectionDeviceDataOrg1
    return nil
}


// ----------------------- Trade ------------------------------------------------

// to be called by seller (only owner can sell their asset)
// creates a trade agreement if seller is owner
func (s *SmartContract) AgreeToSell(ctx contractapi.TransactionContextInterface, deviceId string) error {
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

    deviceAsBytes, err := ctx.GetStub().GetPrivateData(marketplaceCollection, deviceId)
    if err != nil {}

    var device DevicePublicDetails
    err = json.Unmarshal(deviceAsBytes, &device)
    ownerOrgId := device.Owner
    peerOrgId, err := shim.GetMSPID()
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
    bidderId, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {}

    // DealsCollection -> where all the deals are stored
    dealsCollection, err := getDealsCollection() // required to generate private-data hash for the bidder's agreement collection:tradeID

    // create Interesttoken
    interestToken := InterestToken{
        ID: interestTokenInput.ID,
        BidderID: bidderId,
        DealsCollection: dealsCollection,
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



// ============================ UTILS =========================================

func generateKeyForInterestToken(deviceId string) string {
    return "TRADE_" + deviceId
}

func generateKeyForDevice(deviceId string) string {
    return "DEVICE_" + deviceId
}

func getMarketplaceCollection() (string, error) {
    return "collection_Marketplace", nil
}

func getDealsCollection() (string, error) {
    msp, err := shim.GetMSPID()
    if err != nil {return "", err}

    return "collection_" + msp + "dealsCollection", nil
}

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


func setDeviceStateBasedEndorsement(ctx contractapi.TransactionContextInterface, deviceId string, orgId string, collection string) error {
    // create a new state based policy for key = deviceId
    ep, err := statebased.NewStateEP(nil)
    if err != nil {}

    // issue roles, here the owner org for a device
    err = ep.AddOrgs(statebased.RoleTypePeer, orgId)
    if err != nil {}

    policy, err := ep.Policy()
    if err != nil {}

    err = ctx.GetStub().SetPrivateDataValidationParameter(collection, deviceId, policy)
    return nil
}


// updateDeviceDescription - error
// addDeviceData

// createTradeAgreement
// createInterestToken
// InvokeDataSharing(BiddersInterestToken) <--> asset transfer
