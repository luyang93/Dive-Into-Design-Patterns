package main

import "fmt"

type Coffee interface {
	Cost() float64
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2
}

type CoffeeDecorator func(Coffee) Coffee

func WithMilk(coffee Coffee) Coffee {
	return &MilkDecorator{coffee}
}

func WithSugar(coffee Coffee) Coffee {
	return &SugarDecorator{coffee}
}

type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.25
}

func main() {
	coffee := &SimpleCoffee{}
	coffeeWithMilk := WithMilk(coffee)
	coffeeWithSugar := WithSugar(coffee)
	coffeeWithMilkAndSugar := WithSugar(coffeeWithMilk)

	fmt.Println("Coffee Cost:", coffee.Cost())
	fmt.Println("Coffee with Milk Cost:", coffeeWithMilk.Cost())
	fmt.Println("Coffee with Sugar Cost:", coffeeWithSugar.Cost())
	fmt.Println("Coffee with Milk and Sugar Cost:", coffeeWithMilkAndSugar.Cost())
}
