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

	fmt.Println("Alô Mundo")

	for student, value := range students {
		sum += value%10
		fmt.Println(student)
	}

	fmt.Println("Soma dos últimos números da matricula dos estudantes:", sum)
}