package abstract_factory

/*
Multiple implementations from Vehicle and Car.
*/

type FamilyCar struct{}

func (l *FamilyCar) NumDoors() int {
	return 5
}

func (l *FamilyCar) NumSeats() int {
	return 5
}

func (l *FamilyCar) NumWheels() int {
	return 4
}
