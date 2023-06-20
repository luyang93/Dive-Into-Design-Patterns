package main

import "fmt"

type State interface {
	Walk()
	Run()
	Jump()
	Idle()
}

type Character struct {
	state        State
	IdleState    State
	WalkingState State
	RunningState State
	JumpingState State
}

func NewCharacter() *Character {
	c := &Character{}
	c.IdleState = &IdleState{c}
	c.WalkingState = &WalkingState{c}
	c.RunningState = &RunningState{c}
	c.JumpingState = &JumpingState{c}
	c.state = c.IdleState
	return c
}

func (c *Character) Walk() {
	c.state.Walk()
}

func (c *Character) Run() {
	c.state.Run()
}

func (c *Character) Jump() {
	c.state.Jump()
}

func (c *Character) Idle() {
	c.state.Idle()
}

func (c *Character) setState(s State) {
	c.state = s
}

type IdleState struct {
	character *Character
}

func (s *IdleState) Walk() {
	s.character.setState(s.character.WalkingState)
	fmt.Println("The character is now walking.")
}

func (s *IdleState) Run() {
	fmt.Println("The character cannot run right now.")
}

func (s *IdleState) Jump() {
	fmt.Println("The character cannot jump right now.")
}

func (s *IdleState) Idle() {
	fmt.Println("The character is already idle.")
}

type WalkingState struct {
	character *Character
}

func (s *WalkingState) Walk() {
	fmt.Println("The character is already walking.")
}

func (s *WalkingState) Run() {
	s.character.setState(s.character.RunningState)
	fmt.Println("The character is now running.")
}

func (s *WalkingState) Jump() {
	fmt.Println("The character cannot jump right now.")
}

func (s *WalkingState) Idle() {
	s.character.setState(s.character.IdleState)
	fmt.Println("The character is now idle.")
}

type RunningState struct {
	character *Character
}

func (s *RunningState) Walk() {
	fmt.Println("The character cannot walk right now.")
}

func (s *RunningState) Run() {
	fmt.Println("The character is already running.")
}

func (s *RunningState) Jump() {
	s.character.setState(s.character.JumpingState)
	fmt.Println("The character is now jumping.")
}

func (s *RunningState) Idle() {
	s.character.setState(s.character.IdleState)
	fmt.Println("The character is now idle.")
}

type JumpingState struct {
	character *Character
}

func (s *JumpingState) Walk() {
	fmt.Println("The character cannot walk right now.")
}

func (s *JumpingState) Run() {
	fmt.Println("The character cannot run right now.")
}

func (s *JumpingState) Jump() {
	fmt.Println("The character is already jumping.")
}

func (s *JumpingState) Idle() {
	s.character.setState(s.character.IdleState)
	fmt.Println("The character is now idle.")
}

func main() {
	character := NewCharacter()

	character.Walk() // The character is now walking.
	character.Run()  // The character is now running.
	character.Jump() // The character is now jumping.
	character.Idle() // The character is now idle.

	character.Run()  // The character cannot run right now.
	character.Jump() // The character cannot jump right now.
}
