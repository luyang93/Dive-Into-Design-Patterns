package main

import "fmt"

type IPizza interface {
	GetPrice() int
}

type VeggeMania struct{}

func (p *VeggeMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.pizza.GetPrice()
	return pizzaPrice + 10
}

func main() {

	pizza := &VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
