package main

import "fmt"

//Print even number
func evenNumber(){
	fmt.Println("Even numbers are:")
	for i:=0; i<=10; i = i+2{
		fmt.Println(i)
	}
}

//Print even number using condition
func evenUsingCondition(){
	for i := 1; i<=10; i++{
		if i%2 == 0{
			fmt.Println(i)
		}
	}
}

//1-10 Numbers Sum
func sum(){
	sum:=0
	for i:=1; i<=10; i++{
		sum = sum + i
	}
	fmt.Printf("Sum = %v\n",sum)
}

//Numbers summation by using user input
func summation(){
	var startNumber, endNumber int
	sum:=0
	fmt.Println("Enter your starting number:")
	fmt.Scan(&startNumber)
	fmt.Println("Enter your end number:")
	fmt.Scan(&endNumber)

	for i:=startNumber; i<=endNumber; i++{
		sum = sum + i
	}
	fmt.Printf("Sum = %v",sum)
}




func main(){
	evenNumber()
	evenUsingCondition()
	sum()
	summation()
}