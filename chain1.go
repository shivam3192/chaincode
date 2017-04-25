package main

import (
        "errors"
        "fmt"
       // "strconv"
        "encoding/json"
        "github.com/hyperledger/fabric/core/chaincode/shim"
)
type CrowdFundChaincode struct {
}
type Info struct {

        StudentRollNo string   `json:"studentrollno"`
        StudentName string `json:"StudentName"`
        StudentMarksSem1 string   `json:"studentmarkssem1"`
		StudentMarksSem2 string   `json:"studentmarkssem2"`
		StudentMarksSem3 string   `json:"studentmarkssem3"`
		StudentMarksSem4 string   `json:"studentmarkssem4"`
		BadgeInfo 
}
	type BadgeInfo struct {
	
        BadgeName       []string   `json:"Badgeame"`
        BadgeUrl        []string `json:"Badgeurl"`
        BadgeIssuedBy   []string   `json:"Badgeissuedby"`
        BadgeIssuedTo   []string `json:"Badgeissuedto"`
        //time 
}
func (t *CrowdFundChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        var err error

        if len(args) != 2 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }

     if err!=nil {
                        return nil, err
                }
         record := Info{}
        record.StudentRollNo="12"
        record.StudentName = "assa"
        record.StudentMarksSem1 = "99";
		record.StudentMarksSem1 = "98";
		record.StudentMarksSem1 = "97";
		record.StudentMarksSem1 = "96";
        
	    newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
                err=stub.PutState("default",newrecordByte);
         if err!=nil {
                        return nil, err
                }

        return nil, nil
}

func (t *CrowdFundChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    
var account string

        var err error

        if len(args) != 6 {
                return nil, errors.New("Incorrect number of arguments. Expecting 2.")
        }
          account = args[0]

         recordByte, err := stub.GetState(account);
        fmt.Println(recordByte);
        if err != nil {

            return nil, err
        }
        record := Info{}
        if recordByte != nil {
        errrecordmarshal := json.Unmarshal(recordByte,&record);
        if errrecordmarshal != nil {
            return nil, errrecordmarshal
        }            
        }
        record.StudentRollNo   =args[0];
        record.StudentName     =args[1];
        record.StudentMarksSem1=args[2];
		record.StudentMarksSem2=args[3];
		record.StudentMarksSem3=args[4];
		record.StudentMarksSem4=args[5];
        newrecordByte, err := json.Marshal(record);
        if err!=nil {

            return nil, err
        }
        err =stub.PutState(account,newrecordByte);
        if err != nil {

            return nil, err;
        } 
        return nil, nil
}
func (t *CrowdFundChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if function != "read" {
                return nil, errors.New("Invalid query function name. Expecting \"query\".")
        }
        var err error

         if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
        }

     var   account = args[0]
        accountValueBytes ,err := stub.GetState(account)
        if err != nil {
                 return nil, err
        }
        return accountValueBytes, nil
}

func main() {
        err := shim.Start(new(CrowdFundChaincode))

        if err != nil {
                fmt.Printf("Error starting CrowdFundChaincode: %s", err)
        }
}

