package main

type Transport interface {
	Deliver() string
}

type Ship struct {
}

func (s Ship) Deliver() string {
	return "Delivery by sea is in progress."
}

type Trunk struct {
}

func (t *Trunk) Deliver() string {
	return "Delivery by road is in progress."
}
