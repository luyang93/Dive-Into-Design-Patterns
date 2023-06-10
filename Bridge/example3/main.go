package main

import "fmt"

type Payment interface {
	ProcessPayment()
}

type CreditCardPayment struct{}

func (c CreditCardPayment) ProcessPayment() {
	fmt.Println("Processing credit card payment")
}

type WeChatPayment struct{}

func (w WeChatPayment) ProcessPayment() {
	fmt.Println("Processing wechat payment")
}

type CashPayment struct{}

func (c CashPayment) ProcessPayment() {
	fmt.Println("Processing cash payment")
}

type Product interface {
	Purchase()
}

type Book struct {
	payment Payment
}

func NewBook(payment Payment) *Book {
	return &Book{payment: payment}
}

func (b *Book) Purchase() {
	b.payment.ProcessPayment()
	fmt.Println("Purchased book")
}

type Electronics struct {
	payment Payment
}

func NewElectronics(payment Payment) *Electronics {
	return &Electronics{payment: payment}
}

func (e *Electronics) Purchase() {
	e.payment.ProcessPayment()
	fmt.Println("Purchased electronics")
}

func main() {
	creditCardPayment := &CreditCardPayment{}
	wechatPayment := &WeChatPayment{}
	cashPayment := &CashPayment{}

	book := NewBook(creditCardPayment)
	book.Purchase()

	electronics := NewElectronics(creditCardPayment)
	electronics.Purchase()

	book2 := NewBook(wechatPayment)
	book2.Purchase()

	book3 := NewBook(cashPayment)
	book3.Purchase()
}
