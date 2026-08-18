//if-else

package main
import "fmt"

func main(){

	var marks int = 129
	
	//if-else

	if marks >= 80 && marks <= 100{
		fmt.Println("A+")
	} else if marks >= 70 && marks < 80 {
		fmt.Println("A")
	} else if marks >= 60 && marks < 70 {
		fmt.Println("A-")
	} else if marks >= 50 && marks < 60 {
		fmt.Println("B")
	} else if marks >= 40 && marks < 50 {
		fmt.Println("C")
	} else if marks >= 33 && marks < 40 {
		fmt.Println("D")
	} else if marks >= 0 && marks < 33 {
		fmt.Println("F")
	} else{
		fmt.Println("Didn't appear in the exam!")
	}


	//switch-case

	var today string = "Frday"

	switch today {
	case "Saturday":
		fmt.Println("Last Relaxing Day")
	case "Sunday":
		fmt.Println("First Working Day")
	case "Monday":
		fmt.Println("Second Working Day")
	case "Tuesday":
		fmt.Println("Third Working Day")
	case "Wednesday":
		fmt.Println("Fourth Working Day")
	case "Thursday":
		fmt.Println("Fifth Working Day")
	case "Friday":
		fmt.Println("First Relaxing Day")
	default:
		fmt.Println("Vacation Time")
	}
}




