```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class RemoteControl {
  device: Device
  TogglePower()
  volumeDown()
  volumeUp()
  channelDown()
  channelUp()
}

class AdvancedRemoteControl {
  Mute()
}

interface Device {
  IsEnabled(): bool
  Enable()
  Disable()
  GetVolume(): int
  SetVolume(percent: int)
  GetChannel(): int
  SetChannel(channel: int)
}

class TV {
  isEnabled: bool
  volume: int
  channel: int
  IsEnabled(): bool
  Enable()
  Disable()
  GetVolume(): int
  SetVolume(volume: int)
  GetChannel(): int
  SetChannel(channel: int)
}

class Radio {
  isEnabled: bool
  volume: int
  channel: int
  IsEnabled(): bool
  Enable()
  Disable()
  GetVolume(): int
  SetVolume(volume: int)
  GetChannel(): int
  SetChannel(channel: int)
}

Device <|.. TV
Device <|.. Radio
RemoteControl o-- Device
AdvancedRemoteControl --|> RemoteControl
```