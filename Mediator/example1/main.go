package main

import "fmt"

type Mediator interface {
	Notify(Sender Component, Event string)
}

type Component interface {
	Click()
	KeyPress()
}

type Button struct {
	Dialog Mediator
}

func (c *Button) Click() {
	c.Dialog.Notify(c, "Click")
}

func (c *Button) KeyPress() {
	c.Dialog.Notify(c, "KeyPress")
}

type TextBox struct {
	Dialog Mediator
	Text   string
}

func (c *TextBox) Click() {
	c.Dialog.Notify(c, "Click")
}

func (c *TextBox) KeyPress() {
	c.Dialog.Notify(c, "KeyPress")
}

type CheckBox struct {
	Dialog  Mediator
	Checked bool
}

func (c *CheckBox) Check() {
	c.Dialog.Notify(c, "Check")
}

func (c *CheckBox) Click() {
	c.Checked = !c.Checked
	c.Dialog.Notify(c, "Click")
}

func (c *CheckBox) KeyPress() {
	c.Dialog.Notify(c, "KeyPress")
}

type AuthenticationDialog struct {
	Title                                                         string
	LoginOrRegisterCheckBox                                       *CheckBox
	LoginUsername, LoginPassword                                  *TextBox
	RegistrationUsername, RegistrationPassword, RegistrationEmail *TextBox
	OkBtn, CancelBtn                                              *Button
}

func NewAuthenticationDialog() *AuthenticationDialog {
	dialog := &AuthenticationDialog{}
	dialog.LoginOrRegisterCheckBox = &CheckBox{Dialog: dialog}
	dialog.LoginUsername = &TextBox{Dialog: dialog}
	dialog.LoginPassword = &TextBox{Dialog: dialog}
	dialog.OkBtn = &Button{Dialog: dialog}
	dialog.CancelBtn = &Button{Dialog: dialog}
	return dialog
}

func (a *AuthenticationDialog) Notify(Sender Component, Event string) {
	if Sender == a.LoginOrRegisterCheckBox && Event == "Check" {
		if a.LoginOrRegisterCheckBox.Checked {
			a.Title = "Log in"
			fmt.Println("Log in Check")
			// 1. Show login form components.
			// 2. Hide registration form components.
		} else {
			a.Title = "Register"
			fmt.Println("Register Check")
			// 1. Show registration form components.
			// 2. Hide login form components
		}
	}
	if Sender == a.OkBtn && Event == "Click" {
		if a.LoginOrRegisterCheckBox.Checked {
			// Try to find a user using login credentials.
			found := findUser(a.LoginUsername, a.LoginPassword)
			if !found {
				// Show an error message above the login field.
				fmt.Println("Error: User not found")
			}
		} else {
			// 1. Create a user account using data from the registration fields.
			// 2. Log that user in.
			a.RegistrationUsername = &TextBox{Dialog: a}
			a.RegistrationPassword = &TextBox{Dialog: a}
			a.RegistrationEmail = &TextBox{Dialog: a}
			registerUser(a.RegistrationUsername, a.RegistrationPassword, a.RegistrationEmail)
			fmt.Println("User registered and logged in")
		}
	}
}

// mock function for user validation
func findUser(username, password *TextBox) bool {
	// replace this with actual user validation logic
	return username.Text == "user" && password.Text == "password"
}

// mock function for user registration
func registerUser(username, password, email *TextBox) {
	// replace this with actual user registration logic
	fmt.Printf("Registering user %s with email %s\n", username.Text, email.Text)
}

func main() {
	dialog := NewAuthenticationDialog()
	dialog.LoginOrRegisterCheckBox.Click()
	dialog.OkBtn.Click()
	dialog.LoginOrRegisterCheckBox.Click()
	dialog.OkBtn.Click()
}
