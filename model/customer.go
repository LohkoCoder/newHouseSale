package model

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	localUtils "github.com/newHouseSale/utils"
)

type Customer struct {
	Id string `json:"id"`// 身份证
	Name string `json:"name"`
	PhoneNum string `json:"phoneNum"`
	// 客户的购房状态：0、注册；1、开发商确认；2、交易中，支付首付；3、开发商确认收到首付金额；4、购房全款到账；5、开发商确认全款到账
	Status localUtils.CustomerStatus `json:"status"`
}

func (customer *Customer) Insert(ctx contractapi.TransactionContextInterface) (*Customer, error) {

	customerAsBytes, err := json.Marshal(customer)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal json array: %s ", err)
	}

	err = ctx.GetStub().PutState(customer.Id, customerAsBytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to put into world state: %s ", err)
	}

	return customer, nil
}

func (customer *Customer) Get(ctx contractapi.TransactionContextInterface) (*Customer, error) {

	customerAsBytes, err := ctx.GetStub().GetState(customer.Id)
	if err != nil {
		return nil, fmt.Errorf("Failed to query the world state: %s ", err.Error())
	}

	err = json.Unmarshal(customerAsBytes, customer)
	if err != nil {
		return nil, fmt.Errorf("Failed to query the world state: %s ", err.Error())
	}
	return customer, nil
}
