package main

import "fmt"

// Product interface
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Concrete product
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// Concrete product A
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Concrete product B
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// Factory
func getGun(gunType string) (IGun, error) {
	if gunType == "Ak47" {
		return newAk47(), nil
	}
	if gunType == "Musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")

}

// Client
func main() {
	ak47, _ := getGun("Ak47")
	musket, _ := getGun("Musket")

	printGunDetails(ak47)
	printGunDetails(musket)
}

func printGunDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
