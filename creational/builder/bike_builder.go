package builder

// Inherited from BuildProcess
type BikeBuilder struct {
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
