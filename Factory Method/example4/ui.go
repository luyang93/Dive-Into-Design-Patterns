package main

// UIFramework is now a struct that has a Button
type UIFramework struct {
	Button
}

func (u *UIFramework) CreateButton() {
	u.Button = &SquareButton{}
}

// UIWithRoundButtons embeds UIFramework
type UIWithRoundButtons struct {
	UIFramework
}

// Overriding CreateButton method in UIWithRoundButtons
func (u *UIWithRoundButtons) CreateButton() {
	u.Button = &RoundButton{}
}
