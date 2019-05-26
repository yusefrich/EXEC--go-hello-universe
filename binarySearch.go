//busca binaria simples em Golang
package main

import (
	"fmt"
)

func binarySearch(agulha int, palheiro []int) bool {
	minimo := 0
	maximo := len(palheiro) - 1

	for minimo <= maximo {
		media := (minimo + maximo) / 2

		if palheiro[media] < agulha {
			minimo = media + 1
		} else {
			maximo = media - 1
		}
	}

	if minimo == len(palheiro) || palheiro[minimo] != agulha {
		return false
	}

	return true
}

func main() {
	itens := []int{1, 2, 10, 21, 36, 46, 76, 90, 126}
	//função de sorting para agrupar os parametros da array
	fmt.Println(binarySearch(36, itens))
}
