package main

func main() {
	client := &Client{}
	macComputer := &Mac{}

	client.InsertLightningConnectorIntoComputer(macComputer)

	windowsComputer := &Windows{}
	WindowsComputerAdapter := &WindowsAdapter{
		windowsComputer: windowsComputer,
	}

	client.InsertLightningConnectorIntoComputer(WindowsComputerAdapter)
}
