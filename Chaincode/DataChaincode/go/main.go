package main

import (
	"fmt"
	"chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func printError (err error) {
	fmt.Println(err.Error())
}

func main() {
	mc, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	printError(err)

	err = mc.Start()
	printError(err)
}
