
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {	
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initStringSha" {
		return t.initStringSha(stub, args)
	} else if function == "initFileSha" {
		return t.initFileSha(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) initStringSha(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0       
	// "stringSha"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	// ==== Input sanitation ====
	fmt.Println("- start init StringSha")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	dataId := args[0]

	// ==== Create stringSha object and marshal to JSON ====
	ObjectType := "stringSha"
	stringSha := &StringSha{ObjectType, dataId}
	stringShaJSONasBytes, err := json.Marshal(stringSha)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save stringSha to state ===
	err = stub.PutState(dataId, stringShaJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Order saved and indexed. Return success ====
	fmt.Println("- end init stringSha")
	return shim.Success(nil)
}

func (t *SimpleChaincode) initFileSha(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0       
	// "fileSha"
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	// ==== Input sanitation ====
	fmt.Println("- start init FileSha")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	dataId := args[0]

	// ==== Create fileSha object and marshal to JSON ====
	ObjectType := "fileSha"
	fileSha := &FileSha{ObjectType, dataId}
	fileShaJSONasBytes, err := json.Marshal(fileSha)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save fileSha to state ===
	err = stub.PutState(dataId, fileShaJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Order saved and indexed. Return success ====
	fmt.Println("- end init fileSha")
	return shim.Success(nil)
}
