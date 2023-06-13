package main

import "fmt"

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Turning on the projector...")
}

func (p *Projector) Off() {
	fmt.Println("Turning off the projector...")
}

type DVDPlayer struct{}

func (d *DVDPlayer) Play() {
	fmt.Println("Playing the movie...")
}

func (d *DVDPlayer) Stop() {
	fmt.Println("Stopping the movie...")
}

type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Turning on the sound system...")
}

func (s *SoundSystem) Off() {
	fmt.Println("Turning off the sound system...")
}

type HomeTheaterFacade struct {
	projector   *Projector
	dvdPlayer   *DVDPlayer
	soundSystem *SoundSystem
}

func NewHomeTheaterFacade(projector *Projector, dvdPlayer *DVDPlayer, soundSystem *SoundSystem) *HomeTheaterFacade {
	return &HomeTheaterFacade{projector: projector, dvdPlayer: dvdPlayer, soundSystem: soundSystem}
}

func (h *HomeTheaterFacade) WatchMovie() {
	h.projector.On()
	h.dvdPlayer.Play()
	h.soundSystem.On()
}

func (h *HomeTheaterFacade) StopMovie() {
	h.projector.Off()
	h.dvdPlayer.Stop()
	h.soundSystem.Off()
}

func main() {
	projector := &Projector{}
	dvdPlayer := &DVDPlayer{}
	soundSystem := &SoundSystem{}

	homeTheater := NewHomeTheaterFacade(projector, dvdPlayer, soundSystem)
	homeTheater.WatchMovie()
	homeTheater.StopMovie()
}
