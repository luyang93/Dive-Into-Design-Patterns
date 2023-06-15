package main

import "fmt"

type Light interface {
	On()
	Off()
}

type KitchenLight struct{}

func (l *KitchenLight) On() {
	fmt.Println("Kitchen light is on")
}

func (l *KitchenLight) Off() {
	fmt.Println("Kitchen light is on")
}

type LivingRoomLight struct{}

func (l LivingRoomLight) On() {
	fmt.Println("Living room light is on")
}

func (l LivingRoomLight) Off() {
	fmt.Println("Living room light is off")
}

type Command interface {
	Execute()
	Undo()
}

type LightOnCommand struct {
	light Light
}

func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.On()
}

type LightOffCommand struct {
	light Light
}

func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func main() {
	livingRoomLight := &LivingRoomLight{}
	kitchenLight := &KitchenLight{}

	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)

	remote := &RemoteControl{}

	remote.SetCommand(livingRoomLightOn)
	remote.PressButton() // Living room light is on
	remote.PressUndo()   // Living room light is off

	remote.SetCommand(kitchenLightOff)
	remote.PressButton() // Kitchen light is off
	remote.PressUndo()   // Kitchen light is on
}
