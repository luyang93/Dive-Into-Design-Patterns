```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class AudioPlayer {
  -State state
  +clickLock()
  +clickPlay()
  +clickNext()
  +clickPrevious()
  +changeState(state: State)
}

interface State {
  +clickLock(p: AudioPlayer)
  +clickPlay(p: AudioPlayer)
  +clickNext(p: AudioPlayer)
  +clickPrevious(p: AudioPlayer)
}

class LockedState {
  +clickLock(p: AudioPlayer)
  +clickPlay(p: AudioPlayer)
  +clickNext(p: AudioPlayer)
  +clickPrevious(p: AudioPlayer)
}

class ReadyState {
  +clickLock(p: AudioPlayer)
  +clickPlay(p: AudioPlayer)
  +clickNext(p: AudioPlayer)
  +clickPrevious(p: AudioPlayer)
}

class PlayingState {
  +clickLock(p: AudioPlayer)
  +clickPlay(p: AudioPlayer)
  +clickNext(p: AudioPlayer)
  +clickPrevious(p: AudioPlayer)
}

AudioPlayer --> State : has
State <|.. LockedState
State <|.. ReadyState
State <|.. PlayingState
```