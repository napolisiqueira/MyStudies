package main

import "fmt"

func main() {
	var vitGremio int64
	var vitInter int64
	var Empate int64
	var golGremio int64
	var golInter int64
	var numJogos int64
	for {

		fmt.Scanln(&golInter, &golGremio)
		numJogos += 1
		switch {
		case golInter > golGremio:
			vitInter += 1
		case golGremio > golInter:
			vitGremio += 1
		case golInter == golGremio:
			Empate += 1
		}

		var resp string
		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scanln(&resp)
		if resp == "1" {
			continue
		} else {
			fmt.Printf("%d grenais\n", numJogos)
			fmt.Printf("Inter:%d\n", vitInter)
			fmt.Printf("Gremio:%d\n", vitGremio)
			fmt.Printf("Empates:%d\n", Empate)
			switch {
			case vitGremio > vitInter:
				fmt.Println("Gremio venceu mais")
			case vitInter > vitGremio:
				fmt.Println("Inter venceu mais")
			default:
				fmt.Println("Nao houve vencedor")
			}
		}
		return
	}

}
