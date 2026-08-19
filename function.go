package main
import "fmt"

func welcomeMessage(){
	fmt.Println("Welcome to Our Application!")
}

//username input
func getName() string{
	var name string
	fmt.Print("Enter your Name here: ")
	fmt.Scanln(&name)

	return name

}

//number input
func getNumbers()(int , int){
	var num1, num2 int

	fmt.Println("Enter First Number-")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second Number-")
	fmt.Scanln(&num2)

	return num1, num2
}

//add 
func add(num1, num2 int) int{
	sum := num1 + num2
	return sum
}

func display(name string, sum int){
	fmt.Println("Hello", name)
	fmt.Println("Your summation is: ",sum)
	
}

func goodbye(){
	fmt.Println("Thank you for using our application!!")
}

func main()  {
	welcomeMessage()
	name := getName()
	num1, num2 := getNumbers()
	sum := add(num1,num2)
	display(name, sum)
	goodbye()

}