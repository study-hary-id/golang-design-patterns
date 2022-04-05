package abstract_factory

import (
	"errors"
	"fmt"
)

// Car is an abstraction for car structures.
type Car interface {
	NumDoors() int
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

// CarFactory used to build any available Car structure.
type CarFactory struct{}

// NewVehicle returns chosen Car type.
func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("error: vehicle of type %d isn't recognized", v),
		)
	}
}
