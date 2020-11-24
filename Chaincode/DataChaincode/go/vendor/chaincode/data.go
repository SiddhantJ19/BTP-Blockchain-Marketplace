package chaincode

import (
	"encoding/json"
    "github.com/hyperledger/fabric-chaincode-go/shim"

    //"fmt"
	"time"
	"errors"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) AddDeviceData(ctx contractapi.TransactionContextInterface) error {
	//marketplaceCollection, _ := getMarketplaceCollection()
	privateDetailsCollection, _ := getPrivateDetailsCollectionName()

	// 1. get transient map
	transientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
	}

	// 2.1 get Device from transientMap
	dataInputAsBytes := transientMap["_Data"]
	if dataInputAsBytes == nil {
		return errors.New("No input")
	}

	type DeviceDataInput struct {
		// DeviceId
		// Data - JSON object
		ID   string `json:"deviceId"`
		Data string `json:"dataJSON"` // JSON Data -> string
	}
	// 2.2 unmarshal json to an object

	var deviceInput DeviceDataInput
	err = json.Unmarshal(dataInputAsBytes, &deviceInput)
	if err != nil {
	}

	// 2.3 validate non empty fields

	//3. verify if clientMSP = peerMSP

	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
	}

	/*
	key := generateKeyForDevice(deviceInput.ID)
	deviceAsBytes, err := ctx.GetStub().GetPrivateData(marketplaceCollection, key)
	if err != nil {
		return fmt.Errorf("device %v does not exist \n %v", key, err.Error())
	}
	// ----------------- add Data -------------------

	// 4. getPrivateDetailsCollectionName(ctx)


	//4. check if device exist
	deviceAsBytes, err = ctx.GetStub().GetPrivateData(privateDetailsCollection, deviceInput.ID)
	if err != nil {
	}
	*/
	devicedataKey := generateKeyForDevicedata(deviceInput.ID)
	deviceDataAsBytes, err := ctx.GetStub().GetPrivateData(privateDetailsCollection, devicedataKey)
	if err != nil {
		deviceDataAsBytes = []byte("{}")
	}
	var deviceAllData DeviceData;
	err = json.Unmarshal(deviceDataAsBytes,&deviceAllData)
	if err != nil {

	}

	newDataEntry := DeviceDataObject{
		Timestamp: time.Now(),
		Data: deviceInput.Data,
	}
	newDataEntryAsBytes, err := json.Marshal(newDataEntry)
	deviceAllData.ID = deviceInput.ID

	deviceAllData.Data = append(deviceAllData.Data, newDataEntry )

	//if len(deviceAllData.Data) > 5 {
	//	deviceAllData.Data = deviceAllData.Data[1:]
	//}

	deviceAllDataAsBytes, err := json.Marshal(deviceAllData)
	if err != nil {
		return err;
	}
	err = ctx.GetStub().PutPrivateData(privateDetailsCollection, devicedataKey, deviceAllDataAsBytes)
	if err != nil {
		return err;
	}

    // make collections
    // copy to collections after checking ACL
    aclCollection, _ := getACLCollection()
    deviceACLAsBytes, err := ctx.GetStub().GetPrivateData(aclCollection, deviceInput.ID)
    var deviceACL DeviceACL
    err = json.Unmarshal(deviceACLAsBytes, &deviceACL)

    ownerMSP, err := shim.GetMSPID();
    for _, aclObject := range deviceACL.List {
        sharingCollection, _ := getSharingCollection(ownerMSP, aclObject.BuyerId)
        err = ctx.GetStub().PutPrivateData(sharingCollection, devicedataKey, newDataEntryAsBytes)
    }
	return nil
}

func (s *SmartContract) GetDeviceAllData(ctx contractapi.TransactionContextInterface, deviceId string) ([]DeviceDataObject, error) {
	privateDetailsCollection, _ := getPrivateDetailsCollectionName()
	deviceDataKey := generateKeyForDevicedata(deviceId)

	deviceDataAsBytes, err := ctx.GetStub().GetPrivateData(privateDetailsCollection, deviceDataKey)
	var dataObjects []DeviceDataObject
	if err != nil {
		return nil, errors.New("Device or data does not exist")
	}

	var deviceAllData DeviceData;
	err = json.Unmarshal(deviceDataAsBytes,&deviceAllData)
	if err != nil {
		return nil, errors.New("Device or data does not exist")
	}

	dataObjects = deviceAllData.Data;

	return dataObjects, nil

}

func (s *SmartContract) GetDeviceLatestData(ctx contractapi.TransactionContextInterface, deviceId string) (DeviceDataObject, error) {
	privateDetailsCollection, _ := getPrivateDetailsCollectionName()
	deviceDataKey := generateKeyForDevicedata(deviceId)

	deviceDataAsBytes, err := ctx.GetStub().GetPrivateData(privateDetailsCollection, deviceDataKey)
	var data DeviceDataObject
	if err != nil {
		return data, errors.New("Device or data does not exist")
	}

	var deviceAllData DeviceData;
	err = json.Unmarshal(deviceDataAsBytes,&deviceAllData)
	if err != nil {
		return data, errors.New("Device or data does not exist")
	}

	dataObjects := deviceAllData.Data;

	return dataObjects[len(dataObjects) - 1], nil

}

type ArbitaryData struct{
	Data interface{}
}

func (s *SmartContract) QueryPrivateData(ctx contractapi.TransactionContextInterface, queryString string) ([]*ArbitaryData,error) {

	privateDetailsCollection, _ := getPrivateDetailsCollectionName()



	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(privateDetailsCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*ArbitaryData
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset ArbitaryData
		err = json.Unmarshal(queryResult.Value, &asset.Data)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil

}
