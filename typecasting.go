package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Println("ENTER THE VALUE")
	fmt.Scan(&x)
	fmt.Println("THE TYPECASTED VALUE", change(x))
	fmt.Printf("THE TYPECASTED VALUE %.1f\n", change(x))
	fmt.Println("THE DOUBLE VALUE OF x =", double(x))
	fmt.Println("ENTER THE NAME")
	var s string
	fmt.Scan(&s)
	greet(s)
	fmt.Println("ENTER THE RADIUS OF CIRCLE")
	var radius float64
	fmt.Scan(&radius)

	fmt.Printf("THE PERIMETER OF CIRCLE IS %.2f\n", circlePerimeter(radius))

	fmt.Println("ENTER THE SIDE OF SQUARE")
	var sq float64

	fmt.Scan(&sq)
	fmt.Printf("THE PERIMETER OF SQUARE %.2f\n", sqPerimeter(sq))

}
func sqPerimeter(x float64) float64 {

	return 4 * x
}
func circlePerimeter(x float64) float64 {

	return 2 * (math.Pi) * x
}

func change(x int) float64 {

	return float64(x)
}
func double(x int) int {

	return x * 2
}
func greet(s string) {
	fmt.Println("Hello", s)
}
