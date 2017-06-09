package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/sterlingbates/blockchain/lib"
)

func main() {
	err := shim.Start(new(shared.Blocktrust))
	if err != nil {
		fmt.Printf("Error starting Blocktrust chaincode: %s", err)
	}
}
