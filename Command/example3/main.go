package main

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute()
	Undo()
}

type InsertTextCommand struct {
	TextToInsert string
	TextEditor   *strings.Builder
	Position     int
}

func NewInsertTextCommand(text string, editor *strings.Builder, position int) *InsertTextCommand {
	return &InsertTextCommand{
		TextToInsert: text,
		TextEditor:   editor,
		Position:     position,
	}
}

func (c *InsertTextCommand) Execute() {
	s := c.TextEditor.String()
	c.TextEditor.Reset()
	c.TextEditor.WriteString(s[:c.Position] + c.TextToInsert + s[c.Position:])
}

func (c *InsertTextCommand) Undo() {
	s := c.TextEditor.String()
	c.TextEditor.Reset()
	c.TextEditor.WriteString(strings.Replace(s, c.TextToInsert, "", 1))
}

type DeleteTextCommand struct {
	DeletedText string
	TextEditor  *strings.Builder
	Position    int
}

func NewDeleteTextCommand(editor *strings.Builder, position, length int) *DeleteTextCommand {
	s := editor.String()
	return &DeleteTextCommand{
		DeletedText: s[position : position+length],
		TextEditor:  editor,
		Position:    position,
	}
}

func (c *DeleteTextCommand) Execute() {
	s := c.TextEditor.String()
	c.TextEditor.Reset()
	c.TextEditor.WriteString(strings.Replace(s, c.DeletedText, "", 1))
}

func (c *DeleteTextCommand) Undo() {
	s := c.TextEditor.String()
	c.TextEditor.Reset()
	c.TextEditor.WriteString(s[:c.Position] + c.DeletedText + s[c.Position:])
}

type CommandInvoker struct {
	CommandStack []Command
	UndoStack    []Command
}

func NewCommandInvoker() *CommandInvoker {
	return &CommandInvoker{}
}

func (ci *CommandInvoker) Execute(c Command) {
	c.Execute()
	ci.CommandStack = append(ci.CommandStack, c)
}

func (ci *CommandInvoker) Undo() {
	if len(ci.CommandStack) > 0 {
		c := ci.CommandStack[len(ci.CommandStack)-1]
		ci.CommandStack = ci.CommandStack[:len(ci.CommandStack)-1]
		c.Undo()
		ci.UndoStack = append(ci.UndoStack, c)
	}
}

func (ci *CommandInvoker) Redo() {
	if len(ci.UndoStack) > 0 {
		c := ci.UndoStack[len(ci.UndoStack)-1]
		ci.UndoStack = ci.UndoStack[:len(ci.UndoStack)-1]
		c.Execute()
		ci.CommandStack = append(ci.CommandStack, c)
	}
}

func main() {
	textEditor := &strings.Builder{}
	invoker := NewCommandInvoker()

	invoker.Execute(NewInsertTextCommand("Hello, ", textEditor, 0))
	invoker.Execute(NewInsertTextCommand("world!", textEditor, 7))
	invoker.Execute(NewDeleteTextCommand(textEditor, 5, 2))

	fmt.Println("Current text: ", textEditor.String())

	invoker.Undo()
	fmt.Println("After undo: ", textEditor.String())

	invoker.Undo()
	fmt.Println("After undo: ", textEditor.String())

	invoker.Redo()
	fmt.Println("After redo: ", textEditor.String())
}
