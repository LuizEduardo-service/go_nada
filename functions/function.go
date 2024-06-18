package functions

import "fmt"

func PrintArray(list []int) {
	fmt.Println("Imprimindo valores de um slice")
	for _, value := range list {
		fmt.Printf("%d\t", value)
	}
}

func RemoverNumerosPares(arr []int) (numerosPares []int) {

	fmt.Println("removendo valores pares de um slice")
	for _, value := range arr {

		if value%2 != 0 {
			numerosPares = append(numerosPares, value)
		}

	}
	return
}

func InvertendoLista(arr []int) []int {

	fmt.Println("invertendo lista")
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}

func MovendoZerosParaOFinalDaLista(arr []int) []int {
	fmt.Println("Movendo Zeros para o final da lista")
	j := 0
	for i := 0; i < len(arr); i++ {

		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		if arr[j] != 0 {
			j++
		}

	}
	return arr
}
