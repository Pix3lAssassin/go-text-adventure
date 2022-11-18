package state

import "github.com/Pix3lAssassin/go-text-adventure/entities"

type GameState struct {
	Player   *entities.Player
	Factions []Faction
}

type Faction struct {
	Name        string
	Description string
	Leader      *entities.Leader
	Reputation  map[entities.Combatant]float64
}
