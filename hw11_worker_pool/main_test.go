package main

import (
	"testing"
)

func TestCounterIncrement(t *testing.T) {
	counter := &Counter{}

	// Инкрементируем счетчик 1000 раз.
	for i := 0; i < 1000; i++ {
		counter.Increment()
	}

	if counter.GetValue() != 1000 {
		t.Errorf("Expected counter value to be 1000, got %d", counter.GetValue())
	}
}

func TestRunGoroutines(t *testing.T) {
	counter := &Counter{}
	const numGoroutines = 5
	const incrementsPerGoroutine = 1000

	RunGoroutines(counter, numGoroutines, incrementsPerGoroutine)

	expectedValue := numGoroutines * incrementsPerGoroutine
	if counter.GetValue() != expectedValue {
		t.Errorf("Expected counter value to be %d, got %d", expectedValue, counter.GetValue())
	}
}
