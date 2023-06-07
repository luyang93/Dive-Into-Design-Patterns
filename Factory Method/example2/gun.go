package main

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
