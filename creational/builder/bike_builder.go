package builder

// BikeBuilder contains a set of build processes to build a bike.
type BikeBuilder struct {
	vehicle VehicleProduct
}

// SetSeats used to set vehicle's seats.
func (c *BikeBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 2
	return c
}

// SetWheels used to set vehicle's wheels.
func (c *BikeBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 2
	return c
}

// SetStructure used to set vehicle's structure.
func (c *BikeBuilder) SetStructure() BuildProcess {
	c.vehicle.Structure = "Bike"
	return c
}

// Build used to build constructed vehicle from ManufacturingDirector.
func (c *BikeBuilder) Build() VehicleProduct {
	return c.vehicle
}
