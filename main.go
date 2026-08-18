package main
import "fmt"

func main()  {

//Arithmetic Operator: +, -, *, /, %
var num1,num2 int

fmt.Printf("Enter first number: ")
//Scan take input from user
fmt.Scan(&num1)
fmt.Printf("Enter second number: ")
fmt.Scan(&num2)

result:= num1 + num2
fmt.Printf("%v + %v = %v\n",num1,num2,result)

result= num1 - num2
fmt.Printf("%v - %v = %v\n",num1,num2,result)

result= num1 * num2
fmt.Printf("%v * %v = %v\n",num1,num2,result)

result= num1 / num2
fmt.Printf("%v / %v = %v\n",num1,num2,result)

result2:= float32(num1) / float32(num2)
fmt.Printf("%v / %v = %.2f\n",num1,num2,result2)

result= num1 % num2
fmt.Printf("%v %% %v = %v\n",num1,num2,result)

}