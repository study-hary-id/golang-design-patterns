package abstract_factory

import "testing"

func TestSportMotorbikeType(t *testing.T) {
	// Get the MotorbikeFactory.
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get the SportMotorbike from MotorbikeFactory.
	motorbikeVehicle, err := motorbikeFactory.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"Motorbike vehicle seats: %d, wheels: %d",
		motorbikeVehicle.NumSeats(),
		motorbikeVehicle.NumWheels(),
	)

	// Assert SportMotorbike as Motorbike type.
	sportMotorbike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed.")
	}
	t.Logf("Sport motorbike has type: %d", sportMotorbike.GetMotorbikeType())
}

func TestCruiseMotorbikeType(t *testing.T) {
	// Get the MotorbikeFactory.
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get the CruiseMotorbike from MotorbikeFactory.
	motorbikeVehicle, err := motorbikeFactory.NewVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"Motorbike vehicle seats: %d, wheels: %d",
		motorbikeVehicle.NumSeats(),
		motorbikeVehicle.NumWheels(),
	)

	// Assert CruiseMotorbike as Motorbike type.
	cruiseMotorbike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed.")
	}
	t.Logf("Sport motorbike has type: %d", cruiseMotorbike.GetMotorbikeType())
}
