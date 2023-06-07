package main

import "fmt"

// getGun is a factory function that creates an instance of a gun based on the type passed
func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		// Returns an error if the gunType passed is not recognized
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
