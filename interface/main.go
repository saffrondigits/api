package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

type PayPal struct {
	Email string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using credit card ending with %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using using PayPal account %s", amount, pp.Email)
}

func ProcessPayment(pm PaymentMethod, amount float64) {
	fmt.Println(pm.Pay(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234567890"}
	pp := PayPal{Email: "dave@example.com"}

	ProcessPayment(cc, 100)
	ProcessPayment(pp, 130)
}

// Learn methods in Golang
// Learn interface in Golang
