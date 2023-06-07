package main

import "fmt"

// main is the entry point of the application
func main() {
	ak47, _ := getGun("ak47")
	if ak47 != nil {
		printGunDetails(ak47)
	}

	musket, _ := getGun("musket")
	if musket != nil {
		printGunDetails(musket)
	}
}

// printGunDetails is a function that prints the details of a gun
func printGunDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
