package singleton

import "testing"

// TestGetInstance calls GetInstance and increments the instance,
// create another instance and ensure that the new instance
// is the same as the previous one.
func TestGetInstance(t *testing.T) {
	// Get the singleton/counter instance.
	counter := GetInstance()
	if counter == nil {
		t.Errorf("GetInstance() = %v, nil", counter)
	}

	// Increment the counter using AddOne method.
	currentCount := counter.AddOne()
	if currentCount != 1 {
		t.Errorf("counter.AddOne() = %v, want match for 1, nil", currentCount)
	}

	// Expected same instance from the first counter.
	secondCounter := GetInstance()
	if secondCounter != counter {
		t.Errorf(
			"GetInstance() = %v, want match for %v, nil",
			secondCounter,
			counter,
		)
	}

	// Increment `secondCounter` since `counter` and `secondCounter`
	// are the same reference.
	currentCount = secondCounter.AddOne()
	if currentCount != 2 {
		t.Errorf(
			"secondCounter.AddOne() = %v, want match for 2, nil",
			currentCount,
		)
	}
}
