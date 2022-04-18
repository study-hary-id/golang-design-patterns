package builder

type CarBuilder struct {
	vehicle VehicleProduct
}

// SetSeats used to set vehicle's seats.
func (c *CarBuilder) SetSeats() VehicleBuildProcess {
	c.vehicle.Seats = 5
	return c
}

// SetWheels used to set vehicle's wheels.
func (c *CarBuilder) SetWheels() VehicleBuildProcess {
	c.vehicle.Wheels = 4
	return c
}

// SetStructure used to set vehicle's structure.
func (c *CarBuilder) SetStructure() VehicleBuildProcess {
	c.vehicle.Structure = "Car"
	return c
}

// Build used to build a car product,
// returns new instance of VehicleProduct.
func (c *CarBuilder) Build() VehicleProduct {
	return c.vehicle
}
