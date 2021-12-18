package main

import (
	"gorpctest/calc"
	"log"
	"net/rpc"
)

type CalcClient struct {
	*rpc.Client
}

var _ calc.ServiceInterface = (*CalcClient)(nil)

func DialCalcService(network, address string) (*CalcClient, error) {
	c, err := rpc.DialHTTP(network, address)
	if err != nil {
		return nil, err
	}
	return &CalcClient{Client: c}, nil
}

// CalcTwoNumber 对两数进行运算
func (c *CalcClient) CalcTwoNumber(request calc.Calc, reply *float64) error {
	return c.Client.Call(calc.ServiceName+".CalcTwoNumber", request, reply)
}

func (c *CalcClient) GetOperators(request struct{}, reply *[]string) error {
	return c.Client.Call(calc.ServiceName+".GetOperators", request, reply)
}

func main() {
	client, err := DialCalcService("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Err Dial Client:", err)
	}

	// test GetOperators
	var opers []string
	err = client.GetOperators(struct{}{}, &opers)
	if err != nil {
		log.Println(err)
	}
	log.Println(opers)

	// test CalcTwoNumber
	testAdd := calc.Calc{
		Number1:  2.0,
		Number2:  3.14,
		Operator: "*",
	}
	var result float64
	err1 := client.CalcTwoNumber(testAdd, &result)
	if err1 != nil {
		log.Println(err1)
	}
	log.Println(result)

}
