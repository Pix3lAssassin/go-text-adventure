package items

import "github.com/Pix3lAssassin/go-text-adventure/stats"

var _ ItemI = (*Weapon)(nil)

type ArmorBuilder struct {
	ID int
	ArmorType ItemType
	Name string
	Description string
	Weight float64
	Value float64
	Absorbtion float64
	Stats stats.CombatantStats
}

type Armor struct {
	id int
	armorType ItemType
	name string
	description string
	weight float64
	value float64
	absorbtion float64
	stats stats.CombatantStats
}

func NewArmor(a ArmorBuilder) *Armor {
	return &Armor{
		id: a.ID,
		armorType: a.ArmorType,
		name: a.Name,
		description: a.Description,
		weight: a.Weight,
		value: a.Value,
		absorbtion: a.Absorbtion,
		stats: a.Stats,
	}
}

func (a *Armor) Absorbtion() float64 {
	return a.absorbtion
}

func (a *Armor) Stats() stats.CombatantStats {
	return a.stats
}

func (a *Armor) ID() {
	return a.id
}

func (a *Armor) Type() ItemType {
	return a.armorType
}

func (a *Armor) Name() string {
	return a.name
}

func (a *Armor) Description() string {
	return a.description
}

func (a *Armor) Weight() float64 {
	return a.weight
}

func (a *Armor) Value() float64 {
	return a.value
}

func (a *Armor) Use(c *Combatant) bool {
	c.

	return true
}

