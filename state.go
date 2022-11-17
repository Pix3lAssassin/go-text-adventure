package main

type State struct {
	Player   *Player
	Factions []Faction
}

type Faction struct {
	Data       string
	Leader     *Leader
	Reputation map[Combatant]float64
}
