package main

import "fmt"

func main() {

	// no go tem 2 maneiras de declarar variaveis explícitos e implícitos

	// implícita
	var variavel1 string = "Variável 1"
	variavel2 := "Variavel 2 " // implícita (nome tecnico é "inferência de tipo") || apenas coloco o nome e o valor e o go esta ferindo o tipo da variável baseado no valor dela

	fmt.Println("Variável : ", variavel1)
	println()
	fmt.Println("Variável :", variavel2)
	println()

	// tipos explícitos
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println("Variável :", variavel3, "Variável :", variavel4)
	println()

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println("Variável :", variavel5, ", Variável :", variavel6)
	println()

	const constante1 string = "Constante 1" // constantye e parecido com variavel porem nao pode se trocar o valor
	fmt.Println("Constante :", constante1)
	println()

	// Invertendo os valores das variaveis
	variavel5, variavel6 = "Variavel 6", "Variavel 5"
	fmt.Println("Variável 5 agora é :", variavel5, ", Variável 6 bagora é :", variavel6)
	println()

}
