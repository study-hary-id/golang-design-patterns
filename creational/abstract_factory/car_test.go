package abstract_factory

import "testing"

// TestLuxuryCarType executes a sets of sequence to create a LuxuryCar.
func TestLuxuryCarType(t *testing.T) {
	// Get the CarFactory, implementation of VehicleFactory.
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get the LuxuryCar from CarFactory, implementation of Vehicle.
	carVehicle, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"LuxuryCar has seats: %d, wheels: %d",
		carVehicle.NumSeats(),
		carVehicle.NumWheels(),
	)

	// Assert LuxuryCar to Car type, needs to invoke .NumDoors() method.
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed.")
	}
	t.Logf("Luxury car has doors: %d", luxuryCar.NumDoors())
}

// TestFamilyCarType executes a sets of sequence to create a FamilyCar.
func TestFamilyCarType(t *testing.T) {
	// Get the CarFactory
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get the FamilyCar from CarFactory
	carVehicle, err := carFactory.NewVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"FamilyCar has seats: %d, wheels: %d",
		carVehicle.NumSeats(),
		carVehicle.NumWheels(),
	)

	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed.")
	}
	t.Logf("Luxury car has doors: %d", familyCar.NumDoors())
}
