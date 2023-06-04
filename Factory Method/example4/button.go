package main

import "fmt"

type Button interface {
	Click()
}

// SquareButton is a type of button
type SquareButton struct{}

func (sb *SquareButton) Click() {
	fmt.Println("Square button click!")
}

// RoundButton is another type of button
type RoundButton struct{}

func (rb *RoundButton) Click() {
	fmt.Println("Round button click!")
}
