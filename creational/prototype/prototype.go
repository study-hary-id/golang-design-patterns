package prototype

import (
	"errors"
	"fmt"
)

/*
Cloner interface
*/

// ShirtCloner is an abstraction to any cloner.
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

// GetShirtCloner returns the instance of ShirtCloner.
func GetShirtCloner() ShirtCloner {
	return new(ShirtCache)
}

/*
Cloner implementation
*/

// ShirtCache used to cache any Shirt, or clone.
type ShirtCache struct{}

// GetClone returns cloned shirt.
func (sc *ShirtCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("error: shirt model not recognized")
	}
}

/*
Product interface
*/

// ItemInfoGetter is an abstraction to any prototype instance.
type ItemInfoGetter interface {
	GetInfo() string
}

/*
Product implementation with ShirtColor type.
*/

// Enumeration color for Shirt.
const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtColor byte

// Shirt is an entity used to clone or to prototyping.
type Shirt struct {
	SKU   string
	Price float32
	Color ShirtColor
}

// GetInfo returns information of the shirt.
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf(
		"Shirt with SKU: %s, color: %d, price: %0.2f",
		s.SKU,
		s.Color,
		s.Price,
	)
}

// GetPrice returns the price of the shirt.
func (s *Shirt) GetPrice() float32 {
	return s.Price
}

/*
Prototype instance, used to be clone.
*/

var whitePrototype *Shirt = &Shirt{
	SKU:   "empty",
	Price: 15.00,
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	SKU:   "empty",
	Price: 16.00,
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	SKU:   "empty",
	Price: 17.00,
	Color: Blue,
}
