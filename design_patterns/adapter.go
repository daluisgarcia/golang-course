package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying with cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

// Creating a new struct that implements Payment interface differently
type BankPayment struct{}

func (BankPayment) Pay(bankAccount string) { // This method mismatch the interface method definition
	fmt.Printf("Paying with bank account: %s\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	BankAccount string
}

func (b BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.BankAccount)
}

// func main() {
// 	cash := &CashPayment{}
// 	ProcessPayment(cash)

// 	bank := &BankPayment{}
// 	// ProcessPayment(bank) // This will not work because the method signature mismatch

// 	bpa := BankPaymentAdapter{
// 		BankPayment: bank,
// 		BankAccount: "123456",
// 	}
// 	ProcessPayment(bpa) // This will work because the adapter struct implements the interface correctly

// }
