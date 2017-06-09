package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"blocktrust.org/lib"
)

func main() {
	err := shim.Start(new(Blocktrust))
	if err != nil {
		fmt.Printf("Error starting Blocktrust chaincode: %s", err)
	}
}
