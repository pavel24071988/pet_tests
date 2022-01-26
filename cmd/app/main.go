package main

import (
	"fmt"
	"pet_tests/internal/helper"
	"pet_tests/internal/service"
)

func main() {
	fmt.Println(helper.Max([]int{2, 4, 1, 8, 4, 5, 15, 6}))

	squareArea := service.Square{5.5}
	circleArea := service.Circle{3.4}

	service.PrintShapeArea(squareArea)
	service.PrintShapeArea(circleArea)
}
