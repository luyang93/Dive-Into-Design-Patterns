package main

import "fmt"

type Calculator struct {
	result int
}

func (c *Calculator) add(value int) {
	c.result += value
}

func (c *Calculator) subtract(value int) {
	c.result -= value
}

func (c *Calculator) save() CalculatorMemento {
	return CalculatorMemento{result: c.result}
}

func (c *Calculator) restore(m CalculatorMemento) {
	c.result = m.result
}

func (c *Calculator) getResult() int {
	return c.result
}

type CalculatorMemento struct {
	result int
}

type CalculatorHistory struct {
	history []CalculatorMemento
}

func (ch *CalculatorHistory) save(c Calculator) {
	ch.history = append(ch.history, c.save())
}

func (ch *CalculatorHistory) undo(c *Calculator) {
	if len(ch.history) > 0 {
		c.restore(ch.history[len(ch.history)-1])
		ch.history = ch.history[:len(ch.history)-1]
	}
}

func main() {
	calculator := Calculator{}
	history := CalculatorHistory{}

	calculator.add(5)
	calculator.subtract(3)
	history.save(calculator)                       // Save state
	fmt.Println("Result:", calculator.getResult()) // Output: Result: 2

	calculator.add(8)
	fmt.Println("Result:", calculator.getResult()) // Output: Result: 10

	history.undo(&calculator)                      // Undo to the previous saved state
	fmt.Println("Result:", calculator.getResult()) // Output: Result: 2
}
