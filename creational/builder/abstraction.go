package builder

// VehicleBuildProcess is an abstraction for each builder.
type VehicleBuildProcess interface {
	SetSeats() VehicleBuildProcess
	SetWheels() VehicleBuildProcess
	SetStructure() VehicleBuildProcess
	Build() VehicleProduct
}

// VehicleProduct represents the actual product,
// built by builder that implements VehicleBuildProcess.
type VehicleProduct struct {
	Seats     int
	Wheels    int
	Structure string
}
