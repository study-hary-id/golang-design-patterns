package abstract_factory

import (
	"errors"
	"fmt"
)

// Vehicle is an abstraction for each vehicle structures,
// each vehicles must implement this interface.
type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

// VehicleFactory is an abstraction for each vehicle factory,
// each factory structures should use Factory suffix.
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

// BuildFactory used to choose which VehicleFactory do you want to use.
func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("error: factory with id %d not recognized", f),
		)
	}
}
