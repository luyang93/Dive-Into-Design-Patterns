package main

import "fmt"

// createDialog function creates an instance of Dialog based on the operating system type
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
	// Create a WindowsDialog instance and render it
	dialog, _ := createDialog("Windows")
	dialog.render()
}
