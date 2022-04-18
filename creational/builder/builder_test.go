package builder

import (
	"testing"
)

// TestCarBuilder build a car product using CarBuilder.
func TestCarBuilder(t *testing.T) {
	var (
		manufacturingComplex = ManufacturingDirector{}
		carBuilder           = &CarBuilder{}
	)

	// Need to pass a reference because manufacturingComplex needs to
	// call CarBuilder's method within it.
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
		t.Errorf("car.Structure = %q, want match for \"Car\", nil", car.Structure)
	}
}

// TestBikeBuilder build a bike product using BikeBuilder,
func TestBikeBuilder(t *testing.T) {
	var (
		manufacturingComplex = ManufacturingDirector{}
		bikeBuilder          = &BikeBuilder{}
	)

	// Need to pass a reference because manufacturingComplex needs to
	// call CarBuilder's method within it.
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.Build()

	if bike.Seats != 2 {
		t.Errorf("bike.Seats = %v, want match for 2, nil", bike.Seats)
	}
	if bike.Wheels != 2 {
		t.Errorf("bike.Wheels = %v, want match for 2, nil", bike.Wheels)
	}
	if bike.Structure != "Bike" {
		t.Errorf("bike.Structure = %q, want match for \"Bike\", nil", bike.Structure)
	}
}

// TestCarBuilderWithoutDirector build a bike without ManufacturingDirector instance.
func TestCarBuilderWithoutDirector(t *testing.T) {
	carBuilder := &CarBuilder{}

	// Constructing a car without director.
	carBuilder.SetSeats().SetStructure().SetWheels()
	car := carBuilder.Build()

	if car.Seats != 5 {
		t.Errorf("car.Seats = %v, want match for 5, nil", car.Seats)
	}
	if car.Wheels != 4 {
		t.Errorf("car.Wheels = %v, want match for 4, nil", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("car.Structure = %q, want match for \"Car\", nil", car.Structure)
	}
}
