package main

// IGun is an interface that specifies the structure of a Gun
type IGun interface {
	setName(name string) // method to set name of the gun
	setPower(power int)  // method to set power of the gun
	getName() string     // method to get name of the gun
	getPower() int       // method to get power of the gun
}
