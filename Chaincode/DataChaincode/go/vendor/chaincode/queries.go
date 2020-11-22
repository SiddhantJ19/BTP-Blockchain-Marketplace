package chaincode

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

//import "github.com/hyperledger/fabric-contract-api-go/contractapi"
//

//collection = MArketplace

//// queryOnSaleDataMarketplace -> list of DevicePublicDetails onSale
func (s *SmartContract) queryOnSaleDataMarketplace(ctx contractapi.TransactionContextInterface) ([]*DevicePublicDetails, error) {
   queryString := fmt.Sprintf(`{"selector":{"onSale":true}}`)
   resultsIterator, err := getQueryResultForQueryString(ctx, queryString)
   if err != nil {
		return nil, err
	}
   return constructPublicDevicesQueryResponseFromIterator(resultsIterator)
}

func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) (shim.StateQueryIteratorInterface, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

   return resultsIterator, nil
}

func constructPublicDevicesQueryResponseFromIterator(resultsIterator) ([]*DevicePublicDetails, error) {
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

//collection = MArketplace
// key
// queryBidders -> list of all InterestTokens for a tradeId


