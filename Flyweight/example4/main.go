package main

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if dress, ok := d.dressMap[dressType]; ok {
		return dress, nil
	} else {
		switch dressType {
		case CounterTerrroristDressType:
			dress = newCounterTerroristDress()
			d.dressMap[dressType] = dress
			return dress, nil
		case TerroristDressType:
			dress = newTerroristDress()
			d.dressMap[dressType] = dress
			return dress, nil
		default:
			return nil, fmt.Errorf("Wrong dress type passed")
		}
	}
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

type Dress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

type player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

func main() {
	game := newGame()
	for i := 0; i < 10; i++ {
		game.addTerrorist(TerroristDressType)
	}
	for i := 0; i < 5; i++ {
		game.addTerrorist(CounterTerrroristDressType)
	}

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
