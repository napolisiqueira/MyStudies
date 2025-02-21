package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func Calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 + n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}

	var txt string
	fmt.Scanln(&txt)
	retorno := f(txt)
	fmt.Println("Esse foi o seu retorno: ", retorno)

	resultadoSoma, resultadoSubtracao := Calculos(12, 23)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
