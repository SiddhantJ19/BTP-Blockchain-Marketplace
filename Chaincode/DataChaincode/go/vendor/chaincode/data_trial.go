package chaincode

import (
	"encoding/json"
	//"fmt"
	"time"
	"errors"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//default_data_limit=2
func (s *SmartContract) AddDeviceDataP1(ctx contractapi.TransactionContextInterface) error {
	//marketplaceCollection, _ := getMarketplaceCollection()
	//privateDetailsCollection, _ := getPrivateDetailsCollectionName()

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
	/*
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

	devicedataKey := generateKeyForDevicedata(deviceInput.ID)
	deviceDataAsBytes, _ := ctx.GetStub().GetPrivateData(privateDetailsCollection, devicedataKey)

	var deviceAllData []*DeviceData;
	err = json.Unmarshal(deviceDataAsBytes,&deviceAllData)
	if err != nil {
		return err;
	}

	newDeviceData := DeviceData{
		Data: deviceInput.Data,
	}
	deviceAllData = append(deviceAllData, &newDeviceData )

	deviceAllDataAsBytes, err := json.Marshal(deviceAllData)
	if err != nil {
		return err;
	}
	err = ctx.GetStub().PutPrivateData(privateDetailsCollection, devicedataKey, deviceAllDataAsBytes)
	if err != nil {
		return err;
	}
	// todo = key for collection
	// push to collectionDeviceDataOrg1


	*/
	return nil
}

func (s *SmartContract) AddDeviceDataP2(ctx contractapi.TransactionContextInterface) error {
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
	deviceAllData.ID = deviceInput.ID

	deviceAllData.Data = append(deviceAllData.Data, newDataEntry )
	if len(deviceAllData.Data) > 5 {
		deviceAllData.Data = deviceAllData.Data[1:]
	}
	deviceAllDataAsBytes, err := json.Marshal(deviceAllData)
	if err != nil {
		return err;
	}
	err = ctx.GetStub().PutPrivateData(privateDetailsCollection, devicedataKey, deviceAllDataAsBytes)
	if err != nil {
		return err;
	}



	return nil
}