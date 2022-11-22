package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ABstore struct {
	contractapi.Contract
}

type BlockerContract struct {
	Hash       string `json:"hash"`
	Contractor string `json:"contractor"`
	Date       string `json:"date"`
}

type BlockerCancleContract struct {
	Hash        string `json:"hash"`
	Cancle_Hash string `json:"cancle_hash"`
	Contractor  string `json:"contractor"`
	Date        string `json:"date"`
}

func (t *ABstore) Init(ctx contractapi.TransactionContextInterface, input_hash string, contractor string, date string) error {
	var err error
	var b_contract = BlockerContract{
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

func (t *ABstore) Query(ctx contractapi.TransactionContextInterface, input_hash string) (string, error) {
	var err error
	Avalbytes, err := ctx.GetStub().GetState(input_hash)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + input_hash + "\"}"
		return "", errors.New(jsonResp)
	}
	b_contract := BlockerContract{}
	json.Unmarshal(Avalbytes, &b_contract)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}

	buffer.WriteString("{\"Hash\":")
	buffer.WriteString("\"")
	buffer.WriteString(b_contract.Hash)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Contractor\":")
	buffer.WriteString("\"")
	buffer.WriteString(b_contract.Contractor)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Date\":")
	buffer.WriteString("\"")
	buffer.WriteString(b_contract.Date)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	jsonResp := "{\"Hash\":\"" + input_hash + "\",\"contrat_infor\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	return string(buffer.Bytes()), nil
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
