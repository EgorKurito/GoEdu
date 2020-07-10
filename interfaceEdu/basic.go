package main

import (
	"fmt"
)

// ----------------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money in wallet")
	}
	w.Cash -= amount
	return nil
}

// ----------------------

type Card struct {
	Balance		int
	ValidUntil	string
	Cardholder	string
	CVV			string
	Number		string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough money on card")
	}
	c.Balance -= amount
	return nil
}

// ----------------------

type ApplePay struct {
	Money	int
	AppleID	string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enough money on account")
	}
	a.Money -= amount
	return nil
}

// ----------------------

type Payer interface {
	Pay(int) error
}

// ----------------------

func Buy(p Payer)  {
	switch p.(type) {
	case *Wallet:
		fmt.Printf("Cash")
	case *Card:
		plasticCard, _ := p.(*Card)
		//if !ok {
		//	fmt.Println("To type *Card - fail")
		//}
		fmt.Printf("Insert Card, %s\n", plasticCard.Cardholder)
	default:
		fmt.Println("New!!!")
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Transaction error with - %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Thank for shopping with - %T\n\n", p)
}

// ----------------------

func main()  {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "EKurito"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}

