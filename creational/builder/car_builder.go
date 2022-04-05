package builder

// CarBuilder contains a set of build processes to build a car.
type CarBuilder struct {
	vehicle VehicleProduct
}

// SetSeats used to set vehicle's seats.
func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 5
	return c
}

// SetWheels used to set vehicle's wheels.
func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 4
	return c
}

// SetStructure used to set vehicle's structure.
func (c *CarBuilder) SetStructure() BuildProcess {
	c.vehicle.Structure = "Car"
	return c
}

// Build used to build constructed vehicle from ManufacturingDirector.
func (c *CarBuilder) Build() VehicleProduct {
	return c.vehicle
}
