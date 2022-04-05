package builder

// ManufacturingDirector uses as a dependency to construct any builder.
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct will construct a vehicle based on current builder.
func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder used to set which builder to use.
func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}
