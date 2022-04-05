package builder

// BuildProcess is an abstraction to each builder struct.
type BuildProcess interface {
	SetSeats() BuildProcess
	SetWheels() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

// VehicleProduct represents the actual product,
// built by builder struct.
type VehicleProduct struct {
	Seats int
	Wheels int
	Structure string
}
