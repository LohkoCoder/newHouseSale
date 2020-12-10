package ChainCode

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	. "github.com/newHouseSale/model"
	localUtils "github.com/newHouseSale/utils"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) string {
	return "init"
}

func (s *SmartContract) QueryOrder(ctx contractapi.TransactionContextInterface) (*Order, error) {
	order := new(Order)

	order, err := order.ParseArgsWithCustomer(ctx)
	if err != nil {
		return nil, err
	}
	return order.Get(ctx)
}

func (s *SmartContract) RegisterOrder(ctx contractapi.TransactionContextInterface) (*Order, error) {
	order := new(Order)

	order, err := order.ParseArgsWithCustomer(ctx)
	if err != nil {
		return nil, err
	}

	order.Customer.Status = localUtils.Registered
	return order.Insert(ctx)
}

func (s *SmartContract) ConfirmOrder(ctx contractapi.TransactionContextInterface) (*Order, error)  {
	order := new(Order)

	order, err := order.ParseArgsWithCustomer(ctx)
	if err != nil {
		return nil, err
	}

	order, err = order.Get(ctx)
	if err != nil {
		return nil, err
	}

	if order.Customer.Status != localUtils.Registered {
		return nil, fmt.Errorf("customer status not valid, should be '%s' ", localUtils.Registered)
	}

	order.Customer.Status = localUtils.ConfirmRegistered
	return order.Insert(ctx)
}

func (s *SmartContract) DownPayment(ctx contractapi.TransactionContextInterface) (*Order, error)  {
	order := new(Order)

	order, err := order.ParseArgsWithCustomerAndHouse(ctx)
	if err != nil {
		return nil, err
	}

	_, err = order.Get(ctx)
	if err != nil {
		return nil, err
	}

	if order.Customer.Status != localUtils.ConfirmRegistered {
		return nil, fmt.Errorf("customer status not valid, should be '%s' ", localUtils.ConfirmRegistered)
	}

	order, _ = order.ParseArgsWithCustomerAndHouse(ctx)
	order.Customer.Status = localUtils.DownPayment
	return order.Insert(ctx)
}

func (s *SmartContract) ConfirmDownPayment(ctx contractapi.TransactionContextInterface) (*Order, error)  {
	order := new(Order)
	order, err := order.ParseArgsWithCustomer(ctx)
	if err != nil {
		return nil, err
	}

	order, err = order.Get(ctx)
	if err != nil {
		return nil, err
	}

	if order.Customer.Status != localUtils.DownPayment {
		return nil, fmt.Errorf("customer status not valid, should be '%s' ", localUtils.DownPayment)
	}

	order.Customer.Status = localUtils.ConfirmDownPayment
	return order.Insert(ctx)
}

func (s *SmartContract) FullPayment(ctx contractapi.TransactionContextInterface) (*Order, error)  {
	order := new(Order)

	order, err := order.ParseArgsWithCustomer(ctx)
	if err != nil {
		return nil, err
	}

	order, err = order.Get(ctx)
	if err != nil {
		return nil, err
	}
	if order.Customer.Status != localUtils.ConfirmDownPayment {
		return nil, fmt.Errorf("customer status not valid, should be '%s' ", localUtils.ConfirmDownPayment)
	}

	order.Customer.Status = localUtils.FullPayment
	return order.Insert(ctx)
}

func (s *SmartContract) ConfirmFullPayment(ctx contractapi.TransactionContextInterface) (*Order, error)  {
	order := new(Order)

	order, err := order.ParseArgsWithCustomer(ctx)
	if err != nil {
		return nil, err
	}

	order, err = order.Get(ctx)
	if err != nil {
		return nil, err
	}
	if order.Customer.Status != localUtils.FullPayment {
		return nil, fmt.Errorf("customer status not valid, should be '%s' ", localUtils.FullPayment)
	}

	order.Customer.Status = localUtils.ConfirmFullPayment

	return order.Insert(ctx)
}
