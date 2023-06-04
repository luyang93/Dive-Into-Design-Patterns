package main

// Logistics is an interface that declares the methods for planning delivery and creating transportation.
type Logistics interface {
	createTransport() Transport
	planDelivery() string
}

// RoadLogistics is a struct that represents logistics conducted on roads.
type RoadLogistics struct{}

// createTransport returns a new instance of a Transport type.
func (rl *RoadLogistics) createTransport() Transport {
	return &Trunk{}
}

// planDelivery orchestrates the delivery process and returns a delivery message.
func (rl *RoadLogistics) planDelivery() string {
	transport := rl.createTransport()
	return transport.Deliver()
}

// SeaLogistics is a struct that represents logistics conducted by sea.
type SeaLogistics struct{}

// createTransport returns a new instance of a Transport type.
func (sl *SeaLogistics) createTransport() Transport {
	return &Ship{}
}

// planDelivery orchestrates the delivery process and returns a delivery message.
func (sl *SeaLogistics) planDelivery() string {
	transport := sl.createTransport()
	return transport.Deliver()
}
