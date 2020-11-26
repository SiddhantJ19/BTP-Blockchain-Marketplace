package chaincode

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "errors"
)

func (s *SmartContract) CreateDevice(ctx contractapi.TransactionContextInterface) error {

    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil {return err}

    // 2.1 get Device from transientMap
    deviceAsBytes := transientMap["_Device"]
    if deviceAsBytes == nil {return fmt.Errorf("_Device not present in transient map")}

    // 2.2 unmarshal json to an object
    type deviceTransientInput struct {
        ID          string `json:"deviceId"`
        Data        string `json:"dataDescription"`
        Description string `json:"description"`
        Secret      string `json:"deviceSecret"`
    }

    var deviceInput deviceTransientInput
    err = json.Unmarshal(deviceAsBytes, &deviceInput)
    if err != nil {}

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}


    // 4. check if device already exists
    // collection = ORG1_DevicePrivateDetails, key = deviceID

    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

    deviceKey := generateKeyForDevice(deviceInput.ID)

    deviceAsBytes, err = ctx.GetStub().GetPrivateData(marketplaceCollection, deviceKey)
    if err != nil {}
    if deviceAsBytes != nil { return fmt.Errorf("device with ID %v already exist", deviceInput.ID)}


    //5. public details
    clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {}

    devicePublicDetails := DevicePublicDetails{
        ID:	deviceInput.ID,
        Owner: clientOrgID,
        Data:	deviceInput.Data,
        Description: deviceInput.Description,
        OnSale: true,
    }

    deviceAsBytes, err = json.Marshal(devicePublicDetails)
    if err != nil {}

    // 5.1 save to collection
    // Marketplace => key : deviceID

    fmt.Println("Create Device : Putting into market collection")
    err = ctx.GetStub().PutPrivateData(marketplaceCollection, deviceKey, deviceAsBytes)
    if err != nil {}

    //// 5.5 set the endorsement policy such that an owner's endorsement is required to update marketplace details of an asset
    //// this is to prevent asset loss because of more number of nodes wanting to change the asset details on a public marketplace
    err = setDeviceStateBasedEndorsement(ctx, deviceKey, clientOrgID, marketplaceCollection)
    if err != nil {}

    // 6. private details
    devicePrivateDetails := DevicePrivateDetails{
        ID: deviceInput.ID,
        Secret: deviceInput.Secret,
    }
    deviceAsBytes, err = json.Marshal(devicePrivateDetails)
    if err != nil {}

    // 6.1 save to db
    privateDetailsCollection, err := getPrivateDetailsCollectionName(ctx)
    if err != nil {
        return err
    }

    fmt.Println("Create Device : Putting into private collection")
    err = ctx.GetStub().PutPrivateData(privateDetailsCollection, deviceInput.ID, deviceAsBytes)
    if err != nil {
        return err
    }

    return nil
}

func (s *SmartContract) UpdateDeviceDetails(ctx contractapi.TransactionContextInterface) error {
    // 1. get transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil { }

    // 2.1 get Device from transientMap
    deviceAsBytes := transientMap["_Device"]
    if deviceAsBytes == nil {
        return errors.New("No input")
    }

    // 2.2 unmarshal json to an object
    type DeviceTransientInput struct {
        ID          string `json:"deviceId"`
        Description string `json:"description"`
        OnSale      bool   `json:"on_sale"`
    }

    var deviceInput DeviceTransientInput
    err = json.Unmarshal(deviceAsBytes, &deviceInput)
    if err != nil {
        return err
    }

    // 2.3 validate non empty fields

    //3. verify if clientMSP = peerMSP
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {}

    // ---------------- update description ----------------
    // get devicePublicDetails
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}

    fmt.Println("Update Device : getting from private collection")
    key := generateKeyForDevice(deviceInput.ID)
    deviceAsBytes, err = ctx.GetStub().GetPrivateData(marketplaceCollection, key)
    if err != nil {
        return fmt.Errorf("device %v does not exist \n %v" , key, err.Error())
    }

    // unmarshall to DevicePublicDetails
    var deviceMarketplace DevicePublicDetails
    err = json.Unmarshal(deviceAsBytes,&deviceMarketplace)
    if err != nil {
        return fmt.Errorf("device %v does not exist from unmarshal\n %v" , key, err.Error())
    }

    // change the description if device's owner == clientOrgId -> done by the state based ep
    deviceMarketplace.Description = deviceInput.Description
    deviceMarketplace.OnSale = deviceInput.OnSale

    // marshall the device
    deviceAsBytes, err = json.Marshal(deviceMarketplace)
    if err != nil {}


    // put in the db
    fmt.Println("Update Device : Putting into private collection")
    err = ctx.GetStub().PutPrivateData(marketplaceCollection, key, deviceAsBytes)

    return nil
}

func (s *SmartContract) DeleteDevice(ctx contractapi.TransactionContextInterface, deviceId string) (DevicePublicDetails, error) {
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}
    var deviceMarketplace DevicePublicDetails

    key := generateKeyForDevice(deviceId)
    deviceAsBytes, derr := ctx.GetStub().GetPrivateData(marketplaceCollection, key)
    if derr != nil {
        return deviceMarketplace, fmt.Errorf("device %v does not exist \n %v" , key, err.Error())
    }

    err = json.Unmarshal(deviceAsBytes,&deviceMarketplace)
    if err != nil {}

    err = ctx.GetStub().DelPrivateData(marketplaceCollection, key)
    if err!= nil {
        return deviceMarketplace, err
    }
    return deviceMarketplace, nil
}

func (s *SmartContract) GetDeviceDetails(ctx contractapi.TransactionContextInterface, deviceId string) (*DevicePublicDetails, error) {
    marketplaceCollection, err := getMarketplaceCollection()
    if err != nil {}
    var deviceMarketplace DevicePublicDetails

    key := generateKeyForDevice(deviceId)
    deviceAsBytes, derr := ctx.GetStub().GetPrivateData(marketplaceCollection, key)
    if derr != nil {
        return nil, fmt.Errorf("device %v does not exist \n %v" , key, err.Error())
    }


    err = json.Unmarshal(deviceAsBytes,&deviceMarketplace)
    if err != nil {
        return nil, fmt.Errorf("device %v does not exist \n %v" , key, err.Error())
    }

    return &deviceMarketplace, nil


}