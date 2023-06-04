package main

type Dialog interface {
	createButton() Button
	render()
}

type WebDialog struct {
}

func (hd *WebDialog) createButton() Button {
	return &HTMLButton{}
}

func (hd *WebDialog) render() {
	button := hd.createButton()
	button.render()
}

type WindowsDialog struct {
}

func (wd *WindowsDialog) createButton() Button {
	return &WindowsButton{}
}

func (wd *WindowsDialog) render() {
	button := wd.createButton()
	button.render()
}
