package main

import "fmt"

type Memento struct {
	content string
}

func (m *Memento) getContent() string {
	return m.content
}

type TextEditor struct {
	content string
}

func (te *TextEditor) write(text string) {
	te.content += text
	fmt.Println("Current content: ", te.content)
}

func (te *TextEditor) save() *Memento {
	return &Memento{content: te.content}
}

func (te *TextEditor) restore(m *Memento) {
	te.content = m.content
}

func main() {
	// Create a text editor
	editor := &TextEditor{}

	// Edit text
	editor.write("This is the first sentence.")
	saved := editor.save() // Save current state

	editor.write(" This is the second sentence.")
	fmt.Println("Before restore: ", editor.content) // Print out current text

	// Restore to previous state
	editor.restore(saved)
	fmt.Println("After restore: ", editor.content) // Print out current text
}
