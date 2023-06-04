package main

import "fmt"

type GUIFactory interface {
	createButton() Button
	createCheckBox() Checkbox
}

func getGUIFactory(platform string) (GUIFactory, error) {
	switch platform {
	case "MacOS":
		return &MacOSFactory{}, nil
	case "Windows":
		return &WindowsFactory{}, nil
	default:
		return nil, fmt.Errorf("")
	}
}

type MacOSFactory struct {
}

func (mf *MacOSFactory) createButton() Button {
	return &MacOSButton{}
}

func (mf *MacOSFactory) createCheckBox() Checkbox {
	return &MacOSCheckbox{}
}

type WindowsFactory struct {
}

func (wf *WindowsFactory) createButton() Button {
	return &WindowsButton{}
}

func (wf *WindowsFactory) createCheckBox() Checkbox {
	return &WindowsCheckbox{}
}

type Button interface {
	paint()
}

type MacOSButton struct {
}

func (mb *MacOSButton) paint() {
	fmt.Println("You have created MacOSButton.")
}

type WindowsButton struct {
}

func (wb *WindowsButton) paint() {
	fmt.Println("You have created WindowsButton.")
}

type Checkbox interface {
	paint()
}

type MacOSCheckbox struct {
}

func (mc *MacOSCheckbox) paint() {
	fmt.Println("You have created MacOSCheckbox.")
}

type WindowsCheckbox struct {
}

func (wc *WindowsCheckbox) paint() {
	fmt.Println("You have created WindowsCheckbox.")
}

func main() {
	factory, _ := getGUIFactory("Windows")

	button := factory.createButton()
	button.paint()

	checkbox := factory.createCheckBox()
	checkbox.paint()
}
