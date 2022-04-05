package builder

import "testing"

// TestCarBuilder tries to build a Car Product using CarBuilder.
func TestCarBuilder(t *testing.T) {
	var (
		manufacturingComplex = ManufacturingDirector{}
		carBuilder           = &CarBuilder{}
	)
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()
	if car.Seats != 5 {
		t.Errorf("car.Seats = %v, want match for 5, nil", car.Seats)
	}
	if car.Wheels != 4 {
		t.Errorf("car.Wheels = %v, want match for 4, nil", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("car.Structure = '%v', want match for 'Car', nil", car.Structure)
	}
}

// TestBikeBuilder tries to build a Bike Product using BikeBuilder,
// without using ManufacturingDirector.
func TestBikeBuilder(t *testing.T) {
	var (
		bikeBuilder = &BikeBuilder{}
	)

	bikeBuilder.SetSeats().SetStructure().SetWheels()
	bike := bikeBuilder.Build()

	if bike.Seats != 2 {
		t.Errorf("bike.Seats = %v, want match for 2, nil", bike.Seats)
	}
	if bike.Wheels != 2 {
		t.Errorf("bike.Wheels = %v, want match for 2, nil", bike.Wheels)
	}
	if bike.Structure != "Bike" {
		t.Errorf("bike.Structure = '%v', want match for 'Bike', nil", bike.Structure)
	}
}
