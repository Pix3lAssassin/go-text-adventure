package items

import "github.com/Pix3lAssassin/go-text-adventure/entities"

type ItemI interface {
	ID() int
	Type() ItemType
	Name() string
	Description() string
	Weight() float64
	Value() float64
	Use(c *entities.Combatant) bool
}

type Inventory interface {
	Get() []*ItemI
	AddItem(i *ItemI, quantity int) bool
	RemoveItem(i *ItemI, quantity int) bool
	UseItem(i *ItemI) bool
}

type ItemType int

const (
	Sword ItemType = iota
)
