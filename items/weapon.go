package items

import "github.com/Pix3lAssassin/go-text-adventure/stats"

var _ ItemI = (*Weapon)(nil)

type WeaponBuilder struct {
	ID          int
	WeaponType  ItemType
	Name        string
	Description string
	Weight      float64
	Value       float64
	Power       float64
	Stats       stats.CombatantStats
}

type Weapon struct {
	id          int
	weaponType  ItemType
	name        string
	description string
	weight      float64
	value       float64
	power       float64
	stats       stats.CombatantStats
}

func NewWeapon(w WeaponBuilder) *Weapon {
	return &Weapon{
		id:          w.ID,
		weaponType:  w.WeaponType,
		name:        w.Name,
		description: w.Description,
		weight:      w.Weight,
		value:       w.Value,
		power:       w.Power,
		stats:       w.Stats,
	}
}

func (w *Weapon) Damage() float64 {
	return w.power
}

func (a *Armor) Absorbtion() float64 {

}

func (a *Armor) Stats() stats.CombatantStats {

}

func (a *Armor) ID() {

}

func (a *Armor) Type() ItemType {

}

func (a *Armor) Name() string {

}

func (a *Armor) Description() string {

}

func (a *Armor) Weight() float64 {

}

func (a *Armor) Value() float64 {

}

func (a *Armor) Use() bool {

}
