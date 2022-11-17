package main

var _ Combatant = (*Player)(nil)

type Player struct {
	health        float64
	mana          float64
	weapon        Weapon
	helmet        Armor
	armor         Armor
	gloves        Armor
	boots         Armor
	strength      float64
	agility       float64
	stealth       float64
	intelligence  float64
	combatantType CombatantType
	inventory     Inventory
}

func (p *Player) Health() float64 {
	return p.health
}

func (p *Player) SetHealth(health float64) bool {
	p.health = health
	return p.health > 0
}

func (p *Player) AddHealth(health float64) float64 {
	p.health += health
	return p.health
}

func (p *Player) SubtractHealth(health float64) float64 {
	p.health -= health
	return p.health
}

func (p *Player) Mana() float64 {
	return p.mana
}

func (p *Player) SetMana(mana float64) bool {
	p.mana = mana
	return p.mana > 0
}

func (p *Player) AddMana(mana float64) float64 {
	p.mana += mana
	return p.mana
}

func (p *Player) SubtractMana(mana float64) float64 {
	p.mana -= mana
	return p.mana
}

func (p *Player) Attack(c Combatant) float64 {
	return c.SubtractHealth(p.Weapon().Damage())
}

func (p *Player) Weapon() Weapon {
	return p.weapon
}

func (p *Player) Helmet() Armor {
	return p.helmet
}

func (p *Player) Armor() Armor {
	return p.armor
}

func (p *Player) Gloves() Armor {
	return p.gloves
}

func (p *Player) Boots() Armor {
	return p.boots
}

func (p *Player) Strength() float64 {
	return p.strength
}

func (p *Player) Agility() float64 {
	return p.agility
}

func (p *Player) Stealth() float64 {
	return p.stealth
}

func (p *Player) Intelligence() float64 {
	return p.intelligence
}

func (p *Player) Type() CombatantType {
	return p.combatantType
}

func (p *Player) Inventory() Inventory {
	return p.inventory
}
