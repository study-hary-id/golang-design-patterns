package abstract_factory

type CruiseMotorbike struct {}

func (s *CruiseMotorbike) NumSeats() int {
	return 1
}

func (s *CruiseMotorbike) NumWheels() int {
	return 2
}

func (s *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
