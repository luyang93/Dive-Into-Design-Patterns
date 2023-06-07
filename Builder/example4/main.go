package main

import "fmt"

// Builder is the interface that specifies the steps to build the product.
type Builder interface {
	Reset()
	SetSeats(int)
	SetEngine(string)
	SetTripComputer(bool)
	SetGPS(bool)
}

// Car represents a product to be built.
type Car struct {
	Seats        int
	Engine       string
	TripComputer bool
	GPS          bool
}

// CarBuilder is a concrete builder that implements the Builder interface.
type CarBuilder struct {
	car *Car
}

// Reset method clears the current building process.
func (cb *CarBuilder) Reset() {
	cb.car = &Car{}
}

// SetSeats method sets the number of seats in the car
func (cb *CarBuilder) SetSeats(seats int) {
	cb.car.Seats = seats
}

// SetEngine method installs a specific engine
func (cb *CarBuilder) SetEngine(engine string) {
	cb.car.Engine = engine
}

// SetTripComputer method installs a trip computer
func (cb *CarBuilder) SetTripComputer(hasTripComputer bool) {
	cb.car.TripComputer = hasTripComputer
}

// SetGPS method installs a GPS
func (cb *CarBuilder) SetGPS(hasGPS bool) {
	cb.car.GPS = hasGPS
}

// GetProduct method returns the result of the building process.
func (cb *CarBuilder) GetProduct() *Car {
	result := cb.car
	cb.Reset()
	return result
}

type Manual struct {
	Description string
}

type ManualBuilder struct {
	manual *Manual
}

func (mb *ManualBuilder) Reset() {
	mb.manual = &Manual{}
}

func (mb *ManualBuilder) SetSeats(seats int) {
	mb.manual.Description += fmt.Sprintf("This car has %d seats.\n", seats)
}

func (mb *ManualBuilder) SetEngine(engine string) {
	mb.manual.Description += fmt.Sprintf("This car has a %s engine.\n", engine)
}

func (mb *ManualBuilder) SetTripComputer(hasTripComputer bool) {
	if hasTripComputer {
		mb.manual.Description += "This car has a trip computer.\n"
	}
}

func (mb *ManualBuilder) SetGPS(hasGPS bool) {
	if hasGPS {
		mb.manual.Description += "This car has a GPS.\n"
	}
}

func (mb *ManualBuilder) GetProduct() *Manual {
	result := mb.manual
	mb.Reset()
	return result
}

// Director is responsible for executing the building steps in a particular sequence.
type Director struct {
	builder Builder
}

// SetBuilder method sets the builder that the Director will use.
func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

// ConstructSportsCar method constructs a sports car.
func (d *Director) ConstructSportsCar() {
	d.builder.Reset()
	d.builder.SetSeats(2)
	d.builder.SetEngine("SportEngine")
	d.builder.SetTripComputer(true)
	d.builder.SetGPS(true)
}

// ConstructSUV method constructs an SUV.
func (d *Director) ConstructSUV() {
	d.builder.Reset()
	d.builder.SetSeats(5)
	d.builder.SetEngine("SUVEngine")
	d.builder.SetTripComputer(false)
	d.builder.SetGPS(true)
}

// Application simulates client code which uses Director and CarBuilderImpl
func main() {
	director := &Director{}
	carBuilder := &CarBuilder{}
	manualBuilder := &ManualBuilder{}

	// car
	director.SetBuilder(carBuilder)

	// Construct sports car
	director.ConstructSportsCar()
	sportsCar := carBuilder.GetProduct()
	fmt.Println(sportsCar)

	// Construct SUV
	director.ConstructSUV()
	SUV := carBuilder.GetProduct()
	fmt.Println(SUV)

	// manual
	director.SetBuilder(manualBuilder)

	// Construct sports car
	director.ConstructSportsCar()
	sportsCarManual := manualBuilder.GetProduct()
	fmt.Println(sportsCarManual)

	// Construct SUV
	director.ConstructSUV()
	SUVManual := manualBuilder.GetProduct()
	fmt.Println(SUVManual)
}
