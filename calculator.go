package main

import "fmt"



//addition
func add(x, y float32) float32{
	return x + y
}

//subtraction
func sub(x, y float32) float32{
	return x - y
}

//multiplication
func mul(x, y float32) float32{
	return x * y
}

//division
func div(x, y float32) float32 {
	return x / y
}

func main(){

	var num1, num2, result float32
	var option string
	i:= true


for i == true{
	//Taking input from user
	fmt.Print("Enter your first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter your second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter your option (+, -, *, /): ")
	fmt.Scan(&option)

	switch option {
	case "+":
		result = add(num1,num2)
	
	case "-":
		result = sub(num1,num2)
	
	case "*":
		result = mul(num1,num2)

	case "/":
		result = div(num1,num2)

	default:
		fmt.Println("Invalid Option")
		continue
	}
	fmt.Printf("Result = %0.2f\n", result)
}

}