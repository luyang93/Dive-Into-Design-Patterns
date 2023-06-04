package main

import "fmt"

func createDialog(OS string) (Dialog, error) {
	switch OS {
	case "Windows":
		return &WindowsDialog{}, nil
	case "Web":
		return &WebDialog{}, nil
	default:
		return nil, fmt.Errorf("")
	}
}

func main() {
	dialog, _ := createDialog("Windows")
	dialog.render()
}
