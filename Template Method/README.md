```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class Position {
  +CoordinateX: int
  +CoordinateY: int
}

class Structure {
  +Resource: int
}

interface GameAI {
  +Turn()
  +CollectResources()
  +BuildStructures()
  +BuildUnits()
  +Attack()
  +SendScouts(position: Position)
  +SendWarriors(position: Position)
}

class BaseGameAI {
  -BuiltStructures: []Structure
  +Turn()
  +CollectResources()
  +BuildStructures()
  +BuildUnits()
  +Attack()
  +SendScouts(position: Position)
  +SendWarriors(position: Position)
}

class OrcsAI {
  -Resources: int
  -Scouts: []int
  -Warriors: []int
  +BuildStructures()
  +BuildUnits()
  +SendScouts(position: Position)
  +SendWarriors(position: Position)
}

class MonstersAI {
  +CollectResources()
  +BuildStructures()
  +BuildUnits()
}

BaseGameAI ..|> GameAI
OrcsAI --|> BaseGameAI
MonstersAI --|> BaseGameAI
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Game {
  +initializeGame()
  +playGame()
  +endGame()
  +play()
}

class Chess {
  +initializeGame()
  +playGame()
  +endGame()
  +play()
}

class Football {
  +initializeGame()
  +playGame()
  +endGame()
  +play()
}

Game <|.. Chess
Game <|.. Football
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface IOtp {
  + genRandomOTP(int) : string
  + saveOTPCache(string)
  + getMessage(string) : string
  + sendNotification(string) : error
}

class Otp {
  - iOtp: IOtp
  + genAndSendOTP(int) : error
}

class Sms {
  + genRandomOTP(int) : string
  + saveOTPCache(string)
  + getMessage(string) : string
  + sendNotification(string) : error
}

class Email {
  + genRandomOTP(int) : string
  + saveOTPCache(string)
  + getMessage(string) : string
  + sendNotification(string) : error
}

Otp -> IOtp : Uses
Sms --|> Otp : Extends
Sms ..|> IOtp : Implements
Email --|> Otp : Extends
Email ..|> IOtp : Implements
```