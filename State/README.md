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
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class Character {
  +state: State
  +IdleState: State
  +WalkingState: State
  +RunningState: State
  +JumpingState: State
  +Walk()
  +Run()
  +Jump()
  +Idle()
  +setState(s State)
}

interface State {
  +Walk()
  +Run()
  +Jump()
  +Idle()
}

class IdleState {
  -character: Character
  +Walk()
  +Run()
  +Jump()
  +Idle()
}

class WalkingState {
  -character: Character
  +Walk()
  +Run()
  +Jump()
  +Idle()
}

class RunningState {
  -character: Character
  +Walk()
  +Run()
  +Jump()
  +Idle()
}

class JumpingState {
  -character: Character
  +Walk()
  +Run()
  +Jump()
  +Idle()
}

Character -up-|> State : Implements
IdleState -up-|> State : Implements
WalkingState -up-|> State : Implements
RunningState -up-|> State : Implements
JumpingState -up-|> State : Implements

Character "1" o--> "1" IdleState : Has
Character "1" o--> "1" WalkingState : Has
Character "1" o--> "1" RunningState : Has
Character "1" o--> "1" JumpingState : Has
```