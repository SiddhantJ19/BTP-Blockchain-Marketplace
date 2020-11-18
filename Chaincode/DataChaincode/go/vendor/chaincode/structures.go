package chaincode

import (
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// DevicePublicDetails ...
type DevicePublicDetails struct {
	Owner       string `json:"owner"`
	ID          string `json:"deviceId"`        // uniqueId = DEVICE_{ID} on collection_Marketplace
	Data        string `json:"dataDescription"`
	Description string `json:"description"`
	OnSale      bool   `json:"onSale"`
}

type DevicePrivateDetails struct {          // Device Meta data
	ID     string `json:"deviceId"`         // uniqueId on collection_devicePrivatedetails
	Secret string `json:"deviceSecret"`
}

// Agreement
type TradeAgreement struct { // the hash of respective trade agreements should match
	ID       string `json:"tradeId"`    // unique key on collection_TradeAgreement
	DeviceId string `json:"deviceId"`   // search all trades for this device
	Price    int `json:"tradePrice"`
}

type InterestToken struct { // token of interest passed by the bidder
	ID              string `json:"tradeId"`         // unique key as TRADE_{deviceID} on Collection_Marketplace
	deviceId        string `json:"deviceId"`        // search all biddings for this device
	BidderID        string `json:"bidderId"`
    TradeAgreementCollection string `json:"dealsCollection"` // required to generate private-data hash for the bidder's agreement collection:tradeID
}

// Data
type DeviceData struct {
	// DeviceId
	// Data - JSON object
	ID        string    `json:"deviceId"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"dataJSON"` // JSON Data -> string
}

// ACL
type DeviceACL struct {
	// Device ID
	// TradeID
	ID      string `json:"deviceId"`
	TradeID string `json:"tradeId"`
}
