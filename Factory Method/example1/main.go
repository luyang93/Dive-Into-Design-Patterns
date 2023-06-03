package main

import "fmt"

// Factory
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
	logistics := createLogistics("sea")
	fmt.Println(logistics.planDelivery())
}
