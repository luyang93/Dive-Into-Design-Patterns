package main

import "fmt"

// Button interface defines methods that all buttons should have
type Button interface {
	render()
	onClick()
}

// WindowsButton is a concrete implementation of the Button interface, representing a button in the Windows environment
type WindowsButton struct{}

// render method for WindowsButton
func (wb *WindowsButton) render() {
	fmt.Println("Windows button render logic")
}

// onClick method for WindowsButton
func (wb *WindowsButton) onClick() {
	fmt.Println("Windows button onClick logic")
}

// HTMLButton is another concrete implementation of the Button interface, representing a button in the Web environment
type HTMLButton struct{}

// render method for HTMLButton
func (hb *HTMLButton) render() {
	fmt.Println("HTML button render logic")
}

// onClick method for HTMLButton
func (hb *HTMLButton) onClick() {
	fmt.Println("HTML button onClick logic")
}
