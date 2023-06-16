```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Light {
    +On()
    +Off()
    +IncreaseBrightness()
}

class KitchenLight {
}

class LivingRoomLight {
}

Light <|.. KitchenLight
Light <|.. LivingRoomLight

interface Command {
    +Execute()
    +Undo()
}

class LightOnCommand {
    -light: Light
    +Execute()
    +Undo()
}

class LightOffCommand {
    -light: Light
    +Execute()
    +Undo()
}

class LightIncreaseCommand {
    -light: Light
    +Execute()
    +Undo()
}

class FanOnCommand {
    -fan: Fan
    +Execute()
    +Undo()
}

class AllLightsOffCommand {
    -lights: Light[]
    +Execute()
    +Undo()
}

Command <|.. LightOnCommand
Command <|.. LightOffCommand
Command <|.. LightIncreaseCommand
Command <|.. FanOnCommand
Command <|.. AllLightsOffCommand

class RemoteControl {
    -command: Command
    +SetCommand(command: Command)
    +PressButton()
    +PressUndo()
}

class Fan {
    +On()
    +Off()
}
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class Application {
    -Clipboard: string
    -Editors: Editor[]
    -ActiveEditor: Editor
    -History: CommandHistory
    +ExecuteCommand(command: Command)
    +Undo()
}

class Editor {
    -Text: string
    +GetSelection(): string
    +DeleteSelection()
    +ReplaceSelection(text: string)
}

interface Command {
    +Execute(): bool
    +Undo()
}

class CopyCommand {
    -App: Application
    -Editor: Editor
    -Backup: string
}

class CutCommand {
    -App: Application
    -Editor: Editor
    -Backup: string
}

class PasteCommand {
    -App: Application
    -Editor: Editor
    -Backup: string
}

class UndoCommand {
    -App: Application
    -Editor: Editor
    -Backup: string
}

class CommandHistory {
    -History: Command[]
    +Push(c: Command)
    +Pop(): Command
}

Application --> Editor
Application --> CommandHistory
CopyCommand ..|> Command
CutCommand ..|> Command
PasteCommand ..|> Command
UndoCommand ..|> Command
CommandHistory o--> Command
```