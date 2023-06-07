package main

import "fmt"

type Builder interface {
	Reset()
	BuildStep1(string)
	BuildStep2(string)
	BuildStep3(string)
}

type Product1 struct {
	FeatureA1 string
	FeatureB1 string
	FeatureC1 string
}

type ConcreteBuilder1 struct {
	Product *Product1
}

// Reset resets the product to its initial state.
func (cb *ConcreteBuilder1) Reset() {
	cb.Product = &Product1{}
}

// BuildStep1 builds the first feature of the product.
func (cb *ConcreteBuilder1) BuildStep1(feature string) {
	cb.Product.FeatureA1 = feature + "1"
}

// BuildStep2 builds the second feature of the product.
func (cb *ConcreteBuilder1) BuildStep2(feature string) {
	cb.Product.FeatureB1 = feature + "1"
}

// BuildStep3 builds the third feature of the product.
func (cb *ConcreteBuilder1) BuildStep3(feature string) {
	cb.Product.FeatureC1 = feature + "1"
}

// GetProduct returns the product built by the builder.
func (cb *ConcreteBuilder1) GetProduct() *Product1 {
	return cb.Product
}

type Product2 struct {
	FeatureA2 string
	FeatureB2 string
	FeatureC2 string
}

type ConcreteBuilder2 struct {
	Product *Product2
}

// Reset resets the product to its initial state.
func (cb *ConcreteBuilder2) Reset() {
	cb.Product = &Product2{}
}

// BuildStep1 builds the first feature of the product.
func (cb *ConcreteBuilder2) BuildStep1(feature string) {
	cb.Product.FeatureA2 = feature + "2"
}

// BuildStep2 builds the second feature of the product.
func (cb *ConcreteBuilder2) BuildStep2(feature string) {
	cb.Product.FeatureB2 = feature + "2"
}

// BuildStep3 builds the third feature of the product.
func (cb *ConcreteBuilder2) BuildStep3(feature string) {
	cb.Product.FeatureC2 = feature + "2"
}

// GetProduct returns the product built by the builder.
func (cb *ConcreteBuilder2) GetProduct() *Product2 {
	return cb.Product
}

type Director struct {
	Builder Builder
}

// NewDirector creates a new Director instance with the given builder.
func NewDirector(Builder Builder) *Director {
	return &Director{
		Builder: Builder,
	}
}

// SetBuilder sets the builder for the director.
func (d *Director) SetBuilder(Builder Builder) {
	d.Builder = Builder
}

// MakeProduct orchestrates the building process for the product.
func (d *Director) MakeProduct() {
	d.Builder.Reset()
	d.Builder.BuildStep1("Feature-A-")
	d.Builder.BuildStep2("Feature-B-")
	d.Builder.BuildStep3("Feature-C-")
}

func main() {
	director := &Director{}
	product1Builder := &ConcreteBuilder1{}
	product2Builder := &ConcreteBuilder2{}

	director.SetBuilder(product1Builder)
	director.MakeProduct()
	product1 := product1Builder.GetProduct()
	fmt.Println(*product1)

	director.SetBuilder(product2Builder)
	director.MakeProduct()
	product2 := product2Builder.GetProduct()
	fmt.Println(*product2)
}
