package abstract_factory

import "testing"

// TestCarFactory
func TestCarFactory(t *testing.T) {
	// Get the CarFactory.
	carFactory, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get the LuxuryCar with Vehicle type.
	carVehicle, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"Car vehicle seats: %d, wheels: %d\n",
		carVehicle.NumSeats(),
		carVehicle.NumWheels(),
	)

	// Assert LuxuryCar to Car type.
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has doors: %d\n", luxuryCar.NumDoors())
}
