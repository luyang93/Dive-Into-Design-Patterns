```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface EventListener {
  + Update(filename, eventType: string)
}

class EventManager {
  - listeners: map[string][]EventListener
  + Subscribe(eventType: string, listener: EventListener)
  + Notify(eventType: string, data: string)
}

class Editor {
  - Events: EventManager
  - File: string
  + NewEditor()
  + OpenFile(path: string)
  + SaveFile()
}

class LoggingListener {
  - Log: string
  - Message: string
  + Update(eventType, filename: string)
}

class EmailAlertsListener {
  - Email: string
  - Message: string
  + Update(eventType, filename: string)
}

EventManager o--> "0..*" EventListener: notifies >
Editor --> EventManager: uses >
LoggingListener ..|> EventListener
EmailAlertsListener ..|> EventListener
```