package main

import "fmt"

func main() {

	/* tipo int :
	int = o sistema reconhece de arcordo com a arquitetura do computador
	int8 = aceita ate 8 bytes
	int16 = aceita ate 16 bytes
	int32 = aceita ate 32 bytes
	int64 = aceita ate 64 bytes
	int aceita numeros negativos
	*/
	// var int8 int8 = 1111 || errado nao suporta o valor | cannot use 1111 (untyped int constant) as int8 value in variable declaration (overflows)
	var int8 int8 = 111
	var int16 int16 = 22222
	var int32 int32 = 333333333
	var int64 int64 = 4444444444444444444
	var int int = 5555555555555555555
	fmt.Println("Int8 : ", int8)
	fmt.Println("Int16 : ", int16)
	fmt.Println("Int32 : ", int32)
	fmt.Println("Int64 : ", int64)
	fmt.Println("Int : ", int)
	println()

}
