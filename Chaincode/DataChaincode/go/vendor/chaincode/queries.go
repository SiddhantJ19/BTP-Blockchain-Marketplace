package chaincode

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "time"
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

func (s *SmartContract) QuerySharedDevices(ctx contractapi.TransactionContextInterface, ownerOrg string) ([]string, error) {
	selfMsp, mspErr :=ctx.GetClientIdentity().GetMSPID()
	if mspErr != nil {

		return nil, mspErr
	}
	sharedDevicesDetailsCollection, _ := getSharingCollection(ownerOrg,selfMsp)

	queryString := fmt.Sprintf(`{"selector":{"_id":{"$regex":"DATA*"}}}`)

	resultsIterator, err := getQueryResultForQueryString(ctx, sharedDevicesDetailsCollection, queryString)
	if err != nil {
		return nil, err
	}

	fullData, err := constructDevicesDataQueryResponseFromIterator(resultsIterator)

	var devicesList []string

	for d := range fullData {
		devicesList = append(devicesList, fullData[d].ID)
	}

	return devicesList, nil
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
        return nil, fmt.Errorf("No Interest Tokens for such device")
    }
   return constructInterestTokensQueryResponseFromIterator(resultsIterator)
}

func (s *SmartContract) QueryTradeAgreementsForDevice(ctx contractapi.TransactionContextInterface, deviceId string) ([]*TradeAgreement, error) {
	tradeAgreementCollection, _ := getTradeAgreementCollection(ctx)

	queryString := fmt.Sprintf(`{"selector":{"deviceId":"%s"}}`, deviceId)

	resultsIterator, err := getQueryResultForQueryString(ctx, tradeAgreementCollection, queryString)
	if err != nil {
		return nil, err
	}
	return constructTradeAgreementsQueryResponseFromIterator(resultsIterator)
}

func (s *SmartContract) QueryACLForDevice(ctx contractapi.TransactionContextInterface, deviceId string) (*DeviceACL, error) {
    aclCollection, _ := getACLCollection(ctx)

    aclAsBytes, err := ctx.GetStub().GetPrivateData(aclCollection, deviceId)
    if err != nil {
        fmt.Println("No ACL for the Device")
        fmt.Println(err.Error())
        return nil, fmt.Errorf("No ACL for the device", err.Error())
    }
    if aclAsBytes ==  nil {
        fmt.Println("Empty ACL")
        return nil, fmt.Errorf("Empty ACL")
    }
    var acl DeviceACL
    err = json.Unmarshal(aclAsBytes, &acl)
    if err != nil {}
    fmt.Println(acl)
    return &acl, nil
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

func constructTradeAgreementsQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*TradeAgreement, error) {
	var assets []*TradeAgreement
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset TradeAgreement
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func constructDevicesDataQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]DeviceData, error) {
	var assets []DeviceData
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset DeviceData
		err = json.Unmarshal(queryResult.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (s *SmartContract) GetRevokeTime(ctx contractapi.TransactionContextInterface, tradeId string) (time.Time, error) {
    tradeAgreementCollection, err := getTradeAgreementCollection(ctx)
    if err != nil {}
    tradeAgreementAsBytes, err := ctx.GetStub().GetPrivateData(tradeAgreementCollection, tradeId)
    if err != nil {}
    var tradeAgreement TradeAgreement
    err = json.Unmarshal(tradeAgreementAsBytes, &tradeAgreement)
    fmt.Println("\nTrade Agreement")
    fmt.Println(tradeAgreement)
    if err != nil {}
    return tradeAgreement.RevokeTime, nil
}
