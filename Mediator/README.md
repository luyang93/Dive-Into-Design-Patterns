```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface Mediator {
  +Notify(Sender : Component, Event : string)
}

interface Component {
  +Click()
  +KeyPress()
}

class Button {
  -Dialog : Mediator
  +Click()
  +KeyPress()
}

class TextBox {
  -Dialog : Mediator
  -Text : string
  +Click()
  +KeyPress()
}

class CheckBox {
  -Dialog : Mediator
  -Checked : bool
  +Check()
  +Click()
  +KeyPress()
}

class AuthenticationDialog {
  -Title : string
  -LoginOrRegisterCheckBox : CheckBox
  -LoginUsername, LoginPassword : TextBox
  -RegistrationUsername, RegistrationPassword, RegistrationEmail : TextBox
  -OkBtn, CancelBtn : Button
  +Notify(Sender : Component, Event : string)
}

Mediator <|.. AuthenticationDialog
Component <|.. Button
Component <|.. TextBox
Component <|.. CheckBox
Button ..> Mediator
TextBox ..> Mediator
CheckBox ..> Mediator
AuthenticationDialog ..> Button
AuthenticationDialog ..> TextBox
AuthenticationDialog ..> CheckBox
```