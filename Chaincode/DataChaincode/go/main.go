
package main

import (
	"fmt"
	"chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


func main() {
	mc, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil { fmt.Println(err.Error()) }

	err = mc.Start()
	if err != nil { fmt.Println(err.Error()) }
}
