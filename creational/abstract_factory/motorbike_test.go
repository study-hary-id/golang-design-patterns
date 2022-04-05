package abstract_factory

import "testing"

// TestMotorbikeFactory
func TestMotorbikeFactory(t *testing.T) {
	// Get the MotorbikeFactory.
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get the SportMotorbike as Vehicle type.
	motorbikeVehicle, err := motorbikeFactory.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"Motorbike vehicle seats: %d, wheels: %d\n",
		motorbikeVehicle.NumSeats(),
		motorbikeVehicle.NumWheels(),
	)

	// Assert SportMotorbike as Motorbike type.
	sportMotorbike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type: %d\n", sportMotorbike.GetMotorbikeType())
}
