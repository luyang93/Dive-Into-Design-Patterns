package main

import "fmt"

type WindowsAdapter struct {
	Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.InsertIntoUSBPort()
}