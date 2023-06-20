package main

import "fmt"

type Strategy interface {
	execute(a, b int) int
}

// ConcreteStrategyAdd implements Strategy interface
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) execute(a, b int) int {
	return a + b
}

// ConcreteStrategySubtract implements Strategy interface
type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) execute(a, b int) int {
	return a - b
}

// ConcreteStrategyMultiply implements Strategy interface
type ConcreteStrategyMultiply struct{}

func (s *ConcreteStrategyMultiply) execute(a, b int) int {
	return a * b
}

// Context holds a reference to one of the concrete strategies and delegates it executing the behavior
type Context struct {
	strategy Strategy
}

func (c *Context) setStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) executeStrategy(a, b int) int {
	return c.strategy.execute(a, b)
}

func main() {
	context := &Context{}

	// Client code would go here. For this example, let's consider two numbers a and b
	a := 5
	b := 3

	// Add strategy
	context.setStrategy(&ConcreteStrategyAdd{})
	fmt.Println("Result of addition:", context.executeStrategy(a, b))

	// Subtract strategy
	context.setStrategy(&ConcreteStrategySubtract{})
	fmt.Println("Result of subtraction:", context.executeStrategy(a, b))

	// Multiply strategy
	context.setStrategy(&ConcreteStrategyMultiply{})
	fmt.Println("Result of multiplication:", context.executeStrategy(a, b))
}
