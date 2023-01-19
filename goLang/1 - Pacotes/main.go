package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	//	auxiliar.escrever() || não pode ser chamada pois esta com letrra minuscula || escrever not declared by package auxiliar
	println() // para pular uma linha

	erro := checkmail.ValidateFormat("123") // chamada invalida
	fmt.Println(erro)
	println()

	erro = checkmail.ValidateFormat("devBook@gmail.com") // chamada correta
	fmt.Println(erro)

}
