package main

import (
	"fmt"
	auxiliar "modulo/Auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo um arquivo main")
	auxiliar.Escrever()
	auxiliar.EscreverO()
	erro := checkmail.ValidateFormat("imfelipenapoli@gmail.com")
	fmt.Println(erro)
}
