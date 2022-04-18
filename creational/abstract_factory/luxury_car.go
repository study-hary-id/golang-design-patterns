package abstract_factory

/*
Multiple implementations from Vehicle and Car.
*/

type LuxuryCar struct{}

func (l *LuxuryCar) NumDoors() int {
	return 4
}

func (l *LuxuryCar) NumSeats() int {
	return 5
}

func (l *LuxuryCar) NumWheels() int {
	return 4
}
