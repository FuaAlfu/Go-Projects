package main

import(
	"fmt"
)

var digitOne, digitTwo float64
var operator string

func main(){
	fmt.Print("Enter first digit: ")
	fmt.Scanln(&digitOne)

	fmt.Print("Enter second digit: ")
	fmt.Scanln(&digitTwo)

	fmt.Print("Enter the operator: + x - / ")
	fmt.Scanln(&operator)
	switch operator{
	case "+":
		fmt.Printf("%f %s %f = %f", digitOne,operator,digitTwo,digitOne + digitTwo)
		break
	case "x":
		fmt.Printf("%f %s %f = %f", digitOne,operator,digitTwo,digitOne * digitTwo)
		break
	case "-":
		fmt.Printf("%f %s %f = %f", digitOne,operator,digitTwo,digitOne - digitTwo)
		break
	case "/":
		if digitTwo == 0.0{
			fmt.Println("divided by 0")
		}else{
			fmt.Printf("%f %s %f = %f", digitOne,operator,digitTwo,digitOne / digitTwo)
		}
	default:
		fmt.Println("invalid operator.. !")
	}

}