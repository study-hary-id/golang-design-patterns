package builder

// Abstraction layer
type BuildProcess interface {
	SetSeats() BuildProcess
	SetWheels() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}

// Actual product
type VehicleProduct struct {
	Seats int
	Wheels int
	Structure string
}
