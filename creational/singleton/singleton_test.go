package singleton

import "testing"

// TestGetInstance calls GetInstance and increments the instance.
func TestGetInstance(t *testing.T) {
	// Get the counter instance.
	counter := GetInstance()
	if counter == nil {
		t.Errorf("GetInstance() = %v, want match for pointer, nil", counter)
	}
	expectedCounter := counter

	// Increment the counter using AddOne method.
	currentCount := counter.AddOne()
	if currentCount != 1 {
		t.Errorf("counter.AddOne() = %v, want match for 1, nil", currentCount)
	}

	// Expected same instance from the first counter.
	secondCounter := GetInstance()
	if secondCounter != expectedCounter {
		t.Errorf(
			"GetInstance() = %v, want match for %v, nil",
			secondCounter,
			expectedCounter,
		)
	}

	// Increment the same counter since counter and secondCounter
	// are the same reference.
	currentCount = secondCounter.AddOne()
	if currentCount != 2 {
		t.Errorf(
			"secondCounter.AddOne() = %v, want match for 2, nil",
			currentCount,
		)
	}
}
