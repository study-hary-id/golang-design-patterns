package builder

// ManufacturingDirector is a dependency to construct any builder.
type ManufacturingDirector struct {
	builder VehicleBuildProcess
}

// Construct sets properties of VehicleProduct.
func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder sets a builder to ManufacturingDirector.
func (m *ManufacturingDirector) SetBuilder(b VehicleBuildProcess) {
	m.builder = b
}
