package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

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

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "setContract" {
		return s.setContract(APIstub, args)
	} else if function == "setCancleContract" {
		return s.setCancleContract(APIstub, args)
	} else if function == "func_verification" {
		return s.setCancleContract(APIstub, args)
	}
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func (s *SmartContract) setContract(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var b_contract = Blocker_contract{
		Hash:       args[0],
		Contractor: args[1],
		Date:       args[2],
	}
	ctrAsByte, _ := json.Marshal(b_contract)
	APIstub.PutState(args[0], ctrAsByte)
	return shim.Success(nil)
}

func (s *SmartContract) setCancleContract(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var b_contract = Blocker_cancle_contract{
		Hash:        args[0],
		Cancle_Hash: args[1],
		Contractor:  args[2],
		Date:        args[3],
	}
	ctrAsByte, _ := json.Marshal(b_contract)
	APIstub.PutState(args[1], ctrAsByte)
	return shim.Success(nil)
}

func (s *SmartContract) func_verification(APIstub shim.ChaincodeStubInterface, target_hash string) pb.Response {
	Avalbytes, err := APIstub.GetState(target_hash)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to verification this contract(" + target_hash + ")\"}"
		return shim.Error("Failed to create asset " + jsonResp)
	}

	jsonResp := "{\"hash\":\"" + target_hash + "\",\"json\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return string(Avalbytes), nil
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
