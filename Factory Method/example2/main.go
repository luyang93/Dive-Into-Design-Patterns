package main

import "fmt"

// IGun is an interface that specifies the structure of a Gun
type IGun interface {
	setName(name string) // method to set name of the gun
	setPower(power int)  // method to set power of the gun
	getName() string     // method to get name of the gun
	getPower() int       // method to get power of the gun
}

// Gun is a struct that implements IGun interface
type Gun struct {
	name  string // name of the gun
	power int    // power of the gun
}

// setName is a method to set the name of the gun
func (g *Gun) setName(name string) {
	g.name = name
}

// getName is a method to return the name of the gun
func (g *Gun) getName() string {
	return g.name
}

// setPower is a method to set the power of the gun
func (g *Gun) setPower(power int) {
	g.power = power
}

// getPower is a method to return the power of the gun
func (g *Gun) getPower() int {
	return g.power
}

// Ak47 is a struct that embeds the Gun struct
type Ak47 struct {
	Gun
}

// newAk47 is a function to create an instance of AK47 type
func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Musket is a struct that embeds the Gun struct
type Musket struct {
	Gun
}

// newMusket is a function to create an instance of Musket type
func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// getGun is a factory function that creates an instance of a gun based on the type passed
func getGun(gunType string) (IGun, error) {
	if gunType == "Ak47" {
		return newAk47(), nil
	}
	if gunType == "Musket" {
		return newMusket(), nil
	}
	// Returns an error if the gunType passed is not recognized
	return nil, fmt.Errorf("Wrong gun type passed")
}

// main is the entry point of the application
func main() {
	ak47, _ := getGun("Ak47")
	musket, _ := getGun("Musket")

	printGunDetails(ak47)
	printGunDetails(musket)
}

// printGunDetails is a function that prints the details of a gun
func printGunDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
