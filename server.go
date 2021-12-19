package main

import (
	"fmt"
	calc2 "fpcxm911/rpccalc/server"
	"net/http"
	"net/rpc"

	"github.com/fpcxm911/rpccalc/calc"
)

//type Operation func(Number1, Number2 float64) float64
//
//// Add 是加法Operation的实现
//func Add(Number1, Number2 float64) float64 {
//	return Number1 + Number2
//}
//
//func Sub(Number1, Number2 float64) float64 {
//	return Number1 - Number2
//}
//
//func Mul(Number1, Number2 float64) float64 {
//	return Number1 * Number2
//}
//
//func Div(Number1, Number2 float64) float64 {
//	return Number1 / Number2
//}
//
//var Operators = map[string]Operation{
//	"+": Add,
//	"-": Sub,
//	"*": Mul,
//	"/": Div,
//}
//
//func CreateOperation(operator string) (Operation, error) {
//	var oper Operation
//	if oper, ok := Operators[operator]; ok {
//		return oper, nil
//	}
//	return oper, errors.New("Illegal Operator")
//}

// CalcService 是计算器rpc服务的实现
type CalcService struct{}

func (c *CalcService) CalcTwoNumber(request calc.Calc, reply *float64) error {
	oper, err := calc2.CreateOperation(request.Operator)
	if err != nil {
		return err
	}
	*reply = oper(request.Number1, request.Number2)
	return nil
}

func (c *CalcService) GetOperators(request struct{}, reply *[]string) error {
	opers := make([]string, 0, len(calc2.Operators))
	for key := range calc2.Operators {
		opers = append(opers, key)
	}
	*reply = opers
	return nil
}

func main() {
	err := calc.RegisterCalcService(new(CalcService))
	if err != nil {
		fmt.Println(err)
	}
	rpc.HandleHTTP()
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		fmt.Println(err2)
	}
}
