package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ABstore struct {
	contractapi.Contract
}

type Blocker_contract struct {
	Hash       string `json:"hash"`
	Contractor string `json:"contractor"`
	Date       string `json:"date"`
}

type Blocker_cancle_contract struct {
	Hash        string `json:"hash"`
	Cancle_Hash string `json:"cancle_hash"`
	Contractor  string `json:"contractor"`
	Date        string `json:"date"`
}

func (t *ABstore) setContract(ctx contractapi.TransactionContextInterface, input_hash string, contractor string, date string) error {
	var err error
	var b_contract = Blocker_contract{
		Hash:       input_hash,
		Contractor: contractor,
		Date:       date,
	}
	ctrAsByte, _ := json.Marshal(b_contract)

	err = ctx.GetStub().PutState(input_hash, ctrAsByte)
	if err != nil {
		return err
	}

	return nil
}

func (t *ABstore) setCancleContract(ctx contractapi.TransactionContextInterface, input_hash string, input_cancle_hash string, contractor string, date string) error {
	var err error
	var b_contract = Blocker_cancle_contract{
		Hash:        input_hash,
		Cancle_Hash: input_cancle_hash,
		Contractor:  contractor,
		Date:        date,
	}
	ctrAsByte, _ := json.Marshal(b_contract)

	err = ctx.GetStub().PutState(input_cancle_hash, ctrAsByte)
	if err != nil {
		return err
	}

	return nil
}

func (t *ABstore) func_verification(ctx contractapi.TransactionContextInterface, target_hash string) (string, error) {
	var err error
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetState(target_hash)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to verification this contract(" + target_hash + ")\"}"
		return "", errors.New(jsonResp)
	}

	jsonResp := "{\"hash\":\"" + target_hash + "\",\"json\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return string(Avalbytes), nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(ABstore))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting ABstore chaincode: %s", err)
	}
}
