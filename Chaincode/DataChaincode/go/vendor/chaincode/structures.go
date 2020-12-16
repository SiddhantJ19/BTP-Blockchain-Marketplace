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
	RevokeTime  time.Time   `json:"revoke_time"`
}

type InterestToken struct { // token of interest passed by the bidder
	ID              string `json:"tradeId"`         // search all biddings for this device
	DeviceId        string `json:"deviceId"`        // unique key as TRADE_{deviceID} on Collection_Marketplace
	BidderID        string `json:"bidderId"`
	SellerId        string `json:"seller_id"`
    TradeAgreementCollection string `json:"dealsCollection"` // required to generate private-data hash for the bidder's agreement collection:tradeID
}
// to be returned via event
type Receipt struct {
	TimeStamp     time.Time `json:"time_stamp"`
	Seller        string    `json:"seller"`
	Buyer         string    `json:"buyer"`
	TransactionId string    `json:"trade_confirmation_transaction_id"`
	TradeId       string    `json:"trade_id"`
	Type          string    `json:"type"`
	RevokeTime    time.Time `json:"revoke_time"`
}

// on the blockchain
type TradeConfirmation struct {
    Type    string  `json:"type"`
    SellerAgreementHash string  `json:"seller_agreement_hash"`
    BuyerAgreementHash  string  `json:"buyer_agreement_hash"`
    RevokeTime          time.Time   `json:"revoke_time"`
}

// temp object to be returned from verifyTradeAgreements
type AgreementDetails struct  {
    TradeId string
    BuyerID string
    RevokeTime time.Time
    SellerAgreementHash string
    BuyerAgreementHash  string
}


type DeviceDataObject struct {
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"dataJSON"` // JSON Data -> string
	TransactionId string `json:"transactionId"`
}
// Data
type DeviceData struct {
	// DeviceId
	// Data - JSON object
	ID        string    `json:"deviceId"`
	Data      []DeviceDataObject    `json:"dataJSON"` // JSON Data -> string
}

// ACL
type ACLObject struct {
    BuyerId     string          `json:"buyerId"`
    TradeID     string          `json:"tradeId"`
    RevokeTime time.Time    `json:"revoke_time"`
}

type DeviceACL struct {
	// Device ID
	// TradeID
	ID      string `json:"deviceId"`
    List     []ACLObject `json:"acl"`
}
