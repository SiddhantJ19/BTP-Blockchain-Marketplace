package chaincode

import (
   "github.com/hyperledger/fabric-contract-api-go/contractapi"
   "encoding/json"
   "fmt"
   "github.com/hyperledger/fabric-chaincode-go/shim"
)

const assetCollection = "assetCollection"
//import "github.com/hyperledger/fabric-contract-api-go/contractapi"
//

//collection = MArketplace

//// queryOnSaleDataMarketplace -> list of DevicePublicDetails onSale
func (s *SmartContract) QueryOnSaleDataMarketplace(ctx contractapi.TransactionContextInterface) ([]*DevicePublicDetails, error) {
   marketplaceCollection, _ := getMarketplaceCollection()

   queryString := fmt.Sprintf(`{"selector":{"onSale":true, "_id":{"$regex":"DEVICE*"}}}`)
   resultsIterator, err := getQueryResultForQueryString(ctx, marketplaceCollection, queryString)
   if err != nil {
		return nil, err
	}
   return constructPublicDevicesQueryResponseFromIterator(resultsIterator)
}


func (t *SmartContract) QueryDevices(ctx contractapi.TransactionContextInterface, queryString string) ([]*DevicePublicDetails, error) {
   marketplaceCollection, _ := getMarketplaceCollection()
   resultsIterator, err := getQueryResultForQueryString(ctx, marketplaceCollection, queryString)

   if err != nil {
		return nil, err
	}
   return constructPublicDevicesQueryResponseFromIterator(resultsIterator)
}

//collection = MArketplace
// key
// queryBidders -> InterestToken for a tradeId
func (s *SmartContract) QueryInterestTokenFromTradeId(ctx contractapi.TransactionContextInterface, tradeId string) ([]*InterestToken, error) {
    marketplaceCollection, _ := getMarketplaceCollection()

    queryString := fmt.Sprintf(`{"selector":{"tradeId":"%s", "_id":{"$regex":"TRADE*"}}}`, tradeId)
    resultsIterator, err := getQueryResultForQueryString(ctx, marketplaceCollection, queryString)
    if err != nil {
        return nil, err
    }
    return constructInterestTokensQueryResponseFromIterator(resultsIterator)
}
// queryBidders -> list of all InterestTokens for a tradeId

func (s *SmartContract) QueryInterestTokensForDevice(ctx contractapi.TransactionContextInterface, deviceId string) ([]*InterestToken, error) {
   marketplaceCollection, _ := getMarketplaceCollection()

   queryString := fmt.Sprintf(`{"selector":{"deviceId":"%s", "_id":{"$regex":"TRADE*"}}}`, deviceId)

   resultsIterator, err := getQueryResultForQueryString(ctx, marketplaceCollection, queryString)
   if err != nil {
		return nil, err
	}
   return constructInterestTokensQueryResponseFromIterator(resultsIterator)
}

func (t *SmartContract) QueryInterestTokens(ctx contractapi.TransactionContextInterface, queryString string) ([]*InterestToken, error) {
   marketplaceCollection, _ := getMarketplaceCollection()
   resultsIterator, err := getQueryResultForQueryString(ctx, marketplaceCollection, queryString)

   if err != nil {
		return nil, err
	}
   return constructInterestTokensQueryResponseFromIterator(resultsIterator)
}



func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, collectionName string, queryString string) (shim.StateQueryIteratorInterface, error) {

   resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collectionName, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

   return resultsIterator, nil
}

func constructPublicDevicesQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*DevicePublicDetails, error) {
	var assets []*DevicePublicDetails
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset DevicePublicDetails
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func constructInterestTokensQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*InterestToken, error) {
	var assets []*InterestToken
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset InterestToken
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

