package main

import "fmt"

type CelsiusTemperature interface {
	GetCelsius() float64
}

type FahrenheitTemprature struct {
	Fahrenheit float64
}

func (f *FahrenheitTemprature) GetFahrenheit() float64 {
	return f.Fahrenheit
}

type FahrenheitToCelsiusAdapter struct {
	FahrenheitTemprature
}

func (f *FahrenheitToCelsiusAdapter) GetCelsius() float64 {
	return (f.GetFahrenheit() - 32) * 5 / 9
}

func main() {
	f := &FahrenheitTemprature{100}
	c := &FahrenheitToCelsiusAdapter{*f}
	fmt.Printf("The temperature is %.2f degrees Celsius.\n", c.GetCelsius())
}
