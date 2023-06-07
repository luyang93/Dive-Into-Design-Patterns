package main

import "fmt"

type Sedan interface {
	turnOnHeadLight()
}

type SUV interface {
	turnOnHeadLight()
}

type BMWM5 struct{}

func (b *BMWM5) turnOnHeadLight() {
	fmt.Println("BMW M5 headlight")
}

type BMWX5 struct{}

func (b *BMWX5) turnOnHeadLight() {
	fmt.Println("BMW X5 headlight")
}

type TeselaModelS struct{}

func (t *TeselaModelS) turnOnHeadLight() {
	fmt.Println("Tesela ModelS headlight")
}

type TeselaModelY struct{}

func (t *TeselaModelY) turnOnHeadLight() {
	fmt.Println("Tesela ModelY headlight")
}

type CarFactory interface {
	createSedan() Sedan
	createSUV() SUV
}

type BMWFactory struct{}

func (b *BMWFactory) createSedan() Sedan {
	return &BMWM5{}
}

func (b *BMWFactory) createSUV() SUV {
	return &BMWX5{}
}

type TeslaFactory struct{}

func (t *TeslaFactory) createSedan() Sedan {
	return &TeselaModelS{}
}

func (t *TeslaFactory) createSUV() SUV {
	return &TeselaModelY{}
}

type BrandBooth struct {
	Sedan Sedan
	SUV   SUV
}

func NewBrandBooth(factory CarFactory) *BrandBooth {
	return &BrandBooth{
		Sedan: factory.createSedan(),
		SUV:   factory.createSUV(),
	}
}

func (b *BrandBooth) ShowTime() {
	b.Sedan.turnOnHeadLight()
	b.SUV.turnOnHeadLight()
}

func main() {
	bmwFactory := &BMWFactory{}
	bmwBooth := NewBrandBooth(bmwFactory)
	bmwBooth.ShowTime()

	teslaFactory := &TeslaFactory{}
	teslaBooth := NewBrandBooth(teslaFactory)
	teslaBooth.ShowTime()
}
