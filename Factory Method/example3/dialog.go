package main

// Dialog interface defines methods that all dialogs should have
type Dialog interface {
	createButton() Button
	render()
}

// WebDialog is a concrete implementation of the Dialog interface, representing a dialog in the Web environment
type WebDialog struct{}

// createButton method for WebDialog, creates a HTMLButton
func (hd *WebDialog) createButton() Button {
	return &HTMLButton{}
}

// render method for WebDialog, renders the button in the dialog
func (hd *WebDialog) render() {
	button := hd.createButton()
	button.render()
}

// WindowsDialog is another concrete implementation of the Dialog interface, representing a dialog in the Windows environment
type WindowsDialog struct{}

// createButton method for WindowsDialog, creates a WindowsButton
func (wd *WindowsDialog) createButton() Button {
	return &WindowsButton{}
}

// render method for WindowsDialog, renders the button in the dialog
func (wd *WindowsDialog) render() {
	button := wd.createButton()
	button.render()
}
