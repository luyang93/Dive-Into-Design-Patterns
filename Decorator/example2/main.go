package main

import "fmt"

type IPizza interface {
	GetPrice() int
}

type VeggeMania struct{}

func (p *VeggeMania) GetPrice() int {
	return 15
}

type peppyPaneer struct {
}

func (p *peppyPaneer) GetPrice() int {
	return 20
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

	veggiePizza := &VeggeMania{}

	//Add cheese topping
	veggiePizzaWithCheese := &CheeseTopping{
		pizza: veggiePizza,
	}

	//Add tomato topping
	veggiePizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.GetPrice())

	peppyPaneerPizza := &peppyPaneer{}

	//Add cheese topping
	peppyPaneerPizzaWithCheese := &CheeseTopping{
		pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithCheese.GetPrice())
}
