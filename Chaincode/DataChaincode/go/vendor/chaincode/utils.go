package chaincode

import (
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/pkg/statebased"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ============================ UTILS =========================================

// ---------------------------keys for collection -------------------------

func generateKeyForInterestToken(tradeId string) string {
    return "TRADE_" + tradeId
}

func generateKeyForDevice(deviceId string) string {
    return "DEVICE_" + deviceId
}

// ----------------------Collection names---------------------------

func getMarketplaceCollection() (string, error) {
    return "collection_Marketplace", nil
}

//func getDealsCollection() (string, error) {
//    msp, err := shim.GetMSPID()
//    if err != nil {return "", err}
//
//    return msp + "_dealsCollection", nil
//}

func getTradeAgreementCollection() (string, error) {
    msp, err := shim.GetMSPID()
    if err != nil {return "", err}

    return msp + "_tradeAgreementCollection", nil
}

func getPrivateDetailsCollectionName() (string, error) {
    msp, err := shim.GetMSPID()
    if err != nil {return "", err}

    return msp + "_privateDetailsCollection", nil
}
func getACLCollection() (string, error) {
    msp, err := shim.GetMSPID()
    if err != nil {return "", err}

    return msp + "_aclCollection", nil
}

// ------------------------------------------------------------------------

func verifyClientOrgMatchesPeerOrg(ctx contractapi.TransactionContextInterface) error {
	clientMSP, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {}

	peerMSP, err := shim.GetMSPID();
	if err != nil {}

	if clientMSP != peerMSP {
		return fmt.Errorf("client MSP %v does not match PeerMSP %v", clientMSP, peerMSP)
	}
	return nil
}


func setDeviceStateBasedEndorsement(ctx contractapi.TransactionContextInterface, deviceKey string, orgId string, collection string) error {
    // create a new state based policy for key = deviceId
    ep, err := statebased.NewStateEP(nil)
    if err != nil {}

    // issue roles, here the owner org for a device
    err = ep.AddOrgs(statebased.RoleTypePeer, orgId)
    if err != nil {}

    policy, err := ep.Policy()
    if err != nil {}

    err = ctx.GetStub().SetPrivateDataValidationParameter(collection, deviceKey, policy)
    return nil
}


// updateDeviceDescription - error
// addDeviceData

// createTradeAgreement
// createInterestToken
// InvokeDataSharing(BiddersInterestToken) <--> asset transfer
