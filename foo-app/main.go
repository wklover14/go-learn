package main

import (
	"fmt"

	"math"
)

// Variadic functions ~ functions with an undefined number of parameters
func sum(nums ...int) {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("le total est de ", total)
}

// methods
type Circle struct {
	x_center, y_center float64
	ray                float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.ray, 2)
}

func main() {
	sum(1, 3, 5, 7, 9)
	c1 := Circle{x_center: 0.0, y_center: 0.0, ray: 4.0}
	fmt.Println(c1, "has an Area of", c1.Area())
}
