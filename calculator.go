package main

import "fmt"

func add(a, b int) int {
	//Insert code here
	return a + b
}

func subtract(a, b int) int {
	//Insert code here
	return a - b
}

func multiply(a, b int) int {
	//Insert code here
	return a * b
}

func divide(a, b int) int {
	//Insert code here
	//consider for b = 0
	if b != 0 {
		return a / b
	} else {
		fmt.Println("Error!")
		return 0
	}
}

func main() {
	var a, b int
	var process string
	fmt.Printf("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Printf("Enter process (add, sub, mul, div): ")
	fmt.Scanln(&process)
	fmt.Printf("Enter an integer: ")
	fmt.Scanln(&b)

	//Insert code here
	if process == "add" {
		result := add(a, b)
		fmt.Println("Result:", result)
	} else if process == "sub" {
		result := subtract(a, b)
		fmt.Println("Result:", result)
	} else if process == "mul" {
		result := multiply(a, b)
		fmt.Println("Result:", result)
	} else if process == "div" && b != 0 {
		result := divide(a, b)
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error!")
	}
}
