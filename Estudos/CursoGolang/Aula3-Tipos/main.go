package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var inicializar string //Dá pra inicializar uma variavel sem valor, porem sempre vai ser 0
	fmt.Println(inicializar)

	var numero int32 = -12345
	fmt.Println(numero)

	var numero_sem_sinal uint8 = 123
	fmt.Println(numero_sem_sinal)

	var numero_Aski rune = 123456
	fmt.Println(numero_Aski)

	var numero_decimal float32 = 1.2233122
	fmt.Println(numero_decimal)

	var texto string = "Clarinha"
	fmt.Println(texto)

	char := 'A' //Aspas simples refere a um tipo "char", que remete a posição daquela string na tabela Asky
	fmt.Println(char)

	var Joia bool = true
	fmt.Println(Joia)

	var Errinho error
	fmt.Println(Errinho)

}
