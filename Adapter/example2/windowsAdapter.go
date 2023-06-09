package main

import "fmt"

type WindowsAdapter struct {
	windowComputer *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowComputer.InsertIntoUSBPort()
}
