package main

import "fmt"

type AudioPlayer struct {
	State State
	// UI, volume, playlist, currentSong // Here add fields as needed
}

func NewAudioPlayer() *AudioPlayer {
	return &AudioPlayer{
		State: &ReadyState{},
	}
}

func (p *AudioPlayer) SetState(s State) {
	p.State = s
}

func (p *AudioPlayer) ClickLock() {
	p.State.ClickLock(p)
}

func (p *AudioPlayer) ClickPlay() {
	p.State.ClickPlay(p)
}

func (p *AudioPlayer) ClickNext() {
	p.State.ClickNext(p)
}

func (p *AudioPlayer) ClickPrevious() {
	p.State.ClickPrevious(p)
}

type State interface {
	ClickLock(p *AudioPlayer)
	ClickPlay(p *AudioPlayer)
	ClickNext(p *AudioPlayer)
	ClickPrevious(p *AudioPlayer)
}

type LockedState struct{}

func (s *LockedState) ClickLock(p *AudioPlayer) {
	fmt.Println("Unlocking player")
	p.SetState(&ReadyState{})
}

func (s *LockedState) ClickPlay(p *AudioPlayer) {
	fmt.Println("Player is locked, can't play")
}

func (s *LockedState) ClickNext(p *AudioPlayer) {
	fmt.Println("Player is locked, can't change song")
}

func (s *LockedState) ClickPrevious(p *AudioPlayer) {
	fmt.Println("Player is locked, can't change song")
}

type ReadyState struct{}

func (s *ReadyState) ClickLock(p *AudioPlayer) {
	fmt.Println("Locking player")
	p.SetState(&LockedState{})
}

func (s *ReadyState) ClickPlay(p *AudioPlayer) {
	fmt.Println("Starting playback")
	p.SetState(&PlayingState{})
}

func (s *ReadyState) ClickNext(p *AudioPlayer) {
	fmt.Println("Changing to next song")
}

func (s *ReadyState) ClickPrevious(p *AudioPlayer) {
	fmt.Println("Changing to previous song")
}

type PlayingState struct{}

func (s *PlayingState) ClickLock(p *AudioPlayer) {
	fmt.Println("Locking player")
	p.SetState(&LockedState{})
}

func (s *PlayingState) ClickPlay(p *AudioPlayer) {
	fmt.Println("Stopping playback")
	p.SetState(&ReadyState{})
}

func (s *PlayingState) ClickNext(p *AudioPlayer) {
	fmt.Println("Changing to next song")
}

func (s *PlayingState) ClickPrevious(p *AudioPlayer) {
	fmt.Println("Changing to previous song")
}

func main() {
	player := NewAudioPlayer()
	player.ClickPlay()
	player.ClickLock()
	player.ClickPlay()
}
