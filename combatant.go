package main

type CombatantType int

const (
	PlayerType CombatantType = iota
	Bandit
	Bear
)

type Combatant interface {
	Health() float64
	SetHealth(float64) bool
	SubtractHealth(float64) float64
	AddHealth(float64) float64
	Mana() float64
	SetMana(float64) bool
	SubtractMana(float64) float64
	AddMana(float64) float64
	Attack(Combatant) float64
	Weapon() Weapon
	Helmet() Armor
	Armor() Armor
	Gloves() Armor
	Boots() Armor
	Strength() float64
	Agility() float64
	Stealth() float64
	Intelligence() float64
	Type() CombatantType
	Inventory() Inventory
}

type Leader struct {
	Combatant
	Subordinates []Combatant
}

type Stats struct {
	Health       float64
	Mana         float64
	Strength     float64
	Agility      float64
	Stealth      float64
	Intelligence float64
}
