package main

import (
	"fmt"

	"github.com/LuizEduardo-service/go_nada/functions"
)

func main() {

	slice_ := []int{}
	slice_ = append(slice_, 5)
	slice_ = append(slice_, 4)
	slice_ = append(slice_, 0)
	slice_ = append(slice_, 3)
	slice_ = append(slice_, 0)
	slice_ = append(slice_, 10)

	functions.PrintArray(slice_)
	fmt.Println(functions.RemoverNumerosPares(slice_))
	fmt.Println(functions.InvertendoLista(slice_))
	fmt.Println(functions.MovendoZerosParaOFinalDaLista(slice_))

}
