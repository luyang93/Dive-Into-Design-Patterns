package main

import "fmt"

// MilkTea is a struct representing a milk tea, with price, topping, type of tea and amount of sugar.
type MilkTea struct {
	price   float64 // Price of the milk tea
	topping string  // Topping added to the milk tea
	tea     string  // Type of tea used in the milk tea
	sugar   int     // Amount of sugar in the milk tea
}

// MilkTeaBuilder is an interface that outlines the necessary steps to build a milk tea.
type MilkTeaBuilder interface {
	Reset()               // Reset the milk tea builder
	AddTopping()          // Add a topping to the milk tea
	AddTea()              // Add a type of tea to the milk tea
	AddSugar()            // Add sugar to the milk tea
	GetProduct() *MilkTea // Get the final milk tea product
}

// SignatureMilkTeaBuilder is a type that implements the MilkTeaBuilder interface for building a signature milk tea.
type SignatureMilkTeaBuilder struct {
	MilkTea *MilkTea // The milk tea that is being built
}

// Implement the interface methods for SignatureMilkTeaBuilder.
// Reset method initializes the SignatureMilkTeaBuilder with a new MilkTea object.
func (s *SignatureMilkTeaBuilder) Reset() {
	s.MilkTea = &MilkTea{
		price: 5.7,
	}
}

// AddTopping method sets the topping of the milk tea to "boba".
func (s *SignatureMilkTeaBuilder) AddTopping() {
	s.MilkTea.topping = "boba"
}

// AddTea method sets the tea of the milk tea to "signature tea".
func (s *SignatureMilkTeaBuilder) AddTea() {
	s.MilkTea.tea = "signature tea"
}

// AddSugar method sets the sugar level of the milk tea to 100.
func (s *SignatureMilkTeaBuilder) AddSugar() {
	s.MilkTea.sugar = 100
}

// GetProduct method returns the milk tea and prints the tea details.
func (s *SignatureMilkTeaBuilder) GetProduct() *MilkTea {
	fmt.Printf("Signature Milk Tea: %s %s %d\n", s.MilkTea.topping, s.MilkTea.tea, s.MilkTea.sugar)
	return s.MilkTea
}

// Similar steps are repeated for OolongMilkTeaBuilder and CustomizedMilkTeaBuilder
type OolongMilkTeaBuilder struct {
	MilkTea *MilkTea
}

func (o *OolongMilkTeaBuilder) Reset() {
	o.MilkTea = &MilkTea{
		price: 4.5,
	}
}

func (o *OolongMilkTeaBuilder) AddTopping() {
	o.MilkTea.topping = "boba"
}

func (o *OolongMilkTeaBuilder) AddTea() {
	o.MilkTea.tea = "oolong"
}

func (o *OolongMilkTeaBuilder) AddSugar() {
	o.MilkTea.sugar = 100
}

func (o *OolongMilkTeaBuilder) GetProduct() *MilkTea {
	fmt.Printf("Oolong Milk Tea: %s %s %d\n", o.MilkTea.topping, o.MilkTea.tea, o.MilkTea.sugar)
	return o.MilkTea
}

type CustomizedMilkTeaBuilder struct {
	MilkTea *MilkTea
}

func (c *CustomizedMilkTeaBuilder) Reset() {
	c.MilkTea = &MilkTea{
		price: 4.5,
	}
}

func (c *CustomizedMilkTeaBuilder) AddTopping(topping string) {
	c.MilkTea.topping = topping
}

func (c *CustomizedMilkTeaBuilder) AddTea(tea string) {
	c.MilkTea.tea = tea
}

func (c *CustomizedMilkTeaBuilder) AddSugar(sugar int) {
	c.MilkTea.sugar = sugar
}

func (c *CustomizedMilkTeaBuilder) GetProduct() *MilkTea {
	fmt.Printf("Customeized Milk Tea: %s %s %d\n", c.MilkTea.topping, c.MilkTea.tea, c.MilkTea.sugar)
	return c.MilkTea
}

type Director struct {
	builder MilkTeaBuilder // The MilkTeaBuilder that the Director is currently using
}

// NewDirector creates a new Director with the specified MilkTeaBuilder.
func NewDirector(builder MilkTeaBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// ChangeBuilder allows the Director to change its current MilkTeaBuilder.
func (d *Director) ChangeBuilder(builder MilkTeaBuilder) {
	d.builder = builder
}

// Make creates a new milk tea of the specified type.
// If the tea type is not "signature" or "oolong", it defaults to a custom milk tea.
func (d *Director) Make(teaType string) *MilkTea {
	switch teaType {
	case "signature":
		d.ChangeBuilder(&SignatureMilkTeaBuilder{})
	case "oolong":
		d.ChangeBuilder(&OolongMilkTeaBuilder{})
	default:
		// In the case of any other tea type, create a custom milk tea with "custom" topping and tea, and 10 sugar level.
		customBuilder := &CustomizedMilkTeaBuilder{}
		customBuilder.Reset()
		customBuilder.AddTopping("custom")
		customBuilder.AddTea("custom")
		customBuilder.AddSugar(10)
		return customBuilder.GetProduct()
	}
	// If the tea type is "signature" or "oolong", call the MakeMilkTea function to create the milk tea.
	return d.MakeMilkTea()
}

// MakeMilkTea creates a new milk tea using the current builder.
func (d *Director) MakeMilkTea() *MilkTea {
	d.builder.Reset()             // Reset the builder
	d.builder.AddTopping()        // Add the topping
	d.builder.AddTea()            // Add the tea
	d.builder.AddSugar()          // Add sugar
	return d.builder.GetProduct() // Get the final product
}

func main() {
	// Create a new Director with a SignatureMilkTeaBuilder.
	director := NewDirector(&SignatureMilkTeaBuilder{})
	director.MakeMilkTea() // Make a signature milk tea

	director.ChangeBuilder(&OolongMilkTeaBuilder{}) // Change the builder to OolongMilkTeaBuilder
	director.MakeMilkTea()                          // Make an oolong milk tea

	director.Make("signature") // Make a signature milk tea
	director.Make("custom")    // Make a custom milk tea

	// Create a customized milk tea with "green" topping and tea, and 50 sugar level.
	customBuilder := &CustomizedMilkTeaBuilder{}
	customBuilder.Reset()
	customBuilder.AddTopping("green")
	customBuilder.AddTea("green")
	customBuilder.AddSugar(50)
	customBuilder.GetProduct()
}
