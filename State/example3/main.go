package main

import "fmt"

type TrafficLightState interface {
	handle(*TrafficLight)
}

type RedState struct{}

func (r *RedState) handle(trafficLight *TrafficLight) {
	fmt.Println("Red Light: Stopped")
	trafficLight.setState(&GreenState{})
}

type YellowState struct{}

func (y *YellowState) handle(trafficLight *TrafficLight) {
	fmt.Println("Yellow Light: Be prepared to stop")
	trafficLight.setState(&RedState{})
}

type GreenState struct{}

func (g *GreenState) handle(trafficLight *TrafficLight) {
	fmt.Println("Green Light: Go")
	trafficLight.setState(&YellowState{})
}

type TrafficLight struct {
	state TrafficLightState
}

func (tl *TrafficLight) setState(state TrafficLightState) {
	tl.state = state
}

func (tl *TrafficLight) change() {
	tl.state.handle(tl)
}

func main() {
	trafficLight := &TrafficLight{state: &RedState{}}

	for i := 0; i < 6; i++ {
		trafficLight.change()
	}
}
