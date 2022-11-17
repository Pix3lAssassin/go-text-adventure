package main

type Item interface {
	ID()
	Type() ItemType
	Name() string
	Description() string
	Weight() int
	Value() int
	Use() bool
}

type Inventory interface {
	Get() []*Item
	AddItem(i *Item, quantity int) bool
	RemoveItem(i *Item, quantity int) bool
	UseItem(i *Item) bool
}

type ItemType int

const (
	Sword ItemType = iota
)

type Weapon interface {
	Damage() float64
}

type Armor interface {
	Absorbtion() float64
	Stats() Stats
}
