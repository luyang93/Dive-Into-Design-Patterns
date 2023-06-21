package main

import "fmt"

type Game interface {
	initializeGame()
	playGame()
	endGame()
	play()
}

type Chess struct{}

func (c *Chess) initializeGame() {
	fmt.Println("Initializing Chess Game")
}

func (c *Chess) playGame() {
	fmt.Println("Playing Chess Game")
}

func (c *Chess) endGame() {
	fmt.Println("Ending Chess Game")
}

func (c *Chess) play() {
	c.initializeGame()
	c.playGame()
	c.endGame()
}

type Football struct{}

func (f *Football) initializeGame() {
	fmt.Println("Initializing Football Game")
}

func (f *Football) playGame() {
	fmt.Println("Playing Football Game")
}

func (f *Football) endGame() {
	fmt.Println("Ending Football Game")
}

func (f *Football) play() {
	f.initializeGame()
	f.playGame()
	f.endGame()
}

func main() {
	chess := &Chess{}
	chess.play()

	football := &Football{}
	football.play()
}
