package test_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pet_tests/internal/helper"
	"testing"
)

func TestMax(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 4, -4, 2, 5, 3, 1},
			expected: 5,
		},
		{
			numbers:  []int{1, 4, -4, 20, 5, 3, 1},
			expected: 20,
		},
		{
			numbers:  []int{1, 4, -4, 2, 1, 3, 1},
			expected: 4,
		},
		{
			numbers:  []int{1, 4, -4, 2, 1, 3, 1, 7, 8, 2},
			expected: 8,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := helper.Max(testCase.numbers)
		expected := testCase.expected

		t.Logf("Calling Max(%v), result %d\n", testCase.numbers, result)

		// Assert
		assert.EqualValuesf(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result. Expect %d, got %d", expected, result))
	}
}

func TestMaxIndex(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 4, -4, 2, 5, 3, 1},
			expected: 4,
		},
		{
			numbers:  []int{1, 4, -4, 20, 5, 3, 1},
			expected: 3,
		},
		{
			numbers:  []int{1, 4, -4, 2, 1, 3, 1},
			expected: 1,
		},
		{
			numbers:  []int{1, 4, -4, 2, 1, 3, 1, 7, 8, 2},
			expected: 8,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := helper.MaxIndex(testCase.numbers)
		expected := testCase.expected

		t.Logf("Calling Max(%v), result %d\n", testCase.numbers, result)

		// Assert
		if result != expected {
			t.Errorf("Incorrect result. Expect %d, got %d", expected, result)
		}
	}
}
