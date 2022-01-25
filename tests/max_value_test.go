package tests

import (
	"pet_tests/pkg/helper"
	"testing"
)

func TestMax(t *testing.T) {
	// Arrange
	numbers := []int{1, 4, -4, 2, 5, 3, 1}
	expected := 5

	// Act
	result := helper.Max(numbers)

	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d", expected, result)
	}
}

func TestMaxIndex(t *testing.T) {
	// Arrange
	numbers := []int{1, 4, -4, 2, 5, 3, 1}
	expected := 4

	// Act
	result := helper.MaxIndex(numbers)

	// Assert
	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d", expected, result)
	}
}
