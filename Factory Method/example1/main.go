package main

import "fmt"

// createLogistics is a factory function that creates and returns a Logistics type based on the logisticsType input.
func createLogistics(logisticsType string) Logistics {
	switch logisticsType {
	case "road":
		return &RoadLogistics{}
	case "sea":
		return &SeaLogistics{}
	default:
		return nil
	}
}

// Client
func main() {
	logistics := createLogistics("sea")   // create a sea logistics instance.
	fmt.Println(logistics.planDelivery()) // plan delivery using the created logistics instance.
}
