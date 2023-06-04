package main

import "fmt"

// Abstract Product
type Chair interface {
	SitOn()
}

// Concreate Product
type VictorianChair struct{}

func (v *VictorianChair) SitOn() {
	fmt.Println("Sitting on a Victorian chair")
}

type ModernChair struct{}

func (m *ModernChair) SitOn() {
	fmt.Println("Sitting on a Modern chair")
}

// Abstract Product
type Sofa interface {
	SitOn()
}

// Concreate Product
type VictorianSofa struct{}

func (v *VictorianSofa) SitOn() {
	fmt.Println("Sitting on a Victorian sofa")
}

type ModernSofa struct{}

func (m *ModernSofa) SitOn() {
	fmt.Println("Sitting on a Victorian sofa")
}

// Abstract Product
type CoffeeTable interface {
	PlaceOn()
}

// Concreate Product
type VictorianCoffeeTable struct{}

func (v *VictorianCoffeeTable) PlaceOn() {
	fmt.Println("Placing on a Victorian coffee table")
}

type ModernCoffeeTable struct{}

func (m *ModernCoffeeTable) PlaceOn() {
	fmt.Println("Placing on a modern coffee table")
}

// Abstract Factory
type FurnitureFactory interface {
	createChair() Chair
	createSofa() Sofa
	createCoffeeTable() CoffeeTable
}

// Concreate Factory
type VictorianFurnitureFactory struct{}

func (v *VictorianFurnitureFactory) createChair() Chair {
	return &VictorianChair{}
}

func (v *VictorianFurnitureFactory) createSofa() Sofa {
	return &VictorianSofa{}
}

func (v *VictorianFurnitureFactory) createCoffeeTable() CoffeeTable {
	return &VictorianCoffeeTable{}
}

// Concreate Factory
type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) createChair() Chair {
	return &ModernChair{}
}

func (m *ModernFurnitureFactory) createSofa() Sofa {
	return &ModernSofa{}
}

func (m *ModernFurnitureFactory) createCoffeeTable() CoffeeTable {
	return &ModernCoffeeTable{}
}

func getFurnitureFactory(style string) FurnitureFactory {
	switch style {
	case "Victorian":
		return &VictorianFurnitureFactory{}
	case "Modern":
		return &ModernFurnitureFactory{}
	default:
		return nil
	}
}

func main() {
	factory := getFurnitureFactory("Modern")

	chair := factory.createChair()
	chair.SitOn()

	sofa := factory.createSofa()
	sofa.SitOn()

	table := factory.createCoffeeTable()
	table.PlaceOn()
}
