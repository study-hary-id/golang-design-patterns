package singleton

// Singleton is an abstraction Singleton pattern.
type Singleton interface {
	AddOne() int
}

// singleton is struct for counter instance.
type singleton struct {
	count int
}

// A pre-defined singleton instance.
// A pre-defined counter instance.
var instance *singleton

// GetInstance returns a pointer of singleton instance.
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne increments `count`, 1 number up.
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
