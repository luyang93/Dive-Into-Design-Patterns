package main

import "fmt"

type Observer interface {
	update(temperature float64)
}

type CurrentTemperatureDisplay struct{}

func (d *CurrentTemperatureDisplay) update(temperature float64) {
	fmt.Printf("Current Temperature: %.2f\n", temperature)
}

type TemperatureHistoryDisplay struct{}

func (t *TemperatureHistoryDisplay) update(temperature float64) {
	fmt.Printf("Temperature History Updated with: %.2f\n", temperature)
}

type Subject interface {
	addObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers()
}

type WeatherStation struct {
	temperature float64
	observers   []Observer
}

func (w *WeatherStation) addObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherStation) removeObserver(observer Observer) {
	for i, o := range w.observers {
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherStation) setTemperature(temperature float64) {
	w.temperature = temperature
	w.notifyObservers()
}

func (w *WeatherStation) notifyObservers() {
	for _, observer := range w.observers {
		observer.update(w.temperature)
	}
}

func main() {
	weatherStation := &WeatherStation{}
	currentTemperatureDisplay := &CurrentTemperatureDisplay{}
	temperatureHistoryDisplay := &TemperatureHistoryDisplay{}

	weatherStation.addObserver(currentTemperatureDisplay)
	weatherStation.addObserver(temperatureHistoryDisplay)

	weatherStation.setTemperature(25)
	weatherStation.setTemperature(26)
}
