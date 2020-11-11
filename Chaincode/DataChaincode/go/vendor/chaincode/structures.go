package chaincode

import (
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// Device
type DevicePublicDetails struct {
	Owner       string `json:"owner"`
	ID          string `json:"deviceId"`
	Data        string `json:"dataDescription"`
	Description string `json:"description"`
}

type DevicePrivateDetails struct {
	ID     string `json:"deviceId"`
	Secret string `json:"deviceSecret"`
}

// Agreement
type TradeAgreement struct { // the hash of respective trade agreements should match
	ID       string `json:"tradeId"`
	DeviceId string `json:"deviceId"`
	Price    string `json:"tradePrice"`
}

type InterestToken struct { // token of interest passed by the bidder
	ID              string `json:"tradeId"`
	BidderID        string `json:"bidderId"`
	DealsCollection string `json:"dealsCollection"` // required to generate private-data hash for the bidder's agreement collection:tradeID
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
