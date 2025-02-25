package main

import "fmt"

func main() {
	var num1 float64 = -1
	var num2 float64 = -1

	for 0 > num1 || num1 > 10 {
		fmt.Scanln(&num1)

		if 0.0 >= num1 || num1 > 10.0 {
			fmt.Println("valor invalido")
			num1 = -1
		}
	}

	for 0 > num2 || num2 > 10 {
		fmt.Scanln(&num2)

		if 0.0 >= num2 || num2 > 10.0 {
			fmt.Println("valor invalido")
			num2 = -1
		}
	}
	var media float64 = (num1 + num2) / 2
	fmt.Printf("media = %.2f\n", media)

}
