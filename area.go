package main

import "fmt"


func main(){

	var base, height, triangleArea, circleArea float32
	
	fmt.Print("Enter the base of the triangle: ")
	fmt.Scan(&base)
	
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scan(&height)

	triangleArea =  0.5 * base * height
	circleArea = 3.1416 * base * base

	fmt.Printf("The area of triangle is: %v\n",triangleArea)
	fmt.Printf("The area of circle is: %.2f",circleArea)
}