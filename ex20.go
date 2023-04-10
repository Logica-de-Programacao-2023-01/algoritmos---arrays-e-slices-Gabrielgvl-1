package main

import "fmt"

// Crie um programa que leia um array
//de inteiros e verifique se ele está ordenado
//em ordem crescente.

func main() {
	slice := []int{1, 2, 3, 9, 5, 6, 7, 8}

	crescente := true
	for i := 1; i < len(slice); i++ {
		atual := slice[i]
		anterior := slice[i-1]

		if anterior >= atual {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("O slice está ordenado")
	} else {
		fmt.Println("O slice não está ordenado")
	}

}
