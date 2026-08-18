package main

import "fmt"


func main(){

	var base, height, radius, triangleArea, circleArea float32
	
	fmt.Print("Enter the base of the triangle: ")
	fmt.Scan(&base)
	
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scan(&height)

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)

	triangleArea =  0.5 * base * height
	circleArea = 3.1416 * radius * radius

	fmt.Printf("The area of triangle is: %v\n",triangleArea)
	fmt.Printf("The area of circle is: %.2f",circleArea)
}