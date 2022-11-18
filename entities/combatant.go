package entities

import "github.com/Pix3lAssassin/go-text-adventure/items"

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
	Weapon() *items.Weapon
	Helmet() *items.Armor
	Armor() *items.Armor
	Gloves() *items.Armor
	Boots() *items.Armor
	Strength() float64
	Agility() float64
	Stealth() float64
	Intelligence() float64
	Type() CombatantType
	Inventory() *items.Inventory
}

type Leader struct {
	Combatant
	Subordinates []Combatant
}
