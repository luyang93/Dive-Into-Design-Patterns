package main

import (
	"fmt"
	"sync"
)

type EventListener interface {
	Update(filename, eventType string)
}

type EventManager struct {
	listeners map[string][]EventListener
	mux       sync.Mutex
}

func (e *EventManager) Subscribe(eventType string, listener EventListener) {
	e.mux.Lock()
	defer e.mux.Unlock()

	e.listeners[eventType] = append(e.listeners[eventType], listener)
}

func (e *EventManager) Notify(eventType string, data string) {
	for _, listener := range e.listeners[eventType] {
		listener.Update(eventType, data)
	}
}

type Editor struct {
	Events *EventManager
	File   string
}

func NewEditor() *Editor {
	return &Editor{
		Events: &EventManager{
			listeners: make(map[string][]EventListener),
		},
	}
}

func (e *Editor) OpenFile(path string) {
	e.File = path
	e.Events.Notify("open", path)
}

func (e *Editor) SaveFile() {
	// Save file logic here
	e.Events.Notify("save", e.File)
}

type LoggingListener struct {
	Log     string
	Message string
}

func (l *LoggingListener) Update(eventType, filename string) {
	fmt.Println("LoggingListener")
	fmt.Printf(l.Message+"\n", eventType, filename)
}

type EmailAlertsListener struct {
	Email   string
	Message string
}

func (e *EmailAlertsListener) Update(eventType, filename string) {
	fmt.Println("EmailAlertsListener")
	fmt.Printf(e.Message+"\n", eventType, filename)
}

func main() {
	editor := NewEditor()

	logger := &LoggingListener{
		Log:     "/path/to/log.txt",
		Message: "Someone has %s the file: %s",
	}
	editor.Events.Subscribe("open", logger)

	emailAlerts := &EmailAlertsListener{
		Email:   "admin@example.com",
		Message: "Someone has %s the file: %s",
	}
	editor.Events.Subscribe("save", emailAlerts)
	editor.Events.Subscribe("save", logger)

	editor.OpenFile("test.txt")
	editor.SaveFile()
}
