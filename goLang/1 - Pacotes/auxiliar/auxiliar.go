package auxiliar

import "fmt"

// escreve da função auxiliar || como se fosse pyblico letra minuscula
func Escrever() {

	fmt.Println("Escrevendo do Pacote Auxiliar")
	escrever2() // mesmo pacote então conseguinos utilizar ela e como esta na mesma pasta não precisa fazer o importe
}
