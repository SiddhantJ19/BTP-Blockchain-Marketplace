package chaincode

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

//import "github.com/hyperledger/fabric-contract-api-go/contractapi"
//

//collection = MArketplace

//// queryOnSaleDataMarketplace -> list of DevicePublicDetails onSale
func (s *SmartContract) queryOnSaleDataMarketplace(ctx contractapi.TransactionContextInterface) ([]*DevicePublicDetails, error) {

   return nil, nil
}

//collection = MArketplace
// key
// queryBidders -> InterestToken for a tradeId
func (s *SmartContract) QueryBidderInterestToken(ctx contractapi.TransactionContextInterface, tradeId string) (InterestToken, error) {
    return InterestToken{
        DeviceId: "dev001",
        ID: "tradeId001",
        TradeAgreementCollection: "not-required",
        BidderID: "Org2Msp",
    }, nil
}

