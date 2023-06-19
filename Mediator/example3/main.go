package main

import "fmt"

type Mediator interface {
	send(Message string, Sender *Airplane)
}

type AirTrafficControl struct {
	Airplanes []*Airplane
}

func (a *AirTrafficControl) send(Message string, Sender *Airplane) {
	for _, airplane := range a.Airplanes {
		if airplane != Sender {
			airplane.Receive(Message)
		}
	}
}

func (a *AirTrafficControl) RegisterAirplane(airplane *Airplane) {
	a.Airplanes = append(a.Airplanes, airplane)
}

type Airplane struct {
	Name     string
	Mediator Mediator
}

func (a *Airplane) Send(Message string) {
	fmt.Println(a.Name, "sends Message:", Message)
	a.Mediator.send(Message, a)
}

func (a *Airplane) Receive(Message string) {
	fmt.Println(a.Name, "receives Message:", Message)
}

func main() {
	atc := &AirTrafficControl{}

	airplane1 := &Airplane{"Airplane1", atc}
	airplane2 := &Airplane{"Airplane2", atc}
	airplane3 := &Airplane{"Airplane3", atc}

	atc.RegisterAirplane(airplane1)
	atc.RegisterAirplane(airplane2)
	atc.RegisterAirplane(airplane3)

	airplane1.Send("Hello, everyone!")
	airplane2.Send("Hello, everyone!")
	airplane3.Send("Hello, everyone!")
}
