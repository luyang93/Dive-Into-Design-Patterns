package main

// Musket is a struct that embeds the Gun struct
type Musket struct {
	Gun
}

// newMusket is a function to create an instance of Musket type
func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
