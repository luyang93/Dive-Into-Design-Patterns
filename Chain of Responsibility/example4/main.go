package main

import "fmt"

type DispenseChain interface {
	SetNextChain(nextChain DispenseChain)
	Dispense(currency *Currency)
}

type Currency struct {
	Amount int
}

func NewCurrency(amount int) *Currency {
	return &Currency{Amount: amount}
}

type Dollar50Dispenser struct {
	Chain DispenseChain
}

func (d *Dollar50Dispenser) SetNextChain(nextChain DispenseChain) {
	d.Chain = nextChain
}

func (d *Dollar50Dispenser) Dispense(currency *Currency) {
	numberOfNotes, remainder := currency.Amount/50, currency.Amount%50
	if numberOfNotes >= 1 {
		fmt.Printf("Dispensing %d 50$ note\n", numberOfNotes)
		if remainder != 0 {
			d.Chain.Dispense(NewCurrency(remainder))
		}
	} else {
		d.Chain.Dispense(currency)
	}
}

type Dollar20Dispenser struct {
	Chain DispenseChain
}

func (d *Dollar20Dispenser) SetNextChain(nextChain DispenseChain) {
	d.Chain = nextChain
}

func (d *Dollar20Dispenser) Dispense(currency *Currency) {
	numberOfNotes, remainder := currency.Amount/20, currency.Amount%20
	if numberOfNotes >= 1 {
		fmt.Printf("Dispensing %d 20$ note\n", numberOfNotes)
		if remainder != 0 {
			d.Chain.Dispense(NewCurrency(remainder))
		}
	} else {
		d.Chain.Dispense(currency)
	}
}

type Dollar10Dispenser struct {
	Chain DispenseChain
}

func (d *Dollar10Dispenser) SetNextChain(nextChain DispenseChain) {
	d.Chain = nextChain
}

func (d *Dollar10Dispenser) Dispense(currency *Currency) {
	if currency.Amount >= 10 {
		numberOfNotes := currency.Amount / 10
		remainder := currency.Amount % 10
		fmt.Printf("Dispensing %d 10$ note\n", numberOfNotes)
		if remainder != 0 {
			d.Chain.Dispense(NewCurrency(remainder))
		}
	} else {
		d.Chain.Dispense(currency)
	}
}

type ATMDispenser struct {
	C1 DispenseChain
}

func NewATMDispenser() *ATMDispenser {
	d := &ATMDispenser{}
	c1 := &Dollar50Dispenser{}
	c2 := &Dollar20Dispenser{}
	c3 := &Dollar10Dispenser{}

	c1.SetNextChain(c2)
	c2.SetNextChain(c3)

	d.C1 = c1
	return d
}

func (d *ATMDispenser) DispenseMoney(currency *Currency) {
	if currency.Amount%10 != 0 {
		fmt.Println("Amount should be in multiple of 10.")
		return
	}
	d.C1.Dispense(currency)
}

func main() {
	atmDispenser := NewATMDispenser()
	currency := NewCurrency(380)
	atmDispenser.DispenseMoney(currency)
}
