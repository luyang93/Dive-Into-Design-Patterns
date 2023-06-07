package main

import "fmt"

type MilkTea struct {
	price   float64
	topping string
	tea     string
	sugar   int
}

type MilkTeaBuilder interface {
	Reset()
	AddTopping()
	AddTea()
	AddSugar()
	GetProduct() *MilkTea
}

type SignatureMilkTeaBuilder struct {
	MilkTea *MilkTea
}

func (s *SignatureMilkTeaBuilder) Reset() {
	s.MilkTea = &MilkTea{
		price: 5.7,
	}
}

func (s *SignatureMilkTeaBuilder) AddTopping() {
	s.MilkTea.topping = "boba"
}

func (s *SignatureMilkTeaBuilder) AddTea() {
	s.MilkTea.tea = "signature tea"
}

func (s *SignatureMilkTeaBuilder) AddSugar() {
	s.MilkTea.sugar = 100
}

func (s *SignatureMilkTeaBuilder) GetProduct() *MilkTea {
	fmt.Printf("Signature Milk Tea: %s %s %d\n", s.MilkTea.topping, s.MilkTea.tea, s.MilkTea.sugar)
	return s.MilkTea
}

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
	builder MilkTeaBuilder
}

func NewDirector(builder MilkTeaBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) ChangeBuilder(builder MilkTeaBuilder) {
	d.builder = builder
}

func (d *Director) Make(teaType string) *MilkTea {
	switch teaType {
	case "signature":
		d.ChangeBuilder(&SignatureMilkTeaBuilder{})
	case "oolong":
		d.ChangeBuilder(&OolongMilkTeaBuilder{})
	default:
		panic("Invalid tea type!")
		return nil
	}
	return d.MakeMilkTea()
}

func (d *Director) MakeMilkTea() *MilkTea {
	d.builder.Reset()
	d.builder.AddTopping()
	d.builder.AddTea()
	d.builder.AddSugar()
	return d.builder.GetProduct()
}

func main() {
	director := NewDirector(&SignatureMilkTeaBuilder{})
	director.MakeMilkTea()

	director.ChangeBuilder(&OolongMilkTeaBuilder{})
	director.MakeMilkTea()

	director.Make("signature")
	director.Make("custom")

	customBuilder := &CustomizedMilkTeaBuilder{}
	customBuilder.Reset()
	customBuilder.AddTopping("custom")
	customBuilder.AddTea("custom")
	customBuilder.AddSugar(10)
	customBuilder.GetProduct()
}
