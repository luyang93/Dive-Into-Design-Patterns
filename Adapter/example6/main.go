package main

import "fmt"

type CelsiusTemperature interface {
	GetCelsiusTemperature() float64
}

type FahrenheitTemperature interface {
	GetFahrenheitTemperature() float64
}

type Celsius struct {
	temperature float64
}

func (c *Celsius) GetCelsiusTemperature() float64 {
	return c.temperature
}

type Fahrenheit struct {
	temperature float64
}

func (f *Fahrenheit) GetFahrenheitTemperature() float64 {
	return f.temperature
}

type TwoWayAdapter struct {
	CelsiusTemperature
	FahrenheitTemperature
}

func NewTwoWayAdapterCelsius(c *Celsius) TwoWayAdapter {
	f := &Fahrenheit{c.temperature*9/5 + 32}
	return TwoWayAdapter{
		c,
		f,
	}
}

func NewTwoWayAdapterFahrenheit(f *Fahrenheit) TwoWayAdapter {
	c := &Celsius{(f.temperature - 32) * 5 / 9}
	return TwoWayAdapter{
		c,
		f,
	}
}

func main() {
	celsius := &Celsius{25}
	fahrenheit := &Fahrenheit{100}

	adapter1 := NewTwoWayAdapterCelsius(celsius)
	adapter2 := NewTwoWayAdapterFahrenheit(fahrenheit)

	fmt.Println("Celsius temperature:", adapter1.GetCelsiusTemperature())
	fmt.Println("Fahrenheit temperature:", adapter1.GetFahrenheitTemperature())

	fmt.Println("Celsius temperature:", adapter2.GetCelsiusTemperature())
	fmt.Println("Fahrenheit temperature:", adapter2.GetFahrenheitTemperature())
}
