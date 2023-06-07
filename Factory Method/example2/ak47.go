package main

// Ak47 is a struct that embeds the Gun struct
type Ak47 struct {
	Gun
}

// newAk47 is a function to create an instance of AK47 type
func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
