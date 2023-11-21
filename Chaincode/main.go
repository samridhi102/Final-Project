package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/kba-chf/final-project/chaincode/contracts"
)

func main() {
	ticketContract := new(contracts.TicketContract)

	// orderContract := new(contracts.OrderContract)

	chaincode, err := contractapi.NewChaincode(ticketContract)

	if err != nil {
		log.Panicf("Could not create chaincode." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		log.Panicf("Failed to start chaincode. " + err.Error())
	}
}
