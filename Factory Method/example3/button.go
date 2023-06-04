package main

import "fmt"

type Button interface {
	render()
	onClick()
}

type WindowsButton struct {
}

func (wb *WindowsButton) render() {
	fmt.Println("Windows button render logic")
}
func (wb *WindowsButton) onClick() {
	fmt.Println("Windows button onClick logic")
}

type HTMLButton struct {
}

func (hb *HTMLButton) render() {
	fmt.Println("HTML button render logic")
}
func (hb *HTMLButton) onClick() {
	fmt.Println("HTML button onClick logic")
}
