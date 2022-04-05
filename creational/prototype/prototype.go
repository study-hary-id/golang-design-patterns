package prototype

import (
	"errors"
	"fmt"
)

// ShirtCloner is an abstraction to any cloner structures.
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

// GetShirtCloner returns the instance of ShirtCloner.
func GetShirtCloner() ShirtCloner {
	return new(ShirtCache)
}

// ShirtCache used to cache any Shirt,
// implemented from ShirtCloner
type ShirtCache struct{}

// GetClone gives a clone Shirt.
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

// ItemInfoGetter is an abstraction to any prototype object.
type ItemInfoGetter interface {
	GetInfo() string
}

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
