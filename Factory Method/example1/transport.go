package main

// Transport is an interface that declares a method for delivering goods.
type Transport interface {
	Deliver() string
}

// Ship is a struct that represents a mode of transport over sea.
type Ship struct {
}

// Deliver orchestrates the delivery process and returns a delivery message.
func (s Ship) Deliver() string {
	return "Delivery by sea is in progress."
}

// Trunk is a struct that represents a mode of transport over road.
type Trunk struct {
}

// Deliver orchestrates the delivery process and returns a delivery message.
func (t *Trunk) Deliver() string {
	return "Delivery by road is in progress."
}
