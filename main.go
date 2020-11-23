package main

import "fmt"

func main() {

	students := map[string]int{
		"Julio Cesar Vieira Cruz": 201904229,
		"Walysson Cristino Paiva": 201907516,
		"Hyago Coelho Teodoro de Rezende": 201804661,
		"Jair Souza Meira Rodrigues": 201910884,
	}

	var sum int

	for _, value := range students {
		sum += value%10
	}

	fmt.Println("Hello World")
	fmt.Println(fmt.Sprintln(`Sum of each last student's registration number:`, sum))
}