package main

type Logistics interface {
	createTransport() Transport
	planDelivery() string
}

type RoadLogistics struct {
}

func (rl *RoadLogistics) createTransport() Transport {
	return &Trunk{}
}

func (rl *RoadLogistics) planDelivery() string {
	transport := rl.createTransport()
	return transport.Deliver()
}

type SeaLogistics struct {
}

func (sl *SeaLogistics) createTransport() Transport {
	return &Ship{}
}

func (sl *SeaLogistics) planDelivery() string {
	transport := sl.createTransport()
	return transport.Deliver()
}
