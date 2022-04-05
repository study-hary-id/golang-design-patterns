package builder

// Inherited from BuildProcess
type CarBuilder struct {
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
