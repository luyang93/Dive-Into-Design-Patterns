package main

import "fmt"

type Item interface {
	Accept(visitor Visitor)
	GetPrice() float64
}

type Book struct {
	price float64
}

func (b *Book) GetPrice() float64 {
	return b.price
}

func (b *Book) Accept(visitor Visitor) {
	visitor.visitBook(b)
}

type Electronics struct {
	price float64
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) Accept(visitor Visitor) {
	visitor.visitElectronics(e)
}

type Visitor interface {
	visitBook(book *Book)
	visitElectronics(electronics *Electronics)
}

type RegularPriceVisitor struct {
	totalCost float64
}

func (v *RegularPriceVisitor) visitBook(book *Book) {
	v.totalCost += book.GetPrice()
}

func (v *RegularPriceVisitor) visitElectronics(electronics *Electronics) {
	v.totalCost += electronics.GetPrice()
}

func (v *RegularPriceVisitor) getTotalCost() float64 {
	return v.totalCost
}

type DiscountPriceVisitor struct {
	totalCost float64
}

func (v *DiscountPriceVisitor) visitBook(book *Book) {
	v.totalCost += book.GetPrice() * 0.8
}

func (v *DiscountPriceVisitor) visitElectronics(electronics *Electronics) {
	v.totalCost += electronics.GetPrice() * 0.9
}

func (v *DiscountPriceVisitor) getTotalCost() float64 {
	return v.totalCost
}

func main() {
	items := []Item{
		&Book{50},
		&Book{30},
		&Electronics{100},
	}

	regularVisitor := &RegularPriceVisitor{}
	discountVisitor := &DiscountPriceVisitor{}

	for _, item := range items {
		item.Accept(regularVisitor)
		item.Accept(discountVisitor)
	}

	fmt.Println("Regular total cost: ", regularVisitor.totalCost)
	fmt.Println("Discount total cost: ", discountVisitor.totalCost)
}
