package abstract_factory

import (
	"errors"
	"fmt"
)

// Motorbike is another abstraction for vehicle product.
type Motorbike interface {
	GetMotorbikeType() int
}

// Enumeration of specific Motorbike type.
const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

// MotorbikeFactory used to build any available Motorbike type.
type MotorbikeFactory struct{}

// NewVehicle returns chosen Motorbike type.
func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("error: motorbike of type %d isn't recognized", v),
		)
	}
}
