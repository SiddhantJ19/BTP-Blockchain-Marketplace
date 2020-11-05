package main

import (
	"encoding/json"
	"fmt"
	// "strings"
	// "time"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"math/rand"
)


type Device struct {
	ObjectType string `json:"docType"`
	Device string `json:"device"`
	Owner string `json:"owner"`
	Description string `json:"description"`
}

type DeviceData struct {
	ObjectType string `json:"docType"`
	Device string `json:"device"`
	TimeStamp string `json:"timestamp"` // TODO: change to time.Time
	Data int `json:"data"`
}

type DevicePrivateDetails struct {
	ObjectType string `json:"docType"`
	Device string `json:"device"`
	Price string `json:"price"`
}

type SmartContract struct {
	contractapi.Contract
}

// Collections
var o1o2_marketPlace = "ORG1-ORG2-marketPlace"
var o1o2_dataSharing = "ORG1-ORG2-dataSharing"
var o1 = "collection-ORG1"
var o2 = "collection-ORG2"
var o1acl = "collection-ORG1-ACL"

func (s *SmartContract) InitDevice(ctx contractapi.TransactionContextInterface) error{
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("Error getting transient: " + err.Error())
	}
	// retrieve device data from transient
	transientDeviceJSON, ok := transMap["device"]
	if !ok {
		return fmt.Errorf("device not found in transient map")
	}

	type DeviceTransientInput struct {
		Device string `json:"device"`
		Owner string `json:"owner"`
		Description string `json:"description"`
	}

	var deviceDetails DeviceTransientInput
	err = json.Unmarshal(transientDeviceJSON, &deviceDetails)
	if err != nil {
		return fmt.Errorf("failed to unmarshall json: %s", err.Error())
	}
	// TODO: Input data validation

	// check if device already exists
	deviceAsBytes, err := ctx.GetStub().GetPrivateData( o1o2_marketPlace, deviceDetails.Device)
	if err != nil {
		return fmt.Errorf("Failed to get device: " + err.Error())
	} else if deviceAsBytes != nil {
		return fmt.Errorf("This device already exists: " + deviceDetails.Device)
	}

	// create device, marshal to JSON, save to state
	device := &Device {
		ObjectType: "Device",
		Device: deviceDetails.Device,
		Description: deviceDetails.Description,
		Owner: deviceDetails.Owner,
	}
	deviceJSONasBytes, err := json.Marshal(device)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// save to state
	err = ctx.GetStub().PutPrivateData(o1o2_marketPlace, deviceDetails.Device, deviceJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put Device: %s", err.Error())
	}

	return nil
}


func (s *SmartContract) ChangeDescription(ctx contractapi.TransactionContextInterface, newDescription, deviceId string) error {
	deviceJSON, err := ctx.GetStub().GetPrivateData(o1o2_marketPlace, deviceId)
	if err != nil {
		return fmt.Errorf("Failed to get device " + err.Error())
	}
	if deviceJSON == nil {
		return fmt.Errorf("No device registered with Id %s, Error: %s", deviceId, err.Error())
	}
	var device Device 
	err = json.Unmarshal(deviceJSON, &device)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	device.Description = newDescription
	updatedDevice, err := json.Marshal(device)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	err = ctx.GetStub().PutPrivateData(o1o2_marketPlace, deviceId, updatedDevice)
	return nil
}

func (s *SmartContract) RequestData(ctx contractapi.TransactionContextInterface, deviceId string) error {
	key := "request" + string(rand.Intn(100))
	err := ctx.GetStub().PutPrivateData(o1o2_dataSharing, key, []byte(deviceId))
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}


func (s *SmartContract) AddData(ctx contractapi.TransactionContextInterface) error {
	// 1. get data from transient
	transMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("Error getting transient: " + err.Error())
	}

	transientDataJSON, ok := transMap["data"]
	if !ok {
		return fmt.Errorf("")
	}

	// 2. binary (getTransient returns byte) to json
	type dataTransientInput struct {
		Timestamp string `json:"timestamp"`
		Device string `json:"device"`
		Data int `json:"data"`
	}

	var dataInput dataTransientInput
	err = json.Unmarshal(transientDataJSON, &dataInput)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}

	// 3. check if device is registered on o1o2_marketPlace
	deviceAsBytes, err := ctx.GetStub().GetPrivateData(o1o2_marketPlace, dataInput.Device)
	if err != nil {
		return fmt.Errorf("Failed to get existing device: " + err.Error())
	}
	if deviceAsBytes == nil {
		return fmt.Errorf("The device is not registered: " + dataInput.Device)
	}

	// 4. create data object, marshal to JSON, try to save to o1 and o2
	deviceData := &DeviceData {
		ObjectType: "Data",
		TimeStamp: dataInput.Timestamp,
		Device: dataInput.Device,
		Data: dataInput.Data,
	}
	deviceDataJSONasBytes, err := json.Marshal(deviceData)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// try to put on o1
	err = ctx.GetStub().PutPrivateData(o1, dataInput.Timestamp, deviceDataJSONasBytes)
	if err != nil {
		fmt.Println("failed to put data private details: %s", err.Error())
	}

	// try to put in o2
	err = ctx.GetStub().PutPrivateData(o2, dataInput.Timestamp, deviceDataJSONasBytes)
	if err != nil {
		fmt.Println("failed to put data private details: %s", err.Error())
	}

	// 5. check if data sharing is on and share, put on o1o2_marketPlace --> only suuport org1 for now
	isSharingAsBytes, err := ctx.GetStub().GetPrivateData(o1acl, "isSharing")
	if err != nil {
		return fmt.Errorf("addData: Error fetching ACL: " + err.Error())
	}
	if isSharingAsBytes != nil && string(isSharingAsBytes) == "true" {
		err = ctx.GetStub().PutPrivateData(o1o2_dataSharing, dataInput.Timestamp, deviceDataJSONasBytes)
		if err != nil {
			fmt.Printf("Error while putting data on data sharing state: " + err.Error())
		}
	}

	// TODO: add indexes
	return nil

}

func (s *SmartContract) InvokeDataSharing(ctx contractapi.TransactionContextInterface) error {
	err := ctx.GetStub().PutPrivateData(o1acl, "isSharing", []byte("true"))
	if err != nil {
		return fmt.Errorf("Error while fetching from ACL: " + err.Error())
	}
	return nil
}

func (s *SmartContract) RevokeDataSharing(ctx contractapi.TransactionContextInterface) error {
	err := ctx.GetStub().PutPrivateData(o1acl, "isSharing", []byte("false"))
	if err != nil {
		return fmt.Errorf("Error while fetching from ACL: " + err.Error())
	}
	return nil
}


func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Println("Error creating private_data_share chaincode : %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error in starting the private data sharing chaincode: %s", err.Error())
	}
}