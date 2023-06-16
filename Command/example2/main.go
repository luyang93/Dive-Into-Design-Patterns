package main

import "fmt"

type Application struct {
	Clipboard    string
	Editors      []Editor
	ActiveEditor *Editor
	History      CommandHistory
}

type Editor struct {
	Text string
}

// Editor methods
func (e *Editor) GetSelection() string {
	// return selected text
	return e.Text
}

func (e *Editor) DeleteSelection() {
	// delete selected text
	e.Text = ""
}

func (e *Editor) ReplaceSelection(text string) {
	// replace selected text
	e.Text = text
}

type CommandHistory struct {
	History []Command
}

type Command interface {
	Execute() bool
	Undo()
}

type CopyCommand struct {
	App    *Application
	Editor *Editor
	Backup string
}

// CopyCommand methods
func (c *CopyCommand) Execute() bool {
	c.App.Clipboard = c.Editor.GetSelection()
	return false
}

func (c *CopyCommand) Undo() {
	c.Editor.Text = c.Backup
}

type CutCommand struct {
	App    *Application
	Editor *Editor
	Backup string
}

// CutCommand methods
func (c *CutCommand) Execute() bool {
	c.Backup = c.Editor.Text
	c.App.Clipboard = c.Editor.GetSelection()
	c.Editor.DeleteSelection()
	return true
}

func (c *CutCommand) Undo() {
	c.Editor.Text = c.Backup
}

type PasteCommand struct {
	App    *Application
	Editor *Editor
	Backup string
}

// PasteCommand methods
func (c *PasteCommand) Execute() bool {
	c.Backup = c.Editor.Text
	c.Editor.ReplaceSelection(c.App.Clipboard)
	return true
}

func (c *PasteCommand) Undo() {
	c.Editor.Text = c.Backup
}

type UndoCommand struct {
	App    *Application
	Editor *Editor
	Backup string
}

// UndoCommand methods
func (u *UndoCommand) Execute() bool {
	u.App.Undo()
	return false
}

func (u *UndoCommand) Undo() {
	u.Editor.Text = u.Backup
}

// CommandHistory methods
func (h *CommandHistory) Push(c Command) {
	h.History = append(h.History, c)
}

func (h *CommandHistory) Pop() Command {
	length := len(h.History)
	if length == 0 {
		return nil
	}

	lastCommand := h.History[length-1]
	h.History = h.History[:length-1]
	return lastCommand
}

func (app *Application) ExecuteCommand(command Command) {
	if command.Execute() {
		app.History.Push(command)
	}
}

func (app *Application) Undo() {
	command := app.History.Pop()
	if command != nil {
		command.Undo()
	}
}

func main() {
	app := Application{
		Editors:      []Editor{{Text: "Hello, world!"}},
		ActiveEditor: &Editor{Text: "Hello, world!"},
		History:      CommandHistory{},
	}

	copyCommand := &CopyCommand{
		App:    &app,
		Editor: app.ActiveEditor,
	}
	app.ExecuteCommand(copyCommand)
	fmt.Println("Clipboard:", app.Clipboard)

	cutCommand := &CutCommand{
		App:    &app,
		Editor: app.ActiveEditor,
	}
	app.ExecuteCommand(cutCommand)
	fmt.Println("Clipboard:", app.Clipboard)
	fmt.Println("ActiveEditor:", app.ActiveEditor.Text)

	pasteCommand := &PasteCommand{
		App:    &app,
		Editor: app.ActiveEditor,
	}
	app.ExecuteCommand(pasteCommand)
	fmt.Println("ActiveEditor:", app.ActiveEditor.Text)

	undoCommand := &UndoCommand{
		App:    &app,
		Editor: app.ActiveEditor,
	}
	app.ExecuteCommand(undoCommand)
	fmt.Println("ActiveEditor:", app.ActiveEditor.Text)
}
