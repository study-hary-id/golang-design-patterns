package abstract_factory

import (
	"errors"
	"fmt"
)

// Motorbike is an abstraction for motorbike structures.
type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

// MotorbikeFactory used to build any available Motorbike structure.
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
			fmt.Sprintf("error: vehicle of type %d isn't recognized", v),
		)
	}
}
