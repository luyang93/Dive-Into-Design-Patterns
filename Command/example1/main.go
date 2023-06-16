package main

import "fmt"

type Light interface {
	On()
	Off()
	IncreaseBrightness()
}

type KitchenLight struct{}

func (l *KitchenLight) On() {
	fmt.Println("Kitchen light is on")
}

func (l *KitchenLight) Off() {
	fmt.Println("Kitchen light is on")
}

func (l *KitchenLight) IncreaseBrightness() {
	fmt.Println("Kitchen light brightness increased")
}

type LivingRoomLight struct{}

func (l *LivingRoomLight) On() {
	fmt.Println("Living room light is on")
}

func (l *LivingRoomLight) Off() {
	fmt.Println("Living room light is off")
}

func (l *LivingRoomLight) IncreaseBrightness() {
	fmt.Println("Living room light brightness increased")
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
	c.light.Off()
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

type LightIncreaseCommand struct {
	light Light
}

func NewLightIncreaseCommand(light Light) *LightIncreaseCommand {
	return &LightIncreaseCommand{light: light}
}

func (c *LightIncreaseCommand) Execute() {
	c.light.IncreaseBrightness()
}

func (c *LightIncreaseCommand) Undo() {
	fmt.Println("TODO")
}

type Fan struct{}

func (f *Fan) On() {
	fmt.Println("Fan is on")
}

func (f *Fan) Off() {
	fmt.Println("Fan is off")
}

type FanOnCommand struct {
	fan *Fan
}

func NewFanOnCommand(fan *Fan) *FanOnCommand {
	return &FanOnCommand{fan: fan}
}

func (c *FanOnCommand) Execute() {
	c.fan.On()
}

func (c *FanOnCommand) Undo() {
	c.fan.Off()
}

type AllLightsOffCommand struct {
	lights []Light
}

func NewAllLightsOffCommand(lights []Light) *AllLightsOffCommand {
	return &AllLightsOffCommand{lights: lights}
}

func (c AllLightsOffCommand) Execute() {
	for _, light := range c.lights {
		light.On()
	}
}

func (c AllLightsOffCommand) Undo() {
	for _, light := range c.lights {
		light.Off()
	}
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
	fan := &Fan{}

	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)
	kitchenLightIncrease := NewLightIncreaseCommand(kitchenLight)
	fanOn := NewFanOnCommand(fan)
	allLightsOn := NewAllLightsOffCommand([]Light{livingRoomLight, kitchenLight})
	remote := &RemoteControl{}

	remote.SetCommand(livingRoomLightOn)
	remote.PressButton() // Living room light is on
	remote.PressUndo()   // Living room light is off

	remote.SetCommand(kitchenLightOff)
	remote.PressButton() // Kitchen light is off
	remote.PressUndo()   // Kitchen light is on

	remote.SetCommand(kitchenLightIncrease)
	remote.PressButton()

	remote.SetCommand(fanOn)
	remote.PressButton()

	remote.SetCommand(allLightsOn)
	remote.PressButton()
}
