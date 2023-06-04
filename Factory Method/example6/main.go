package main

import "fmt"

// Mail is the base class with the method deliver().
type Mail interface {
	Deliver()
}

// Transport represents different types of transport for mails.
type Transport interface {
	Transport()
}

// Plane implements Transport interface.
type Plane struct{}

func (p *Plane) Transport() {
	fmt.Println("Transporting by air.")
}

// Truck implements Transport interface.
type Truck struct{}

func (t *Truck) Transport() {
	fmt.Println("Transporting by ground - truck.")
}

// Train implements Transport interface.
type Train struct{}

func (t *Train) Transport() {
	fmt.Println("Transporting by ground - train.")
}

// AirMail is a subclass of Mail that uses Plane as transport.
type AirMail struct {
	transport Transport
}

func (am *AirMail) Deliver() {
	am.transport.Transport()
}

// GroundMail is a subclass of Mail that can use both Truck and Train as transport.
type GroundMail struct {
	transport Transport
}

func (gm *GroundMail) Deliver() {
	gm.transport.Transport()
}

// Factory method to create AirMail and GroundMail.
func NewMail(transportType string) Mail {
	switch transportType {
	case "AirMail":
		return &AirMail{&Plane{}}
	case "GroundMailTruck":
		return &GroundMail{&Truck{}}
	case "GroundMailTrain":
		return &GroundMail{&Train{}}
	default:
		return nil
	}
}

func main() {
	mailTypes := []string{"AirMail", "GroundMailTruck", "GroundMailTrain"}

	for _, mailType := range mailTypes {
		mail := NewMail(mailType)
		mail.Deliver()
	}
}
