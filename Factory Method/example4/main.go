package main

func main() {
	ui := UIFramework{}

	// Create a square button in UIFramework
	ui.CreateButton()
	ui.Button.Click() // Prints "Square button click!"

	roundUI := UIWithRoundButtons{}

	// Create a round button in UIWithRoundButtons
	roundUI.CreateButton()
	roundUI.Button.Click() // Prints "Round button click!"
}
