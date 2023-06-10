package main

import "fmt"

type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(percent int)
	GetChannel() int
	SetChannel(channel int)
}

type RemoteControl struct {
	device Device
}

func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{device: device}
}

func (r *RemoteControl) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *RemoteControl) volumeDown() {
	r.device.SetVolume(r.device.GetVolume() - 10)
}

func (r *RemoteControl) volumeUp() {
	r.device.SetVolume(r.device.GetVolume() + 10)
}

func (r *RemoteControl) channelDown() {
	r.device.SetChannel(r.device.GetChannel() - 1)
}

func (r *RemoteControl) channelUp() {
	r.device.SetChannel(r.device.GetChannel() + 1)
}

type AdvancedRemoteControl struct {
	RemoteControl
}

func NewAdvancedRemoteControl(device Device) *AdvancedRemoteControl {
	return &AdvancedRemoteControl{
		RemoteControl{
			device: device,
		},
	}
}

func (a *AdvancedRemoteControl) Mute() {
	a.device.SetVolume(0)
}

type TV struct {
	isEnabled bool
	volume    int
	channel   int
}

func (t *TV) IsEnabled() bool {
	return t.isEnabled
}

func (t *TV) Enable() {
	t.isEnabled = true
}

func (t *TV) Disable() {
	t.isEnabled = false
}

func (t *TV) GetVolume() int {
	return t.volume
}

func (t *TV) SetVolume(volume int) {
	t.volume = volume
}

func (t *TV) GetChannel() int {
	return t.channel
}

func (t *TV) SetChannel(channel int) {
	t.channel = channel
}

type Radio struct {
	isEnabled bool
	volume    int
	channel   int
}

func (r *Radio) IsEnabled() bool {
	return r.isEnabled
}

func (r *Radio) Enable() {
	r.isEnabled = true
}

func (r *Radio) Disable() {
	r.isEnabled = false
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
}

func (r *Radio) GetChannel() int {
	return r.channel
}

func (r *Radio) SetChannel(channel int) {
	r.channel = channel
}

func main() {
	tv := &TV{}
	remote := NewRemoteControl(tv)
	remote.TogglePower()

	fmt.Println(tv.IsEnabled()) // Should print true

	radio := &Radio{}
	advancedRemote := NewAdvancedRemoteControl(radio)
	advancedRemote.Mute()

	fmt.Println(radio.GetVolume()) // Should print 0
}
