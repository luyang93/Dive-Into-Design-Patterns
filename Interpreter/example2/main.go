package main

import "fmt"

type Expression interface {
	Interpret(*Context) int
}

type NumberExpression struct {
	Number int
}

func (n *NumberExpression) Interpret(c *Context) int {
	return n.Number
}

type VariableExpression struct {
	VariableName string
}

func (v *VariableExpression) Interpret(c *Context) int {
	return c.Get(v.VariableName)
}

type AddExpression struct {
	left  Expression
	right Expression
}

func (a *AddExpression) Interpret(c *Context) int {
	return a.left.Interpret(c) + a.right.Interpret(c)
}

type Context struct {
	Variables map[string]int
}

func (c *Context) Set(name string, value int) {
	c.Variables[name] = value
}

func (c *Context) Get(name string) int {
	return c.Variables[name]
}

func main() {
	context := &Context{Variables: make(map[string]int)}
	context.Set("x", 5)
	context.Set("y", 7)

	x := &VariableExpression{VariableName: "x"}
	y := &VariableExpression{VariableName: "y"}
	sum := &AddExpression{left: x, right: y}

	// Interpret the expression
	result := sum.Interpret(context)

	fmt.Printf("x + y = %d\n", result) // Output: x + y = 12
}
