package builder

type BikeBuilder struct {
	vehicle VehicleProduct
}

// SetSeats used to set vehicle's seats.
func (c *BikeBuilder) SetSeats() VehicleBuildProcess {
	c.vehicle.Seats = 2
	return c
}

// SetWheels used to set vehicle's wheels.
func (c *BikeBuilder) SetWheels() VehicleBuildProcess {
	c.vehicle.Wheels = 2
	return c
}

// SetStructure used to set vehicle's structure.
func (c *BikeBuilder) SetStructure() VehicleBuildProcess {
	c.vehicle.Structure = "Bike"
	return c
}

// Build used to build a bike product,
// returns new instance of VehicleProduct.
func (c *BikeBuilder) Build() VehicleProduct {
	return c.vehicle
}
