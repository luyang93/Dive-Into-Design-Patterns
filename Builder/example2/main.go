package main

import "fmt"

// Car represents a product to be built.
type Car struct {
	Seats        int
	Engine       string
	TripComputer bool
	GPS          bool
}

// Builder is the interface that specifies the steps to build the product.
type Builder interface {
	Reset()
	SetSeats(int)
	SetEngine(string)
	SetTripComputer(bool)
	SetGPS(bool)
	GetProduct() interface{}
}

// CarBuilder is a concrete builder that implements the Builder interface.
type CarBuilder struct {
	car *Car
}

// Reset method clears the current building process.
func (b *CarBuilder) Reset() {
	b.car = &Car{}
}

// SetSeats method sets the number of seats in the car
func (b *CarBuilder) SetSeats(seats int) {
	b.car.Seats = seats
}

// SetEngine method installs a specific engine
func (b *CarBuilder) SetEngine(engine string) {
	b.car.Engine = engine
}

// SetTripComputer method installs a trip computer
func (b *CarBuilder) SetTripComputer(hasTripComputer bool) {
	b.car.TripComputer = hasTripComputer
}

// SetGPS method installs a GPS
func (b *CarBuilder) SetGPS(hasGPS bool) {
	b.car.GPS = hasGPS
}

// GetProduct method returns the result of the building process.
func (b *CarBuilder) GetProduct() interface{} {
	result := b.car
	b.Reset()
	return result
}

type Manual struct {
	Description string
}

type ManualBuilder struct {
	manual *Manual
}

func (b *ManualBuilder) Reset() {
	b.manual = &Manual{}
}

func (b *ManualBuilder) SetSeats(seats int) {
	b.manual.Description += fmt.Sprintf("This car has %d seats.\n", seats)
}

func (b *ManualBuilder) SetEngine(engine string) {
	b.manual.Description += fmt.Sprintf("This car has a %s engine.\n", engine)
}

func (b *ManualBuilder) SetTripComputer(hasTripComputer bool) {
	if hasTripComputer {
		b.manual.Description += "This car has a trip computer.\n"
	}
}

func (b *ManualBuilder) SetGPS(hasGPS bool) {
	if hasGPS {
		b.manual.Description += "This car has a GPS.\n"
	}
}

func (b *ManualBuilder) GetProduct() interface{} {
	result := b.manual
	b.Reset()
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
	director.SetBuilder(carBuilder)

	// Construct sports car
	director.ConstructSportsCar()
	sportsCar := carBuilder.GetProduct()
	fmt.Println(sportsCar)

	// Construct SUV
	director.ConstructSUV()
	SUV := carBuilder.GetProduct()
	fmt.Println(SUV)

	manualBuilder := &ManualBuilder{}
	director.SetBuilder(manualBuilder)

	// Construct manual for sports car
	director.ConstructSportsCar()
	sportsCarManual := manualBuilder.GetProduct()
	fmt.Println(sportsCarManual)

	// Construct manual for SUV
	director.ConstructSUV()
	SUVManual := manualBuilder.GetProduct()
	fmt.Println(SUVManual)
}
