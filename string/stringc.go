package main

import "fmt"

type CreditCard string

type ConfidentialCustomer struct {
	CustomerId int64
	CreditCard CreditCard
}

func (c CreditCard) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func main() {
	c := ConfidentialCustomer{
		CustomerId: 1,
		CreditCard: "1234-5678-1234-5678",
	}

	fmt.Printf("%#v\n", c)
}
