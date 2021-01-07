package chaincode

import (
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) Test(ctx contractapi.TransactionContextInterface, arg string) (string, error) {
    fmt.Println(arg)
    argAsBytes := []byte(arg)
    fmt.Println(argAsBytes)
    err := ctx.GetStub().PutState("Key", argAsBytes)
    if err != nil {
        return "ERROR - ", fmt.Errorf(err.Error())
    }

    txId := ctx.GetStub().GetTxID()
    fmt.Println(txId)

    return txId, ctx.GetStub().SetEvent("EVENT", argAsBytes)
}
