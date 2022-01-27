package tests_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pet_tests/internal/service"
	"testing"
)

func TestShape(t *testing.T) {
	area := float32(5.5)
	expect := float32(30.25)

	testSqure := service.Square{SideLenght: area}
	result := testSqure.Area()
	assert.EqualValuesf(t, expect, result,
		fmt.Sprintf("Incorrect result. Expect %.2f centimeterse, got %.2f centimeterse", expect, result))
}

func TestCircle(t *testing.T) {
	area := float32(5.5)
	expect := float32(17.27876)

	testCircle := service.Circle{Radius: area}
	result := testCircle.Area()
	assert.EqualValuesf(t, expect, result,
		fmt.Sprintf("Incorrect result. Expect %.2f centimeterse, got %.2f centimeterse", expect, result))
}
