package main

import "fmt"

func main() {
	var (
		variavel1 string = "Abracadabra"
		variavel2 string = "Pilimpimpim"
		variavel3 string = "Alakazam"
	)

	um := 1
	dois := 2
	tres := 3

	const nome string = "Felipe"

	fmt.Printf("feito por %s \n", nome)
	fmt.Println(um, dois, tres)
	fmt.Println(variavel1, variavel3, variavel2, "mostra o peito pra mim")

	variavel1, variavel2 = variavel2, variavel1

}
